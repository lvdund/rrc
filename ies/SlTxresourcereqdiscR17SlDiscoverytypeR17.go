package ies

import "rrc/utils"

// SL-TxResourceReqDisc-r17-sl-DiscoveryType-r17 ::= ENUMERATED
type SlTxresourcereqdiscR17SlDiscoverytypeR17 struct {
	Value utils.ENUMERATED
}

const (
	SlTxresourcereqdiscR17SlDiscoverytypeR17EnumeratedNothing = iota
	SlTxresourcereqdiscR17SlDiscoverytypeR17EnumeratedRelay
	SlTxresourcereqdiscR17SlDiscoverytypeR17EnumeratedNon_Relay
)
