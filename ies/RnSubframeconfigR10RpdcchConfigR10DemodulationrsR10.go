package ies

import "rrc/utils"

// RN-SubframeConfig-r10-rpdcch-Config-r10-demodulationRS-r10 ::= CHOICE
const (
	RnSubframeconfigR10RpdcchConfigR10DemodulationrsR10ChoiceNothing = iota
	RnSubframeconfigR10RpdcchConfigR10DemodulationrsR10ChoiceInterleavingR10
	RnSubframeconfigR10RpdcchConfigR10DemodulationrsR10ChoiceNointerleavingR10
)

type RnSubframeconfigR10RpdcchConfigR10DemodulationrsR10 struct {
	Choice            uint64
	InterleavingR10   *utils.ENUMERATED
	NointerleavingR10 *utils.ENUMERATED
}
