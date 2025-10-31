package ies

// SuccessHO-Report-r17 ::= SEQUENCE
// Extensible
type SuccesshoReportR17 struct {
	SourcecellinfoR17         *SuccesshoReportR17SourcecellinfoR17
	TargetcellinfoR17         *SuccesshoReportR17TargetcellinfoR17
	MeasresultneighcellsR17   *SuccesshoReportR17MeasresultneighcellsR17
	LocationinfoR17           *LocationinfoR16
	TimesincechoReconfigR17   *TimesincechoReconfigR17
	ShrCauseR17               *ShrCauseR17
	RaInformationcommonR17    *RaInformationcommonR16
	UpinterruptiontimeathoR17 *UpinterruptiontimeathoR17
	CRntiR17                  *RntiValue
}
