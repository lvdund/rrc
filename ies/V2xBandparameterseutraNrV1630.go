package ies

import "rrc/utils"

// V2X-BandParametersEUTRA-NR-v1630 ::= CHOICE
type V2xBandparameterseutraNrV1630 interface {
	isV2xBandparameterseutraNrV1630()
}

type V2xBandparameterseutraNrV1630Eutra struct {
	Value struct{}
}

func (*V2xBandparameterseutraNrV1630Eutra) isV2xBandparameterseutraNrV1630() {}

type V2xBandparameterseutraNrV1630Nr struct {
	Value interface{}
}

func (*V2xBandparameterseutraNrV1630Nr) isV2xBandparameterseutraNrV1630() {}
