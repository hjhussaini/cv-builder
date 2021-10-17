package resume

import (
	"github.com/jung-kurt/gofpdf"
)

func AddSkills(
	pdf *gofpdf.Fpdf,
	x float64,
	y float64,
	width float64,
	skills []string,
) (float64, float64) {
	x, y = AddSectionCaption(pdf, x, y, "Skills")

	pdf.SetFont("Times", "B", 12)
	_, lineHeight := pdf.GetFontSize()
	for _, text := range skills {
		y = y + lineHeight + 0.1
		pdf.Text(x, y, text)
	}

	return x, y
}
