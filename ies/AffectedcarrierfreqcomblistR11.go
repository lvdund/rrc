package ies

import "rrc/utils"

// AffectedCarrierFreqCombList-r11 ::= SEQUENCE OF AffectedCarrierFreqComb-r11
// SIZE (1..maxCombIDC-r11)
type AffectedcarrierfreqcomblistR11 struct {
	Value utils.Sequence[AffectedcarrierfreqcombR11]
}
