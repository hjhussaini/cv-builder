package resume

import (
	"github.com/hjhussaini/cv-builder/common"
	"github.com/hjhussaini/cv-builder/models"
	"github.com/jung-kurt/gofpdf"
)

func Languages(
	pdf *gofpdf.Fpdf,
	x float64,
	y float64,
	languages []models.Language,
) float64 {
	number := len(languages)
	if number == 0 {
		return y
	}

	indent := 1.2
	y = common.SubsectionTittle(pdf, x, y, indent, number, "Languages")
	lineHeight := common.SetFont(pdf, "Times", "", 12)
	x = x + indent
	for _, language := range languages {
		pdf.SetFontStyle("B")
		y = y + lineHeight
		pdf.Text(x, y, language.Language)

		y = y + lineHeight
		pdf.SetFontStyle("")
		pdf.Text(x, y, language.Proficiency+" proficiency")

		y = y + 0.5
	}

	return y
}
