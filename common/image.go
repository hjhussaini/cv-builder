package common

import (
	"github.com/jung-kurt/gofpdf"
)

func Image(
	pdf *gofpdf.Fpdf,
	x float64,
	y float64,
	width float64,
	height float64,
	imageFile string,
) {
	pdf.ImageOptions(
		imageFile,
		x, y,
		width, height,
		false,
		gofpdf.ImageOptions{
			ReadDpi: true,
		},
		0,
		"",
	)
}
