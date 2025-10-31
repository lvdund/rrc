package ies

import "rrc/utils"

// CSI-RS-ProcFrameworkForSRS ::= SEQUENCE
type CsiRsProcframeworkforsrs struct {
	MaxnumberperiodicsrsAssoccsiRsPerbwp  utils.INTEGER `lb:0,ub:4`
	MaxnumberaperiodicsrsAssoccsiRsPerbwp utils.INTEGER `lb:0,ub:4`
	MaxnumberspSrsAssoccsiRsPerbwp        utils.INTEGER `lb:0,ub:4`
	SimultaneoussrsAssoccsiRsPercc        utils.INTEGER `lb:0,ub:8`
}
