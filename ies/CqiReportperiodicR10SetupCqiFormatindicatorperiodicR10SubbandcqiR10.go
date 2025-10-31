package ies

import "rrc/utils"

// CQI-ReportPeriodic-r10-setup-cqi-FormatIndicatorPeriodic-r10-subbandCQI-r10 ::= SEQUENCE
type CqiReportperiodicR10SetupCqiFormatindicatorperiodicR10SubbandcqiR10 struct {
	K                    utils.INTEGER `lb:0,ub:4`
	PeriodicityfactorR10 CqiReportperiodicR10SetupCqiFormatindicatorperiodicR10SubbandcqiR10PeriodicityfactorR10
}
