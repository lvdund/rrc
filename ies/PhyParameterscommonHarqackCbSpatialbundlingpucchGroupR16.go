package ies

import "rrc/utils"

// Phy-ParametersCommon-harqACK-CB-SpatialBundlingPUCCH-Group-r16 ::= ENUMERATED
type PhyParameterscommonHarqackCbSpatialbundlingpucchGroupR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonHarqackCbSpatialbundlingpucchGroupR16EnumeratedNothing = iota
	PhyParameterscommonHarqackCbSpatialbundlingpucchGroupR16EnumeratedSupported
)
