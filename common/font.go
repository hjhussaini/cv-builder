package common

import (
	"github.com/jung-kurt/gofpdf"
)

func SetFont(pdf *gofpdf.Fpdf, family string, style string, size float64) float64 {
	pdf.SetFont(family, style, size)
	_, lineHeight := pdf.GetFontSize()

	return lineHeight
}
