package ies

import "rrc/utils"

// SL-V2X-PacketDuplicationConfig-r15 ::= SEQUENCE
// Extensible
type SlV2xPacketduplicationconfigR15 struct {
	ThreshslReliabilityR15      SlReliabilityR15
	AllowedcarrierfreqconfigR15 *SlPpprDestCarrierfreqlistR15
}
