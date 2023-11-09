package main

import (
	"github.com/go-pdf/fpdf"
	"fmt"
	"strings"
)

func main() {

	var pp PageParams
	
	pp = pp.new(20,15,7)

	var paragraph1 strings.Builder

	paragraph1.WriteString("Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?")

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.SetLeftMargin(pp.mgl)
	pdf.SetRightMargin(pp.mgr)
	pdf.AddPage()
	pdf.SetFont("Times", "B", 16)
	pdf.CellFormat(pp.pgw, pp.tht, "Title", "", 2, "CT", false, 0, "")
	pdf.SetFont("Times", "", 14)
	
	add_paragraph(pdf, pp, paragraph1.String())

	pdf.MultiCell(pp.pgw, pp.lht, paragraph1.String(),"","J",false)
	pdf.MultiCell(pp.pgw, pp.lht, paragraph1.String(),"","L",false)

	pdf.AddPage()
	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		fmt.Print("Can't create file")
	}
}

type PageParams struct {
	mgl float64
	mgr float64
	pgw float64
	spw0 float64
	lht float64
	tht float64
}

func (pp PageParams)new(mgl float64, mgr float64, lht float64) PageParams {
	return PageParams {
		mgl: mgl,
		mgr: mgr,
		pgw: 210.0 - mgl - mgr,
		spw0: 0.1, // do not change
		lht: lht,
		tht: 2.0*lht,
	}
}

func add_paragraph(pdf *fpdf.Fpdf, pp PageParams, text string) {
	lines := pdf.SplitText(text, pp.pgw)
	lnm := len(lines)
	for j := range lines {
		spn := float64(strings.Count(lines[j], " "))
		len := pdf.GetStringWidth(lines[j])
		spl := pp.spw0 * spn
		dl := pp.pgw-len
		spw := pp.spw0 * (spl + dl)/spl
		if j == lnm-1 {
			spw = pp.spw0
		}
		pdf.SetWordSpacing(spw)
		pdf.CellFormat(pp.pgw, pp.lht, lines[j], "", 2, "LT", false, 0, "")
	}
}