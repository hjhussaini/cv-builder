package common

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func SubsectionTittle(
	pdf *gofpdf.Fpdf,
	x float64,
	y float64,
	indent float64,
	number int,
	tittle string,
) float64 {
	red, green, blue := pdf.GetTextColor()
	pdf.SetTextColor(0, 0, 200)

	lineHeight := SetFont(pdf, "Times", "B", 32)
	y = y + lineHeight
	if number < 10 {
		pdf.Text(x+indent/2.0, y, fmt.Sprintf("%d", number))
	} else {
		pdf.Text(x, y, tittle)
	}

	pdf.SetFontSize(14)
	pdf.Text(x+indent, y-lineHeight/3.0, tittle)

	pdf.SetTextColor(red, green, blue)

	return y
}
