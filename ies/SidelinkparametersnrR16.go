package ies

// SidelinkParametersNR-r16 ::= SEQUENCE
// Extensible
type SidelinkparametersnrR16 struct {
	RlcParameterssidelinkR16        *RlcParameterssidelinkR16
	MacParameterssidelinkR16        *MacParameterssidelinkR16
	FddAddUeSidelinkCapabilitiesR16 *UeSidelinkcapabilityaddxddModeR16
	TddAddUeSidelinkCapabilitiesR16 *UeSidelinkcapabilityaddxddModeR16
	SupportedbandlistsidelinkR16    *[]BandsidelinkR16                        `lb:1,ub:maxBands`
	RelayparametersR17              *RelayparametersR17                       `ext`
	P0OlpcSidelinkR17               *SidelinkparametersnrR16P0OlpcSidelinkR17 `ext`
}
