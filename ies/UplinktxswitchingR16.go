package ies

import "rrc/utils"

// UplinkTxSwitching-r16 ::= SEQUENCE
type UplinktxswitchingR16 struct {
	UplinktxswitchingperiodlocationR16 utils.BOOLEAN
	UplinktxswitchingcarrierR16        UplinktxswitchingR16UplinktxswitchingcarrierR16
}
