package ies

import "rrc/utils"

// MeasAndMobParametersCommon-ssb-RLM ::= ENUMERATED
type MeasandmobparameterscommonSsbRlm struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonSsbRlmEnumeratedNothing = iota
	MeasandmobparameterscommonSsbRlmEnumeratedSupported
)
