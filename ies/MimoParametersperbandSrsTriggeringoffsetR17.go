package ies

import "rrc/utils"

// MIMO-ParametersPerBand-srs-TriggeringOffset-r17 ::= ENUMERATED
type MimoParametersperbandSrsTriggeringoffsetR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSrsTriggeringoffsetR17EnumeratedNothing = iota
	MimoParametersperbandSrsTriggeringoffsetR17EnumeratedN1
	MimoParametersperbandSrsTriggeringoffsetR17EnumeratedN2
	MimoParametersperbandSrsTriggeringoffsetR17EnumeratedN4
)
