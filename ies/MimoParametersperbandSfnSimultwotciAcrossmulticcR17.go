package ies

import "rrc/utils"

// MIMO-ParametersPerBand-sfn-SimulTwoTCI-AcrossMultiCC-r17 ::= ENUMERATED
type MimoParametersperbandSfnSimultwotciAcrossmulticcR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSfnSimultwotciAcrossmulticcR17EnumeratedNothing = iota
	MimoParametersperbandSfnSimultwotciAcrossmulticcR17EnumeratedSupported
)
