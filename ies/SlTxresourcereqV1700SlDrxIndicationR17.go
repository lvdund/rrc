package ies

import "rrc/utils"

// SL-TxResourceReq-v1700-sl-DRX-Indication-r17 ::= ENUMERATED
type SlTxresourcereqV1700SlDrxIndicationR17 struct {
	Value utils.ENUMERATED
}

const (
	SlTxresourcereqV1700SlDrxIndicationR17EnumeratedNothing = iota
	SlTxresourcereqV1700SlDrxIndicationR17EnumeratedOn
	SlTxresourcereqV1700SlDrxIndicationR17EnumeratedOff
)
