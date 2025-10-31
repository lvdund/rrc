package ies

import "rrc/utils"

// PRACH-ConfigSIB-v1310-mpdcch-startSF-CSS-RA-r13 ::= CHOICE
const (
	PrachConfigsibV1310MpdcchStartsfCssRaR13ChoiceNothing = iota
	PrachConfigsibV1310MpdcchStartsfCssRaR13ChoiceFddR13
	PrachConfigsibV1310MpdcchStartsfCssRaR13ChoiceTddR13
)

type PrachConfigsibV1310MpdcchStartsfCssRaR13 struct {
	Choice uint64
	FddR13 *utils.ENUMERATED
	TddR13 *utils.ENUMERATED
}
