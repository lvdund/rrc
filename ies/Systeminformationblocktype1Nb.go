package ies

import "rrc/utils"

// SystemInformationBlockType1-NB ::= SEQUENCE
type Systeminformationblocktype1Nb struct {
	HypersfnMsbR13            utils.BITSTRING
	CellaccessrelatedinfoR13  interface{}
	CellselectioninfoR13      interface{}
	PMaxR13                   *PMax
	FreqbandindicatorR13      FreqbandindicatorNbR13
	FreqbandinfoR13           *NsPmaxlistNbR13
	MultibandinfolistR13      *MultibandinfolistNbR13
	DownlinkbitmapR13         *DlBitmapNbR13
	EutracontrolregionsizeR13 *utils.ENUMERATED
	NrsCrsPoweroffsetR13      *utils.ENUMERATED
	SchedulinginfolistR13     SchedulinginfolistNbR13
	SiWindowlengthR13         utils.ENUMERATED
	SiRadioframeoffsetR13     *utils.INTEGER
	SysteminfovaluetaglistR13 *SysteminfovaluetaglistNbR13
	Latenoncriticalextension  *utils.OCTETSTRING
	Noncriticalextension      *Systeminformationblocktype1NbV1350
}
