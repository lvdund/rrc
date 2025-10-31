package ies

import "rrc/utils"

// MIMO-ParametersPerBand-srs-TriggeringDCI-r17 ::= ENUMERATED
type MimoParametersperbandSrsTriggeringdciR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSrsTriggeringdciR17EnumeratedNothing = iota
	MimoParametersperbandSrsTriggeringdciR17EnumeratedSupported
)
