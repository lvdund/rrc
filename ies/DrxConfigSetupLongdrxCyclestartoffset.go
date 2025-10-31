package ies

import "rrc/utils"

// DRX-Config-setup-longDRX-CycleStartOffset ::= CHOICE
const (
	DrxConfigSetupLongdrxCyclestartoffsetChoiceNothing = iota
	DrxConfigSetupLongdrxCyclestartoffsetChoiceSf10
	DrxConfigSetupLongdrxCyclestartoffsetChoiceSf20
	DrxConfigSetupLongdrxCyclestartoffsetChoiceSf32
	DrxConfigSetupLongdrxCyclestartoffsetChoiceSf40
	DrxConfigSetupLongdrxCyclestartoffsetChoiceSf64
	DrxConfigSetupLongdrxCyclestartoffsetChoiceSf80
	DrxConfigSetupLongdrxCyclestartoffsetChoiceSf128
	DrxConfigSetupLongdrxCyclestartoffsetChoiceSf160
	DrxConfigSetupLongdrxCyclestartoffsetChoiceSf256
	DrxConfigSetupLongdrxCyclestartoffsetChoiceSf320
	DrxConfigSetupLongdrxCyclestartoffsetChoiceSf512
	DrxConfigSetupLongdrxCyclestartoffsetChoiceSf640
	DrxConfigSetupLongdrxCyclestartoffsetChoiceSf1024
	DrxConfigSetupLongdrxCyclestartoffsetChoiceSf1280
	DrxConfigSetupLongdrxCyclestartoffsetChoiceSf2048
	DrxConfigSetupLongdrxCyclestartoffsetChoiceSf2560
)

type DrxConfigSetupLongdrxCyclestartoffset struct {
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
	Sf1280 *utils.INTEGER `lb:0,ub:1279`
	Sf2048 *utils.INTEGER `lb:0,ub:2047`
	Sf2560 *utils.INTEGER `lb:0,ub:2559`
}
