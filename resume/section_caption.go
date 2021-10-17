package resume

import (
	"github.com/jung-kurt/gofpdf"
)

func AddSectionCaption(
	pdf *gofpdf.Fpdf,
	x float64,
	y float64,
	caption string,
) (float64, float64) {
	pdf.SetFont("Times", "B", 14)
	_, lineHeight := pdf.GetFontSize()
	y = y + 1.0 + lineHeight
	pdf.Text(x, y, caption)

	pdf.Rect(x, y+0.2, 5.0, 0.15, "F")

	return x, y + lineHeight
}
