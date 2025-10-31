package ies

// PUSCH-PowerControl-v1610 ::= SEQUENCE
// Extensible
type PuschPowercontrolV1610 struct {
	PathlossreferencerstoaddmodlistsizeextV1610  *[]PuschPathlossreferencersR16     `lb:1,ub:maxNrofPUSCHPathlossreferencerssdiffR16`
	PathlossreferencerstoreleaselistsizeextV1610 *[]PuschPathlossreferencersIdV1610 `lb:1,ub:maxNrofPUSCHPathlossreferencerssdiffR16`
	P0PuschSetlistR16                            *[]P0PuschSetR16                   `lb:1,ub:maxNrofSRIPuschMappings`
	OlpcParameterset                             *PuschPowercontrolV1610OlpcParameterset
	SriPuschMappingtoaddmodlist2R17              *[]SriPuschPowercontrol          `lb:1,ub:maxNrofSRIPuschMappings,ext`
	SriPuschMappingtoreleaselist2R17             *[]SriPuschPowercontrolid        `lb:1,ub:maxNrofSRIPuschMappings,ext`
	P0PuschSetlist2R17                           *[]P0PuschSetR16                 `lb:1,ub:maxNrofSRIPuschMappings,ext`
	Dummy                                        *[]DummypathlossreferencersV1710 `lb:1,ub:maxNrofPUSCHPathlossreferencerssR16,ext`
}
