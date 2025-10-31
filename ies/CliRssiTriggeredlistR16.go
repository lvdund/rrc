package ies

// CLI-RSSI-TriggeredList-r16 ::= SEQUENCE OF RSSI-ResourceId-r16
// SIZE (1.. maxNrofCLI-RSSI-Resources-r16)
type CliRssiTriggeredlistR16 struct {
	Value []RssiResourceidR16 `lb:1,ub:maxNrofCLIRssiResourcesR16`
}
