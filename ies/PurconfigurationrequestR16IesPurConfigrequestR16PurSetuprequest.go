package ies

// PURConfigurationRequest-r16-IEs-pur-ConfigRequest-r16-pur-SetupRequest ::= SEQUENCE
type PurconfigurationrequestR16IesPurConfigrequestR16PurSetuprequest struct {
	RequestednumoccasionsR16         PurconfigurationrequestR16IesPurConfigrequestR16PurSetuprequestRequestednumoccasionsR16
	RequestedperiodicityandoffsetR16 *PurPeriodicityandoffsetR16
	RequestedtbsR16                  PurconfigurationrequestR16IesPurConfigrequestR16PurSetuprequestRequestedtbsR16
	RrcAckR16                        *bool
}
