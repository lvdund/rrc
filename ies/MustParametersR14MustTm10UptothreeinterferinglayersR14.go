package ies

import "rrc/utils"

// MUST-Parameters-r14-must-TM10-UpToThreeInterferingLayers-r14 ::= ENUMERATED
type MustParametersR14MustTm10UptothreeinterferinglayersR14 struct {
	Value utils.ENUMERATED
}

const (
	MustParametersR14MustTm10UptothreeinterferinglayersR14EnumeratedNothing = iota
	MustParametersR14MustTm10UptothreeinterferinglayersR14EnumeratedSupported
)
