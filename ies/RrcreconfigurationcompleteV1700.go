package ies

// RRCReconfigurationComplete-v1700-IEs ::= SEQUENCE
type RrcreconfigurationcompleteV1700 struct {
	NeedforgapncsgInfonrR17    *NeedforgapncsgInfonrR17
	NeedforgapncsgInfoeutraR17 *NeedforgapncsgInfoeutraR17
	SelectedcondrrcreconfigR17 *CondreconfigidR16
	Noncriticalextension       *RrcreconfigurationcompleteV1720
}
