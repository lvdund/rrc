package ies

import "rrc/utils"

// RLF-Report-r16-nr-RLF-Report-r16 ::= SEQUENCE
// Extensible
type RlfReportR16NrRlfReportR16 struct {
	MeasresultlastservcellR16    MeasresultrlfnrR16
	MeasresultneighcellsR16      *RlfReportR16NrRlfReportR16MeasresultneighcellsR16
	CRntiR16                     RntiValue
	PreviouspcellidR16           *RlfReportR16NrRlfReportR16PreviouspcellidR16
	FailedpcellidR16             RlfReportR16NrRlfReportR16FailedpcellidR16
	ReconnectcellidR16           *RlfReportR16NrRlfReportR16ReconnectcellidR16
	TimeuntilreconnectionR16     *TimeuntilreconnectionR16
	ReestablishmentcellidR16     *CgiInfoLoggingR16
	TimeconnfailureR16           *utils.INTEGER `lb:0,ub:1023`
	TimesincefailureR16          TimesincefailureR16
	ConnectionfailuretypeR16     RlfReportR16NrRlfReportR16ConnectionfailuretypeR16
	RlfCauseR16                  RlfReportR16NrRlfReportR16RlfCauseR16
	LocationinfoR16              *LocationinfoR16
	NosuitablecellfoundR16       *utils.BOOLEAN
	RaInformationcommonR16       *RaInformationcommonR16
	TimeconnsourcedapsFailureR17 *TimeconnsourcedapsFailureR17
	TimesincechoReconfigR17      *TimesincechoReconfigR17
	ChocellidR17                 *RlfReportR16NrRlfReportR16ChocellidR17
	ChocandidatecelllistR17      *ChocandidatecelllistR17
}
