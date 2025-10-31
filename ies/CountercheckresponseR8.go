package ies

// CounterCheckResponse-r8-IEs ::= SEQUENCE
type CountercheckresponseR8 struct {
	DrbCountinfolist     DrbCountinfolist
	Noncriticalextension *CountercheckresponseV8a0
}
