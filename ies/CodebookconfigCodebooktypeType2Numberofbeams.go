package ies

import "rrc/utils"

// CodebookConfig-codebookType-type2-numberOfBeams ::= ENUMERATED
type CodebookconfigCodebooktypeType2Numberofbeams struct {
	Value utils.ENUMERATED
}

const (
	CodebookconfigCodebooktypeType2NumberofbeamsEnumeratedNothing = iota
	CodebookconfigCodebooktypeType2NumberofbeamsEnumeratedTwo
	CodebookconfigCodebooktypeType2NumberofbeamsEnumeratedThree
	CodebookconfigCodebooktypeType2NumberofbeamsEnumeratedFour
)
