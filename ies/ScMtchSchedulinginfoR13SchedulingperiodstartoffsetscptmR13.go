package ies

import "rrc/utils"

// SC-MTCH-SchedulingInfo-r13-schedulingPeriodStartOffsetSCPTM-r13 ::= CHOICE
const (
	ScMtchSchedulinginfoR13SchedulingperiodstartoffsetscptmR13ChoiceNothing = iota
	ScMtchSchedulinginfoR13SchedulingperiodstartoffsetscptmR13ChoiceSf10
	ScMtchSchedulinginfoR13SchedulingperiodstartoffsetscptmR13ChoiceSf20
	ScMtchSchedulinginfoR13SchedulingperiodstartoffsetscptmR13ChoiceSf32
	ScMtchSchedulinginfoR13SchedulingperiodstartoffsetscptmR13ChoiceSf40
	ScMtchSchedulinginfoR13SchedulingperiodstartoffsetscptmR13ChoiceSf64
	ScMtchSchedulinginfoR13SchedulingperiodstartoffsetscptmR13ChoiceSf80
	ScMtchSchedulinginfoR13SchedulingperiodstartoffsetscptmR13ChoiceSf128
	ScMtchSchedulinginfoR13SchedulingperiodstartoffsetscptmR13ChoiceSf160
	ScMtchSchedulinginfoR13SchedulingperiodstartoffsetscptmR13ChoiceSf256
	ScMtchSchedulinginfoR13SchedulingperiodstartoffsetscptmR13ChoiceSf320
	ScMtchSchedulinginfoR13SchedulingperiodstartoffsetscptmR13ChoiceSf512
	ScMtchSchedulinginfoR13SchedulingperiodstartoffsetscptmR13ChoiceSf640
	ScMtchSchedulinginfoR13SchedulingperiodstartoffsetscptmR13ChoiceSf1024
	ScMtchSchedulinginfoR13SchedulingperiodstartoffsetscptmR13ChoiceSf2048
	ScMtchSchedulinginfoR13SchedulingperiodstartoffsetscptmR13ChoiceSf4096
	ScMtchSchedulinginfoR13SchedulingperiodstartoffsetscptmR13ChoiceSf8192
)

type ScMtchSchedulinginfoR13SchedulingperiodstartoffsetscptmR13 struct {
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
	Sf2048 *utils.INTEGER `lb:0,ub:2048`
	Sf4096 *utils.INTEGER `lb:0,ub:4096`
	Sf8192 *utils.INTEGER `lb:0,ub:8192`
}
