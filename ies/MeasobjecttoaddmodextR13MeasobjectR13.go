package ies

// MeasObjectToAddModExt-r13-measObject-r13 ::= CHOICE
// Extensible
const (
	MeasobjecttoaddmodextR13MeasobjectR13ChoiceNothing = iota
	MeasobjecttoaddmodextR13MeasobjectR13ChoiceMeasobjecteutraR13
	MeasobjecttoaddmodextR13MeasobjectR13ChoiceMeasobjectutraR13
	MeasobjecttoaddmodextR13MeasobjectR13ChoiceMeasobjectgeranR13
	MeasobjecttoaddmodextR13MeasobjectR13ChoiceMeasobjectcdma2000R13
	MeasobjecttoaddmodextR13MeasobjectR13ChoiceMeasobjectwlanV1320
	MeasobjecttoaddmodextR13MeasobjectR13ChoiceMeasobjectnrR15
)

type MeasobjecttoaddmodextR13MeasobjectR13 struct {
	Choice                uint64
	MeasobjecteutraR13    *Measobjecteutra
	MeasobjectutraR13     *Measobjectutra
	MeasobjectgeranR13    *Measobjectgeran
	Measobjectcdma2000R13 *Measobjectcdma2000
	MeasobjectwlanV1320   *MeasobjectwlanR13
	MeasobjectnrR15       *MeasobjectnrR15
}
