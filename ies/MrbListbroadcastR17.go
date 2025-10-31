package ies

// MRB-ListBroadcast-r17 ::= SEQUENCE OF MRB-InfoBroadcast-r17
// SIZE (1..maxNrofMRB-Broadcast-r17)
type MrbListbroadcastR17 struct {
	Value []MrbInfobroadcastR17 `lb:1,ub:maxNrofMRBBroadcastR17`
}
