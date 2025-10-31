package ies

import "rrc/utils"

// MIMO-ParametersPerBand-sfn-QCL-TypeD-Collision-twoTCI-r17 ::= ENUMERATED
type MimoParametersperbandSfnQclTypedCollisionTwotciR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSfnQclTypedCollisionTwotciR17EnumeratedNothing = iota
	MimoParametersperbandSfnQclTypedCollisionTwotciR17EnumeratedSupported
)
