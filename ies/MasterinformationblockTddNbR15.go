package ies

import "rrc/utils"

// MasterInformationBlock-TDD-NB-r15 ::= SEQUENCE
type MasterinformationblockTddNbR15 struct {
	SystemframenumberMsbR15 utils.BITSTRING `lb:4,ub:4`
	HypersfnLsbR15          utils.BITSTRING `lb:2,ub:2`
	Schedulinginfosib1R15   utils.INTEGER   `lb:0,ub:15`
	SysteminfovaluetagR15   utils.INTEGER   `lb:0,ub:31`
	AbEnabledR15            utils.BOOLEAN
	OperationmodeinfoR15    MasterinformationblockTddNbR15OperationmodeinfoR15
	Sib1CarrierinfoR15      MasterinformationblockTddNbR15Sib1CarrierinfoR15
	AbEnabled5gcR16         utils.BOOLEAN
	Spare                   utils.BITSTRING `lb:8,ub:8`
}
