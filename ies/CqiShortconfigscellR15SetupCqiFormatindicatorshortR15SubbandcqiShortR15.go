package ies

import "rrc/utils"

// CQI-ShortConfigSCell-r15-setup-cqi-FormatIndicatorShort-r15-subbandCQI-Short-r15 ::= SEQUENCE
type CqiShortconfigscellR15SetupCqiFormatindicatorshortR15SubbandcqiShortR15 struct {
	KR15                 utils.INTEGER `lb:0,ub:4`
	PeriodicityfactorR15 CqiShortconfigscellR15SetupCqiFormatindicatorshortR15SubbandcqiShortR15PeriodicityfactorR15
}
