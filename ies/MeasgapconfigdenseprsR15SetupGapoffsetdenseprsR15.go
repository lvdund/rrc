package ies

import "rrc/utils"

// MeasGapConfigDensePRS-r15-setup-gapOffsetDensePRS-r15 ::= CHOICE
// Extensible
const (
	MeasgapconfigdenseprsR15SetupGapoffsetdenseprsR15ChoiceNothing = iota
	MeasgapconfigdenseprsR15SetupGapoffsetdenseprsR15ChoiceRstd0R15
	MeasgapconfigdenseprsR15SetupGapoffsetdenseprsR15ChoiceRstd1R15
	MeasgapconfigdenseprsR15SetupGapoffsetdenseprsR15ChoiceRstd2R15
	MeasgapconfigdenseprsR15SetupGapoffsetdenseprsR15ChoiceRstd3R15
	MeasgapconfigdenseprsR15SetupGapoffsetdenseprsR15ChoiceRstd4R15
	MeasgapconfigdenseprsR15SetupGapoffsetdenseprsR15ChoiceRstd5R15
	MeasgapconfigdenseprsR15SetupGapoffsetdenseprsR15ChoiceRstd6R15
	MeasgapconfigdenseprsR15SetupGapoffsetdenseprsR15ChoiceRstd7R15
	MeasgapconfigdenseprsR15SetupGapoffsetdenseprsR15ChoiceRstd8R15
	MeasgapconfigdenseprsR15SetupGapoffsetdenseprsR15ChoiceRstd9R15
	MeasgapconfigdenseprsR15SetupGapoffsetdenseprsR15ChoiceRstd10R15
	MeasgapconfigdenseprsR15SetupGapoffsetdenseprsR15ChoiceRstd11R15
	MeasgapconfigdenseprsR15SetupGapoffsetdenseprsR15ChoiceRstd12R15
	MeasgapconfigdenseprsR15SetupGapoffsetdenseprsR15ChoiceRstd13R15
	MeasgapconfigdenseprsR15SetupGapoffsetdenseprsR15ChoiceRstd14R15
	MeasgapconfigdenseprsR15SetupGapoffsetdenseprsR15ChoiceRstd15R15
	MeasgapconfigdenseprsR15SetupGapoffsetdenseprsR15ChoiceRstd16R15
	MeasgapconfigdenseprsR15SetupGapoffsetdenseprsR15ChoiceRstd17R15
	MeasgapconfigdenseprsR15SetupGapoffsetdenseprsR15ChoiceRstd18R15
	MeasgapconfigdenseprsR15SetupGapoffsetdenseprsR15ChoiceRstd19R15
	MeasgapconfigdenseprsR15SetupGapoffsetdenseprsR15ChoiceRstd20R15
)

type MeasgapconfigdenseprsR15SetupGapoffsetdenseprsR15 struct {
	Choice    uint64
	Rstd0R15  *utils.INTEGER `lb:0,ub:79`
	Rstd1R15  *utils.INTEGER `lb:0,ub:159`
	Rstd2R15  *utils.INTEGER `lb:0,ub:319`
	Rstd3R15  *utils.INTEGER `lb:0,ub:639`
	Rstd4R15  *utils.INTEGER `lb:0,ub:1279`
	Rstd5R15  *utils.INTEGER `lb:0,ub:159`
	Rstd6R15  *utils.INTEGER `lb:0,ub:319`
	Rstd7R15  *utils.INTEGER `lb:0,ub:639`
	Rstd8R15  *utils.INTEGER `lb:0,ub:1279`
	Rstd9R15  *utils.INTEGER `lb:0,ub:319`
	Rstd10R15 *utils.INTEGER `lb:0,ub:639`
	Rstd11R15 *utils.INTEGER `lb:0,ub:1279`
	Rstd12R15 *utils.INTEGER `lb:0,ub:319`
	Rstd13R15 *utils.INTEGER `lb:0,ub:639`
	Rstd14R15 *utils.INTEGER `lb:0,ub:1279`
	Rstd15R15 *utils.INTEGER `lb:0,ub:639`
	Rstd16R15 *utils.INTEGER `lb:0,ub:1279`
	Rstd17R15 *utils.INTEGER `lb:0,ub:639`
	Rstd18R15 *utils.INTEGER `lb:0,ub:1279`
	Rstd19R15 *utils.INTEGER `lb:0,ub:639`
	Rstd20R15 *utils.INTEGER `lb:0,ub:1279`
}
