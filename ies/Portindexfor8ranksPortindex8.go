package ies

// PortIndexFor8Ranks-portIndex8 ::= SEQUENCE
type Portindexfor8ranksPortindex8 struct {
	Rank18 *Portindex8
	Rank28 *[]Portindex8 `lb:2,ub:2`
	Rank38 *[]Portindex8 `lb:3,ub:3`
	Rank48 *[]Portindex8 `lb:4,ub:4`
	Rank58 *[]Portindex8 `lb:5,ub:5`
	Rank68 *[]Portindex8 `lb:6,ub:6`
	Rank78 *[]Portindex8 `lb:7,ub:7`
	Rank88 *[]Portindex8 `lb:8,ub:8`
}
