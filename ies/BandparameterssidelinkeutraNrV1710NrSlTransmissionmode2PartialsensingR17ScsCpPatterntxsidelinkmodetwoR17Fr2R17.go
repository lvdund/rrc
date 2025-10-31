package ies

import "rrc/utils"

// BandParametersSidelinkEUTRA-NR-v1710-nr-sl-TransmissionMode2-PartialSensing-r17-scs-CP-PatternTxSidelinkModeTwo-r17-fr2-r17 ::= SEQUENCE
type BandparameterssidelinkeutraNrV1710NrSlTransmissionmode2PartialsensingR17ScsCpPatterntxsidelinkmodetwoR17Fr2R17 struct {
	Scs60khzR17  *utils.BITSTRING `lb:16,ub:16`
	Scs120khzR17 *utils.BITSTRING `lb:16,ub:16`
}
