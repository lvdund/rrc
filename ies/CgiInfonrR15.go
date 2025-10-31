package ies

// CGI-InfoNR-r15 ::= SEQUENCE
// Extensible
type CgiInfonrR15 struct {
	PlmnIdentityinfolistR15 *PlmnIdentityinfolistnrR15
	FrequencybandlistR15    *MultifrequencybandlistnrR15
	Nosib1R15               *CgiInfonrR15Nosib1R15
}
