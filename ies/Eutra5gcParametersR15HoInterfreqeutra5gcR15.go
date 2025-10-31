package ies

import "rrc/utils"

// EUTRA-5GC-Parameters-r15-ho-InterfreqEUTRA-5GC-r15 ::= ENUMERATED
type Eutra5gcParametersR15HoInterfreqeutra5gcR15 struct {
	Value utils.ENUMERATED
}

const (
	Eutra5gcParametersR15HoInterfreqeutra5gcR15EnumeratedNothing = iota
	Eutra5gcParametersR15HoInterfreqeutra5gcR15EnumeratedSupported
)
