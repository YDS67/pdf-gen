package main

import (
	"fmt"
	"github.com/go-pdf/fpdf"
	"strings"
)

func main() {

	var pp PageParams

	pp = pp.new(20, 15, 20, 7)

	var part1 strings.Builder
	var chapter11 strings.Builder
	var paragraph111 strings.Builder
	var paragraph112 strings.Builder
	var paragraph113 strings.Builder
	var paragraph114 strings.Builder
	var paragraph115 strings.Builder
	var paragraph116 strings.Builder
	var paragraph117 strings.Builder
	var paragraph118 strings.Builder
	var paragraph119 strings.Builder
	var paragraph1110 strings.Builder
	var paragraph1111 strings.Builder
	var paragraph1112 strings.Builder

	part1.WriteString("PART 1 - DROPPED FROM THE CLOUDS ")
	chapter11.WriteString("Chapter 1")
	paragraph111.WriteString("\"Are we rising again?\" \"No. On the contrary.\" \"Are we descending?\" \"Worse than that, captain! we are falling!\" \"For Heaven's sake heave out the ballast!\" \"There! the last sack is empty!\" \"Does the balloon rise?\" \"No!\" \"I hear a noise like the dashing of waves. The sea is below the car! It cannot be more than 500 feet from us!\" \"Overboard with every weight! ... everything!\"")
	paragraph112.WriteString("Such were the loud and startling words which resounded through the air, above the vast watery desert of the Pacific, about four o'clock in the evening of the 23rd of March, 1865.")
	paragraph113.WriteString("Few can possibly have forgotten the terrible storm from the northeast, in the middle of the equinox of that year. The tempest raged without intermission from the 18th to the 26th of March. Its ravages were terrible in America, Europe, and Asia, covering a distance of eighteen hundred miles, and extending obliquely to the equator from the thirty-fifth north parallel to the fortieth south parallel. Towns were overthrown, forests uprooted, coasts devastated by the mountains of water which were precipitated on them, vessels cast on the shore, which the published accounts numbered by hundreds, whole districts leveled by waterspouts which destroyed everything they passed over, several thousand people crushed on land or drowned at sea; such were the traces of its fury, left by this devastating tempest. It surpassed in disasters those which so frightfully ravaged Havana and Guadalupe, one on the 25th of October, 1810, the other on the 26th of July, 1825.")
	paragraph114.WriteString("But while so many catastrophes were taking place on land and at sea, a drama not less exciting was being enacted in the agitated air.")
	paragraph115.WriteString("In fact, a balloon, as a ball might be carried on the summit of a waterspout, had been taken into the circling movement of a column of air and had traversed space at the rate of ninety miles an hour, turning round and round as if seized by some aerial maelstrom.")
	paragraph116.WriteString("Beneath the lower point of the balloon swung a car, containing five passengers, scarcely visible in the midst of the thick vapor mingled with spray which hung over the surface of the ocean.")
	paragraph117.WriteString("Whence, it may be asked, had come that plaything of the tempest? From what part of the world did it rise? It surely could not have started during the storm. But the storm had raged five days already, and the first symptoms were manifested on the 18th. It cannot be doubted that the balloon came from a great distance, for it could not have traveled less than two thousand miles in twenty-four hours.")
	paragraph118.WriteString("At any rate the passengers, destitute of all marks for their guidance, could not have possessed the means of reckoning the route traversed since their departure. It was a remarkable fact that, although in the very midst of the furious tempest, they did not suffer from it. They were thrown about and whirled round and round without feeling the rotation in the slightest degree, or being sensible that they were removed from a horizontal position.")
	paragraph119.WriteString("Their eyes could not pierce through the thick mist which had gathered beneath the car. Dark vapor was all around them. Such was the density of the atmosphere that they could not be certain whether it was day or night. No reflection of light, no sound from inhabited land, no roaring of the ocean could have reached them, through the obscurity, while suspended in those elevated zones. Their rapid descent alone had informed them of the dangers which they ran from the waves. However, the balloon, lightened of heavy articles, such as ammunition, arms, and provisions, had risen into the higher layers of the atmosphere, to a height of 4,500 feet. ")
	paragraph119.WriteString("The voyagers, after having discovered that the sea extended beneath them, and thinking the dangers above less dreadful than those below, did not hesitate to throw overboard even their most useful articles, while they endeavored to lose no more of that fluid, the life of their enterprise, which sustained them above the abyss.")
	paragraph1110.WriteString("The night passed in the midst of alarms which would have been death to less energetic souls. Again the day appeared and with it the tempest began to moderate. From the beginning of that day, the 24th of March, it showed symptoms of abating. At dawn, some of the lighter clouds had risen into the more lofty regions of the air. In a few hours the wind had changed from a hurricane to a fresh breeze, that is to say, the rate of the transit of the atmospheric layers was diminished by half. It was still what sailors call \"a close-reefed topsail breeze,\" but the commotion in the elements had none the less considerably diminished.")
	paragraph1111.WriteString("Towards eleven o'clock, the lower region of the air was sensibly clearer. The atmosphere threw off that chilly dampness which is felt after the passage of a great meteor. The storm did not seem to have gone farther to the west. It appeared to have exhausted itself. Could it have passed away in electric sheets, as is sometimes the case with regard to the typhoons of the Indian Ocean?")
	paragraph1112.WriteString("But at the same time, it was also evident that the balloon was again slowly descending with a regular movement. It appeared as if it were, little by little, collapsing, and that its case was lengthening and extending, passing from a spherical to an oval form. Towards midday the balloon was hovering above the sea at a height of only 2,000 feet. It contained 50,000 cubic feet of gas, and, thanks to its capacity, it could maintain itself a long time in the air, although it should reach a great altitude or might be thrown into a horizontal position.")

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.SetLeftMargin(pp.mgl)
	pdf.SetRightMargin(pp.mgr)
	pdf.SetTopMargin(pp.mgt)
	pdf.AddPage()

	pdf.SetTextColor(0,55,165)
	add_main_title(pdf, pp, part1.String())

	pdf.SetTextColor(0,0,0)
	add_title(pdf, pp, chapter11.String())
	add_paragraph(pdf, pp, paragraph111.String())
	add_paragraph(pdf, pp, paragraph112.String())
	add_paragraph(pdf, pp, paragraph113.String())
	add_paragraph(pdf, pp, paragraph114.String())
	add_paragraph(pdf, pp, paragraph115.String())
	add_paragraph(pdf, pp, paragraph116.String())
	add_paragraph(pdf, pp, paragraph117.String())
	add_paragraph(pdf, pp, paragraph118.String())
	add_paragraph(pdf, pp, paragraph119.String())
	add_paragraph(pdf, pp, paragraph1110.String())
	add_paragraph(pdf, pp, paragraph1111.String())
	add_paragraph(pdf, pp, paragraph1112.String())

	err := pdf.OutputFileAndClose("MysteriousIsland.pdf")
	if err != nil {
		fmt.Print("Can't create file")
	}
}

