package ies

import "rrc/utils"

// RRCEarlyDataRequest-5GC-NB-r16-IEs ::= SEQUENCE
type Rrcearlydatarequest5gcNbR16Ies struct {
	Ng5gSTmsiR16             Ng5gSTmsiR15
	EstablishmentcauseR16    utils.ENUMERATED
	CqiNpdcchR16             *CqiNpdcchNbR14
	DedicatedinfonasR16      Dedicatedinfonas
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
