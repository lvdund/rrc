package ies

import "rrc/utils"

// PhysicalConfigDedicatedSTTI-r15-setup ::= SEQUENCE
type PhysicalconfigdedicatedsttiR15Setup struct {
	AntennainfodedicatedsttiR15        *AntennainfodedicatedsttiR15
	AntennainfoulSttiR15               *AntennainfoulSttiR15
	PucchConfigdedicatedV1530          *PucchConfigdedicatedV1530
	SchedulingrequestconfigV1530       *SchedulingrequestconfigV1530
	UplinkpowercontroldedicatedsttiR15 *UplinkpowercontroldedicatedsttiR15
	CqiReportconfigR15                 *CqiReportconfigR15
	CsiRsConfigR15                     *CsiRsConfigR15
	CsiRsConfignzptoreleaselistR15     *CsiRsConfignzptoreleaselistR15
	CsiRsConfignzptoaddmodlistR15      *CsiRsConfignzptoaddmodlistR15
	CsiRsConfigzptoreleaselistR15      *CsiRsConfigzptoreleaselistR11
	CsiRsConfigzptoaddmodlistR11       *CsiRsConfigzptoaddmodlistR11
	CsiRsConfigzpAplistR15             *CsiRsConfigzpAplistR14
	EimtaMainconfigR12                 *EimtaMainconfigR12
	EimtaMainconfigservcellR15         *EimtaMainconfigservcellR12
	SemiopenloopsttiR15                utils.BOOLEAN
	SlotorsubslotpdschConfigR15        *SlotorsubslotpdschConfigR15
	SlotorsubslotpuschConfigR15        *SlotorsubslotpuschConfigR15
	SpdcchConfigR15                    *SpdcchConfigR15
	SpucchConfigR15                    *SpucchConfigR15
	SrsDci7TriggeringconfigR15         utils.BOOLEAN
	ShortprocessingtimeR15             utils.BOOLEAN
	ShortttiR15                        *ShortttiR15
}
