package ies

import "rrc/utils"

// UplinkTxSwitching-r16-uplinkTxSwitchingCarrier-r16 ::= ENUMERATED
type UplinktxswitchingR16UplinktxswitchingcarrierR16 struct {
	Value utils.ENUMERATED
}

const (
	UplinktxswitchingR16UplinktxswitchingcarrierR16EnumeratedNothing = iota
	UplinktxswitchingR16UplinktxswitchingcarrierR16EnumeratedCarrier1
	UplinktxswitchingR16UplinktxswitchingcarrierR16EnumeratedCarrier2
)
