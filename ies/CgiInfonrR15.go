package ies

import "rrc/utils"

// CGI-InfoNR-r15 ::= SEQUENCE
// Extensible
type CgiInfonrR15 struct {
	PlmnIdentityinfolistR15 *PlmnIdentityinfolistnrR15
	FrequencybandlistR15    *MultifrequencybandlistnrR15
	Nosib1R15               *interface{}
}
