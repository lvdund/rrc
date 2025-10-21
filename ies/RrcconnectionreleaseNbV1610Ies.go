package ies

import "rrc/utils"

// RRCConnectionRelease-NB-v1610-IEs ::= SEQUENCE
type RrcconnectionreleaseNbV1610Ies struct {
	ResumeidentityR16    *IRntiR15
	AnrMeasconfigR16     *AnrMeasconfigNbR16
	PurConfigR16         *Setuprelease
	Noncriticalextension *interface{}
}
