package ies

import "rrc/utils"

// SL-TypeTxSync-r16 ::= ENUMERATED
type SlTypetxsyncR16 struct {
	Value utils.ENUMERATED
}

const (
	SlTypetxsyncR16EnumeratedNothing = iota
	SlTypetxsyncR16EnumeratedGnss
	SlTypetxsyncR16EnumeratedGnbenb
	SlTypetxsyncR16EnumeratedUe
)
