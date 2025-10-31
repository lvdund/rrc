package ies

import "rrc/utils"

// CounterCheck-IEs ::= SEQUENCE
type Countercheck struct {
	DrbCountmsbInfolist      DrbCountmsbInfolist
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *CountercheckIesNoncriticalextension
}