type PageParams struct {
	mgl  float64
	mgr  float64
	mgt  float64
	pgw  float64
	pgh  float64
	spw0 float64
	spw1 float64
	lht  float64
	tht  float64
	fntext string
	fnttexts float64
	fntitle string
	fntitles float64
	fnmain string
	fnmains float64
	fntpage string
	fntpages float64
}

func (pp PageParams) new(mgl, mgr, mgt, lht float64) PageParams {
	return PageParams{
		mgl:  mgl,
		mgr:  mgr,
		mgt:  mgt,
		pgw:  210.0 - mgl - mgr,
		pgh:  297.0 - 2.0*mgt,
		spw0: 0.1, // do not change
		spw1: 0.2,
		lht:  lht,
		tht:  1.5 * lht,
		fntext: "Times",
		fnttexts: 14,
		fntitle: "Arial",
		fntitles: 16,
		fnmain: "Arial",
		fnmains: 18,
		fntpage: "Helvetica",
		fntpages: 13,
	}
}

func add_paragraph(pdf *fpdf.Fpdf, pp PageParams, text string) {
	pdf.SetFont(pp.fntext, "", pp.fnttexts)
	lines := pdf.SplitText(text, pp.pgw)
	lnm := len(lines)
	for j := range lines {
		page_num(pdf, &pp, "text")
		spn := float64(strings.Count(lines[j], " "))
		len := pdf.GetStringWidth(lines[j])
		spl := pp.spw0 * spn
		dl := pp.pgw - len
		spw := pp.spw0 * (spl + dl) / spl
		lht := pp.lht
		if j == lnm-1 {
			spw = pp.spw0
			lht = pp.tht
		}
		pdf.SetWordSpacing(spw)
		pdf.CellFormat(pp.pgw, lht, lines[j], "", 2, "LT", false, 0, "")
	}
}

