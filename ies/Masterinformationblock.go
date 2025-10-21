package ies

import "rrc/utils"

// MasterInformationBlock ::= SEQUENCE
type Masterinformationblock struct {
	DlBandwidth              utils.ENUMERATED
	PhichConfig              PhichConfig
	Systemframenumber        utils.BITSTRING
	Schedulinginfosib1BrR13  utils.INTEGER
	SysteminfounchangedBrR15 bool
	Spare                    utils.BITSTRING
}
