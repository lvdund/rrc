package ies

// PortIndexFor8Ranks-portIndex4 ::= SEQUENCE
type Portindexfor8ranksPortindex4 struct {
	Rank14 *Portindex4
	Rank24 *[]Portindex4 `lb:2,ub:2`
	Rank34 *[]Portindex4 `lb:3,ub:3`
	Rank44 *[]Portindex4 `lb:4,ub:4`
}
