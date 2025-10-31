package ies

import "rrc/utils"

// LogMeasInfo-r16 ::= SEQUENCE
// Extensible
type LogmeasinfoR16 struct {
	LocationinfoR16             *LocationinfoR16
	RelativetimestampR16        utils.INTEGER `lb:0,ub:7200`
	ServcellidentityR16         *CgiInfoLoggingR16
	MeasresultservingcellR16    *MeasresultservingcellR16
	MeasresultneighcellsR16     *LogmeasinfoR16MeasresultneighcellsR16
	AnycellselectiondetectedR16 *utils.BOOLEAN
	IndevicecoexdetectedR17     *utils.BOOLEAN `ext`
}
