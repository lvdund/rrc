package ies

// PDCCH-ConfigCommon ::= SEQUENCE
// Extensible
type PdcchConfigcommon struct {
	Controlresourcesetzero                *Controlresourcesetzero
	Commoncontrolresourceset              *Controlresourceset
	Searchspacezero                       *Searchspacezero
	Commonsearchspacelist                 *[]Searchspace `lb:1,ub:4`
	Searchspacesib1                       *Searchspaceid
	Searchspaceothersysteminformation     *Searchspaceid
	Pagingsearchspace                     *Searchspaceid
	RaSearchspace                         *Searchspaceid
	FirstpdcchMonitoringoccasionofpo      *PdcchConfigcommonFirstpdcchMonitoringoccasionofpo      `ext`
	CommonsearchspacelistextR16           *[]SearchspaceextR16                                    `lb:1,ub:4,ext`
	SdtSearchspaceR17                     *PdcchConfigcommonSdtSearchspaceR17                     `ext`
	SearchspacemcchR17                    *Searchspaceid                                          `ext`
	SearchspacemtchR17                    *Searchspaceid                                          `ext`
	Commonsearchspacelistext2R17          *[]SearchspaceextV1700                                  `lb:1,ub:4,ext`
	FirstpdcchMonitoringoccasionofpoV1710 *PdcchConfigcommonFirstpdcchMonitoringoccasionofpoV1710 `ext`
	PeiConfigbwpR17                       *PdcchConfigcommonPeiConfigbwpR17                       `ext`
	FollowunifiedtciStateV1720            *PdcchConfigcommonFollowunifiedtciStateV1720            `ext`
}
