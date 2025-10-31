package ies

import "rrc/utils"

// CDMA2000-Type ::= ENUMERATED
type Cdma2000Type struct {
	Value utils.ENUMERATED
}

const (
	Cdma2000TypeEnumeratedNothing = iota
	Cdma2000TypeEnumeratedType1xrtt
	Cdma2000TypeEnumeratedTypehrpd
)
