package ies

// Phy-ParametersFR1 ::= SEQUENCE
// Extensible
type PhyParametersfr1 struct {
	PdcchMonitoringsingleoccasion         *PhyParametersfr1PdcchMonitoringsingleoccasion
	Scs60khz                              *PhyParametersfr1Scs60khz
	Pdsch256qamFr1                        *PhyParametersfr1Pdsch256qamFr1
	PdschReMappingfr1Persymbol            *PhyParametersfr1PdschReMappingfr1Persymbol
	PdschReMappingfr1Perslot              *PhyParametersfr1PdschReMappingfr1Perslot              `ext`
	PdcchMonitoringsinglespanfirst4symR16 *PhyParametersfr1PdcchMonitoringsinglespanfirst4symR16 `ext`
}
