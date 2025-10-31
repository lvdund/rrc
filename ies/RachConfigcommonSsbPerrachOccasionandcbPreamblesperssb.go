package ies

import "rrc/utils"

// RACH-ConfigCommon-ssb-perRACH-OccasionAndCB-PreamblesPerSSB ::= CHOICE
const (
	RachConfigcommonSsbPerrachOccasionandcbPreamblesperssbChoiceNothing = iota
	RachConfigcommonSsbPerrachOccasionandcbPreamblesperssbChoiceOneeighth
	RachConfigcommonSsbPerrachOccasionandcbPreamblesperssbChoiceOnefourth
	RachConfigcommonSsbPerrachOccasionandcbPreamblesperssbChoiceOnehalf
	RachConfigcommonSsbPerrachOccasionandcbPreamblesperssbChoiceOne
	RachConfigcommonSsbPerrachOccasionandcbPreamblesperssbChoiceTwo
	RachConfigcommonSsbPerrachOccasionandcbPreamblesperssbChoiceFour
	RachConfigcommonSsbPerrachOccasionandcbPreamblesperssbChoiceEight
	RachConfigcommonSsbPerrachOccasionandcbPreamblesperssbChoiceSixteen
)

type RachConfigcommonSsbPerrachOccasionandcbPreamblesperssb struct {
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
