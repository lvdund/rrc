package ies

import "rrc/utils"

// EUTRA-5GC-Parameters-r15-reflectiveQoS-r15 ::= ENUMERATED
type Eutra5gcParametersR15ReflectiveqosR15 struct {
	Value utils.ENUMERATED
}

const (
	Eutra5gcParametersR15ReflectiveqosR15EnumeratedNothing = iota
	Eutra5gcParametersR15ReflectiveqosR15EnumeratedSupported
)
