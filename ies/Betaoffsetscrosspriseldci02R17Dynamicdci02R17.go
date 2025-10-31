package ies

// BetaOffsetsCrossPriSelDCI-0-2-r17-dynamicDCI-0-2-r17 ::= CHOICE
const (
	Betaoffsetscrosspriseldci02R17Dynamicdci02R17ChoiceNothing = iota
	Betaoffsetscrosspriseldci02R17Dynamicdci02R17ChoiceOnebitR17
	Betaoffsetscrosspriseldci02R17Dynamicdci02R17ChoiceTwobitsR17
)

type Betaoffsetscrosspriseldci02R17Dynamicdci02R17 struct {
	Choice     uint64
	OnebitR17  *[]BetaoffsetscrosspriR17 `lb:2,ub:2`
	TwobitsR17 *[]BetaoffsetscrosspriR17 `lb:4,ub:4`
}
