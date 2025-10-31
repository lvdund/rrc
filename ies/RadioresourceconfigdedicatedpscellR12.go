package ies

// RadioResourceConfigDedicatedPSCell-r12 ::= SEQUENCE
// Extensible
type RadioresourceconfigdedicatedpscellR12 struct {
	PhysicalconfigdedicatedpscellR12 *Physicalconfigdedicated
	SpsConfigR12                     *SpsConfig
	NaicsInfoR12                     *NaicsAssistanceinfoR12
}
