package ies

import "rrc/utils"

// RRCResume-v1700-IEs ::= SEQUENCE
type RrcresumeV1700 struct {
	SlConfigdedicatednrR17       *utils.Setuprelease[SlConfigdedicatednrR16]
	SlL2remoteueConfigR17        *utils.Setuprelease[SlL2remoteueConfigR17]
	NeedforgapncsgConfignrR17    *utils.Setuprelease[NeedforgapncsgConfignrR17]
	NeedforgapncsgConfigeutraR17 *utils.Setuprelease[NeedforgapncsgConfigeutraR17]
	ScgStateR17                  *RrcresumeV1700IesScgStateR17
	ApplayermeasconfigR17        *ApplayermeasconfigR17
	Noncriticalextension         *RrcresumeV1700IesNoncriticalextension
}
