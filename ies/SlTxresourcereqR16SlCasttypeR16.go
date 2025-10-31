package ies

import "rrc/utils"

// SL-TxResourceReq-r16-sl-CastType-r16 ::= ENUMERATED
type SlTxresourcereqR16SlCasttypeR16 struct {
	Value utils.ENUMERATED
}

const (
	SlTxresourcereqR16SlCasttypeR16EnumeratedNothing = iota
	SlTxresourcereqR16SlCasttypeR16EnumeratedBroadcast
	SlTxresourcereqR16SlCasttypeR16EnumeratedGroupcast
	SlTxresourcereqR16SlCasttypeR16EnumeratedUnicast
	SlTxresourcereqR16SlCasttypeR16EnumeratedSpare1
)
