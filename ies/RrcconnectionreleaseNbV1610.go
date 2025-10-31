package ies

// RRCConnectionRelease-NB-v1610-IEs ::= SEQUENCE
type RrcconnectionreleaseNbV1610 struct {
	ResumeidentityR16    *IRntiR15
	AnrMeasconfigR16     *AnrMeasconfigNbR16
	PurConfigR16         *SetupreleasePurConfigR16
	Noncriticalextension *RrcconnectionreleaseNbV1610IesNoncriticalextension
}
