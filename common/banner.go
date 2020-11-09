package common

import (
	"github.com/jung-kurt/gofpdf"
)

func Banner(
	pdf *gofpdf.Fpdf,
	x float64,
	y float64,
	width float64,
	height float64,
) float64 {
	pdf.SetFillColor(0, 0, 200)
	pdf.Polygon(
		[]gofpdf.PointType{
			{x, y},
			{width, y},
			{width, y + height},
			{x, y + height/3.0},
		},
		"F",
	)

	return y + height/2.0
}
