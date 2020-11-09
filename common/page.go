package common

import (
	"github.com/jung-kurt/gofpdf"
)

func NewPage(
	pdf *gofpdf.Fpdf,
	leftMargin float64,
	rightMargin float64,
	topMargin float64,
	bottomMargin float64,
	bannerHeight float64,
) (float64, float64, float64, float64) {
	pdf.AddPage()

	x := leftMargin
	y := topMargin
	width, height := pdf.GetPageSize()

	width = width - rightMargin
	y = Banner(pdf, x, y, width, bannerHeight)

	width = width - leftMargin
	height = height - topMargin - bottomMargin

	return x, y, width, height
}
