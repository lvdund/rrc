package ies

// SPS-ConfigUL-setup ::= SEQUENCE
// Extensible
type SpsConfigulSetup struct {
	Semipersistschedintervalul SpsConfigulSetupSemipersistschedintervalul
	Implicitreleaseafter       SpsConfigulSetupImplicitreleaseafter
	P0Persistent               *SpsConfigulSetupP0Persistent
	Twointervalsconfig         *bool
}
