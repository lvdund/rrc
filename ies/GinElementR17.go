package ies

// GIN-Element-r17 ::= SEQUENCE
type GinElementR17 struct {
	PlmnIdentityR17 PlmnIdentity
	NidListR17      []NidR16 `lb:1,ub:maxGINR17`
}
