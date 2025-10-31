package ies

// RRCResumeComplete-v1700-IEs ::= SEQUENCE
type RrcresumecompleteV1700 struct {
	NeedforgapncsgInfonrR17    *NeedforgapncsgInfonrR17
	NeedforgapncsgInfoeutraR17 *NeedforgapncsgInfoeutraR17
	Noncriticalextension       *RrcresumecompleteV1720
}
