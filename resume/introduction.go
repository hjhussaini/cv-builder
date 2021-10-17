package resume

import (
	"github.com/hjhussaini/cv-builder/common"
	"github.com/hjhussaini/cv-builder/models"
	"github.com/jung-kurt/gofpdf"
)

func addIntroduction(
	pdf *gofpdf.Fpdf,
	x float64,
	y float64,
	introduction *models.Introduction,
) (float64, float64) {
	imageWidth := 4.0
	imageHeight := 6.0

	common.Image(pdf, x, y, imageWidth, imageHeight, introduction.Photo)

	y = y + imageHeight

	pdf.SetFont("Times", "B", 16)
	_, lineHeight := pdf.GetFontSize()
	y = y + lineHeight
	pdf.Text(x, y, introduction.FullName())

	pdf.SetFont("Times", "", 14)
	y = y + lineHeight
	pdf.Text(x, y, introduction.Headline)

	_, lineHeight = pdf.GetFontSize()
	y = y + lineHeight
	pdf.SetFont("Times", "", 10)
	pdf.Text(x, y, introduction.Country.Name)

	return x, y
}
