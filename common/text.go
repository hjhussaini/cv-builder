package common

import (
	"github.com/jung-kurt/gofpdf"
)

func MultiLineText(
	pdf *gofpdf.Fpdf,
	x float64,
	y float64,
	width float64,
	text string,
) float64 {
	_, lineHeight := pdf.GetFontSize()
	textWidth := pdf.GetStringWidth(text)
	lines := int(textWidth)/int(width) + 1

	pdf.MoveTo(x, y)
	pdf.MultiCell(
		width, lineHeight,
		text,
		gofpdf.BorderNone,
		gofpdf.AlignLeft,
		false,
	)

	return y + float64(lines)*lineHeight
}
