package ies

import "rrc/utils"

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
	Antennainfo                 *interface{}
	Schedulingrequestconfig     *Schedulingrequestconfig
}
