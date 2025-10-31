package ies

// WLAN-NameList-r16 ::= SEQUENCE OF WLAN-Name-r16
// SIZE (1..maxWLAN-Name-r16)
type WlanNamelistR16 struct {
	Value []WlanNameR16 `lb:1,ub:maxWLANNameR16`
}
