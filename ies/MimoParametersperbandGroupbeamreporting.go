package ies

import "rrc/utils"

// MIMO-ParametersPerBand-groupBeamReporting ::= ENUMERATED
type MimoParametersperbandGroupbeamreporting struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandGroupbeamreportingEnumeratedNothing = iota
	MimoParametersperbandGroupbeamreportingEnumeratedSupported
)
