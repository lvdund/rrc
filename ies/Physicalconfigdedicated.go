package ies

// PhysicalConfigDedicated ::= SEQUENCE
// Extensible
type Physicalconfigdedicated struct {
	PdschConfigdedicated        *PdschConfigdedicated
	PucchConfigdedicated        *PucchConfigdedicated
	PuschConfigdedicated        *PuschConfigdedicated
	Uplinkpowercontroldedicated *Uplinkpowercontroldedicated
	TpcPdcchConfigpucch         *TpcPdcchConfig
	TpcPdcchConfigpusch         *TpcPdcchConfig
	CqiReportconfig             *CqiReportconfig
	SoundingrsUlConfigdedicated *SoundingrsUlConfigdedicated
	Antennainfo                 *PhysicalconfigdedicatedAntennainfo
	Schedulingrequestconfig     *Schedulingrequestconfig
}
