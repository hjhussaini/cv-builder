package resume

import (
	"github.com/hjhussaini/cv-builder/models"
	"github.com/jung-kurt/gofpdf"
)

func education(
	pdf *gofpdf.Fpdf,
	x float64,
	y float64,
	educations []models.Education,
) float64 {
	_, y = AddSectionCaption(pdf, x, y, "Education")

	pdf.SetFont("Times", "B", 12)
	_, lineHeight := pdf.GetFontSize()

	x = x + 2.0
	y = y + lineHeight

	for _, education := range educations {
		pdf.Text(x, y, education.School)

		pdf.SetFontStyle("")
		y = y + lineHeight + 0.1
		pdf.Text(x, y, education.FieldOfStudy())

		y = y + lineHeight + 0.1
		pdf.Text(x, y, education.Date())
	}
	y = y + 1.0

	return y
}
