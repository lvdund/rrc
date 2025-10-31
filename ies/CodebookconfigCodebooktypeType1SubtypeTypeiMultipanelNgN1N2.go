package ies

import "rrc/utils"

// CodebookConfig-codebookType-type1-subType-typeI-MultiPanel-ng-n1-n2 ::= CHOICE
const (
	CodebookconfigCodebooktypeType1SubtypeTypeiMultipanelNgN1N2ChoiceNothing = iota
	CodebookconfigCodebooktypeType1SubtypeTypeiMultipanelNgN1N2ChoiceTwoTwoOneTypeiMultipanelRestriction
	CodebookconfigCodebooktypeType1SubtypeTypeiMultipanelNgN1N2ChoiceTwoFourOneTypeiMultipanelRestriction
	CodebookconfigCodebooktypeType1SubtypeTypeiMultipanelNgN1N2ChoiceFourTwoOneTypeiMultipanelRestriction
	CodebookconfigCodebooktypeType1SubtypeTypeiMultipanelNgN1N2ChoiceTwoTwoTwoTypeiMultipanelRestriction
	CodebookconfigCodebooktypeType1SubtypeTypeiMultipanelNgN1N2ChoiceTwoEightOneTypeiMultipanelRestriction
	CodebookconfigCodebooktypeType1SubtypeTypeiMultipanelNgN1N2ChoiceFourFourOneTypeiMultipanelRestriction
	CodebookconfigCodebooktypeType1SubtypeTypeiMultipanelNgN1N2ChoiceTwoFourTwoTypeiMultipanelRestriction
	CodebookconfigCodebooktypeType1SubtypeTypeiMultipanelNgN1N2ChoiceFourTwoTwoTypeiMultipanelRestriction
)

type CodebookconfigCodebooktypeType1SubtypeTypeiMultipanelNgN1N2 struct {
	Choice                                uint64
	TwoTwoOneTypeiMultipanelRestriction   *utils.BITSTRING `lb:8,ub:8`
	TwoFourOneTypeiMultipanelRestriction  *utils.BITSTRING `lb:16,ub:16`
	FourTwoOneTypeiMultipanelRestriction  *utils.BITSTRING `lb:8,ub:8`
	TwoTwoTwoTypeiMultipanelRestriction   *utils.BITSTRING `lb:64,ub:64`
	TwoEightOneTypeiMultipanelRestriction *utils.BITSTRING `lb:32,ub:32`
	FourFourOneTypeiMultipanelRestriction *utils.BITSTRING `lb:16,ub:16`
	TwoFourTwoTypeiMultipanelRestriction  *utils.BITSTRING `lb:128,ub:128`
	FourTwoTwoTypeiMultipanelRestriction  *utils.BITSTRING `lb:64,ub:64`
}
