package ies

import "rrc/utils"

// Phy-ParametersCommon-dft-S-OFDM-WaveformUL-IAB-r16 ::= ENUMERATED
type PhyParameterscommonDftSOfdmWaveformulIabR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonDftSOfdmWaveformulIabR16EnumeratedNothing = iota
	PhyParameterscommonDftSOfdmWaveformulIabR16EnumeratedSupported
)
