package ies

import "rrc/utils"

// DL-Bitmap-NB-r13 ::= CHOICE
const (
	DlBitmapNbR13ChoiceNothing = iota
	DlBitmapNbR13ChoiceSubframepattern10R13
	DlBitmapNbR13ChoiceSubframepattern40R13
)

type DlBitmapNbR13 struct {
	Choice               uint64
	Subframepattern10R13 *utils.BITSTRING `lb:10,ub:10`
	Subframepattern40R13 *utils.BITSTRING `lb:40,ub:40`
}
