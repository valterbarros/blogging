package markdown

import (
	"strings"
	"testing"
)

func TestRemoveMetadata(t *testing.T) {
	str := RemoveMetadata(templateA)

	if strings.Count(str, "---") != 0 {
		t.Error("It has metadata")
	}
}

var templateA = `
---
layout: post
title: Para além do medo; O impacto da narrativa da IA nos programadores
---

## Introdução

Antes de qualquer coisa gostaria de deixar claro que utilizo IA todos os dias no meu trabalho, como por exemplo: claude, cursor, copilot, dentro de IDEs como cursor e vscode, não sou de forma alguma uma pessoa contrária à IA. O meu intuíto é tentar enxergar o todo de forma consciente sem aquele véu de animação/desânimo com tecnologia X ou Y que às vezes pode nos cegar. Inclusive o post atual foi feito com ajuda de IA. 😀

`
