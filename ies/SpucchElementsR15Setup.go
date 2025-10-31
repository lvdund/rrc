package ies

import "rrc/utils"

// SPUCCH-Elements-r15-setup ::= SEQUENCE
type SpucchElementsR15Setup struct {
	N1subslotspucchAnListR15                  *[]utils.INTEGER       `lb:1,ub:4`
	N1slotspucchFhAnListR15                   *utils.INTEGER         `lb:0,ub:1319`
	N1slotspucchNofhAnListR15                 *utils.INTEGER         `lb:0,ub:3959`
	N3spucchAnListR15                         *utils.INTEGER         `lb:0,ub:549`
	N4spucchslotResourceR15                   *[]N4spucchResourceR15 `lb:1,ub:2`
	N4spucchsubslotResourceR15                *[]N4spucchResourceR15 `lb:1,ub:2`
	N4maxcoderateslotpucchR15                 *utils.INTEGER         `lb:0,ub:7`
	N4maxcoderatesubslotpucchR15              *utils.INTEGER         `lb:0,ub:7`
	N4maxcoderatemultiresourceslotpucchR15    *utils.INTEGER         `lb:0,ub:7`
	N4maxcoderatemultiresourcesubslotpucchR15 *utils.INTEGER         `lb:0,ub:7`
}
