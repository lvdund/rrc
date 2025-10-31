package ies

// PortIndexFor8Ranks-portIndex2 ::= SEQUENCE
type Portindexfor8ranksPortindex2 struct {
	Rank12 *Portindex2
	Rank22 *[]Portindex2 `lb:2,ub:2`
}
