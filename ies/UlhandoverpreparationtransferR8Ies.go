package ies

import "rrc/utils"

// ULHandoverPreparationTransfer-r8-IEs ::= SEQUENCE
type UlhandoverpreparationtransferR8Ies struct {
	Cdma2000Type         Cdma2000Type
	Meid                 *utils.BITSTRING
	Dedicatedinfo        Dedicatedinfocdma2000
	Noncriticalextension *UlhandoverpreparationtransferV8a0Ies
}
