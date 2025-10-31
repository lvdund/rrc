package ies

import "rrc/utils"

// AC-BarringPerPLMN-r12-ac-BarringInfo-r12 ::= SEQUENCE
type AcBarringperplmnR12AcBarringinfoR12 struct {
	AcBarringforemergencyR12    utils.BOOLEAN
	AcBarringformoSignallingR12 *AcBarringconfig
	AcBarringformoDataR12       *AcBarringconfig
}
