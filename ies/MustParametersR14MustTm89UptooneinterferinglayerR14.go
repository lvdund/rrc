package ies

import "rrc/utils"

// MUST-Parameters-r14-must-TM89-UpToOneInterferingLayer-r14 ::= ENUMERATED
type MustParametersR14MustTm89UptooneinterferinglayerR14 struct {
	Value utils.ENUMERATED
}

const (
	MustParametersR14MustTm89UptooneinterferinglayerR14EnumeratedNothing = iota
	MustParametersR14MustTm89UptooneinterferinglayerR14EnumeratedSupported
)
