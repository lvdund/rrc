package ies

import "rrc/utils"

// PDCCH-Config ::= SEQUENCE
// Extensible
type PdcchConfig struct {
	Controlresourcesettoaddmodlist             *[]Controlresourceset   `lb:1,ub:3`
	Controlresourcesettoreleaselist            *[]Controlresourcesetid `lb:1,ub:3`
	Searchspacestoaddmodlist                   *[]Searchspace          `lb:1,ub:10`
	Searchspacestoreleaselist                  *[]Searchspaceid        `lb:1,ub:10`
	Downlinkpreemption                         *utils.Setuprelease[Downlinkpreemption]
	TpcPusch                                   *utils.Setuprelease[PuschTpcCommandconfig]
	TpcPucch                                   *utils.Setuprelease[PucchTpcCommandconfig]
	TpcSrs                                     *utils.Setuprelease[SrsTpcCommandconfig]
	ControlresourcesettoaddmodlistsizeextV1610 *[]Controlresourceset                       `lb:1,ub:2,ext`
	ControlresourcesettoreleaselistsizeextR16  *[]ControlresourcesetidR16                  `lb:1,ub:5,ext`
	SearchspacestoaddmodlistextR16             *[]SearchspaceextR16                        `lb:1,ub:10,ext`
	UplinkcancellationR16                      *utils.Setuprelease[UplinkcancellationR16]  `ext`
	MonitoringcapabilityconfigR16              *PdcchConfigMonitoringcapabilityconfigR16   `ext`
	SearchspaceswitchconfigR16                 *SearchspaceswitchconfigR16                 `ext`
	SearchspacestoaddmodlistextV1700           *[]SearchspaceextV1700                      `lb:1,ub:10,ext`
	MonitoringcapabilityconfigV1710            *PdcchConfigMonitoringcapabilityconfigV1710 `ext`
	SearchspaceswitchconfigR17                 *SearchspaceswitchconfigR17                 `ext`
	PdcchSkippingdurationlistR17               *[]ScsSpecificdurationR17                   `lb:1,ub:3,ext`
}
