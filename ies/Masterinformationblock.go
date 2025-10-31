package ies

import "rrc/utils"

// MasterInformationBlock ::= SEQUENCE
type Masterinformationblock struct {
	DlBandwidth              MasterinformationblockDlBandwidth
	PhichConfig              PhichConfig
	Systemframenumber        utils.BITSTRING `lb:8,ub:8`
	Schedulinginfosib1BrR13  utils.INTEGER   `lb:0,ub:31`
	SysteminfounchangedBrR15 utils.BOOLEAN
	Spare                    utils.BITSTRING `lb:4,ub:4`
}
