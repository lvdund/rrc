package ies

import "rrc/utils"

// SL-TypeTxSync-r14 ::= ENUMERATED
type SlTypetxsyncR14 struct {
	Value utils.ENUMERATED
}

const (
	SlTypetxsyncR14Gnss = 0
	SlTypetxsyncR14Enb  = 1
	SlTypetxsyncR14Ue   = 2
)
