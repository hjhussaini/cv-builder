package resume

import (
	"github.com/hjhussaini/cv-builder/common"
	"github.com/hjhussaini/cv-builder/models"
	"github.com/jung-kurt/gofpdf"
)

const (
	TOP_MARGIN    = 1.0
	LEFT_MARGIN   = 1.0
	RIGHT_MARGIN  = 0.5
	BOTTOM_MARGIN = 0.5
	BANNER_HEIGHT = 1.5
)

func BuildPDF(sections *models.Sections, fileName string) error {
	pdf := gofpdf.New("Portrait", "cm", "A4", "")

	page1(pdf, sections)
	page2(pdf, sections)
	page3(pdf, sections)

	return pdf.OutputFileAndClose(fileName)
}

func page1(pdf *gofpdf.Fpdf, sections *models.Sections) {
	x, y, width, _ := common.NewPage(
		pdf,
		LEFT_MARGIN,
		RIGHT_MARGIN,
		TOP_MARGIN,
		BOTTOM_MARGIN,
		BANNER_HEIGHT,
	)
	x, y = addIntroduction(pdf, x, y, sections.Introduction)
	AddContact(pdf, x, y, width, sections.Contact)
	x, y = AddAbout(pdf, x, y, width, sections.About)
	x, y = AddSkills(pdf, x, y, width, sections.Skills)
	y = AddTechnologies(pdf, x, y, width, sections.Technologies)
	education(pdf, x, y, sections.Educations)
}

func page2(pdf *gofpdf.Fpdf, sections *models.Sections) {
	x, y, width, _ := common.NewPage(
		pdf,
		LEFT_MARGIN,
		RIGHT_MARGIN,
		TOP_MARGIN,
		BOTTOM_MARGIN,
		BANNER_HEIGHT,
	)
	y = common.Header(pdf, x, y, width, sections.Introduction, sections.Contact)
	y = Experience(pdf, x, y, width, sections.Experiences)
	y = Accomplishments(pdf, x, y, sections.Accomplishments)
	y = Languages(pdf, x, y, sections.Accomplishments.Languages)
}

func page3(pdf *gofpdf.Fpdf, sections *models.Sections) {
	x, y, width, _ := common.NewPage(
		pdf,
		LEFT_MARGIN,
		RIGHT_MARGIN,
		TOP_MARGIN,
		BOTTOM_MARGIN,
		BANNER_HEIGHT,
	)
	y = common.Header(pdf, x, y, width, sections.Introduction, sections.Contact)
	y = y + 1.0
	y = Projects(pdf, x, y, width, sections.Accomplishments.Projects)
}
