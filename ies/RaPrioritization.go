package ies

// RA-Prioritization ::= SEQUENCE
// Extensible
type RaPrioritization struct {
	Powerrampingstephighpriority RaPrioritizationPowerrampingstephighpriority
	Scalingfactorbi              *RaPrioritizationScalingfactorbi
}
