package ies

import "rrc/utils"

// MeasAndMobParametersFRX-Diff-csi-RS-RLM ::= ENUMERATED
type MeasandmobparametersfrxDiffCsiRsRlm struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparametersfrxDiffCsiRsRlmEnumeratedNothing = iota
	MeasandmobparametersfrxDiffCsiRsRlmEnumeratedSupported
)
