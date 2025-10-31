package ies

import "rrc/utils"

// MIMO-ParametersPerBand-srs-combEight-r17 ::= ENUMERATED
type MimoParametersperbandSrsCombeightR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSrsCombeightR17EnumeratedNothing = iota
	MimoParametersperbandSrsCombeightR17EnumeratedSupported
)
