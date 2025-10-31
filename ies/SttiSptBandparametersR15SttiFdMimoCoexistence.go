package ies

import "rrc/utils"

// STTI-SPT-BandParameters-r15-sTTI-FD-MIMO-Coexistence ::= ENUMERATED
type SttiSptBandparametersR15SttiFdMimoCoexistence struct {
	Value utils.ENUMERATED
}

const (
	SttiSptBandparametersR15SttiFdMimoCoexistenceEnumeratedNothing = iota
	SttiSptBandparametersR15SttiFdMimoCoexistenceEnumeratedSupported
)
