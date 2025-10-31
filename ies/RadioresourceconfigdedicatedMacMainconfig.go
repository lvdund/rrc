package ies

// RadioResourceConfigDedicated-mac-MainConfig ::= CHOICE
const (
	RadioresourceconfigdedicatedMacMainconfigChoiceNothing = iota
	RadioresourceconfigdedicatedMacMainconfigChoiceExplicitvalue
	RadioresourceconfigdedicatedMacMainconfigChoiceDefaultvalue
)

type RadioresourceconfigdedicatedMacMainconfig struct {
	Choice        uint64
	Explicitvalue *MacMainconfig
	Defaultvalue  *struct{}
}
