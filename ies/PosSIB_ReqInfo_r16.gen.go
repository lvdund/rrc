// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PosSIB-ReqInfo-r16 (line 283).

var posSIBReqInfoR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "gnss-id-r16", Optional: true},
		{Name: "sbas-id-r16", Optional: true},
		{Name: "posSibType-r16"},
	},
}

const (
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType1_1  = 0
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType1_2  = 1
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType1_3  = 2
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType1_4  = 3
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType1_5  = 4
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType1_6  = 5
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType1_7  = 6
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType1_8  = 7
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_1  = 8
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_2  = 9
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_3  = 10
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_4  = 11
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_5  = 12
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_6  = 13
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_7  = 14
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_8  = 15
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_9  = 16
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_10 = 17
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_11 = 18
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_12 = 19
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_13 = 20
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_14 = 21
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_15 = 22
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_16 = 23
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_17 = 24
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_18 = 25
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_19 = 26
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_20 = 27
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_21 = 28
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_22 = 29
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_23 = 30
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType3_1  = 31
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType4_1  = 32
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType5_1  = 33
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType6_1  = 34
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType6_2  = 35
	PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType6_3  = 36
)

var posSIBReqInfoR16PosSibTypeR16Constraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType1_1, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType1_2, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType1_3, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType1_4, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType1_5, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType1_6, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType1_7, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType1_8, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_1, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_2, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_3, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_4, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_5, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_6, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_7, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_8, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_9, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_10, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_11, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_12, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_13, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_14, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_15, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_16, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_17, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_18, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_19, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_20, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_21, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_22, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType2_23, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType3_1, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType4_1, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType5_1, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType6_1, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType6_2, PosSIB_ReqInfo_r16_PosSibType_r16_PosSibType6_3},
	ExtValues:  []int64{37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55},
}

type PosSIB_ReqInfo_r16 struct {
	Gnss_Id_r16    *GNSS_ID_r16
	Sbas_Id_r16    *SBAS_ID_r16
	PosSibType_r16 int64
}

func (ie *PosSIB_ReqInfo_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(posSIBReqInfoR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Gnss_Id_r16 != nil, ie.Sbas_Id_r16 != nil}); err != nil {
		return err
	}

	// 2. gnss-id-r16: ref
	{
		if ie.Gnss_Id_r16 != nil {
			if err := ie.Gnss_Id_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. sbas-id-r16: ref
	{
		if ie.Sbas_Id_r16 != nil {
			if err := ie.Sbas_Id_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. posSibType-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.PosSibType_r16, posSIBReqInfoR16PosSibTypeR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PosSIB_ReqInfo_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(posSIBReqInfoR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. gnss-id-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Gnss_Id_r16 = new(GNSS_ID_r16)
			if err := ie.Gnss_Id_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. sbas-id-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sbas_Id_r16 = new(SBAS_ID_r16)
			if err := ie.Sbas_Id_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. posSibType-r16: enumerated
	{
		v2, err := d.DecodeEnumerated(posSIBReqInfoR16PosSibTypeR16Constraints)
		if err != nil {
			return err
		}
		ie.PosSibType_r16 = v2
	}

	return nil
}
