package ies

import "rrc/utils"

// DL-Bitmap-NB-r13 ::= CHOICE
type DlBitmapNbR13 interface {
	isDlBitmapNbR13()
}

type DlBitmapNbR13Subframepattern10R13 struct {
	Value utils.BITSTRING
}

func (*DlBitmapNbR13Subframepattern10R13) isDlBitmapNbR13() {}

type DlBitmapNbR13Subframepattern40R13 struct {
	Value utils.BITSTRING
}

func (*DlBitmapNbR13Subframepattern40R13) isDlBitmapNbR13() {}
