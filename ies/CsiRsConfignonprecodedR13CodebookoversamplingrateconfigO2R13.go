package ies

import "rrc/utils"

// CSI-RS-ConfigNonPrecoded-r13-codebookOverSamplingRateConfig-O2-r13 ::= ENUMERATED
type CsiRsConfignonprecodedR13CodebookoversamplingrateconfigO2R13 struct {
	Value utils.ENUMERATED
}

const (
	CsiRsConfignonprecodedR13CodebookoversamplingrateconfigO2R13EnumeratedNothing = iota
	CsiRsConfignonprecodedR13CodebookoversamplingrateconfigO2R13EnumeratedN4
	CsiRsConfignonprecodedR13CodebookoversamplingrateconfigO2R13EnumeratedN8
)
