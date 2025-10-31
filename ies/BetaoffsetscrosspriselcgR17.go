package ies

// BetaOffsetsCrossPriSelCG-r17 ::= CHOICE
const (
	BetaoffsetscrosspriselcgR17ChoiceNothing = iota
	BetaoffsetscrosspriselcgR17ChoiceDynamicR17
	BetaoffsetscrosspriselcgR17ChoiceSemistaticR17
)

type BetaoffsetscrosspriselcgR17 struct {
	Choice        uint64
	DynamicR17    *[]BetaoffsetscrosspriR17 `lb:1,ub:4`
	SemistaticR17 *BetaoffsetscrosspriR17
}
