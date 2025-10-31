package ies

import "rrc/utils"

// MeasConfigMN ::= SEQUENCE
// Extensible
type Measconfigmn struct {
	Measuredfrequenciesmn *[]NrFreqinfo `lb:1,ub:maxMeasFreqsMN`
	Measgapconfig         *utils.Setuprelease[Gapconfig]
	Gappurpose            *MeasconfigmnGappurpose
	Measgapconfigfr2      *utils.Setuprelease[Gapconfig] `ext`
	InterfreqnogapR16     *utils.BOOLEAN                 `ext`
}
