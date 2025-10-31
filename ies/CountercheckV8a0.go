package ies

import "rrc/utils"

// CounterCheck-v8a0-IEs ::= SEQUENCE
type CountercheckV8a0 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *CountercheckV1530
}
