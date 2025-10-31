package ies

import "rrc/utils"

// SPS-ConfigUL-STTI-r15-setup-rv-SPS-STTI-UL-Repetitions-r15 ::= ENUMERATED
type SpsConfigulSttiR15SetupRvSpsSttiUlRepetitionsR15 struct {
	Value utils.ENUMERATED
}

const (
	SpsConfigulSttiR15SetupRvSpsSttiUlRepetitionsR15EnumeratedNothing = iota
	SpsConfigulSttiR15SetupRvSpsSttiUlRepetitionsR15EnumeratedUlrvseq1
	SpsConfigulSttiR15SetupRvSpsSttiUlRepetitionsR15EnumeratedUlrvseq2
	SpsConfigulSttiR15SetupRvSpsSttiUlRepetitionsR15EnumeratedUlrvseq3
)
