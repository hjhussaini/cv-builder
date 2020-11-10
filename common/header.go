package common

import (
	"github.com/hjhussaini/cv-builder/models"
	"github.com/jung-kurt/gofpdf"
)

func Header(
	pdf *gofpdf.Fpdf,
	x float64,
	y float64,
	width float64,
	introduction *models.Introduction,
	contact *models.Contact,
) float64 {
	imageWidth := 2.0
	imageHeight := 3.0
	Image(pdf, x, y, imageWidth, imageHeight, introduction.Photo)

	x = x + imageWidth
	y = y + imageHeight

	pdf.SetFont("Times", "", 12)
	pdf.Text(x, y, introduction.Headline)

	pdf.SetFont("Times", "B", 14)
	_, lineHeight := pdf.GetFontSize()
	pdf.Text(x, y-lineHeight, introduction.FullName())

	x = x - imageWidth
	qrCodeSize := 1.5
	QRCode(pdf, contact.LinkedIn, x+width-qrCodeSize, y/2.0, qrCodeSize)

	return y
}
