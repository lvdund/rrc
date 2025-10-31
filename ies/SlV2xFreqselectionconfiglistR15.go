package ies

// SL-V2X-FreqSelectionConfigList-r15 ::= SEQUENCE OF SL-V2X-FreqSelectionConfig-r15
// SIZE (1..8)
type SlV2xFreqselectionconfiglistR15 struct {
	Value []SlV2xFreqselectionconfigR15 `lb:1,ub:8`
}
