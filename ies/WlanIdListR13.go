package ies

// WLAN-Id-List-r13 ::= SEQUENCE OF WLAN-Identifiers-r12
// SIZE (1..maxWLAN-Id-r13)
type WlanIdListR13 struct {
	Value []WlanIdentifiersR12 `lb:1,ub:maxWLANIdR13`
}
