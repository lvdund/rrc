package ies

import "rrc/utils"

// PRACH-Config-v1310-mpdcch-startSF-CSS-RA-r13 ::= CHOICE
const (
	PrachConfigV1310MpdcchStartsfCssRaR13ChoiceNothing = iota
	PrachConfigV1310MpdcchStartsfCssRaR13ChoiceFddR13
	PrachConfigV1310MpdcchStartsfCssRaR13ChoiceTddR13
)

type PrachConfigV1310MpdcchStartsfCssRaR13 struct {
	Choice uint64
	FddR13 *utils.ENUMERATED
	TddR13 *utils.ENUMERATED
}
