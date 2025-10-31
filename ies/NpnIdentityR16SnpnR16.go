package ies

// NPN-Identity-r16-snpn-r16 ::= SEQUENCE
type NpnIdentityR16SnpnR16 struct {
	PlmnIdentityR16 PlmnIdentity
	NidListR16      []NidR16 `lb:1,ub:maxNPNR16`
}
