package ies

import "rrc/utils"

// CounterCheckResponse-IEs ::= SEQUENCE
type Countercheckresponse struct {
	DrbCountinfolist         DrbCountinfolist
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *CountercheckresponseIesNoncriticalextension
}
