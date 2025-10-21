package ies

import "rrc/utils"

// V2X-BandParametersEUTRA-NR-r16 ::= CHOICE
type V2xBandparameterseutraNrR16 interface {
	isV2xBandparameterseutraNrR16()
}

type V2xBandparameterseutraNrR16Eutra struct {
	Value interface{}
}

func (*V2xBandparameterseutraNrR16Eutra) isV2xBandparameterseutraNrR16() {}

type V2xBandparameterseutraNrR16Nr struct {
	Value interface{}
}

func (*V2xBandparameterseutraNrR16Nr) isV2xBandparameterseutraNrR16() {}
