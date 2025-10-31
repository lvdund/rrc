package ies

import "rrc/utils"

// Phy-ParametersCommon-precoderGranularityCORESET ::= ENUMERATED
type PhyParameterscommonPrecodergranularitycoreset struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonPrecodergranularitycoresetEnumeratedNothing = iota
	PhyParameterscommonPrecodergranularitycoresetEnumeratedSupported
)
