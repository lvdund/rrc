package ies

import "rrc/utils"

// RSSI-ResourceConfigCLI-r16 ::= SEQUENCE
// Extensible
type RssiResourceconfigcliR16 struct {
	RssiResourceidR16           RssiResourceidR16
	RssiScsR16                  Subcarrierspacing
	StartprbR16                 utils.INTEGER `lb:0,ub:2169`
	NrofprbsR16                 utils.INTEGER `lb:0,ub:maxNrofPhysicalResourceBlocksPlus1`
	StartpositionR16            utils.INTEGER `lb:0,ub:13`
	NrofsymbolsR16              utils.INTEGER `lb:0,ub:14`
	RssiPeriodicityandoffsetR16 RssiPeriodicityandoffsetR16
	RefservcellindexR16         *Servcellindex
}
