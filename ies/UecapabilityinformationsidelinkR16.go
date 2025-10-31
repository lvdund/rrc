package ies

import "rrc/utils"

// UECapabilityInformationSidelink-r16-IEs ::= SEQUENCE
type UecapabilityinformationsidelinkR16 struct {
	AccessstratumreleasesidelinkR16           AccessstratumreleasesidelinkR16
	PdcpParameterssidelinkR16                 *PdcpParameterssidelinkR16
	RlcParameterssidelinkR16                  *RlcParameterssidelinkR16
	SupportedbandcombinationlistsidelinknrR16 *BandcombinationlistsidelinknrR16
	SupportedbandlistsidelinkR16              *[]Bandsidelinkpc5R16 `lb:1,ub:maxBands`
	AppliedfreqbandlistfilterR16              *Freqbandlist
	Latenoncriticalextension                  *utils.OCTETSTRING
	Noncriticalextension                      *UecapabilityinformationsidelinkV1700
}
