package ies

import "rrc/utils"

// MasterInformationBlock-NB ::= SEQUENCE
type MasterinformationblockNb struct {
	SystemframenumberMsbR13       utils.BITSTRING
	HypersfnLsbR13                utils.BITSTRING
	Schedulinginfosib1R13         utils.INTEGER
	SysteminfovaluetagR13         utils.INTEGER
	AbEnabledR13                  bool
	OperationmodeinfoR13          interface{}
	Additionaltransmissionsib1R15 bool
	AbEnabled5gcR16               bool
	Spare                         utils.BITSTRING
}
