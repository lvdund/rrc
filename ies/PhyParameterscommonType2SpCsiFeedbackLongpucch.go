package ies

import "rrc/utils"

// Phy-ParametersCommon-type2-SP-CSI-Feedback-LongPUCCH ::= ENUMERATED
type PhyParameterscommonType2SpCsiFeedbackLongpucch struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonType2SpCsiFeedbackLongpucchEnumeratedNothing = iota
	PhyParameterscommonType2SpCsiFeedbackLongpucchEnumeratedSupported
)
