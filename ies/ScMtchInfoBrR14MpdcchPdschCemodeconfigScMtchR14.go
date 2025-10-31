package ies

import "rrc/utils"

// SC-MTCH-Info-BR-r14-mpdcch-PDSCH-CEmodeConfig-SC-MTCH-r14 ::= ENUMERATED
type ScMtchInfoBrR14MpdcchPdschCemodeconfigScMtchR14 struct {
	Value utils.ENUMERATED
}

const (
	ScMtchInfoBrR14MpdcchPdschCemodeconfigScMtchR14EnumeratedNothing = iota
	ScMtchInfoBrR14MpdcchPdschCemodeconfigScMtchR14EnumeratedCe_Modea
	ScMtchInfoBrR14MpdcchPdschCemodeconfigScMtchR14EnumeratedCe_Modeb
)
