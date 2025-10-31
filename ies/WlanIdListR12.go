package ies

// WLAN-Id-List-r12 ::= SEQUENCE OF WLAN-Identifiers-r12
// SIZE (1..maxWLAN-Id-r12)
type WlanIdListR12 struct {
	Value []WlanIdentifiersR12 `lb:1,ub:maxWLANIdR12`
}
