package ies

import "rrc/utils"

// SL-CSI-RS-Config-r16-sl-CSI-RS-FreqAllocation-r16 ::= CHOICE
const (
	SlCsiRsConfigR16SlCsiRsFreqallocationR16ChoiceNothing = iota
	SlCsiRsConfigR16SlCsiRsFreqallocationR16ChoiceSlOneantennaportR16
	SlCsiRsConfigR16SlCsiRsFreqallocationR16ChoiceSlTwoantennaportR16
)

type SlCsiRsConfigR16SlCsiRsFreqallocationR16 struct {
	Choice              uint64
	SlOneantennaportR16 *utils.BITSTRING `lb:12,ub:12`
	SlTwoantennaportR16 *utils.BITSTRING `lb:6,ub:6`
}
