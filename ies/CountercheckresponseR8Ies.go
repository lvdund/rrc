package ies

import "rrc/utils"

// CounterCheckResponse-r8-IEs ::= SEQUENCE
type CountercheckresponseR8Ies struct {
	DrbCountinfolist     DrbCountinfolist
	Noncriticalextension *CountercheckresponseV8a0Ies
}
