package ies

import "rrc/utils"

// RRCEarlyDataRequest-5GC-r16-IEs ::= SEQUENCE
type Rrcearlydatarequest5gcR16 struct {
	Ng5gSTmsiR16             Ng5gSTmsiR15
	EstablishmentcauseR16    Rrcearlydatarequest5gcR16IesEstablishmentcauseR16
	DedicatedinfonasR16      Dedicatedinfonas
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *Rrcearlydatarequest5gcR16IesNoncriticalextension
}
