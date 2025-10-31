package ies

import "rrc/utils"

// UplinkTxSwitchingBandParameters-v1700 ::= SEQUENCE
type UplinktxswitchingbandparametersV1700 struct {
	BandindexR17                                utils.INTEGER `lb:0,ub:maxSimultaneousBands`
	Uplinktxswitching2t2tPuschTranscoherenceR17 *UplinktxswitchingbandparametersV1700Uplinktxswitching2t2tPuschTranscoherenceR17
}
