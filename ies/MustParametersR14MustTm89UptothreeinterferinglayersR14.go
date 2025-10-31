package ies

import "rrc/utils"

// MUST-Parameters-r14-must-TM89-UpToThreeInterferingLayers-r14 ::= ENUMERATED
type MustParametersR14MustTm89UptothreeinterferinglayersR14 struct {
	Value utils.ENUMERATED
}

const (
	MustParametersR14MustTm89UptothreeinterferinglayersR14EnumeratedNothing = iota
	MustParametersR14MustTm89UptothreeinterferinglayersR14EnumeratedSupported
)
