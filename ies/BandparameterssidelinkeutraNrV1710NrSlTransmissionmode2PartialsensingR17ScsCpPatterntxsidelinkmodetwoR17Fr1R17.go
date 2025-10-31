package ies

import "rrc/utils"

// BandParametersSidelinkEUTRA-NR-v1710-nr-sl-TransmissionMode2-PartialSensing-r17-scs-CP-PatternTxSidelinkModeTwo-r17-fr1-r17 ::= SEQUENCE
type BandparameterssidelinkeutraNrV1710NrSlTransmissionmode2PartialsensingR17ScsCpPatterntxsidelinkmodetwoR17Fr1R17 struct {
	Scs15khzR17 *utils.BITSTRING `lb:16,ub:16`
	Scs30khzR17 *utils.BITSTRING `lb:16,ub:16`
	Scs60khzR17 *utils.BITSTRING `lb:16,ub:16`
}
