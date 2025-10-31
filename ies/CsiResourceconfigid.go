package ies

import "rrc/utils"

// CSI-ResourceConfigId ::= utils.INTEGER (0..maxNrofCSI-ResourceConfigurations-1)
type CsiResourceconfigid struct {
	Value utils.INTEGER `lb:0,ub:maxNrofCSIResourceconfigurations1`
}
