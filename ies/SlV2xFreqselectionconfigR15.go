package ies

// SL-V2X-FreqSelectionConfig-r15 ::= SEQUENCE
type SlV2xFreqselectionconfigR15 struct {
	PrioritylistR15             SlPrioritylistR13
	ThreshcbrFreqreselectionR15 *SlCbrR14
	ThreshcbrFreqkeepingR15     *SlCbrR14
}
