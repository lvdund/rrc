package ies

// PortIndexFor8Ranks ::= CHOICE
const (
	Portindexfor8ranksChoiceNothing = iota
	Portindexfor8ranksChoicePortindex8
	Portindexfor8ranksChoicePortindex4
	Portindexfor8ranksChoicePortindex2
	Portindexfor8ranksChoicePortindex1
)

type Portindexfor8ranks struct {
	Choice     uint64
	Portindex8 *Portindexfor8ranksPortindex8 `lb:2,ub:2`
	Portindex4 *Portindexfor8ranksPortindex4 `lb:2,ub:2`
	Portindex2 *Portindexfor8ranksPortindex2 `lb:2,ub:2`
	Portindex1 *struct{}
}
