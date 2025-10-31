package ies

import "rrc/utils"

// CSI-RS-Config-r10-csi-RS-r10-setup ::= SEQUENCE
type CsiRsConfigR10CsiRsR10Setup struct {
	AntennaportscountR10 CsiRsConfigR10CsiRsR10SetupAntennaportscountR10
	ResourceconfigR10    utils.INTEGER `lb:0,ub:31`
	SubframeconfigR10    utils.INTEGER `lb:0,ub:154`
	PCR10                utils.INTEGER `lb:0,ub:15`
}
