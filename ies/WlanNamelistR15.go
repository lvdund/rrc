package ies

// WLAN-NameList-r15 ::= SEQUENCE OF WLAN-Name-r15
// SIZE (1..maxWLAN-Name-r15)
type WlanNamelistR15 struct {
	Value []WlanNameR15 `lb:1,ub:maxWLANNameR15`
}
