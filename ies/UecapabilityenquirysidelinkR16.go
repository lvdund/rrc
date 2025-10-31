package ies

import "rrc/utils"

// UECapabilityEnquirySidelink-r16-IEs ::= SEQUENCE
type UecapabilityenquirysidelinkR16 struct {
	FrequencybandlistfiltersidelinkR16 *Freqbandlist
	UeCapabilityinformationsidelinkR16 *utils.OCTETSTRING
	Latenoncriticalextension           *utils.OCTETSTRING
	Noncriticalextension               *UecapabilityenquirysidelinkR16IesNoncriticalextension
}
