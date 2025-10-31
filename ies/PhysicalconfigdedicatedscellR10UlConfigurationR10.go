package ies

// PhysicalConfigDedicatedSCell-r10-ul-Configuration-r10 ::= SEQUENCE
type PhysicalconfigdedicatedscellR10UlConfigurationR10 struct {
	AntennainfoulR10                        *AntennainfoulR10
	PuschConfigdedicatedscellR10            *PuschConfigdedicatedscellR10
	UplinkpowercontroldedicatedscellR10     *UplinkpowercontroldedicatedscellR10
	CqiReportconfigscellR10                 *CqiReportconfigscellR10
	SoundingrsUlConfigdedicatedR10          *SoundingrsUlConfigdedicated
	SoundingrsUlConfigdedicatedV1020        *SoundingrsUlConfigdedicatedV1020
	SoundingrsUlConfigdedicatedaperiodicR10 *SoundingrsUlConfigdedicatedaperiodicR10
}
