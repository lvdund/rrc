package ies

import "rrc/utils"

// ConfiguredGrantConfig-mcs-TableTransformPrecoder ::= ENUMERATED
type ConfiguredgrantconfigMcsTabletransformprecoder struct {
	Value utils.ENUMERATED
}

const (
	ConfiguredgrantconfigMcsTabletransformprecoderEnumeratedNothing = iota
	ConfiguredgrantconfigMcsTabletransformprecoderEnumeratedQam256
	ConfiguredgrantconfigMcsTabletransformprecoderEnumeratedQam64lowse
)
