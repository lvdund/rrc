package ies

import "rrc/utils"

// SchedulingRequestResourceConfigExt-v1700-periodicityAndOffset-r17 ::= CHOICE
const (
	SchedulingrequestresourceconfigextV1700PeriodicityandoffsetR17ChoiceNothing = iota
	SchedulingrequestresourceconfigextV1700PeriodicityandoffsetR17ChoiceSl1280
	SchedulingrequestresourceconfigextV1700PeriodicityandoffsetR17ChoiceSl2560
	SchedulingrequestresourceconfigextV1700PeriodicityandoffsetR17ChoiceSl5120
)

type SchedulingrequestresourceconfigextV1700PeriodicityandoffsetR17 struct {
	Choice uint64
	Sl1280 *utils.INTEGER `lb:0,ub:1279`
	Sl2560 *utils.INTEGER `lb:0,ub:2559`
	Sl5120 *utils.INTEGER `lb:0,ub:5119`
}
