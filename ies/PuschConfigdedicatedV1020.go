package ies

// PUSCH-ConfigDedicated-v1020 ::= SEQUENCE
type PuschConfigdedicatedV1020 struct {
	BetaoffsetmcR10         *PuschConfigdedicatedV1020BetaoffsetmcR10
	GrouphoppingdisabledR10 *bool
	DmrsWithoccActivatedR10 *bool
}
