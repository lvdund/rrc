package ies

import "rrc/utils"

// STTI-SPT-BandParameters-r15 ::= SEQUENCE
// Extensible
type SttiSptBandparametersR15 struct {
	Dl1024qamSlotR15                     *utils.ENUMERATED
	Dl1024qamSubslotta1R15               *utils.ENUMERATED
	Dl1024qamSubslotta2R15               *utils.ENUMERATED
	SimultaneoustxDifferenttxDurationR15 *utils.ENUMERATED
	SttiCaMimoParametersdlR15            *CaMimoParametersdlR15
	SttiCaMimoParametersulR15            CaMimoParametersulR15
	SttiFdMimoCoexistence                *utils.ENUMERATED
	SttiMimoCaParametersperbobcsR15      *MimoCaParametersperbobcR13
	SttiMimoCaParametersperbobcsV1530    *MimoCaParametersperbobcV1430
	SttiSupportedcombinationsR15         *SttiSupportedcombinationsR15
	SttiSupportedcsiProcR15              *utils.ENUMERATED
	Ul256qamSlotR15                      *utils.ENUMERATED
	Ul256qamSubslotR15                   *utils.ENUMERATED
}
