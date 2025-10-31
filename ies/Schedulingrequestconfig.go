package ies

// SchedulingRequestConfig ::= SEQUENCE
type Schedulingrequestconfig struct {
	Schedulingrequesttoaddmodlist  *[]Schedulingrequesttoaddmod `lb:1,ub:maxNrofSRConfigpercellgroup`
	Schedulingrequesttoreleaselist *[]Schedulingrequestid       `lb:1,ub:maxNrofSRConfigpercellgroup`
}
