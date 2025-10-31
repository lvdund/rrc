package ies

import "rrc/utils"

// DummyI ::= SEQUENCE
type Dummyi struct {
	SupportedsrsTxportswitch DummyiSupportedsrsTxportswitch
	Txswitchimpacttorx       *utils.BOOLEAN
}
