package ies

import "rrc/utils"

// EthernetHeaderCompression-r16-ehc-Uplink-r16 ::= SEQUENCE
type EthernetheadercompressionR16EhcUplinkR16 struct {
	MaxcidEhcUlR16      utils.INTEGER `lb:0,ub:32767`
	DrbContinueehcUlR16 *bool
}
