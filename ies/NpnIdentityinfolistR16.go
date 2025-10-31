package ies

// NPN-IdentityInfoList-r16 ::= SEQUENCE OF NPN-IdentityInfo-r16
// SIZE (1..maxNPN-r16)
type NpnIdentityinfolistR16 struct {
	Value []NpnIdentityinfoR16 `lb:1,ub:maxNPNR16`
}
