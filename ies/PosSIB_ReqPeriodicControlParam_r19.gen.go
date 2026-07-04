// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PosSIB-ReqPeriodicControlParam-r19 (line 311).

var posSIBReqPeriodicControlParamR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "deliveryStatusReq-r19"},
	},
}

const (
	PosSIB_ReqPeriodicControlParam_r19_DeliveryStatusReq_r19_Start = 0
	PosSIB_ReqPeriodicControlParam_r19_DeliveryStatusReq_r19_Stop  = 1
)

var posSIBReqPeriodicControlParamR19DeliveryStatusReqR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSIB_ReqPeriodicControlParam_r19_DeliveryStatusReq_r19_Start, PosSIB_ReqPeriodicControlParam_r19_DeliveryStatusReq_r19_Stop},
}

type PosSIB_ReqPeriodicControlParam_r19 struct {
	DeliveryStatusReq_r19 int64
}

func (ie *PosSIB_ReqPeriodicControlParam_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(posSIBReqPeriodicControlParamR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. deliveryStatusReq-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.DeliveryStatusReq_r19, posSIBReqPeriodicControlParamR19DeliveryStatusReqR19Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PosSIB_ReqPeriodicControlParam_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(posSIBReqPeriodicControlParamR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. deliveryStatusReq-r19: enumerated
	{
		v0, err := d.DecodeEnumerated(posSIBReqPeriodicControlParamR19DeliveryStatusReqR19Constraints)
		if err != nil {
			return err
		}
		ie.DeliveryStatusReq_r19 = v0
	}

	return nil
}
