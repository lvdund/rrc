package ies

import "rrc/utils"

// NeighCellsInfo-r12 ::= SEQUENCE
// Extensible
type NeighcellsinfoR12 struct {
	PhyscellidR12           Physcellid
	PBR12                   utils.INTEGER `lb:0,ub:3`
	CrsPortscountR12        NeighcellsinfoR12CrsPortscountR12
	MbsfnSubframeconfigR12  *MbsfnSubframeconfiglist
	PAlistR12               []PA            `lb:1,ub:maxPAPerneighcellR12`
	TransmissionmodelistR12 utils.BITSTRING `lb:8,ub:8`
	ResallocgranularityR12  utils.INTEGER   `lb:0,ub:4`
}
