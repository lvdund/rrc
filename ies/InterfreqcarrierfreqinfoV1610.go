package ies

// InterFreqCarrierFreqInfo-v1610 ::= SEQUENCE
type InterfreqcarrierfreqinfoV1610 struct {
	InterfreqneighcelllistV1610 *InterfreqneighcelllistV1610
	Smtc2LpR16                  *SsbMtc2LpR16
	InterfreqallowedcelllistR16 *InterfreqallowedcelllistR16
	SsbPositionqclCommonR16     *SsbPositionqclRelationR16
	InterfreqcagCelllistR16     *[]InterfreqcagCelllistperplmnR16 `lb:1,ub:maxPLMN`
}
