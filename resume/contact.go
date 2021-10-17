package resume

import (
	"github.com/hjhussaini/cv-builder/common"
	"github.com/hjhussaini/cv-builder/models"
	"github.com/jung-kurt/gofpdf"
)

func AddContact(
	pdf *gofpdf.Fpdf,
	x float64,
	y float64,
	width float64,
	contact *models.Contact,
) {
	pdf.SetFont("Times", "", 12)
	_, lineHeight := pdf.GetFontSize()
	lineHeight = lineHeight + 0.1

	x = width / 2.0
	y = y - lineHeight
	addInfo(pdf, x, y, "Twitter", contact.Twitter)

	y = y - lineHeight
	addInfo(pdf, x, y, "Skype", contact.Skype)

	y = y - lineHeight
	addInfo(pdf, x, y, "Mobile", contact.Mobile)

	y = y - lineHeight
	addInfo(pdf, x, y, "Email", contact.Email)

	y = y - lineHeight
	addInfo(pdf, x, y, "LinkedIn", contact.LinkedIn)

	y = y - 1.0 - lineHeight*2.5
	imageSize := 1.8
	common.QRCode(pdf, contact.LinkedIn, x, y, imageSize)
	pdf.SetFont("Times", "B", 18)
	_, lineHeight = pdf.GetFontSize()
	AddSectionCaption(
		pdf,
		x+imageSize+0.1,
		y+imageSize-lineHeight-1.5,
		"Contact Info",
	)
}

func addInfo(
	pdf *gofpdf.Fpdf,
	x float64,
	y float64,
	label string,
	text string,
) {
	if text == "" {
		return
	}

	pdf.SetFontStyle("B")
	pdf.Text(x, y, label)
	pdf.SetFontStyle("")
	pdf.Text(x+2.0, y, text)
}
