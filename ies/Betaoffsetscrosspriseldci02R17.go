package ies

// BetaOffsetsCrossPriSelDCI-0-2-r17 ::= CHOICE
const (
	Betaoffsetscrosspriseldci02R17ChoiceNothing = iota
	Betaoffsetscrosspriseldci02R17ChoiceDynamicdci02R17
	Betaoffsetscrosspriseldci02R17ChoiceSemistaticdci02R17
)

type Betaoffsetscrosspriseldci02R17 struct {
	Choice             uint64
	Dynamicdci02R17    *Betaoffsetscrosspriseldci02R17Dynamicdci02R17
	Semistaticdci02R17 *BetaoffsetscrosspriR17
}
