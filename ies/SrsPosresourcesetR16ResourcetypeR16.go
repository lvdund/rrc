package ies

// SRS-PosResourceSet-r16-resourceType-r16 ::= CHOICE
// Extensible
const (
	SrsPosresourcesetR16ResourcetypeR16ChoiceNothing = iota
	SrsPosresourcesetR16ResourcetypeR16ChoiceAperiodicR16
	SrsPosresourcesetR16ResourcetypeR16ChoiceSemiPersistentR16
	SrsPosresourcesetR16ResourcetypeR16ChoicePeriodicR16
)

type SrsPosresourcesetR16ResourcetypeR16 struct {
	Choice            uint64
	AperiodicR16      *SrsPosresourcesetR16ResourcetypeR16AperiodicR16
	SemiPersistentR16 *SrsPosresourcesetR16ResourcetypeR16SemiPersistentR16
	PeriodicR16       *SrsPosresourcesetR16ResourcetypeR16PeriodicR16
}
