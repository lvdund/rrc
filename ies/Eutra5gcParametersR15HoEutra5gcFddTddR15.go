package ies

import "rrc/utils"

// EUTRA-5GC-Parameters-r15-ho-EUTRA-5GC-FDD-TDD-r15 ::= ENUMERATED
type Eutra5gcParametersR15HoEutra5gcFddTddR15 struct {
	Value utils.ENUMERATED
}

const (
	Eutra5gcParametersR15HoEutra5gcFddTddR15EnumeratedNothing = iota
	Eutra5gcParametersR15HoEutra5gcFddTddR15EnumeratedSupported
)
