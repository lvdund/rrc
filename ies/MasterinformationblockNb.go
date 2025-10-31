package ies

import "rrc/utils"

// MasterInformationBlock-NB ::= SEQUENCE
type MasterinformationblockNb struct {
	SystemframenumberMsbR13       utils.BITSTRING `lb:4,ub:4`
	HypersfnLsbR13                utils.BITSTRING `lb:2,ub:2`
	Schedulinginfosib1R13         utils.INTEGER   `lb:0,ub:15`
	SysteminfovaluetagR13         utils.INTEGER   `lb:0,ub:31`
	AbEnabledR13                  utils.BOOLEAN
	OperationmodeinfoR13          MasterinformationblockNbOperationmodeinfoR13
	Additionaltransmissionsib1R15 utils.BOOLEAN
	AbEnabled5gcR16               utils.BOOLEAN
	Spare                         utils.BITSTRING `lb:9,ub:9`
}
