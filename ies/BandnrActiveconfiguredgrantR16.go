package ies

import "rrc/utils"

// BandNR-activeConfiguredGrant-r16 ::= SEQUENCE
type BandnrActiveconfiguredgrantR16 struct {
	MaxnumberconfigsperbwpR16 BandnrActiveconfiguredgrantR16MaxnumberconfigsperbwpR16
	MaxnumberconfigsallccR16  utils.INTEGER `lb:0,ub:32`
}
