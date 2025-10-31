package ies

import "rrc/utils"

// RACH-ConfigCommonTwoStepRA-r16-msgA-SSB-PerRACH-OccasionAndCB-PreamblesPerSSB-r16 ::= CHOICE
const (
	RachConfigcommontwostepraR16MsgaSsbPerrachOccasionandcbPreamblesperssbR16ChoiceNothing = iota
	RachConfigcommontwostepraR16MsgaSsbPerrachOccasionandcbPreamblesperssbR16ChoiceOneeighth
	RachConfigcommontwostepraR16MsgaSsbPerrachOccasionandcbPreamblesperssbR16ChoiceOnefourth
	RachConfigcommontwostepraR16MsgaSsbPerrachOccasionandcbPreamblesperssbR16ChoiceOnehalf
	RachConfigcommontwostepraR16MsgaSsbPerrachOccasionandcbPreamblesperssbR16ChoiceOne
	RachConfigcommontwostepraR16MsgaSsbPerrachOccasionandcbPreamblesperssbR16ChoiceTwo
	RachConfigcommontwostepraR16MsgaSsbPerrachOccasionandcbPreamblesperssbR16ChoiceFour
	RachConfigcommontwostepraR16MsgaSsbPerrachOccasionandcbPreamblesperssbR16ChoiceEight
	RachConfigcommontwostepraR16MsgaSsbPerrachOccasionandcbPreamblesperssbR16ChoiceSixteen
)

type RachConfigcommontwostepraR16MsgaSsbPerrachOccasionandcbPreamblesperssbR16 struct {
	Choice    uint64
	Oneeighth *utils.ENUMERATED
	Onefourth *utils.ENUMERATED
	Onehalf   *utils.ENUMERATED
	One       *utils.ENUMERATED
	Two       *utils.ENUMERATED
	Four      *utils.INTEGER `lb:1,ub:16`
	Eight     *utils.INTEGER `lb:1,ub:8`
	Sixteen   *utils.INTEGER `lb:1,ub:4`
}
