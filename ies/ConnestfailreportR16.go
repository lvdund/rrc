package ies

import "rrc/utils"

// ConnEstFailReport-r16 ::= SEQUENCE
// Extensible
type ConnestfailreportR16 struct {
	MeasresultfailedcellR16 MeasresultfailedcellR16
	LocationinfoR16         *LocationinfoR16
	MeasresultneighcellsR16 *ConnestfailreportR16MeasresultneighcellsR16
	NumberofconnfailR16     utils.INTEGER `lb:0,ub:8`
	PerrainfolistR16        PerrainfolistR16
	TimesincefailureR16     TimesincefailureR16
}
