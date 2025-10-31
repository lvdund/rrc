package ies

import "rrc/utils"

// SCG-ConfigInfo-r12-IEs ::= SEQUENCE
type ScgConfiginfoR12 struct {
	RadioresourceconfigdedmcgR12 *Radioresourceconfigdedicated
	ScelltoaddmodlistmcgR12      *ScelltoaddmodlistR10
	MeasgapconfigR12             *Measgapconfig
	PowercoordinationinfoR12     *PowercoordinationinfoR12
	ScgRadioconfigR12            *ScgConfigpartscgR12
	EutraCapabilityinfoR12       *utils.OCTETSTRING
	ScgConfigrestrictinfoR12     *ScgConfigrestrictinfoR12
	MbmsinterestindicationR12    *utils.OCTETSTRING
	MeasresultservcelllistscgR12 *MeasresultservcelllistscgR12
	DrbToaddmodlistscgR12        *DrbInfolistscgR12
	DrbToreleaselistscgR12       *DrbToreleaselist
	ScelltoaddmodlistscgR12      *ScelltoaddmodlistscgR12
	ScelltoreleaselistscgR12     *ScelltoreleaselistR10
	PMaxR12                      *PMax
	Noncriticalextension         *ScgConfiginfoV1310
}
