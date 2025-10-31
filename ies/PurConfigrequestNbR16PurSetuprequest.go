package ies

// PUR-ConfigRequest-NB-r16-pur-SetupRequest ::= SEQUENCE
type PurConfigrequestNbR16PurSetuprequest struct {
	RequestednumoccasionsR16         PurConfigrequestNbR16PurSetuprequestRequestednumoccasionsR16
	RequestedperiodicityandoffsetR16 PurPeriodicityandoffsetNbR16
	RequestedtbsR16                  PurConfigrequestNbR16PurSetuprequestRequestedtbsR16
	RrcAckR16                        *bool
}
