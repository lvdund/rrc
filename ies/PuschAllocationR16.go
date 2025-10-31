package ies

import "rrc/utils"

// PUSCH-Allocation-r16 ::= SEQUENCE
// Extensible
type PuschAllocationR16 struct {
	MappingtypeR16            *PuschAllocationR16MappingtypeR16
	StartsymbolandlengthR16   *utils.INTEGER `lb:0,ub:127`
	StartsymbolR16            *utils.INTEGER `lb:0,ub:13`
	LengthR16                 *utils.INTEGER `lb:0,ub:14`
	NumberofrepetitionsR16    *PuschAllocationR16NumberofrepetitionsR16
	NumberofrepetitionsextR17 *PuschAllocationR16NumberofrepetitionsextR17 `ext`
	NumberofslotstbomsR17     *PuschAllocationR16NumberofslotstbomsR17     `ext`
	Extendedk2R17             *utils.INTEGER                               `lb:0,ub:128,ext`
}
