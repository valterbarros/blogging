import { EditorState } from "https://esm.sh/@codemirror/state";
import { EditorView, keymap } from "https://esm.sh/@codemirror/view";
import { markdown } from "https://esm.sh/@codemirror/lang-markdown";
import { defaultKeymap, history, historyKeymap } from "https://esm.sh/@codemirror/commands";
import { syntaxHighlighting, defaultHighlightStyle } from "https://esm.sh/@codemirror/language";
import { pasteImagem } from "./image-upload.js";

const ws = new WebSocket("ws://localhost:12769/ws");

const listener = window.addEventListener.bind(window);

const getFileFromUrl = () => {
  const urlParams = new URLSearchParams(window.location.search);
  return urlParams.get('name');
};

const debounceFn = (fn, delay) => {
  let timeoutId;
  return function(...args) {
    clearTimeout(timeoutId);
    timeoutId = setTimeout(() => {
      fn.apply(this, args);
    }, delay);
  };
};

function listenUpdatePreview() {
  ws.onmessage = function(event) {
    const preview = document.getElementById('js-preview');
    preview.innerHTML = event.data;
  };
}

const sendChange = debounceFn((value) => ws.send(JSON.stringify({ type: 'change', value: value, file: getFileFromUrl() })), 100);
let editor;

async function mountEditor(body) {
  const updateList = EditorView.updateListener.of((update) => {
  if (update.docChanged) {
      const textoAtual = update.state.doc.toString();
      
      sendChange(textoAtual);
    }
  });

  const theme = EditorView.theme({
    "&": {
      fontSize: "16px",
      border: "1px solid #ccc",
      borderRadius: "6px",
      backgroundColor: "#f5f5f5",
    },
    ".cm-content": {
        fontFamily: "Arial, sans-serif",
    },
    ".cm-line": {
      textAlign: "justify",
    },
    "&.cm-focused": {
      outline: "2px solid #4c9aff",
    },
    ".cm-scroller": {
      minHeight: "200px",
    },
    ".cm-content": {
      padding: "8px",
      caretColor: "#000",
    },
  });

  (async () => {
    const titleEl = document.querySelector("#js-title");
    const title = body.match(/^title:[\S\s](.+$)/gm)
    titleEl.textContent = title
  })()

  const startState = EditorState.create({
    doc: body,
    extensions: [
      history(),
      keymap.of([...defaultKeymap, ...historyKeymap]),
      markdown(),
      EditorView.lineWrapping,
      syntaxHighlighting(defaultHighlightStyle),

      EditorView.contentAttributes.of({
        spellcheck: "true",
        lang: "pt-BR"
      }),
      updateList,
      theme,
      pasteImagem,
    ]
  });

  const parent = document.getElementById("editor-container")
  editor = new EditorView({
    state: startState,
    parent
  });
  editor.focus();
}

function getTextContent(element) {
  return element.content.textContent
}

listener('DOMContentLoaded', async () => {
  const temp = document.getElementById("js-data-article");

  listenUpdatePreview();
  mountEditor(getTextContent(temp));
  sendChange(getTextContent(temp));
});

listener('keydown', (e) => {
  if ((e.ctrlKey || e.metaKey) && e.key === 's') {
    e.preventDefault()
    console.log('save');

    const value = editor.state.doc.toString();
    ws.send(JSON.stringify({ type: 'update', value, file: getFileFromUrl() }));
  }
});
