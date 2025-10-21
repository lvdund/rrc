package ies

import "rrc/utils"

// RRCEarlyDataRequest-5GC-r16-IEs ::= SEQUENCE
type Rrcearlydatarequest5gcR16Ies struct {
	Ng5gSTmsiR16             Ng5gSTmsiR15
	EstablishmentcauseR16    utils.ENUMERATED
	DedicatedinfonasR16      Dedicatedinfonas
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
