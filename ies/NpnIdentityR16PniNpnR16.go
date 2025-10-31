package ies

// NPN-Identity-r16-pni-npn-r16 ::= SEQUENCE
type NpnIdentityR16PniNpnR16 struct {
	PlmnIdentityR16    PlmnIdentity
	CagIdentitylistR16 []CagIdentityinfoR16 `lb:1,ub:maxNPNR16`
}
