package ies

// MeasObjectToAddMod-measObject ::= CHOICE
// Extensible
const (
	MeasobjecttoaddmodMeasobjectChoiceNothing = iota
	MeasobjecttoaddmodMeasobjectChoiceMeasobjecteutra
	MeasobjecttoaddmodMeasobjectChoiceMeasobjectutra
	MeasobjecttoaddmodMeasobjectChoiceMeasobjectgeran
	MeasobjecttoaddmodMeasobjectChoiceMeasobjectcdma2000
	MeasobjecttoaddmodMeasobjectChoiceMeasobjectwlanR13
	MeasobjecttoaddmodMeasobjectChoiceMeasobjectnrR15
)

type MeasobjecttoaddmodMeasobject struct {
	Choice             uint64
	Measobjecteutra    *Measobjecteutra
	Measobjectutra     *Measobjectutra
	Measobjectgeran    *Measobjectgeran
	Measobjectcdma2000 *Measobjectcdma2000
	MeasobjectwlanR13  *MeasobjectwlanR13
	MeasobjectnrR15    *MeasobjectnrR15
}
