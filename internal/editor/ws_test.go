package editor

import (
	"strings"
	"testing"
)

// TODO: maybe move to markdown package
func TestTemplateParse(t *testing.T) {
	t.Run("Should run with filled string", func(t *testing.T) {
		text, err := processImageUrl(`---
	layout: post
	title: O perigo invisível dos projetos legados
	---
	![imagem]({{ "/articles/imgs/Screenshot-de-matematica-aqui.png" | relative_url }})
	
	### Introdução
	
	No começo da minha carreira eu trabalhava como fullstack, me lembro que na época eu precisava lidar com várias tecnologias e linguagens diferentes como sql, css, scss, coffeescript, jquery, js, ruby on rails, etc. Acredito que esse seja um cenário "comum" para muitos devs. Desde então novas tecnologias surgiram e outras morreram, muita coisa que eu fazia nessa época mudou e algumas até radicalmente.
	`)

		if strings.Contains(text, "relative_url") {
			t.Error("It contains the relative_url function", text)
		}

		if err != nil {
			t.Error("An error ocurred", err)
		}
	})
	t.Run("should run with empty string", func(t *testing.T) {
		text, err := processImageUrl("")

		if text != "" {
			t.Error("It contains value", text)
		}

		if err != nil {
			t.Error("An error ocurred", err)
		}
	})
}
