package ies

import "rrc/utils"

// CSI-RS-ConfigNonPrecoded-r13-codebookOverSamplingRateConfig-O1-r13 ::= ENUMERATED
type CsiRsConfignonprecodedR13CodebookoversamplingrateconfigO1R13 struct {
	Value utils.ENUMERATED
}

const (
	CsiRsConfignonprecodedR13CodebookoversamplingrateconfigO1R13EnumeratedNothing = iota
	CsiRsConfignonprecodedR13CodebookoversamplingrateconfigO1R13EnumeratedN4
	CsiRsConfignonprecodedR13CodebookoversamplingrateconfigO1R13EnumeratedN8
)
