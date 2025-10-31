package ies

import "rrc/utils"

// BandParameters-v1540-srs-TxSwitch ::= SEQUENCE
type BandparametersV1540SrsTxswitch struct {
	SupportedsrsTxportswitch BandparametersV1540SrsTxswitchSupportedsrsTxportswitch
	Txswitchimpacttorx       *utils.INTEGER `lb:0,ub:32`
	Txswitchwithanotherband  *utils.INTEGER `lb:0,ub:32`
}
