package ies

import "rrc/utils"

// SystemInformationBlockType1-v1320-IEs-freqHoppingParametersDL-r13 ::= SEQUENCE
type Systeminformationblocktype1V1320IesFreqhoppingparametersdlR13 struct {
	MpdcchPdschHoppingnbR13               *Systeminformationblocktype1V1320IesFreqhoppingparametersdlR13MpdcchPdschHoppingnbR13
	IntervalDlhoppingconfigcommonmodeaR13 *Systeminformationblocktype1V1320IesFreqhoppingparametersdlR13IntervalDlhoppingconfigcommonmodeaR13
	IntervalDlhoppingconfigcommonmodebR13 *Systeminformationblocktype1V1320IesFreqhoppingparametersdlR13IntervalDlhoppingconfigcommonmodebR13
	MpdcchPdschHoppingoffsetR13           *utils.INTEGER `lb:0,ub:maxAvailNarrowBandsR13`
}
