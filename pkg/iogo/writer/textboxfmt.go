package writer

type textboxfmt struct {
	CornerTopLeft     string
	CornerTopRight    string
	CornerBottomLeft  string
	CornerBottomRight string
	WallHeader        string
	WallFooter        string
	WallVertical      string
}

// tbfSimple Renders:
// +--+
// |  |
// +--+
func tbfSimple() textboxfmt {
	return textboxfmt{
		CornerTopLeft:     "+",
		CornerTopRight:    "+",
		CornerBottomLeft:  "+",
		CornerBottomRight: "+",
		WallHeader:        "-",
		WallFooter:        "_",
		WallVertical:      "|",
	}
}

// tbfUnicodeNonSolid Renders:
// ╔══╗
// ║  ║
// ╚══╝
func tbfUnicodeNonSolid() textboxfmt {
	return textboxfmt{
		CornerTopLeft:     "╔",
		CornerTopRight:    "╗",
		CornerBottomLeft:  "╚",
		CornerBottomRight: "╝",
		WallHeader:        "═",
		WallFooter:        "═",
		WallVertical:      "║",
	}
}

// tbfUnicodeSolid Renders:
// ┏━━┓
// ┃  ┃
// ┗━━┛
func tbfUnicodeSolid() textboxfmt {
	return textboxfmt{
		CornerTopLeft:     "┏",
		CornerTopRight:    "┓",
		CornerBottomLeft:  "┗",
		CornerBottomRight: "┛",
		WallHeader:        "━",
		WallFooter:        "━",
		WallVertical:      "┃",
	}
}
