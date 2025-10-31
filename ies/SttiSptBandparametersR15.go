package ies

// STTI-SPT-BandParameters-r15 ::= SEQUENCE
// Extensible
type SttiSptBandparametersR15 struct {
	Dl1024qamSlotR15                     *SttiSptBandparametersR15Dl1024qamSlotR15
	Dl1024qamSubslotta1R15               *SttiSptBandparametersR15Dl1024qamSubslotta1R15
	Dl1024qamSubslotta2R15               *SttiSptBandparametersR15Dl1024qamSubslotta2R15
	SimultaneoustxDifferenttxDurationR15 *SttiSptBandparametersR15SimultaneoustxDifferenttxDurationR15
	SttiCaMimoParametersdlR15            *CaMimoParametersdlR15
	SttiCaMimoParametersulR15            CaMimoParametersulR15
	SttiFdMimoCoexistence                *SttiSptBandparametersR15SttiFdMimoCoexistence
	SttiMimoCaParametersperbobcsR15      *MimoCaParametersperbobcR13
	SttiMimoCaParametersperbobcsV1530    *MimoCaParametersperbobcV1430
	SttiSupportedcombinationsR15         *SttiSupportedcombinationsR15
	SttiSupportedcsiProcR15              *SttiSptBandparametersR15SttiSupportedcsiProcR15
	Ul256qamSlotR15                      *SttiSptBandparametersR15Ul256qamSlotR15
	Ul256qamSubslotR15                   *SttiSptBandparametersR15Ul256qamSubslotR15
}
