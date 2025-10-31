package ies

import "rrc/utils"

// MeasCSI-RS-Id-r12 ::= utils.INTEGER (1..maxCSI-RS-Meas-r12)
type MeascsiRsIdR12 struct {
	Value utils.INTEGER `lb:0,ub:maxCSIRsMeasR12`
}
