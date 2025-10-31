package ies

import "rrc/utils"

// MIMO-ParametersPerBand-mTRP-PUCCH-InterSlot-r17 ::= ENUMERATED
type MimoParametersperbandMtrpPucchInterslotR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandMtrpPucchInterslotR17EnumeratedNothing = iota
	MimoParametersperbandMtrpPucchInterslotR17EnumeratedPf0_2
	MimoParametersperbandMtrpPucchInterslotR17EnumeratedPf1_3_4
	MimoParametersperbandMtrpPucchInterslotR17EnumeratedPf0_4
)
