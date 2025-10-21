package ies

import "rrc/utils"

// SL-CommTxResourceReq-r12 ::= SEQUENCE
type SlCommtxresourcereqR12 struct {
	CarrierfreqR12         *ArfcnValueeutraR9
	DestinationinfolistR12 SlDestinationinfolistR12
}
