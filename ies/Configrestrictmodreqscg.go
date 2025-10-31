package ies

import "rrc/utils"

// ConfigRestrictModReqSCG ::= SEQUENCE
// Extensible
type Configrestrictmodreqscg struct {
	RequestedbcMrdc                   *Bandcombinationinfosn
	RequestedpMaxfr1                  *PMax
	RequestedpdcchBlinddetectionscg   *utils.INTEGER `lb:0,ub:15,ext`
	RequestedpMaxeutra                *PMax          `ext`
	RequestedpMaxfr2R16               *PMax          `ext`
	RequestedmaxinterfreqmeasidscgR16 *utils.INTEGER `lb:0,ub:maxMeasIdentitiesMN,ext`
	RequestedmaxintrafreqmeasidscgR16 *utils.INTEGER `lb:0,ub:maxMeasIdentitiesMN,ext`
	RequestedtoffsetR16               *TOffsetR16    `ext`
}
