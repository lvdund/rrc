package ies

// BandCombinationListSL-Discovery-r17 ::= SEQUENCE OF BandParametersSidelinkDiscovery-r17
// SIZE (1..maxSimultaneousBands)
type BandcombinationlistslDiscoveryR17 struct {
	Value []BandparameterssidelinkdiscoveryR17 `lb:1,ub:maxSimultaneousBands`
}
