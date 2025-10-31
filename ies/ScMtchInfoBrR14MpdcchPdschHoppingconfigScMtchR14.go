package ies

import "rrc/utils"

// SC-MTCH-Info-BR-r14-mpdcch-PDSCH-HoppingConfig-SC-MTCH-r14 ::= ENUMERATED
type ScMtchInfoBrR14MpdcchPdschHoppingconfigScMtchR14 struct {
	Value utils.ENUMERATED
}

const (
	ScMtchInfoBrR14MpdcchPdschHoppingconfigScMtchR14EnumeratedNothing = iota
	ScMtchInfoBrR14MpdcchPdschHoppingconfigScMtchR14EnumeratedOn
	ScMtchInfoBrR14MpdcchPdschHoppingconfigScMtchR14EnumeratedOff
)
