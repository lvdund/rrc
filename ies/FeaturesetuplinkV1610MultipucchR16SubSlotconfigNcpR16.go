package ies

import "rrc/utils"

// FeatureSetUplink-v1610-multiPUCCH-r16-sub-SlotConfig-NCP-r16 ::= ENUMERATED
type FeaturesetuplinkV1610MultipucchR16SubSlotconfigNcpR16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1610MultipucchR16SubSlotconfigNcpR16EnumeratedNothing = iota
	FeaturesetuplinkV1610MultipucchR16SubSlotconfigNcpR16EnumeratedSet1
	FeaturesetuplinkV1610MultipucchR16SubSlotconfigNcpR16EnumeratedSet2
)
