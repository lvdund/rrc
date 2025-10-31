package ies

import "rrc/utils"

// Phy-ParametersCommon-dci-25-AI-RNTI-Support-IAB-r16 ::= ENUMERATED
type PhyParameterscommonDci25AiRntiSupportIabR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonDci25AiRntiSupportIabR16EnumeratedNothing = iota
	PhyParameterscommonDci25AiRntiSupportIabR16EnumeratedSupported
)
