package ies

// PhysicalConfigDedicated-antennaInfo ::= CHOICE
const (
	PhysicalconfigdedicatedAntennainfoChoiceNothing = iota
	PhysicalconfigdedicatedAntennainfoChoiceExplicitvalue
	PhysicalconfigdedicatedAntennainfoChoiceDefaultvalue
)

type PhysicalconfigdedicatedAntennainfo struct {
	Choice        uint64
	Explicitvalue *Antennainfodedicated
	Defaultvalue  *struct{}
}
