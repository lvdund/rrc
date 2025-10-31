package ies

import "rrc/utils"

// RRCReconfiguration-v1700-IEs ::= SEQUENCE
type RrcreconfigurationV1700 struct {
	OtherconfigV1700              *OtherconfigV1700
	SlL2relayueConfigR17          *utils.Setuprelease[SlL2relayueConfigR17]
	SlL2remoteueConfigR17         *utils.Setuprelease[SlL2remoteueConfigR17]
	DedicatedpagingdeliveryR17    *utils.OCTETSTRING
	NeedforgapncsgConfignrR17     *utils.Setuprelease[NeedforgapncsgConfignrR17]
	NeedforgapncsgConfigeutraR17  *utils.Setuprelease[NeedforgapncsgConfigeutraR17]
	MusimGapconfigR17             *utils.Setuprelease[MusimGapconfigR17]
	UlGapfr2ConfigR17             *utils.Setuprelease[UlGapfr2ConfigR17]
	ScgStateR17                   *RrcreconfigurationV1700IesScgStateR17
	ApplayermeasconfigR17         *ApplayermeasconfigR17
	UeTxtegRequestulTdoaConfigR17 *utils.Setuprelease[UeTxtegRequestulTdoaConfigR17]
	Noncriticalextension          *RrcreconfigurationV1700IesNoncriticalextension
}
