package ies

import "rrc/utils"

// SIB1-v1700-IEs ::= SEQUENCE
type Sib1V1700 struct {
	HsdnCellR17                   *utils.BOOLEAN
	UacBarringinfoV1700           *Sib1V1700IesUacBarringinfoV1700
	SdtConfigcommonR17            *SdtConfigcommonsibR17
	RedcapConfigcommonR17         *RedcapConfigcommonsibR17
	FeatureprioritiesR17          *Sib1V1700IesFeatureprioritiesR17
	SiSchedulinginfoV1700         *SiSchedulinginfoV1700
	HypersfnR17                   *utils.BITSTRING `lb:10,ub:10`
	EdrxAllowedidleR17            *utils.BOOLEAN
	EdrxAllowedinactiveR17        *utils.BOOLEAN
	IntrafreqreselectionredcapR17 *Sib1V1700IesIntrafreqreselectionredcapR17
	CellbarredntnR17              *Sib1V1700IesCellbarredntnR17
	Noncriticalextension          *Sib1V1700IesNoncriticalextension
}
