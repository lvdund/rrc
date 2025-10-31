package ies

import "rrc/utils"

// DL-PPW-PreConfig-r17-type-r17 ::= ENUMERATED
type DlPpwPreconfigR17TypeR17 struct {
	Value utils.ENUMERATED
}

const (
	DlPpwPreconfigR17TypeR17EnumeratedNothing = iota
	DlPpwPreconfigR17TypeR17EnumeratedType1a
	DlPpwPreconfigR17TypeR17EnumeratedType1b
	DlPpwPreconfigR17TypeR17EnumeratedType2
)
