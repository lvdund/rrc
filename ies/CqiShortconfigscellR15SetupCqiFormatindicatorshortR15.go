package ies

// CQI-ShortConfigSCell-r15-setup-cqi-FormatIndicatorShort-r15 ::= CHOICE
const (
	CqiShortconfigscellR15SetupCqiFormatindicatorshortR15ChoiceNothing = iota
	CqiShortconfigscellR15SetupCqiFormatindicatorshortR15ChoiceWidebandcqiShortR15
	CqiShortconfigscellR15SetupCqiFormatindicatorshortR15ChoiceSubbandcqiShortR15
)

type CqiShortconfigscellR15SetupCqiFormatindicatorshortR15 struct {
	Choice              uint64
	WidebandcqiShortR15 *CqiShortconfigscellR15SetupCqiFormatindicatorshortR15WidebandcqiShortR15
	SubbandcqiShortR15  *CqiShortconfigscellR15SetupCqiFormatindicatorshortR15SubbandcqiShortR15
}
