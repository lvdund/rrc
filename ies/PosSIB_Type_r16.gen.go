// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PosSIB-Type-r16 (line 4902).

var posSIBTypeR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "encrypted-r16", Optional: true},
		{Name: "gnss-id-r16", Optional: true},
		{Name: "sbas-id-r16", Optional: true},
		{Name: "posSibType-r16"},
		{Name: "areaScope-r16", Optional: true},
	},
}

const (
	PosSIB_Type_r16_Encrypted_r16_True = 0
)

var posSIBTypeR16EncryptedR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSIB_Type_r16_Encrypted_r16_True},
}

const (
	PosSIB_Type_r16_PosSibType_r16_PosSibType1_1  = 0
	PosSIB_Type_r16_PosSibType_r16_PosSibType1_2  = 1
	PosSIB_Type_r16_PosSibType_r16_PosSibType1_3  = 2
	PosSIB_Type_r16_PosSibType_r16_PosSibType1_4  = 3
	PosSIB_Type_r16_PosSibType_r16_PosSibType1_5  = 4
	PosSIB_Type_r16_PosSibType_r16_PosSibType1_6  = 5
	PosSIB_Type_r16_PosSibType_r16_PosSibType1_7  = 6
	PosSIB_Type_r16_PosSibType_r16_PosSibType1_8  = 7
	PosSIB_Type_r16_PosSibType_r16_PosSibType2_1  = 8
	PosSIB_Type_r16_PosSibType_r16_PosSibType2_2  = 9
	PosSIB_Type_r16_PosSibType_r16_PosSibType2_3  = 10
	PosSIB_Type_r16_PosSibType_r16_PosSibType2_4  = 11
	PosSIB_Type_r16_PosSibType_r16_PosSibType2_5  = 12
	PosSIB_Type_r16_PosSibType_r16_PosSibType2_6  = 13
	PosSIB_Type_r16_PosSibType_r16_PosSibType2_7  = 14
	PosSIB_Type_r16_PosSibType_r16_PosSibType2_8  = 15
	PosSIB_Type_r16_PosSibType_r16_PosSibType2_9  = 16
	PosSIB_Type_r16_PosSibType_r16_PosSibType2_10 = 17
	PosSIB_Type_r16_PosSibType_r16_PosSibType2_11 = 18
	PosSIB_Type_r16_PosSibType_r16_PosSibType2_12 = 19
	PosSIB_Type_r16_PosSibType_r16_PosSibType2_13 = 20
	PosSIB_Type_r16_PosSibType_r16_PosSibType2_14 = 21
	PosSIB_Type_r16_PosSibType_r16_PosSibType2_15 = 22
	PosSIB_Type_r16_PosSibType_r16_PosSibType2_16 = 23
	PosSIB_Type_r16_PosSibType_r16_PosSibType2_17 = 24
	PosSIB_Type_r16_PosSibType_r16_PosSibType2_18 = 25
	PosSIB_Type_r16_PosSibType_r16_PosSibType2_19 = 26
	PosSIB_Type_r16_PosSibType_r16_PosSibType2_20 = 27
	PosSIB_Type_r16_PosSibType_r16_PosSibType2_21 = 28
	PosSIB_Type_r16_PosSibType_r16_PosSibType2_22 = 29
	PosSIB_Type_r16_PosSibType_r16_PosSibType2_23 = 30
	PosSIB_Type_r16_PosSibType_r16_PosSibType3_1  = 31
	PosSIB_Type_r16_PosSibType_r16_PosSibType4_1  = 32
	PosSIB_Type_r16_PosSibType_r16_PosSibType5_1  = 33
	PosSIB_Type_r16_PosSibType_r16_PosSibType6_1  = 34
	PosSIB_Type_r16_PosSibType_r16_PosSibType6_2  = 35
	PosSIB_Type_r16_PosSibType_r16_PosSibType6_3  = 36
)

