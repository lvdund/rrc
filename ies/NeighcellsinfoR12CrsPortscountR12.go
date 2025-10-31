package ies

import "rrc/utils"

// NeighCellsInfo-r12-crs-PortsCount-r12 ::= ENUMERATED
type NeighcellsinfoR12CrsPortscountR12 struct {
	Value utils.ENUMERATED
}

const (
	NeighcellsinfoR12CrsPortscountR12EnumeratedNothing = iota
	NeighcellsinfoR12CrsPortscountR12EnumeratedN1
	NeighcellsinfoR12CrsPortscountR12EnumeratedN2
	NeighcellsinfoR12CrsPortscountR12EnumeratedN4
	NeighcellsinfoR12CrsPortscountR12EnumeratedSpare
)
