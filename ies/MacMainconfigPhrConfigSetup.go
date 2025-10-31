package ies

// MAC-MainConfig-phr-Config-setup ::= SEQUENCE
type MacMainconfigPhrConfigSetup struct {
	PeriodicphrTimer MacMainconfigPhrConfigSetupPeriodicphrTimer
	ProhibitphrTimer MacMainconfigPhrConfigSetupProhibitphrTimer
	DlPathlosschange MacMainconfigPhrConfigSetupDlPathlosschange
}
