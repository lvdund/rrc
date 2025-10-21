package ies

import "rrc/utils"

// RadioResourceConfigDedicated ::= SEQUENCE
// Extensible
type Radioresourceconfigdedicated struct {
	SrbToaddmodlist         *SrbToaddmodlist
	DrbToaddmodlist         *DrbToaddmodlist
	DrbToreleaselist        *DrbToreleaselist
	MacMainconfig           *interface{}
	SpsConfig               *SpsConfig
	Physicalconfigdedicated *Physicalconfigdedicated
}
