package ies

import "rrc/utils"

// GroupB-ConfiguredTwoStepRA-r16 ::= SEQUENCE
type GroupbConfiguredtwostepraR16 struct {
	RaMsgaSizegroupa          GroupbConfiguredtwostepraR16RaMsgaSizegroupa
	Messagepoweroffsetgroupb  GroupbConfiguredtwostepraR16Messagepoweroffsetgroupb
	NumberofraPreamblesgroupa utils.INTEGER `lb:0,ub:64`
}
