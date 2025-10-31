package ies

import "rrc/utils"

// SC-MTCH-SchedulingInfo-NB-r14-schedulingPeriodStartOffsetSCPTM-r14 ::= CHOICE
const (
	ScMtchSchedulinginfoNbR14SchedulingperiodstartoffsetscptmR14ChoiceNothing = iota
	ScMtchSchedulinginfoNbR14SchedulingperiodstartoffsetscptmR14ChoiceSf10
	ScMtchSchedulinginfoNbR14SchedulingperiodstartoffsetscptmR14ChoiceSf20
	ScMtchSchedulinginfoNbR14SchedulingperiodstartoffsetscptmR14ChoiceSf32
	ScMtchSchedulinginfoNbR14SchedulingperiodstartoffsetscptmR14ChoiceSf40
	ScMtchSchedulinginfoNbR14SchedulingperiodstartoffsetscptmR14ChoiceSf64
	ScMtchSchedulinginfoNbR14SchedulingperiodstartoffsetscptmR14ChoiceSf80
	ScMtchSchedulinginfoNbR14SchedulingperiodstartoffsetscptmR14ChoiceSf128
	ScMtchSchedulinginfoNbR14SchedulingperiodstartoffsetscptmR14ChoiceSf160
	ScMtchSchedulinginfoNbR14SchedulingperiodstartoffsetscptmR14ChoiceSf256
	ScMtchSchedulinginfoNbR14SchedulingperiodstartoffsetscptmR14ChoiceSf320
	ScMtchSchedulinginfoNbR14SchedulingperiodstartoffsetscptmR14ChoiceSf512
	ScMtchSchedulinginfoNbR14SchedulingperiodstartoffsetscptmR14ChoiceSf640
	ScMtchSchedulinginfoNbR14SchedulingperiodstartoffsetscptmR14ChoiceSf1024
	ScMtchSchedulinginfoNbR14SchedulingperiodstartoffsetscptmR14ChoiceSf2048
	ScMtchSchedulinginfoNbR14SchedulingperiodstartoffsetscptmR14ChoiceSf4096
	ScMtchSchedulinginfoNbR14SchedulingperiodstartoffsetscptmR14ChoiceSf8192
)

type ScMtchSchedulinginfoNbR14SchedulingperiodstartoffsetscptmR14 struct {
	Choice uint64
	Sf10   *utils.INTEGER `lb:0,ub:9`
	Sf20   *utils.INTEGER `lb:0,ub:19`
	Sf32   *utils.INTEGER `lb:0,ub:31`
	Sf40   *utils.INTEGER `lb:0,ub:39`
	Sf64   *utils.INTEGER `lb:0,ub:63`
	Sf80   *utils.INTEGER `lb:0,ub:79`
	Sf128  *utils.INTEGER `lb:0,ub:127`
	Sf160  *utils.INTEGER `lb:0,ub:159`
	Sf256  *utils.INTEGER `lb:0,ub:255`
	Sf320  *utils.INTEGER `lb:0,ub:319`
	Sf512  *utils.INTEGER `lb:0,ub:511`
	Sf640  *utils.INTEGER `lb:0,ub:639`
	Sf1024 *utils.INTEGER `lb:0,ub:1023`
	Sf2048 *utils.INTEGER `lb:0,ub:2047`
	Sf4096 *utils.INTEGER `lb:0,ub:4095`
	Sf8192 *utils.INTEGER `lb:0,ub:8191`
}
