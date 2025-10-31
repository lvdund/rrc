package ies

import "rrc/utils"

// TPC-Index ::= CHOICE
const (
	TpcIndexChoiceNothing = iota
	TpcIndexChoiceIndexofformat3
	TpcIndexChoiceIndexofformat3a
)

type TpcIndex struct {
	Choice          uint64
	Indexofformat3  *utils.INTEGER `lb:1,ub:15`
	Indexofformat3a *utils.INTEGER `lb:1,ub:31`
}
