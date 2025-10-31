package ies

// SRS-PosResource-r16-resourceType-r16 ::= CHOICE
// Extensible
const (
	SrsPosresourceR16ResourcetypeR16ChoiceNothing = iota
	SrsPosresourceR16ResourcetypeR16ChoiceAperiodicR16
	SrsPosresourceR16ResourcetypeR16ChoiceSemiPersistentR16
	SrsPosresourceR16ResourcetypeR16ChoicePeriodicR16
)

type SrsPosresourceR16ResourcetypeR16 struct {
	Choice            uint64
	AperiodicR16      *SrsPosresourceR16ResourcetypeR16AperiodicR16
	SemiPersistentR16 *SrsPosresourceR16ResourcetypeR16SemiPersistentR16
	PeriodicR16       *SrsPosresourceR16ResourcetypeR16PeriodicR16
}
