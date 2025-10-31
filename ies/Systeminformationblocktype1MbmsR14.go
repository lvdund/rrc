package ies

import "rrc/utils"

// SystemInformationBlockType1-MBMS-r14 ::= SEQUENCE
type Systeminformationblocktype1MbmsR14 struct {
	CellaccessrelatedinfoR14        Systeminformationblocktype1MbmsR14CellaccessrelatedinfoR14
	FreqbandindicatorR14            FreqbandindicatorR11
	MultibandinfolistR14            *MultibandinfolistR11
	SchedulinginfolistMbmsR14       SchedulinginfolistMbmsR14
	SiWindowlengthR14               Systeminformationblocktype1MbmsR14SiWindowlengthR14
	SysteminfovaluetagR14           utils.INTEGER `lb:0,ub:31`
	NonmbsfnSubframeconfigR14       *NonmbsfnSubframeconfigR14
	PdschConfigcommonR14            PdschConfigcommon
	Systeminformationblocktype13R14 *Systeminformationblocktype13R9
	CellaccessrelatedinfolistR14    *[]CellaccessrelatedinfoR14 `lb:1,ub:maxPLMN1R14`
	Noncriticalextension            *Systeminformationblocktype1MbmsR14Noncriticalextension
}
