package ies

import "rrc/utils"

// SL-TypeTxSync-r14 ::= ENUMERATED
type SlTypetxsyncR14 struct {
	Value utils.ENUMERATED
}

const (
	SlTypetxsyncR14EnumeratedNothing = iota
	SlTypetxsyncR14EnumeratedGnss
	SlTypetxsyncR14EnumeratedEnb
	SlTypetxsyncR14EnumeratedUe
)
