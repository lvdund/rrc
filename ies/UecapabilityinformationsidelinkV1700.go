package ies

// UECapabilityInformationSidelink-v1700-IEs ::= SEQUENCE
type UecapabilityinformationsidelinkV1700 struct {
	MacParameterssidelinkR17                    *MacParameterssidelinkR17
	SupportedbandcombinationlistsidelinknrV1710 *BandcombinationlistsidelinknrV1710
	Noncriticalextension                        *UecapabilityinformationsidelinkV1700IesNoncriticalextension
}
