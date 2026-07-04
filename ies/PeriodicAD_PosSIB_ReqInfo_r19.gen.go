// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PeriodicAD-PosSIB-ReqInfo-r19 (line 300).

var periodicADPosSIBReqInfoR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "gnss-id-r19", Optional: true},
		{Name: "sbas-id-r19", Optional: true},
		{Name: "periodicAD-posSibType-r19"},
		{Name: "posSIB-ReqPeriodicControlParam-r19"},
	},
}

const (
	PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType1_10  = 0
	PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType1_12  = 1
	PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_12  = 2
	PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_13  = 3
	PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_14  = 4
	PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_15  = 5
	PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_16  = 6
	PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_17  = 7
	PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_17a = 8
	PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_18  = 9
	PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_18a = 10
	PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_19  = 11
	PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_20  = 12
	PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_20a = 13
	PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_21  = 14
	PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_22  = 15
	PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_23  = 16
	PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType6_7   = 17
)

var periodicADPosSIBReqInfoR19PeriodicADPosSibTypeR19Constraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType1_10, PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType1_12, PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_12, PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_13, PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_14, PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_15, PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_16, PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_17, PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_17a, PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_18, PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_18a, PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_19, PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_20, PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_20a, PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_21, PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_22, PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType2_23, PeriodicAD_PosSIB_ReqInfo_r19_PeriodicAD_PosSibType_r19_PosSibType6_7},
}

type PeriodicAD_PosSIB_ReqInfo_r19 struct {
	Gnss_Id_r19                        *GNSS_ID_r16
	Sbas_Id_r19                        *SBAS_ID_r16
	PeriodicAD_PosSibType_r19          int64
	PosSIB_ReqPeriodicControlParam_r19 PosSIB_ReqPeriodicControlParam_r19
}

func (ie *PeriodicAD_PosSIB_ReqInfo_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(periodicADPosSIBReqInfoR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Gnss_Id_r19 != nil, ie.Sbas_Id_r19 != nil}); err != nil {
		return err
	}

	// 3. gnss-id-r19: ref
	{
		if ie.Gnss_Id_r19 != nil {
			if err := ie.Gnss_Id_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sbas-id-r19: ref
	{
		if ie.Sbas_Id_r19 != nil {
			if err := ie.Sbas_Id_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. periodicAD-posSibType-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.PeriodicAD_PosSibType_r19, periodicADPosSIBReqInfoR19PeriodicADPosSibTypeR19Constraints); err != nil {
			return err
		}
	}

	// 6. posSIB-ReqPeriodicControlParam-r19: ref
	{
		if err := ie.PosSIB_ReqPeriodicControlParam_r19.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PeriodicAD_PosSIB_ReqInfo_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(periodicADPosSIBReqInfoR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. gnss-id-r19: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Gnss_Id_r19 = new(GNSS_ID_r16)
			if err := ie.Gnss_Id_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sbas-id-r19: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sbas_Id_r19 = new(SBAS_ID_r16)
			if err := ie.Sbas_Id_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. periodicAD-posSibType-r19: enumerated
	{
		v2, err := d.DecodeEnumerated(periodicADPosSIBReqInfoR19PeriodicADPosSibTypeR19Constraints)
		if err != nil {
			return err
		}
		ie.PeriodicAD_PosSibType_r19 = v2
	}

	// 6. posSIB-ReqPeriodicControlParam-r19: ref
	{
		if err := ie.PosSIB_ReqPeriodicControlParam_r19.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
