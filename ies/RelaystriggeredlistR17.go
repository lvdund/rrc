package ies

// RelaysTriggeredList-r17 ::= SEQUENCE OF SL-SourceIdentity-r17
// SIZE (1.. maxNrofRelayMeas-r17)
type RelaystriggeredlistR17 struct {
	Value []SlSourceidentityR17 `lb:1,ub:maxNrofRelayMeasR17`
}
