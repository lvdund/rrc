package ies

import "rrc/utils"

// CounterCheckResponse-v8a0-IEs ::= SEQUENCE
type CountercheckresponseV8a0 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *CountercheckresponseV1530
}
