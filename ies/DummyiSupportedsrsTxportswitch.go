package ies

import "rrc/utils"

// DummyI-supportedSRS-TxPortSwitch ::= ENUMERATED
type DummyiSupportedsrsTxportswitch struct {
	Value utils.ENUMERATED
}

const (
	DummyiSupportedsrsTxportswitchEnumeratedNothing = iota
	DummyiSupportedsrsTxportswitchEnumeratedT1r2
	DummyiSupportedsrsTxportswitchEnumeratedT1r4
	DummyiSupportedsrsTxportswitchEnumeratedT2r4
	DummyiSupportedsrsTxportswitchEnumeratedT1r4_T2r4
	DummyiSupportedsrsTxportswitchEnumeratedTr_Equal
)
