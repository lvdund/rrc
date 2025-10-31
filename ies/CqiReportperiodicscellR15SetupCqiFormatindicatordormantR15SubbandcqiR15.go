package ies

import "rrc/utils"

// CQI-ReportPeriodicSCell-r15-setup-cqi-FormatIndicatorDormant-r15-subbandCQI-r15 ::= SEQUENCE
type CqiReportperiodicscellR15SetupCqiFormatindicatordormantR15SubbandcqiR15 struct {
	KR15                 utils.INTEGER `lb:0,ub:4`
	PeriodicityfactorR15 CqiReportperiodicscellR15SetupCqiFormatindicatordormantR15SubbandcqiR15PeriodicityfactorR15
}
