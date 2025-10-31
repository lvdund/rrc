package ies

import "rrc/utils"

// UplinkTxDirectCurrentCarrierInfo-r16-servCellInfo-r16 ::= CHOICE
const (
	UplinktxdirectcurrentcarrierinfoR16ServcellinfoR16ChoiceNothing = iota
	UplinktxdirectcurrentcarrierinfoR16ServcellinfoR16ChoiceBwpIdR16
	UplinktxdirectcurrentcarrierinfoR16ServcellinfoR16ChoiceDeactivatedcarrierR16
)

type UplinktxdirectcurrentcarrierinfoR16ServcellinfoR16 struct {
	Choice                uint64
	BwpIdR16              *BwpId
	DeactivatedcarrierR16 *utils.ENUMERATED
}