func add_title(pdf *fpdf.Fpdf, pp PageParams, text string) {
	pdf.SetFont(pp.fntitle, "B", pp.fntitles)
	lines := pdf.SplitText(text, pp.pgw)
	lnm := len(lines)
	for j := range lines {
		page_num(pdf, &pp, "title")
		lht := pp.lht
		if j == lnm-1 {
			lht = pp.tht
		}
		pdf.CellFormat(pp.pgw, lht, lines[j], "", 2, "CT", false, 0, "")
	}
}

func add_main_title(pdf *fpdf.Fpdf, pp PageParams, text string) {
	pdf.SetFont(pp.fnmain, "B", pp.fnmains)
	lines := pdf.SplitText(text, pp.pgw)
	lnm := len(lines)
	for j := range lines {
		page_num(pdf, &pp, "main")
		lht := pp.tht
		if j == lnm-1 {
			lht = 2.0*pp.tht
		}
		pdf.CellFormat(pp.pgw, lht, lines[j], "", 2, "CT", false, 0, "")
	}
}

func add_title_just(pdf *fpdf.Fpdf, pp PageParams, text string) {
	pdf.SetFont(pp.fntitle, "B", pp.fntitles)
	lines := pdf.SplitText(text, pp.pgw)
	lnm := len(lines)
	for j := range lines {
		page_num(pdf, &pp, "title")
		spn := float64(strings.Count(lines[j], " "))
		len := pdf.GetStringWidth(lines[j])
		spl := pp.spw1 * spn
		dl := pp.pgw - len
		spw := pp.spw1 * (spl + dl) / spl
		lht := pp.lht
		if j == lnm-1 {
			spw = pp.spw1
			lht = pp.tht
		}
		pdf.SetWordSpacing(spw)
		pdf.CellFormat(pp.pgw, lht, lines[j], "", 2, "LT", false, 0, "")
	}
}

func add_main_title_just(pdf *fpdf.Fpdf, pp PageParams, text string) {
	pdf.SetFont(pp.fnmain, "B", pp.fnmains)
	lines := pdf.SplitText(text, pp.pgw)
	lnm := len(lines)
	for j := range lines {
		page_num(pdf, &pp, "main")
		spn := float64(strings.Count(lines[j], " "))
		len := pdf.GetStringWidth(lines[j])
		spl := pp.spw0 * spn
		dl := pp.pgw - len
		spw := pp.spw0 * (spl + dl) / spl
		lht := pp.tht
		if j == lnm-1 {
			spw = pp.spw0
			lht = 2.0*pp.tht
		}
		pdf.SetWordSpacing(spw)
		pdf.CellFormat(pp.pgw, lht, lines[j], "", 2, "LT", false, 0, "")
	}
}

func page_num(pdf *fpdf.Fpdf, pp *PageParams, which string) {
	pos := pdf.GetY()
	var fnt string
	var fntd string
	var fnts float64
		switch which {
		case "title":
			fnt = pp.fntitle
			fntd = "B"
			fnts = pp.fntitles
		case "main":
			fnt = pp.fnmain
			fntd = "B"
			fnts = pp.fnmains
		case "text":
			fnt = pp.fntext
			fntd = ""
			fnts = pp.fnttexts
		}
		if pos > pp.pgh-pp.tht {
			cr, cg, cb := pdf.GetTextColor()
			pdf.SetFont(pp.fntpage, "", pp.fntpages)
			pdf.SetTextColor(0,0,0)
			pdf.CellFormat(pp.pgw, pp.tht, fmt.Sprint(pdf.PageNo()), "", 2, "CB", false, 0, "")
			pdf.SetFont(fnt, fntd, fnts)
			pdf.CellFormat(pp.pgw, pp.lht, "", "", 2, "CM", false, 0, "")
			pdf.AddPage()
			pdf.SetTextColor(cr,cg,cb)
		}
}