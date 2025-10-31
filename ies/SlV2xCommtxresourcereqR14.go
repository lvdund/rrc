package ies

import "rrc/utils"

// SL-V2X-CommTxResourceReq-r14 ::= SEQUENCE
type SlV2xCommtxresourcereqR14 struct {
	CarrierfreqcommtxR14      *utils.INTEGER `lb:0,ub:maxFreqV2X1R14`
	V2xTypetxsyncR14          *SlTypetxsyncR14
	V2xDestinationinfolistR14 *SlDestinationinfolistR12
}
