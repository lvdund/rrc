package ies

import "rrc/utils"

// SL-PowerControl-r16 ::= SEQUENCE
// Extensible
type SlPowercontrolR16 struct {
	SlMaxtranspowerR16   utils.INTEGER `lb:0,ub:33`
	SlAlphaPsschPscchR16 *SlPowercontrolR16SlAlphaPsschPscchR16
	DlAlphaPsschPscchR16 *SlPowercontrolR16DlAlphaPsschPscchR16
	SlP0PsschPscchR16    *utils.INTEGER `lb:0,ub:15`
	DlP0PsschPscchR16    *utils.INTEGER `lb:0,ub:15`
	DlAlphaPsfchR16      *SlPowercontrolR16DlAlphaPsfchR16
	DlP0PsfchR16         *utils.INTEGER `lb:0,ub:15`
	DlP0PsschPscchR17    *utils.INTEGER `lb:0,ub:24,ext`
	SlP0PsschPscchR17    *utils.INTEGER `lb:0,ub:24,ext`
	DlP0PsfchR17         *utils.INTEGER `lb:0,ub:24,ext`
}
