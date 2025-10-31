package ies

// RadioResourceConfigDedicated ::= SEQUENCE
// Extensible
type Radioresourceconfigdedicated struct {
	SrbToaddmodlist         *SrbToaddmodlist
	DrbToaddmodlist         *DrbToaddmodlist
	DrbToreleaselist        *DrbToreleaselist
	MacMainconfig           *RadioresourceconfigdedicatedMacMainconfig
	SpsConfig               *SpsConfig
	Physicalconfigdedicated *Physicalconfigdedicated
}
