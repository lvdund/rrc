// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB-TypeInfo-v1700 (line 15100).

var sIBTypeInfoV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sibType-r17"},
		{Name: "valueTag-r17", Optional: true},
		{Name: "areaScope-r17", Optional: true},
	},
}

var sIB_TypeInfo_v1700SibTypeR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "type1-r17"},
		{Name: "type2-r17"},
	},
}

const (
	SIB_TypeInfo_v1700_SibType_r17_Type1_r17 = 0
	SIB_TypeInfo_v1700_SibType_r17_Type2_r17 = 1
)

var sIBTypeInfoV1700ValueTagR17Constraints = per.Constrained(0, 31)

const (
	SIB_TypeInfo_v1700_AreaScope_r17_True = 0
)

var sIBTypeInfoV1700AreaScopeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB_TypeInfo_v1700_AreaScope_r17_True},
}

const (
	SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType15          = 0
	SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType16          = 1
	SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType17          = 2
	SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType18          = 3
	SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType19          = 4
	SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType20          = 5
	SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType21          = 6
	SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType22_v1800    = 7
	SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType23_v1800    = 8
	SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType24_v1800    = 9
	SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType25_v1800    = 10
	SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType17bis_v1820 = 11
	SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType26_v1900    = 12
	SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType27_v1900    = 13
	SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType28_v1910    = 14
	SIB_TypeInfo_v1700_SibType_r17_Type1_r17_Spare1             = 15
)

var sIBTypeInfoV1700SibTypeR17Type1R17Constraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType15, SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType16, SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType17, SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType18, SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType19, SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType20, SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType21, SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType22_v1800, SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType23_v1800, SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType24_v1800, SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType25_v1800, SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType17bis_v1820, SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType26_v1900, SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType27_v1900, SIB_TypeInfo_v1700_SibType_r17_Type1_r17_SibType28_v1910, SIB_TypeInfo_v1700_SibType_r17_Type1_r17_Spare1},
}

var sIBTypeInfoV1700SibTypeR17Type2R17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "posSibType-r17"},
		{Name: "encrypted-r17", Optional: true},
		{Name: "gnss-id-r17", Optional: true},
		{Name: "sbas-id-r17", Optional: true},
	},
}

const (
	SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType1_9         = 0
	SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType1_10        = 1
	SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType2_24        = 2
	SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType2_25        = 3
	SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType6_4         = 4
	SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType6_5         = 5
	SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType6_6         = 6
	SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType2_17a_v1770 = 7
	SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType2_18a_v1770 = 8
	SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType2_20a_v1770 = 9
	SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType1_11_v1800  = 10
	SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType1_12_v1800  = 11
	SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType2_26_v1800  = 12
	SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType2_27_v1800  = 13
	SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType6_7_v1800   = 14
	SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType7_1_v1800   = 15
)

var sIBTypeInfoV1700SibTypeR17Type2R17PosSibTypeR17Constraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType1_9, SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType1_10, SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType2_24, SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType2_25, SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType6_4, SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType6_5, SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType6_6, SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType2_17a_v1770, SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType2_18a_v1770, SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType2_20a_v1770, SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType1_11_v1800, SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType1_12_v1800, SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType2_26_v1800, SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType2_27_v1800, SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType6_7_v1800, SIB_TypeInfo_v1700_SibType_r17_Type2_r17_PosSibType_r17_PosSibType7_1_v1800},
	ExtValues:  []int64{16, 17, 18},
}

const (
	SIB_TypeInfo_v1700_SibType_r17_Type2_r17_Encrypted_r17_True = 0
)

var sIBTypeInfoV1700SibTypeR17Type2R17EncryptedR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB_TypeInfo_v1700_SibType_r17_Type2_r17_Encrypted_r17_True},
}

type SIB_TypeInfo_v1700 struct {
	SibType_r17 struct {
		Choice    int
		Type1_r17 *int64
		Type2_r17 *struct {
			PosSibType_r17 int64
			Encrypted_r17  *int64
			Gnss_Id_r17    *GNSS_ID_r16
			Sbas_Id_r17    *SBAS_ID_r16
		}
	}
	ValueTag_r17  *int64
	AreaScope_r17 *int64
}

