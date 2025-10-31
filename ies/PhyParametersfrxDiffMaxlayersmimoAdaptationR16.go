package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-maxLayersMIMO-Adaptation-r16 ::= ENUMERATED
type PhyParametersfrxDiffMaxlayersmimoAdaptationR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffMaxlayersmimoAdaptationR16EnumeratedNothing = iota
	PhyParametersfrxDiffMaxlayersmimoAdaptationR16EnumeratedSupported
)
