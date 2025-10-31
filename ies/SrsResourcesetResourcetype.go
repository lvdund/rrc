package ies

// SRS-ResourceSet-resourceType ::= CHOICE
// Extensible
const (
	SrsResourcesetResourcetypeChoiceNothing = iota
	SrsResourcesetResourcetypeChoiceAperiodic
	SrsResourcesetResourcetypeChoiceSemiPersistent
	SrsResourcesetResourcetypeChoicePeriodic
)

type SrsResourcesetResourcetype struct {
	Choice         uint64
	Aperiodic      *SrsResourcesetResourcetypeAperiodic
	SemiPersistent *SrsResourcesetResourcetypeSemiPersistent
	Periodic       *SrsResourcesetResourcetypePeriodic
}
