package ies

import "rrc/utils"

// MeasAndMobParametersCommon-maxNumberCLI-RSSI-r16 ::= ENUMERATED
type MeasandmobparameterscommonMaxnumbercliRssiR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasandmobparameterscommonMaxnumbercliRssiR16EnumeratedNothing = iota
	MeasandmobparameterscommonMaxnumbercliRssiR16EnumeratedN8
	MeasandmobparameterscommonMaxnumbercliRssiR16EnumeratedN16
	MeasandmobparameterscommonMaxnumbercliRssiR16EnumeratedN32
	MeasandmobparameterscommonMaxnumbercliRssiR16EnumeratedN64
)
