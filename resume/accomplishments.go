package resume

import (
	"github.com/hjhussaini/cv-builder/models"
	"github.com/jung-kurt/gofpdf"
)

func Accomplishments(
	pdf *gofpdf.Fpdf,
	x float64,
	y float64,
	accomplishments *models.Accomplishments,
) float64 {
	_, y = AddSectionCaption(pdf, x, y, "Accomplishments")

	return y
}
