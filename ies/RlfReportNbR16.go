package ies

import "rrc/utils"

// RLF-Report-NB-r16 ::= SEQUENCE
type RlfReportNbR16 struct {
	FailedpcellidR16          Cellglobalideutra
	ReestablishmentcellidR16  *Cellglobalideutra
	LocationinfoR16           *LocationinfoR10
	MeasresultlastservcellR16 *interface{}
	TimesincefailureR16       *TimesincefailureR11
}
