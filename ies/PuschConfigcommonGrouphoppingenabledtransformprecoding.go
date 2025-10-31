package ies

import "rrc/utils"

// PUSCH-ConfigCommon-groupHoppingEnabledTransformPrecoding ::= ENUMERATED
type PuschConfigcommonGrouphoppingenabledtransformprecoding struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigcommonGrouphoppingenabledtransformprecodingEnumeratedNothing = iota
	PuschConfigcommonGrouphoppingenabledtransformprecodingEnumeratedEnabled
)
