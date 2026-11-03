import { EditorView } from "https://esm.sh/@codemirror/view";

async function sendImage(file) {
  const form = new FormData();
  form.append("file", file);

  const res = await fetch("http://localhost:12769/upload", { method: "POST", body: form });
  if (!res.ok) throw new Error("Falha no upload");
  const { url } = await res.json(); // ajuste conforme sua API
  return url;
}

function replace(view, procurar, novo) {
  const idx = view.state.doc.toString().indexOf(procurar);
  if (idx === -1) return;
  view.dispatch({
    changes: { from: idx, to: idx + procurar.length, insert: novo },
  });
}

function insertTexto(view, texto) {
  const { from, to } = view.state.selection.main;
  view.dispatch({
    changes: { from, to, insert: texto },
    selection: { anchor: from + texto.length },
    scrollIntoView: true,
  });
}

export const pasteImagem = EditorView.domEventHandlers({
  paste(event, view) {
    const imagens = [...(event.clipboardData?.items ?? [])]
      .filter((i) => i.kind === "file" && i.type.startsWith("image/"))
      .map((i) => i.getAsFile())
      .filter(Boolean);

    if (!imagens.length) return false;
    event.preventDefault();

    imagens.forEach(async (file) => {
      console.log('image foreach');
      
      const id = Math.random().toString(36).slice(2, 8);
      const mark = `![enviando-${id}...]()`;
      insertTexto(view, mark + "\n");

      try {
        const url = await sendImage(file);
        console.log(url);

        replace(view, mark, `![imagem](${url})`);
      } catch {
        replace(view, mark, "");
      }
    });

    return true;
  },
});
