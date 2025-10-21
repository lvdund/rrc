package ies

import "rrc/utils"

// UL-ConfigCommonTDD-NB-r15 ::= SEQUENCE
// Extensible
type UlConfigcommontddNbR15 struct {
	TddUlDlAlignmentoffsetR15  TddUlDlAlignmentoffsetNbR15
	NprachParameterslisttddR15 *NprachParameterslisttddNbR15
}
