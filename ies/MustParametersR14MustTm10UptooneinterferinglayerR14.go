package ies

import "rrc/utils"

// MUST-Parameters-r14-must-TM10-UpToOneInterferingLayer-r14 ::= ENUMERATED
type MustParametersR14MustTm10UptooneinterferinglayerR14 struct {
	Value utils.ENUMERATED
}

const (
	MustParametersR14MustTm10UptooneinterferinglayerR14EnumeratedNothing = iota
	MustParametersR14MustTm10UptooneinterferinglayerR14EnumeratedSupported
)
