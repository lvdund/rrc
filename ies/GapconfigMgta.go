package ies

import "rrc/utils"

// GapConfig-mgta ::= ENUMERATED
type GapconfigMgta struct {
	Value utils.ENUMERATED
}

const (
	GapconfigMgtaEnumeratedNothing = iota
	GapconfigMgtaEnumeratedMs0
	GapconfigMgtaEnumeratedMs0dot25
	GapconfigMgtaEnumeratedMs0dot5
)