var posSIBTypeR16PosSibTypeR16Constraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{PosSIB_Type_r16_PosSibType_r16_PosSibType1_1, PosSIB_Type_r16_PosSibType_r16_PosSibType1_2, PosSIB_Type_r16_PosSibType_r16_PosSibType1_3, PosSIB_Type_r16_PosSibType_r16_PosSibType1_4, PosSIB_Type_r16_PosSibType_r16_PosSibType1_5, PosSIB_Type_r16_PosSibType_r16_PosSibType1_6, PosSIB_Type_r16_PosSibType_r16_PosSibType1_7, PosSIB_Type_r16_PosSibType_r16_PosSibType1_8, PosSIB_Type_r16_PosSibType_r16_PosSibType2_1, PosSIB_Type_r16_PosSibType_r16_PosSibType2_2, PosSIB_Type_r16_PosSibType_r16_PosSibType2_3, PosSIB_Type_r16_PosSibType_r16_PosSibType2_4, PosSIB_Type_r16_PosSibType_r16_PosSibType2_5, PosSIB_Type_r16_PosSibType_r16_PosSibType2_6, PosSIB_Type_r16_PosSibType_r16_PosSibType2_7, PosSIB_Type_r16_PosSibType_r16_PosSibType2_8, PosSIB_Type_r16_PosSibType_r16_PosSibType2_9, PosSIB_Type_r16_PosSibType_r16_PosSibType2_10, PosSIB_Type_r16_PosSibType_r16_PosSibType2_11, PosSIB_Type_r16_PosSibType_r16_PosSibType2_12, PosSIB_Type_r16_PosSibType_r16_PosSibType2_13, PosSIB_Type_r16_PosSibType_r16_PosSibType2_14, PosSIB_Type_r16_PosSibType_r16_PosSibType2_15, PosSIB_Type_r16_PosSibType_r16_PosSibType2_16, PosSIB_Type_r16_PosSibType_r16_PosSibType2_17, PosSIB_Type_r16_PosSibType_r16_PosSibType2_18, PosSIB_Type_r16_PosSibType_r16_PosSibType2_19, PosSIB_Type_r16_PosSibType_r16_PosSibType2_20, PosSIB_Type_r16_PosSibType_r16_PosSibType2_21, PosSIB_Type_r16_PosSibType_r16_PosSibType2_22, PosSIB_Type_r16_PosSibType_r16_PosSibType2_23, PosSIB_Type_r16_PosSibType_r16_PosSibType3_1, PosSIB_Type_r16_PosSibType_r16_PosSibType4_1, PosSIB_Type_r16_PosSibType_r16_PosSibType5_1, PosSIB_Type_r16_PosSibType_r16_PosSibType6_1, PosSIB_Type_r16_PosSibType_r16_PosSibType6_2, PosSIB_Type_r16_PosSibType_r16_PosSibType6_3},
}

const (
	PosSIB_Type_r16_AreaScope_r16_True = 0
)

var posSIBTypeR16AreaScopeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PosSIB_Type_r16_AreaScope_r16_True},
}

type PosSIB_Type_r16 struct {
	Encrypted_r16  *int64
	Gnss_Id_r16    *GNSS_ID_r16
	Sbas_Id_r16    *SBAS_ID_r16
	PosSibType_r16 int64
	AreaScope_r16  *int64
}

func (ie *PosSIB_Type_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(posSIBTypeR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Encrypted_r16 != nil, ie.Gnss_Id_r16 != nil, ie.Sbas_Id_r16 != nil, ie.AreaScope_r16 != nil}); err != nil {
		return err
	}

	// 2. encrypted-r16: enumerated
	{
		if ie.Encrypted_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Encrypted_r16, posSIBTypeR16EncryptedR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. gnss-id-r16: ref
	{
		if ie.Gnss_Id_r16 != nil {
			if err := ie.Gnss_Id_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sbas-id-r16: ref
	{
		if ie.Sbas_Id_r16 != nil {
			if err := ie.Sbas_Id_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. posSibType-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.PosSibType_r16, posSIBTypeR16PosSibTypeR16Constraints); err != nil {
			return err
		}
	}

	// 6. areaScope-r16: enumerated
	{
		if ie.AreaScope_r16 != nil {
			if err := e.EncodeEnumerated(*ie.AreaScope_r16, posSIBTypeR16AreaScopeR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PosSIB_Type_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(posSIBTypeR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. encrypted-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(posSIBTypeR16EncryptedR16Constraints)
			if err != nil {
				return err
			}
			ie.Encrypted_r16 = &idx
		}
	}

	// 3. gnss-id-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Gnss_Id_r16 = new(GNSS_ID_r16)
			if err := ie.Gnss_Id_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sbas-id-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sbas_Id_r16 = new(SBAS_ID_r16)
			if err := ie.Sbas_Id_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. posSibType-r16: enumerated
	{
		v3, err := d.DecodeEnumerated(posSIBTypeR16PosSibTypeR16Constraints)
		if err != nil {
			return err
		}
		ie.PosSibType_r16 = v3
	}

	// 6. areaScope-r16: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(posSIBTypeR16AreaScopeR16Constraints)
			if err != nil {
				return err
			}
			ie.AreaScope_r16 = &idx
		}
	}

	return nil
}
