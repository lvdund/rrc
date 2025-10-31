package ies

// Phy-ParametersFR2 ::= SEQUENCE
// Extensible
type PhyParametersfr2 struct {
	Dummy                                             *PhyParametersfr2Dummy
	PdschReMappingfr2Persymbol                        *PhyParametersfr2PdschReMappingfr2Persymbol
	PcellFr2                                          *PhyParametersfr2PcellFr2                                          `ext`
	PdschReMappingfr2Perslot                          *PhyParametersfr2PdschReMappingfr2Perslot                          `ext`
	DefaultspatialrelationpathlossrsR16               *PhyParametersfr2DefaultspatialrelationpathlossrsR16               `ext`
	SpatialrelationupdateapSrsR16                     *PhyParametersfr2SpatialrelationupdateapSrsR16                     `ext`
	MaxnumbersrsPosspatialrelationsallservingcellsR16 *PhyParametersfr2MaxnumbersrsPosspatialrelationsallservingcellsR16 `ext`
}
