package ies

import "rrc/utils"

// RACH-ConfigGeneric ::= SEQUENCE
// Extensible
type RachConfiggeneric struct {
	PrachConfigurationindex               utils.INTEGER `lb:0,ub:255`
	Msg1Fdm                               RachConfiggenericMsg1Fdm
	Msg1Frequencystart                    utils.INTEGER `lb:0,ub:maxNrofPhysicalResourceBlocks1`
	Zerocorrelationzoneconfig             utils.INTEGER `lb:0,ub:15`
	Preamblereceivedtargetpower           utils.INTEGER `lb:0,ub:-60`
	Preambletransmax                      RachConfiggenericPreambletransmax
	Powerrampingstep                      RachConfiggenericPowerrampingstep
	RaResponsewindow                      RachConfiggenericRaResponsewindow
	PrachConfigurationperiodscalingIabR16 *RachConfiggenericPrachConfigurationperiodscalingIabR16 `ext`
	PrachConfigurationframeoffsetIabR16   *utils.INTEGER                                          `lb:0,ub:63,ext`
	PrachConfigurationsoffsetIabR16       *utils.INTEGER                                          `lb:0,ub:39,ext`
	RaResponsewindowV1610                 *RachConfiggenericRaResponsewindowV1610                 `ext`
	PrachConfigurationindexV1610          *utils.INTEGER                                          `lb:0,ub:262,ext`
	RaResponsewindowV1700                 *RachConfiggenericRaResponsewindowV1700                 `ext`
}
