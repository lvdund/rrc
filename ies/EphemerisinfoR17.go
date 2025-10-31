package ies

// EphemerisInfo-r17 ::= CHOICE
const (
	EphemerisinfoR17ChoiceNothing = iota
	EphemerisinfoR17ChoicePositionvelocityR17
	EphemerisinfoR17ChoiceOrbitalR17
)

type EphemerisinfoR17 struct {
	Choice              uint64
	PositionvelocityR17 *PositionvelocityR17
	OrbitalR17          *OrbitalR17
}
