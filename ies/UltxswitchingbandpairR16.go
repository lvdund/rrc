package ies

import "rrc/utils"

// ULTxSwitchingBandPair-r16 ::= SEQUENCE
type UltxswitchingbandpairR16 struct {
	Bandindexul1R16                    utils.INTEGER `lb:0,ub:maxSimultaneousBands`
	Bandindexul2R16                    utils.INTEGER `lb:0,ub:maxSimultaneousBands`
	UplinktxswitchingperiodR16         UltxswitchingbandpairR16UplinktxswitchingperiodR16
	UplinktxswitchingDlInterruptionR16 *utils.BITSTRING `lb:1,ub:maxSimultaneousBands`
}
