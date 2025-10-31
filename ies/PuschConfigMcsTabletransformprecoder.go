package ies

import "rrc/utils"

// PUSCH-Config-mcs-TableTransformPrecoder ::= ENUMERATED
type PuschConfigMcsTabletransformprecoder struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigMcsTabletransformprecoderEnumeratedNothing = iota
	PuschConfigMcsTabletransformprecoderEnumeratedQam256
	PuschConfigMcsTabletransformprecoderEnumeratedQam64lowse
)
