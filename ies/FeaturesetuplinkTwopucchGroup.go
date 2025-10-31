package ies

import "rrc/utils"

// FeatureSetUplink-twoPUCCH-Group ::= ENUMERATED
type FeaturesetuplinkTwopucchGroup struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkTwopucchGroupEnumeratedNothing = iota
	FeaturesetuplinkTwopucchGroupEnumeratedSupported
)
