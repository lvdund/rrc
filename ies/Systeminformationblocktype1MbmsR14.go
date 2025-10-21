package ies

import "rrc/utils"

// SystemInformationBlockType1-MBMS-r14 ::= SEQUENCE
type Systeminformationblocktype1MbmsR14 struct {
	CellaccessrelatedinfoR14        interface{}
	FreqbandindicatorR14            FreqbandindicatorR11
	MultibandinfolistR14            *MultibandinfolistR11
	SchedulinginfolistMbmsR14       SchedulinginfolistMbmsR14
	SiWindowlengthR14               utils.ENUMERATED
	SysteminfovaluetagR14           utils.INTEGER
	NonmbsfnSubframeconfigR14       *NonmbsfnSubframeconfigR14
	PdschConfigcommonR14            PdschConfigcommon
	Systeminformationblocktype13R14 *Systeminformationblocktype13R9
	CellaccessrelatedinfolistR14    *interface{}
	Noncriticalextension            *interface{}
}
