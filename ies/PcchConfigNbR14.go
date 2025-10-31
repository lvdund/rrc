package ies

// PCCH-Config-NB-r14 ::= SEQUENCE
// Extensible
type PcchConfigNbR14 struct {
	NpdcchNumrepetitionpagingR14 *PcchConfigNbR14NpdcchNumrepetitionpagingR14
	PagingweightR14              PagingweightNbR14
}
