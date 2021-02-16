package output

import (
	"dkstar88/mathgen/generator"
	"fmt"
	"github.com/signintech/gopdf"
	log "github.com/sirupsen/logrus"
	"path"
	"strings"
)

func initPdf(pdf *gopdf.GoPdf, header string) {
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	if err := pdf.AddTTFFont("roboto", "fonts/Roboto-Regular.ttf"); err != nil {
		log.Error(err.Error())
		return
	}

	if err := pdf.SetFont("roboto", "", 14); err != nil {
		log.Error(err.Error())
		return
	}

	pdf.SetX(30.0)
	pdf.SetY(40.0)
	CheckErr(pdf.Text(header))

	if err := pdf.SetFont("roboto", "", 11); err != nil {
		log.Error(err.Error())
		return
	}

	pdf.SetMargins(30, 70, 30, 50)
}

func CheckErr(err error) {
	if err != nil {
		log.Error(err.Error())
	}
}

func GenerateTestPaper(questions []generator.QuestionAnswer, title, testPdf, answerPdf string) {

	pdf := gopdf.GoPdf{}
	pdfAnswer := gopdf.GoPdf{}
	initPdf(&pdf, title+", Date: ______________   Name: ______________")
	initPdf(&pdfAnswer, title+", Answer Book")

	margin := gopdf.PageSizeA4
	margin.W = margin.W - pdf.MarginRight() - pdf.MarginLeft()
	margin.H = margin.H - pdf.MarginTop() - pdf.MarginBottom()

	x := pdf.MarginLeft()
	y := pdf.MarginTop()

	for _, q := range questions {
		pdf.SetX(x)
		pdf.SetY(y)
		CheckErr(pdf.Text(q.Question + "="))

		pdfAnswer.SetX(x)
		pdfAnswer.SetY(y)

		CheckErr(pdfAnswer.Text(q.Question + "=" + q.Answer))

		x += 140
		if x >= margin.W {
			x = pdf.MarginLeft()
			y += 25
		}
		if y >= margin.H {
			x = pdf.MarginLeft()
			y = pdf.MarginTop()
			pdf.AddPage()
		}
		fmt.Println(q.Question + "=" + q.Answer)
	}

	testPdf = appendFileExt(testPdf, ".pdf")
	if len(answerPdf) == 0 {
		answerPdf = changeFileExt(testPdf, "-answers.pdf")
	}

	CheckErr(pdf.WritePdf(testPdf))
	CheckErr(pdfAnswer.WritePdf(answerPdf))
}

func changeFileExt(filename string, newExt string) string {
	ext := path.Ext(filename)
	return filename[0:len(filename)-len(ext)] + newExt
}

func appendFileExt(filename string, ext string) string {
	currExt := path.Ext(filename)
	if !strings.EqualFold(currExt, ext) {
		return changeFileExt(filename, ext)
	}
	return filename
}