func (ie *SIB_TypeInfo_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIBTypeInfoV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ValueTag_r17 != nil, ie.AreaScope_r17 != nil}); err != nil {
		return err
	}

	// 2. sibType-r17: choice
	{
		choiceEnc := e.NewChoiceEncoder(sIB_TypeInfo_v1700SibTypeR17Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.SibType_r17.Choice), false, nil); err != nil {
			return err
		}
		switch ie.SibType_r17.Choice {
		case SIB_TypeInfo_v1700_SibType_r17_Type1_r17:
			if err := e.EncodeEnumerated((*ie.SibType_r17.Type1_r17), sIBTypeInfoV1700SibTypeR17Type1R17Constraints); err != nil {
				return err
			}
		case SIB_TypeInfo_v1700_SibType_r17_Type2_r17:
			c := (*ie.SibType_r17.Type2_r17)
			sIBTypeInfoV1700SibTypeR17Type2R17Seq := e.NewSequenceEncoder(sIBTypeInfoV1700SibTypeR17Type2R17Constraints)
			if err := sIBTypeInfoV1700SibTypeR17Type2R17Seq.EncodePreamble([]bool{c.Encrypted_r17 != nil, c.Gnss_Id_r17 != nil, c.Sbas_Id_r17 != nil}); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.PosSibType_r17, sIBTypeInfoV1700SibTypeR17Type2R17PosSibTypeR17Constraints); err != nil {
				return err
			}
			if c.Encrypted_r17 != nil {
				if err := e.EncodeEnumerated((*c.Encrypted_r17), sIBTypeInfoV1700SibTypeR17Type2R17EncryptedR17Constraints); err != nil {
					return err
				}
			}
			if c.Gnss_Id_r17 != nil {
				if err := c.Gnss_Id_r17.Encode(e); err != nil {
					return err
				}
			}
			if c.Sbas_Id_r17 != nil {
				if err := c.Sbas_Id_r17.Encode(e); err != nil {
					return err
				}
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.SibType_r17.Choice), Constraint: "index out of range"}
		}
	}

	// 3. valueTag-r17: integer
	{
		if ie.ValueTag_r17 != nil {
			if err := e.EncodeInteger(*ie.ValueTag_r17, sIBTypeInfoV1700ValueTagR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. areaScope-r17: enumerated
	{
		if ie.AreaScope_r17 != nil {
			if err := e.EncodeEnumerated(*ie.AreaScope_r17, sIBTypeInfoV1700AreaScopeR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SIB_TypeInfo_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIBTypeInfoV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sibType-r17: choice
	{
		choiceDec := d.NewChoiceDecoder(sIB_TypeInfo_v1700SibTypeR17Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.SibType_r17.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SIB_TypeInfo_v1700_SibType_r17_Type1_r17:
			ie.SibType_r17.Type1_r17 = new(int64)
			v, err := d.DecodeEnumerated(sIBTypeInfoV1700SibTypeR17Type1R17Constraints)
			if err != nil {
				return err
			}
			(*ie.SibType_r17.Type1_r17) = v
		case SIB_TypeInfo_v1700_SibType_r17_Type2_r17:
			ie.SibType_r17.Type2_r17 = &struct {
				PosSibType_r17 int64
				Encrypted_r17  *int64
				Gnss_Id_r17    *GNSS_ID_r16
				Sbas_Id_r17    *SBAS_ID_r16
			}{}
			c := (*ie.SibType_r17.Type2_r17)
			sIBTypeInfoV1700SibTypeR17Type2R17Seq := d.NewSequenceDecoder(sIBTypeInfoV1700SibTypeR17Type2R17Constraints)
			if err := sIBTypeInfoV1700SibTypeR17Type2R17Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := d.DecodeEnumerated(sIBTypeInfoV1700SibTypeR17Type2R17PosSibTypeR17Constraints)
				if err != nil {
					return err
				}
				c.PosSibType_r17 = v
			}
			if sIBTypeInfoV1700SibTypeR17Type2R17Seq.IsComponentPresent(1) {
				c.Encrypted_r17 = new(int64)
				v, err := d.DecodeEnumerated(sIBTypeInfoV1700SibTypeR17Type2R17EncryptedR17Constraints)
				if err != nil {
					return err
				}
				(*c.Encrypted_r17) = v
			}
			if sIBTypeInfoV1700SibTypeR17Type2R17Seq.IsComponentPresent(2) {
				c.Gnss_Id_r17 = new(GNSS_ID_r16)
				if err := (*c.Gnss_Id_r17).Decode(d); err != nil {
					return err
				}
			}
			if sIBTypeInfoV1700SibTypeR17Type2R17Seq.IsComponentPresent(3) {
				c.Sbas_Id_r17 = new(SBAS_ID_r16)
				if err := (*c.Sbas_Id_r17).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. valueTag-r17: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(sIBTypeInfoV1700ValueTagR17Constraints)
			if err != nil {
				return err
			}
			ie.ValueTag_r17 = &v
		}
	}

	// 4. areaScope-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sIBTypeInfoV1700AreaScopeR17Constraints)
			if err != nil {
				return err
			}
			ie.AreaScope_r17 = &idx
		}
	}

	return nil
}
