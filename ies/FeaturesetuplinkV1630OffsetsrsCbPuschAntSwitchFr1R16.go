package ies

import "rrc/utils"

// FeatureSetUplink-v1630-offsetSRS-CB-PUSCH-Ant-Switch-fr1-r16 ::= ENUMERATED
type FeaturesetuplinkV1630OffsetsrsCbPuschAntSwitchFr1R16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1630OffsetsrsCbPuschAntSwitchFr1R16EnumeratedNothing = iota
	FeaturesetuplinkV1630OffsetsrsCbPuschAntSwitchFr1R16EnumeratedSupported
)
