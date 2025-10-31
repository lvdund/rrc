package ies

import "rrc/utils"

// MeasResults ::= SEQUENCE
// Extensible
type Measresults struct {
	Measid                         Measid
	Measresultservingmolist        Measresultservmolist
	Measresultneighcells           *MeasresultsMeasresultneighcells
	MeasresultservfreqlisteutraScg *MeasresultservfreqlisteutraScg `ext`
	MeasresultservfreqlistnrScg    *MeasresultservfreqlistnrScg    `ext`
	MeasresultsftdEutra            *MeasresultsftdEutra            `ext`
	MeasresultsftdNr               *MeasresultcellsftdNr           `ext`
	MeasresultcelllistsftdNr       *MeasresultcelllistsftdNr       `ext`
	MeasresultforrssiR16           *MeasresultforrssiR16           `ext`
	LocationinfoR16                *LocationinfoR16                `ext`
	UlPdcpDelayvalueresultlistR16  *UlPdcpDelayvalueresultlistR16  `ext`
	MeasresultsslR16               *MeasresultsslR16               `ext`
	MeasresultcliR16               *MeasresultcliR16               `ext`
	MeasresultrxtxtimediffR17      *MeasresultrxtxtimediffR17      `ext`
	SlMeasresultservingrelayR17    *utils.OCTETSTRING              `ext`
	UlPdcpExcessdelayresultlistR17 *UlPdcpExcessdelayresultlistR17 `ext`
	CoarselocationinfoR17          *utils.OCTETSTRING              `ext`
}
