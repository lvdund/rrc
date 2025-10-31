package ies

// RSSI-ResourceListConfigCLI-r16 ::= SEQUENCE OF RSSI-ResourceConfigCLI-r16
// SIZE (1.. maxNrofCLI-RSSI-Resources-r16)
type RssiResourcelistconfigcliR16 struct {
	Value []RssiResourceconfigcliR16 `lb:1,ub:maxNrofCLIRssiResourcesR16`
}
