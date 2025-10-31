package ies

// PhysicalConfigDedicatedSCell-r10-nonUL-Configuration-r10 ::= SEQUENCE
type PhysicalconfigdedicatedscellR10NonulConfigurationR10 struct {
	AntennainfoR10                  *AntennainfodedicatedR10
	CrosscarrierschedulingconfigR10 *CrosscarrierschedulingconfigR10
	CsiRsConfigR10                  *CsiRsConfigR10
	PdschConfigdedicatedR10         *PdschConfigdedicated
}
