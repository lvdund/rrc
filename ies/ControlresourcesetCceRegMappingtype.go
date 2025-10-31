package ies

// ControlResourceSet-cce-REG-MappingType ::= CHOICE
const (
	ControlresourcesetCceRegMappingtypeChoiceNothing = iota
	ControlresourcesetCceRegMappingtypeChoiceInterleaved
	ControlresourcesetCceRegMappingtypeChoiceNoninterleaved
)

type ControlresourcesetCceRegMappingtype struct {
	Choice         uint64
	Interleaved    *ControlresourcesetCceRegMappingtypeInterleaved
	Noninterleaved *struct{}
}
