package ies

import "rrc/utils"

// FeatureSetUplink-v1610-multiPUCCH-r16-sub-SlotConfig-ECP-r16 ::= ENUMERATED
type FeaturesetuplinkV1610MultipucchR16SubSlotconfigEcpR16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1610MultipucchR16SubSlotconfigEcpR16EnumeratedNothing = iota
	FeaturesetuplinkV1610MultipucchR16SubSlotconfigEcpR16EnumeratedSet1
	FeaturesetuplinkV1610MultipucchR16SubSlotconfigEcpR16EnumeratedSet2
)
