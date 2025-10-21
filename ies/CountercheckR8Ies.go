package ies

import "rrc/utils"

// CounterCheck-r8-IEs ::= SEQUENCE
type CountercheckR8Ies struct {
	DrbCountmsbInfolist  DrbCountmsbInfolist
	Noncriticalextension *CountercheckV8a0Ies
}
