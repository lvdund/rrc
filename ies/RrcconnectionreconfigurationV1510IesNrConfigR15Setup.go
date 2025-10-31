package ies

import "rrc/utils"

// RRCConnectionReconfiguration-v1510-IEs-nr-Config-r15-setup ::= SEQUENCE
type RrcconnectionreconfigurationV1510IesNrConfigR15Setup struct {
	EndcReleaseandaddR15          utils.BOOLEAN
	NrSecondarycellgroupconfigR15 *utils.OCTETSTRING
	PMaxeutraR15                  *PMax
}
