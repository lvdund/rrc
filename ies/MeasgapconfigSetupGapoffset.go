package ies

import "rrc/utils"

// MeasGapConfig-setup-gapOffset ::= CHOICE
// Extensible
const (
	MeasgapconfigSetupGapoffsetChoiceNothing = iota
	MeasgapconfigSetupGapoffsetChoiceGp0
	MeasgapconfigSetupGapoffsetChoiceGp1
	MeasgapconfigSetupGapoffsetChoiceGp2R14
	MeasgapconfigSetupGapoffsetChoiceGp3R14
	MeasgapconfigSetupGapoffsetChoiceGpNcsg0R14
	MeasgapconfigSetupGapoffsetChoiceGpNcsg1R14
	MeasgapconfigSetupGapoffsetChoiceGpNcsg2R14
	MeasgapconfigSetupGapoffsetChoiceGpNcsg3R14
	MeasgapconfigSetupGapoffsetChoiceGpNonuniform1R14
	MeasgapconfigSetupGapoffsetChoiceGpNonuniform2R14
	MeasgapconfigSetupGapoffsetChoiceGpNonuniform3R14
	MeasgapconfigSetupGapoffsetChoiceGpNonuniform4R14
	MeasgapconfigSetupGapoffsetChoiceGp4R15
	MeasgapconfigSetupGapoffsetChoiceGp5R15
	MeasgapconfigSetupGapoffsetChoiceGp6R15
	MeasgapconfigSetupGapoffsetChoiceGp7R15
	MeasgapconfigSetupGapoffsetChoiceGp8R15
	MeasgapconfigSetupGapoffsetChoiceGp9R15
	MeasgapconfigSetupGapoffsetChoiceGp10R15
	MeasgapconfigSetupGapoffsetChoiceGp11R15
)

type MeasgapconfigSetupGapoffset struct {
	Choice           uint64
	Gp0              *utils.INTEGER `lb:0,ub:39`
	Gp1              *utils.INTEGER `lb:0,ub:79`
	Gp2R14           *utils.INTEGER `lb:0,ub:39`
	Gp3R14           *utils.INTEGER `lb:0,ub:79`
	GpNcsg0R14       *utils.INTEGER `lb:0,ub:39`
	GpNcsg1R14       *utils.INTEGER `lb:0,ub:79`
	GpNcsg2R14       *utils.INTEGER `lb:0,ub:39`
	GpNcsg3R14       *utils.INTEGER `lb:0,ub:79`
	GpNonuniform1R14 *utils.INTEGER `lb:0,ub:1279`
	GpNonuniform2R14 *utils.INTEGER `lb:0,ub:2559`
	GpNonuniform3R14 *utils.INTEGER `lb:0,ub:5119`
	GpNonuniform4R14 *utils.INTEGER `lb:0,ub:10239`
	Gp4R15           *utils.INTEGER `lb:0,ub:19`
	Gp5R15           *utils.INTEGER `lb:0,ub:159`
	Gp6R15           *utils.INTEGER `lb:0,ub:19`
	Gp7R15           *utils.INTEGER `lb:0,ub:39`
	Gp8R15           *utils.INTEGER `lb:0,ub:79`
	Gp9R15           *utils.INTEGER `lb:0,ub:159`
	Gp10R15          *utils.INTEGER `lb:0,ub:19`
	Gp11R15          *utils.INTEGER `lb:0,ub:159`
}
