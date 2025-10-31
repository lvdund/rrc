package ies

import "rrc/utils"

// CQI-ReportPeriodicProcExt-r11-cqi-FormatIndicatorPeriodic-r11-subbandCQI-r11 ::= SEQUENCE
type CqiReportperiodicprocextR11CqiFormatindicatorperiodicR11SubbandcqiR11 struct {
	K                    utils.INTEGER `lb:0,ub:4`
	PeriodicityfactorR11 CqiReportperiodicprocextR11CqiFormatindicatorperiodicR11SubbandcqiR11PeriodicityfactorR11
}
