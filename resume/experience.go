package resume

import (
	"github.com/hjhussaini/cv-builder/common"
	"github.com/hjhussaini/cv-builder/models"
	"github.com/jung-kurt/gofpdf"
)

func Experience(
	pdf *gofpdf.Fpdf,
	x float64,
	y float64,
	width float64,
	experiences []models.Experience,
) float64 {
	_, y = AddSectionCaption(pdf, x, y, "Experience")

	pdf.SetFont("Times", "B", 12)
	_, lineHeight := pdf.GetFontSize()

	indent := 2.0
	x = x + indent
	y = y + lineHeight

	for _, experience := range experiences {
		pdf.SetFontStyle("B")
		pdf.Text(x, y, experience.Tittle)
		pdf.SetFontStyle("")
		y = y + lineHeight + 0.1
		pdf.Text(x, y, experience.Company)
		y = y + lineHeight + 0.1
		pdf.Text(x, y, experience.Date())
		y = y + lineHeight + 0.1
		pdf.Text(x, y, experience.Location)
		y = y + 1.0
		for _, description := range experience.Description {
			y = common.MultiLineText(pdf, x, y, width-indent, description)
		}

		y = y + 1.0
	}

	return y
}
