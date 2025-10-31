package ies

// DRX-Config-setup ::= SEQUENCE
type DrxConfigSetup struct {
	Ondurationtimer         DrxConfigSetupOndurationtimer
	DrxInactivitytimer      DrxConfigSetupDrxInactivitytimer
	DrxRetransmissiontimer  DrxConfigSetupDrxRetransmissiontimer
	LongdrxCyclestartoffset DrxConfigSetupLongdrxCyclestartoffset
	Shortdrx                *DrxConfigSetupShortdrx
}
