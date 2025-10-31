package ies

import "rrc/utils"

// SC-MTCH-Info-NB-r14-npdcch-NPDSCH-MaxTBS-SC-MTCH-r14 ::= ENUMERATED
type ScMtchInfoNbR14NpdcchNpdschMaxtbsScMtchR14 struct {
	Value utils.ENUMERATED
}

const (
	ScMtchInfoNbR14NpdcchNpdschMaxtbsScMtchR14EnumeratedNothing = iota
	ScMtchInfoNbR14NpdcchNpdschMaxtbsScMtchR14EnumeratedN680
	ScMtchInfoNbR14NpdcchNpdschMaxtbsScMtchR14EnumeratedN2536
)
