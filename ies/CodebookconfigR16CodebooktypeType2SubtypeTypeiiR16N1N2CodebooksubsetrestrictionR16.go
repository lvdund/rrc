package ies

import "rrc/utils"

// CodebookConfig-r16-codebookType-type2-subType-typeII-r16-n1-n2-codebookSubsetRestriction-r16 ::= CHOICE
const (
	CodebookconfigR16CodebooktypeType2SubtypeTypeiiR16N1N2CodebooksubsetrestrictionR16ChoiceNothing = iota
	CodebookconfigR16CodebooktypeType2SubtypeTypeiiR16N1N2CodebooksubsetrestrictionR16ChoiceTwoOne
	CodebookconfigR16CodebooktypeType2SubtypeTypeiiR16N1N2CodebooksubsetrestrictionR16ChoiceTwoTwo
	CodebookconfigR16CodebooktypeType2SubtypeTypeiiR16N1N2CodebooksubsetrestrictionR16ChoiceFourOne
	CodebookconfigR16CodebooktypeType2SubtypeTypeiiR16N1N2CodebooksubsetrestrictionR16ChoiceThreeTwo
	CodebookconfigR16CodebooktypeType2SubtypeTypeiiR16N1N2CodebooksubsetrestrictionR16ChoiceSixOne
	CodebookconfigR16CodebooktypeType2SubtypeTypeiiR16N1N2CodebooksubsetrestrictionR16ChoiceFourTwo
	CodebookconfigR16CodebooktypeType2SubtypeTypeiiR16N1N2CodebooksubsetrestrictionR16ChoiceEightOne
	CodebookconfigR16CodebooktypeType2SubtypeTypeiiR16N1N2CodebooksubsetrestrictionR16ChoiceFourThree
	CodebookconfigR16CodebooktypeType2SubtypeTypeiiR16N1N2CodebooksubsetrestrictionR16ChoiceSixTwo
	CodebookconfigR16CodebooktypeType2SubtypeTypeiiR16N1N2CodebooksubsetrestrictionR16ChoiceTwelveOne
	CodebookconfigR16CodebooktypeType2SubtypeTypeiiR16N1N2CodebooksubsetrestrictionR16ChoiceFourFour
	CodebookconfigR16CodebooktypeType2SubtypeTypeiiR16N1N2CodebooksubsetrestrictionR16ChoiceEightTwo
	CodebookconfigR16CodebooktypeType2SubtypeTypeiiR16N1N2CodebooksubsetrestrictionR16ChoiceSixteenOne
)

type CodebookconfigR16CodebooktypeType2SubtypeTypeiiR16N1N2CodebooksubsetrestrictionR16 struct {
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
