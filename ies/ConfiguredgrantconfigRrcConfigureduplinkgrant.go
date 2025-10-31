package ies

import "rrc/utils"

// ConfiguredGrantConfig-rrc-ConfiguredUplinkGrant ::= SEQUENCE
// Extensible
type ConfiguredgrantconfigRrcConfigureduplinkgrant struct {
	Timedomainoffset                 utils.INTEGER   `lb:0,ub:5119`
	Timedomainallocation             utils.INTEGER   `lb:0,ub:15`
	Frequencydomainallocation        utils.BITSTRING `lb:18,ub:18`
	Antennaport                      utils.INTEGER   `lb:0,ub:31`
	DmrsSeqinitialization            *utils.INTEGER  `lb:0,ub:1`
	Precodingandnumberoflayers       utils.INTEGER   `lb:0,ub:63`
	SrsResourceindicator             *utils.INTEGER  `lb:0,ub:15`
	Mcsandtbs                        utils.INTEGER   `lb:0,ub:31`
	Frequencyhoppingoffset           *utils.INTEGER  `lb:0,ub:maxNrofPhysicalResourceBlocks1`
	Pathlossreferenceindex           utils.INTEGER   `lb:0,ub:maxNrofPUSCHPathlossreferencerss1`
	PuschReptypeindicatorR16         *ConfiguredgrantconfigRrcConfigureduplinkgrantPuschReptypeindicatorR16
	FrequencyhoppingpuschReptypebR16 *ConfiguredgrantconfigRrcConfigureduplinkgrantFrequencyhoppingpuschReptypebR16
	TimereferencesfnR16              *ConfiguredgrantconfigRrcConfigureduplinkgrantTimereferencesfnR16
	Pathlossreferenceindex2R17       *utils.INTEGER `lb:0,ub:maxNrofPUSCHPathlossreferencerss1`
	SrsResourceindicator2R17         *utils.INTEGER `lb:0,ub:15`
	Precodingandnumberoflayers2R17   *utils.INTEGER `lb:0,ub:63`
	TimedomainallocationV1710        *utils.INTEGER `lb:0,ub:63`
	TimedomainoffsetR17              *utils.INTEGER `lb:0,ub:40959`
	CgSdtConfigurationR17            *CgSdtConfigurationR17
}
