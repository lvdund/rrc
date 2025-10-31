package ies

import "rrc/utils"

// SL-ResourcePool-r16 ::= SEQUENCE
// Extensible
type SlResourcepoolR16 struct {
	SlPscchConfigR16               *utils.Setuprelease[SlPscchConfigR16]
	SlPsschConfigR16               *utils.Setuprelease[SlPsschConfigR16]
	SlPsfchConfigR16               *utils.Setuprelease[SlPsfchConfigR16]
	SlSyncallowedR16               *SlSyncallowedR16
	SlSubchannelsizeR16            *SlResourcepoolR16SlSubchannelsizeR16
	Dummy                          *utils.INTEGER `lb:0,ub:160`
	SlStartrbSubchannelR16         *utils.INTEGER `lb:0,ub:265`
	SlNumsubchannelR16             *utils.INTEGER `lb:0,ub:27`
	SlAdditionalMcsTableR16        *SlResourcepoolR16SlAdditionalMcsTableR16
	SlThreshsRssiCbrR16            *utils.INTEGER `lb:0,ub:45`
	SlTimewindowsizecbrR16         *SlResourcepoolR16SlTimewindowsizecbrR16
	SlTimewindowsizecrR16          *SlResourcepoolR16SlTimewindowsizecrR16
	SlPtrsConfigR16                *SlPtrsConfigR16
	SlUeSelectedconfigrpR16        *SlUeSelectedconfigrpR16
	SlRxparametersncellR16         *SlResourcepoolR16SlRxparametersncellR16
	SlZoneconfigmcrListR16         *[]SlZoneconfigmcrR16 `lb:16,ub:16`
	SlFiltercoefficientR16         *Filtercoefficient
	SlRbNumberR16                  *utils.INTEGER `lb:0,ub:275`
	SlPreemptionenableR16          *SlResourcepoolR16SlPreemptionenableR16
	SlPrioritythresholdUlUrllcR16  *utils.INTEGER `lb:0,ub:9`
	SlPrioritythresholdR16         *utils.INTEGER `lb:0,ub:9`
	SlXOverheadR16                 *SlResourcepoolR16SlXOverheadR16
	SlPowercontrolR16              *SlPowercontrolR16
	SlTxpercentagelistR16          *SlTxpercentagelistR16
	SlMinmaxmcsListR16             *SlMinmaxmcsListR16
	SlTimeresourceR16              *utils.BITSTRING                                    `lb:10,ub:160,ext`
	SlPbpsCpsConfigR17             *utils.Setuprelease[SlPbpsCpsConfigR17]             `ext`
	SlInterueCoordinationconfigR17 *utils.Setuprelease[SlInterueCoordinationconfigR17] `ext`
}
