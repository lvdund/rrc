package ies

import "rrc/utils"

// SL-ConfiguredGrantConfig-r16-rrc-ConfiguredSidelinkGrant-r16 ::= SEQUENCE
type SlConfiguredgrantconfigR16RrcConfiguredsidelinkgrantR16 struct {
	SlTimeresourcecgType1R16    *utils.INTEGER `lb:0,ub:496`
	SlStartsubchannelcgType1R16 *utils.INTEGER `lb:0,ub:26`
	SlFreqresourcecgType1R16    *utils.INTEGER `lb:0,ub:6929`
	SlTimeoffsetcgType1R16      *utils.INTEGER `lb:0,ub:7999`
	SlN1pucchAnR16              *PucchResourceid
	SlPsfchTopucchCgType1R16    *utils.INTEGER `lb:0,ub:15`
	SlResourcepoolidR16         *SlResourcepoolidR16
	SlTimereferencesfnType1R16  *SlConfiguredgrantconfigR16RrcConfiguredsidelinkgrantR16SlTimereferencesfnType1R16
}
