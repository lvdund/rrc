package ies

import "rrc/utils"

// SL-ZoneConfigMCR-r16 ::= SEQUENCE
// Extensible
type SlZoneconfigmcrR16 struct {
	SlZoneconfigmcrIndexR16 utils.INTEGER `lb:0,ub:15`
	SlTransrangeR16         *SlZoneconfigmcrR16SlTransrangeR16
	SlZoneconfigR16         *SlZoneconfigR16
}
