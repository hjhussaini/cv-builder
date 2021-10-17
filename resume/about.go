package resume

import (
	"github.com/hjhussaini/cv-builder/common"
	"github.com/jung-kurt/gofpdf"
)

func AddAbout(
	pdf *gofpdf.Fpdf,
	x float64,
	y float64,
	width float64,
	about []string,
) (float64, float64) {
	x, y = AddSectionCaption(pdf, x, y, "About")

	pdf.SetFont("Times", "", 12)
	for _, text := range about {
		y = common.MultiLineText(pdf, x, y, width, text)
	}

	return x, y
}
