package ies

import "rrc/utils"

// LogMeasInfo-r10 ::= SEQUENCE
// Extensible
type LogmeasinfoR10 struct {
	LocationinfoR10         *LocationinfoR10
	RelativetimestampR10    utils.INTEGER `lb:0,ub:7200`
	ServcellidentityR10     Cellglobalideutra
	MeasresultservcellR10   LogmeasinfoR10MeasresultservcellR10
	MeasresultneighcellsR10 *LogmeasinfoR10MeasresultneighcellsR10
}
