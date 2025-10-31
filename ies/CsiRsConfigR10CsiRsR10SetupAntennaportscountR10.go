package ies

import "rrc/utils"

// CSI-RS-Config-r10-csi-RS-r10-setup-antennaPortsCount-r10 ::= ENUMERATED
type CsiRsConfigR10CsiRsR10SetupAntennaportscountR10 struct {
	Value utils.ENUMERATED
}

const (
	CsiRsConfigR10CsiRsR10SetupAntennaportscountR10EnumeratedNothing = iota
	CsiRsConfigR10CsiRsR10SetupAntennaportscountR10EnumeratedAn1
	CsiRsConfigR10CsiRsR10SetupAntennaportscountR10EnumeratedAn2
	CsiRsConfigR10CsiRsR10SetupAntennaportscountR10EnumeratedAn4
	CsiRsConfigR10CsiRsR10SetupAntennaportscountR10EnumeratedAn8
)
