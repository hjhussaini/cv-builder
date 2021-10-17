package resume

import (
	"github.com/hjhussaini/cv-builder/common"
	"github.com/hjhussaini/cv-builder/models"
	"github.com/jung-kurt/gofpdf"
)

func Projects(
	pdf *gofpdf.Fpdf,
	x float64,
	y float64,
	width float64,
	projects []models.Project,
) float64 {
	number := len(projects)
	if number == 0 {
		return y
	}

	indent := 1.2
	y = common.SubsectionTittle(pdf, x, y, indent, number, "Projects")

	lineHeight := common.SetFont(pdf, "Times", "", 12)
	x = x + indent
	for _, project := range projects {
		y = y + 0.5
		pdf.SetFontStyle("B")
		y = y + lineHeight
		pdf.Text(x, y, project.Name)

		pdf.SetFontStyle("")
		if project.URL != "" {
			y = y + lineHeight
			pdf.Text(x, y, project.URL)
		}
		y = y + 0.5
		for _, description := range project.Description {
			y = common.MultiLineText(pdf, x, y, width-indent, description)
		}
	}

	return y
}
