package ies

import "rrc/utils"

// FeatureSetUplink-v1720-pucch-Repetition-F0-1-2-3-4-RRC-Config-r17 ::= ENUMERATED
type FeaturesetuplinkV1720PucchRepetitionF01234RrcConfigR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1720PucchRepetitionF01234RrcConfigR17EnumeratedNothing = iota
	FeaturesetuplinkV1720PucchRepetitionF01234RrcConfigR17EnumeratedSupported
)
