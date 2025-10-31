package ies

import "rrc/utils"

// ConnEstFailReport-r11 ::= SEQUENCE
// Extensible
type ConnestfailreportR11 struct {
	FailedcellidR11          Cellglobalideutra
	LocationinfoR11          *LocationinfoR10
	MeasresultfailedcellR11  *ConnestfailreportR11MeasresultfailedcellR11
	MeasresultneighcellsR11  *ConnestfailreportR11MeasresultneighcellsR11
	NumberofpreamblessentR11 NumberofpreamblessentR11
	ContentiondetectedR11    utils.BOOLEAN
	MaxtxpowerreachedR11     utils.BOOLEAN
	TimesincefailureR11      TimesincefailureR11
	MeasresultlisteutraV1130 *Measresultlist2eutraV9e0
}
