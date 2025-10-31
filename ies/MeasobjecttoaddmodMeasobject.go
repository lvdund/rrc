package ies

// MeasObjectToAddMod-measObject ::= CHOICE
// Extensible
const (
	MeasobjecttoaddmodMeasobjectChoiceNothing = iota
	MeasobjecttoaddmodMeasobjectChoiceMeasobjectnr
	MeasobjecttoaddmodMeasobjectChoiceMeasobjecteutra
	MeasobjecttoaddmodMeasobjectChoiceMeasobjectutraFddR16
	MeasobjecttoaddmodMeasobjectChoiceMeasobjectnrSlR16
	MeasobjecttoaddmodMeasobjectChoiceMeasobjectcliR16
	MeasobjecttoaddmodMeasobjectChoiceMeasobjectrxtxdiffR17
	MeasobjecttoaddmodMeasobjectChoiceMeasobjectrelayR17
)

type MeasobjecttoaddmodMeasobject struct {
	Choice                uint64
	Measobjectnr          *Measobjectnr
	Measobjecteutra       *Measobjecteutra
	MeasobjectutraFddR16  *MeasobjectutraFddR16
	MeasobjectnrSlR16     *MeasobjectnrSlR16
	MeasobjectcliR16      *MeasobjectcliR16
	MeasobjectrxtxdiffR17 *MeasobjectrxtxdiffR17
	MeasobjectrelayR17    *SlMeasobjectR16
}
