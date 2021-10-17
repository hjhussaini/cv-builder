package resume

import (
	"github.com/jung-kurt/gofpdf"
)

func AddTechnologies(
	pdf *gofpdf.Fpdf,
	x float64,
	y float64,
	width float64,
	technologies []string,
) float64 {
	x, y = AddSectionCaption(pdf, x, y, "Tools & Technologies")

	pdf.SetFont("Times", "B", 12)
	_, lineHeight := pdf.GetFontSize()
	y_ := y

	for index := 0; index < len(technologies); index += 2 {
		y = y + lineHeight + 0.1
		pdf.Text(x, y, technologies[index])
	}

	x_ := width / 2.0
	for index := 1; index < len(technologies); index += 2 {
		y_ = y_ + lineHeight + 0.1
		pdf.Text(x_, y_, technologies[index])
	}

	return y
}
