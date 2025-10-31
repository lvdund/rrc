package ies

import "rrc/utils"

// ULHandoverPreparationTransfer-r8-IEs ::= SEQUENCE
type UlhandoverpreparationtransferR8 struct {
	Cdma2000Type         Cdma2000Type
	Meid                 *utils.BITSTRING `lb:56,ub:56`
	Dedicatedinfo        Dedicatedinfocdma2000
	Noncriticalextension *UlhandoverpreparationtransferV8a0
}
