package ies

import "rrc/utils"

// SystemInformationBlockType1-NB ::= SEQUENCE
type Systeminformationblocktype1Nb struct {
	HypersfnMsbR13            utils.BITSTRING `lb:8,ub:8`
	CellaccessrelatedinfoR13  Systeminformationblocktype1NbCellaccessrelatedinfoR13
	CellselectioninfoR13      Systeminformationblocktype1NbCellselectioninfoR13
	PMaxR13                   *PMax
	FreqbandindicatorR13      FreqbandindicatorNbR13
	FreqbandinfoR13           *NsPmaxlistNbR13
	MultibandinfolistR13      *MultibandinfolistNbR13
	DownlinkbitmapR13         *DlBitmapNbR13
	EutracontrolregionsizeR13 *Systeminformationblocktype1NbEutracontrolregionsizeR13
	NrsCrsPoweroffsetR13      *Systeminformationblocktype1NbNrsCrsPoweroffsetR13
	SchedulinginfolistR13     SchedulinginfolistNbR13
	SiWindowlengthR13         Systeminformationblocktype1NbSiWindowlengthR13
	SiRadioframeoffsetR13     *utils.INTEGER `lb:0,ub:15`
	SysteminfovaluetaglistR13 *SysteminfovaluetaglistNbR13
	Latenoncriticalextension  *utils.OCTETSTRING
	Noncriticalextension      *Systeminformationblocktype1NbV1350
}
