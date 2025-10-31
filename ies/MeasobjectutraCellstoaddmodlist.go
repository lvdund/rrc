package ies

// MeasObjectUTRA-cellsToAddModList ::= CHOICE
const (
	MeasobjectutraCellstoaddmodlistChoiceNothing = iota
	MeasobjectutraCellstoaddmodlistChoiceCellstoaddmodlistutraFdd
	MeasobjectutraCellstoaddmodlistChoiceCellstoaddmodlistutraTdd
)

type MeasobjectutraCellstoaddmodlist struct {
	Choice                   uint64
	CellstoaddmodlistutraFdd *CellstoaddmodlistutraFdd
	CellstoaddmodlistutraTdd *CellstoaddmodlistutraTdd
}
