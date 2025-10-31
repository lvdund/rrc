package ies

// SRS-Resource-resourceType ::= CHOICE
// Extensible
const (
	SrsResourceResourcetypeChoiceNothing = iota
	SrsResourceResourcetypeChoiceAperiodic
	SrsResourceResourcetypeChoiceSemiPersistent
	SrsResourceResourcetypeChoicePeriodic
)

type SrsResourceResourcetype struct {
	Choice         uint64
	Aperiodic      *SrsResourceResourcetypeAperiodic
	SemiPersistent *SrsResourceResourcetypeSemiPersistent
	Periodic       *SrsResourceResourcetypePeriodic
}
