package ies

// RLF-Report-NB-r16 ::= SEQUENCE
type RlfReportNbR16 struct {
	FailedpcellidR16          Cellglobalideutra
	ReestablishmentcellidR16  *Cellglobalideutra
	LocationinfoR16           *LocationinfoR10
	MeasresultlastservcellR16 *RlfReportNbR16MeasresultlastservcellR16
	TimesincefailureR16       *TimesincefailureR11
}
