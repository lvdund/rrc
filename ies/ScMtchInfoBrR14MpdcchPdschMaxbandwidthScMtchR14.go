package ies

import "rrc/utils"

// SC-MTCH-Info-BR-r14-mpdcch-PDSCH-MaxBandwidth-SC-MTCH-r14 ::= ENUMERATED
type ScMtchInfoBrR14MpdcchPdschMaxbandwidthScMtchR14 struct {
	Value utils.ENUMERATED
}

const (
	ScMtchInfoBrR14MpdcchPdschMaxbandwidthScMtchR14EnumeratedNothing = iota
	ScMtchInfoBrR14MpdcchPdschMaxbandwidthScMtchR14EnumeratedBw1dot4
	ScMtchInfoBrR14MpdcchPdschMaxbandwidthScMtchR14EnumeratedBw5
)
