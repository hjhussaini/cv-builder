package common

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/skip2/go-qrcode"
)

func QRCode(
	pdf *gofpdf.Fpdf,
	content string,
	x float64,
	y float64,
	size float64,
) {
	imageFile := "/tmp/qrcode.png"
	err := qrcode.WriteFile(content, qrcode.Medium, 256, imageFile)
	if err != nil {
		return
	}

	Image(pdf, x, y, size, size, imageFile)
}
