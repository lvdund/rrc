package ies

import "rrc/utils"

// SL-TxResourceReqDisc-r17-sl-CastTypeDisc-r17 ::= ENUMERATED
type SlTxresourcereqdiscR17SlCasttypediscR17 struct {
	Value utils.ENUMERATED
}

const (
	SlTxresourcereqdiscR17SlCasttypediscR17EnumeratedNothing = iota
	SlTxresourcereqdiscR17SlCasttypediscR17EnumeratedBroadcast
	SlTxresourcereqdiscR17SlCasttypediscR17EnumeratedGroupcast
	SlTxresourcereqdiscR17SlCasttypediscR17EnumeratedUnicast
	SlTxresourcereqdiscR17SlCasttypediscR17EnumeratedSpare1
)
