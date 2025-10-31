package ies

import "rrc/utils"

// MeasAndMobParametersCommon-ssb-AndCSI-RS-RLM ::= ENUMERATED
type MeasandmobparameterscommonSsbAndcsiRsRlm struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonSsbAndcsiRsRlmEnumeratedNothing = iota
	MeasandmobparameterscommonSsbAndcsiRsRlmEnumeratedSupported
)
