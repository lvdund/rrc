package ies

import "rrc/utils"

// MasterInformationBlock-TDD-NB-r15 ::= SEQUENCE
type MasterinformationblockTddNbR15 struct {
	SystemframenumberMsbR15 utils.BITSTRING
	HypersfnLsbR15          utils.BITSTRING
	Schedulinginfosib1R15   utils.INTEGER
	SysteminfovaluetagR15   utils.INTEGER
	AbEnabledR15            bool
	OperationmodeinfoR15    interface{}
	Sib1CarrierinfoR15      utils.ENUMERATED
	AbEnabled5gcR16         bool
	Spare                   utils.BITSTRING
}
