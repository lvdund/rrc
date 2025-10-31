package ies

// BetaOffsetsCrossPriSel-r17 ::= CHOICE
const (
	BetaoffsetscrosspriselR17ChoiceNothing = iota
	BetaoffsetscrosspriselR17ChoiceDynamicR17
	BetaoffsetscrosspriselR17ChoiceSemistaticR17
)

type BetaoffsetscrosspriselR17 struct {
	Choice        uint64
	DynamicR17    *[]BetaoffsetscrosspriR17 `lb:4,ub:4`
	SemistaticR17 *BetaoffsetscrosspriR17
}
