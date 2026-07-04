// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CarrierFreqIoT-r19 (line 4800).

var carrierFreqIoTR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierFreq-r19"},
		{Name: "carrierFreqOffset-r19", Optional: true},
	},
}

const (
	CarrierFreqIoT_r19_CarrierFreqOffset_r19_V_10    = 0
	CarrierFreqIoT_r19_CarrierFreqOffset_r19_V_9     = 1
	CarrierFreqIoT_r19_CarrierFreqOffset_r19_V_8     = 2
	CarrierFreqIoT_r19_CarrierFreqOffset_r19_V_7     = 3
	CarrierFreqIoT_r19_CarrierFreqOffset_r19_V_6     = 4
	CarrierFreqIoT_r19_CarrierFreqOffset_r19_V_5     = 5
	CarrierFreqIoT_r19_CarrierFreqOffset_r19_V_4     = 6
	CarrierFreqIoT_r19_CarrierFreqOffset_r19_V_3     = 7
	CarrierFreqIoT_r19_CarrierFreqOffset_r19_V_2     = 8
	CarrierFreqIoT_r19_CarrierFreqOffset_r19_V_1     = 9
	CarrierFreqIoT_r19_CarrierFreqOffset_r19_V_0dot5 = 10
	CarrierFreqIoT_r19_CarrierFreqOffset_r19_v0      = 11
	CarrierFreqIoT_r19_CarrierFreqOffset_r19_v1      = 12
	CarrierFreqIoT_r19_CarrierFreqOffset_r19_v2      = 13
	CarrierFreqIoT_r19_CarrierFreqOffset_r19_v3      = 14
	CarrierFreqIoT_r19_CarrierFreqOffset_r19_v4      = 15
	CarrierFreqIoT_r19_CarrierFreqOffset_r19_v5      = 16
	CarrierFreqIoT_r19_CarrierFreqOffset_r19_v6      = 17
	CarrierFreqIoT_r19_CarrierFreqOffset_r19_v7      = 18
	CarrierFreqIoT_r19_CarrierFreqOffset_r19_v8      = 19
	CarrierFreqIoT_r19_CarrierFreqOffset_r19_v9      = 20
)

var carrierFreqIoTR19CarrierFreqOffsetR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CarrierFreqIoT_r19_CarrierFreqOffset_r19_V_10, CarrierFreqIoT_r19_CarrierFreqOffset_r19_V_9, CarrierFreqIoT_r19_CarrierFreqOffset_r19_V_8, CarrierFreqIoT_r19_CarrierFreqOffset_r19_V_7, CarrierFreqIoT_r19_CarrierFreqOffset_r19_V_6, CarrierFreqIoT_r19_CarrierFreqOffset_r19_V_5, CarrierFreqIoT_r19_CarrierFreqOffset_r19_V_4, CarrierFreqIoT_r19_CarrierFreqOffset_r19_V_3, CarrierFreqIoT_r19_CarrierFreqOffset_r19_V_2, CarrierFreqIoT_r19_CarrierFreqOffset_r19_V_1, CarrierFreqIoT_r19_CarrierFreqOffset_r19_V_0dot5, CarrierFreqIoT_r19_CarrierFreqOffset_r19_v0, CarrierFreqIoT_r19_CarrierFreqOffset_r19_v1, CarrierFreqIoT_r19_CarrierFreqOffset_r19_v2, CarrierFreqIoT_r19_CarrierFreqOffset_r19_v3, CarrierFreqIoT_r19_CarrierFreqOffset_r19_v4, CarrierFreqIoT_r19_CarrierFreqOffset_r19_v5, CarrierFreqIoT_r19_CarrierFreqOffset_r19_v6, CarrierFreqIoT_r19_CarrierFreqOffset_r19_v7, CarrierFreqIoT_r19_CarrierFreqOffset_r19_v8, CarrierFreqIoT_r19_CarrierFreqOffset_r19_v9},
}

type CarrierFreqIoT_r19 struct {
	CarrierFreq_r19       ARFCN_ValueEUTRA
	CarrierFreqOffset_r19 *int64
}

func (ie *CarrierFreqIoT_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(carrierFreqIoTR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CarrierFreqOffset_r19 != nil}); err != nil {
		return err
	}

	// 2. carrierFreq-r19: ref
	{
		if err := ie.CarrierFreq_r19.Encode(e); err != nil {
			return err
		}
	}

	// 3. carrierFreqOffset-r19: enumerated
	{
		if ie.CarrierFreqOffset_r19 != nil {
			if err := e.EncodeEnumerated(*ie.CarrierFreqOffset_r19, carrierFreqIoTR19CarrierFreqOffsetR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CarrierFreqIoT_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(carrierFreqIoTR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. carrierFreq-r19: ref
	{
		if err := ie.CarrierFreq_r19.Decode(d); err != nil {
			return err
		}
	}

	// 3. carrierFreqOffset-r19: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(carrierFreqIoTR19CarrierFreqOffsetR19Constraints)
			if err != nil {
				return err
			}
			ie.CarrierFreqOffset_r19 = &idx
		}
	}

	return nil
}
