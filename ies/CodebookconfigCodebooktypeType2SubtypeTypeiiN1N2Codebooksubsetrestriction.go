package ies

import "rrc/utils"

// CodebookConfig-codebookType-type2-subType-typeII-n1-n2-codebookSubsetRestriction ::= CHOICE
const (
	CodebookconfigCodebooktypeType2SubtypeTypeiiN1N2CodebooksubsetrestrictionChoiceNothing = iota
	CodebookconfigCodebooktypeType2SubtypeTypeiiN1N2CodebooksubsetrestrictionChoiceTwoOne
	CodebookconfigCodebooktypeType2SubtypeTypeiiN1N2CodebooksubsetrestrictionChoiceTwoTwo
	CodebookconfigCodebooktypeType2SubtypeTypeiiN1N2CodebooksubsetrestrictionChoiceFourOne
	CodebookconfigCodebooktypeType2SubtypeTypeiiN1N2CodebooksubsetrestrictionChoiceThreeTwo
	CodebookconfigCodebooktypeType2SubtypeTypeiiN1N2CodebooksubsetrestrictionChoiceSixOne
	CodebookconfigCodebooktypeType2SubtypeTypeiiN1N2CodebooksubsetrestrictionChoiceFourTwo
	CodebookconfigCodebooktypeType2SubtypeTypeiiN1N2CodebooksubsetrestrictionChoiceEightOne
	CodebookconfigCodebooktypeType2SubtypeTypeiiN1N2CodebooksubsetrestrictionChoiceFourThree
	CodebookconfigCodebooktypeType2SubtypeTypeiiN1N2CodebooksubsetrestrictionChoiceSixTwo
	CodebookconfigCodebooktypeType2SubtypeTypeiiN1N2CodebooksubsetrestrictionChoiceTwelveOne
	CodebookconfigCodebooktypeType2SubtypeTypeiiN1N2CodebooksubsetrestrictionChoiceFourFour
	CodebookconfigCodebooktypeType2SubtypeTypeiiN1N2CodebooksubsetrestrictionChoiceEightTwo
	CodebookconfigCodebooktypeType2SubtypeTypeiiN1N2CodebooksubsetrestrictionChoiceSixteenOne
)

type CodebookconfigCodebooktypeType2SubtypeTypeiiN1N2Codebooksubsetrestriction struct {
	Choice     uint64
	TwoOne     *utils.BITSTRING `lb:16,ub:16`
	TwoTwo     *utils.BITSTRING `lb:43,ub:43`
	FourOne    *utils.BITSTRING `lb:32,ub:32`
	ThreeTwo   *utils.BITSTRING `lb:59,ub:59`
	SixOne     *utils.BITSTRING `lb:48,ub:48`
	FourTwo    *utils.BITSTRING `lb:75,ub:75`
	EightOne   *utils.BITSTRING `lb:64,ub:64`
	FourThree  *utils.BITSTRING `lb:107,ub:107`
	SixTwo     *utils.BITSTRING `lb:107,ub:107`
	TwelveOne  *utils.BITSTRING `lb:96,ub:96`
	FourFour   *utils.BITSTRING `lb:139,ub:139`
	EightTwo   *utils.BITSTRING `lb:139,ub:139`
	SixteenOne *utils.BITSTRING `lb:128,ub:128`
}
