package ies

import "rrc/utils"

// SPDCCH-Elements-r15-setup ::= SEQUENCE
// Extensible
type SpdcchElementsR15Setup struct {
	SpdcchSetconfigidR15            *utils.INTEGER `lb:0,ub:3`
	SpdcchSetreferencesigR15        *SpdcchElementsR15SetupSpdcchSetreferencesigR15
	TransmissiontypeR15             *SpdcchElementsR15SetupTransmissiontypeR15
	SpdcchNoofsymbolsR15            *utils.INTEGER                                    `lb:0,ub:2`
	DmrsScramblingsequenceintR15    *utils.INTEGER                                    `lb:0,ub:503`
	Dci7CandidatesperalPdcchR15     *[]Dci7CandidatesR15                              `lb:1,ub:4`
	Dci7CandidatesetsperalSpdcchR15 *[]Dci7CandidatesperalSpdcchR15                   `lb:1,ub:2`
	ResourceblockassignmentR15      *SpdcchElementsR15SetupResourceblockassignmentR15 `lb:98,ub:98`
	SubslotapplicabilityR15         *utils.BITSTRING                                  `lb:5,ub:5`
	AlStartingpointspdcchR15        *[]utils.INTEGER                                  `lb:1,ub:4`
	SubframetypeR15                 *SpdcchElementsR15SetupSubframetypeR15
	RatematchingmodeR15             *SpdcchElementsR15SetupRatematchingmodeR15
}
