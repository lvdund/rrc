package ies

import "rrc/utils"

// BandSidelinkPC5-r16-rankTwoReception-r16 ::= ENUMERATED
type Bandsidelinkpc5R16RanktworeceptionR16 struct {
	Value utils.ENUMERATED
}

const (
	Bandsidelinkpc5R16RanktworeceptionR16EnumeratedNothing = iota
	Bandsidelinkpc5R16RanktworeceptionR16EnumeratedSupported
)
