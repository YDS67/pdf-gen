package main

import (
	"github.com/go-pdf/fpdf"
	"fmt"
	"strings"
)

func main() {

	var mgl = 20.0 
	var mgr = 15.0
	var pgw = 210.0 - mgl - mgr
	var spw0 = 0.1 // do not change
	var lnht = 7.0

	var paragraph1 strings.Builder

	paragraph1.WriteString("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	paragraph1.WriteString(" Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?")

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.SetLeftMargin(mgl)
	pdf.SetRightMargin(mgr)
	pdf.SetLineWidth(0.5)
	pdf.AddPage()
	pdf.SetFont("Times", "B", 16)
	pdf.CellFormat(pgw, 2*lnht, "Title", "", 2, "CT", false, 0, "")
	pdf.SetFont("Times", "", 14)
	lines := pdf.SplitText(paragraph1.String(), pgw)
	lnm := len(lines)
	for j := range lines {
		spn := float64(strings.Count(lines[j], " "))
		len := pdf.GetStringWidth(lines[j])
		spl := spw0 * spn
		dl := pgw-len
		spw := spw0 * (spl + dl)/spl
		if j == lnm-1 {
			spw = spw0
		}
		pdf.SetWordSpacing(spw)
		pdf.CellFormat(pgw, lnht, lines[j], "", 2, "LT", false, 0, "")
	}
	pdf.AddPage()
	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		fmt.Print("Can't create file")
	}
}
