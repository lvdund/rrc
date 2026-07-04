// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB-TypeInfo (line 15092).

var sIBTypeInfoConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "type"},
		{Name: "valueTag", Optional: true},
		{Name: "areaScope", Optional: true},
	},
}

const (
	SIB_TypeInfo_Type_SibType2        = 0
	SIB_TypeInfo_Type_SibType3        = 1
	SIB_TypeInfo_Type_SibType4        = 2
	SIB_TypeInfo_Type_SibType5        = 3
	SIB_TypeInfo_Type_SibType6        = 4
	SIB_TypeInfo_Type_SibType7        = 5
	SIB_TypeInfo_Type_SibType8        = 6
	SIB_TypeInfo_Type_SibType9        = 7
	SIB_TypeInfo_Type_SibType10_v1610 = 8
	SIB_TypeInfo_Type_SibType11_v1610 = 9
	SIB_TypeInfo_Type_SibType12_v1610 = 10
	SIB_TypeInfo_Type_SibType13_v1610 = 11
	SIB_TypeInfo_Type_SibType14_v1610 = 12
	SIB_TypeInfo_Type_Spare3          = 13
	SIB_TypeInfo_Type_Spare2          = 14
	SIB_TypeInfo_Type_Spare1          = 15
)

var sIBTypeInfoTypeConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{SIB_TypeInfo_Type_SibType2, SIB_TypeInfo_Type_SibType3, SIB_TypeInfo_Type_SibType4, SIB_TypeInfo_Type_SibType5, SIB_TypeInfo_Type_SibType6, SIB_TypeInfo_Type_SibType7, SIB_TypeInfo_Type_SibType8, SIB_TypeInfo_Type_SibType9, SIB_TypeInfo_Type_SibType10_v1610, SIB_TypeInfo_Type_SibType11_v1610, SIB_TypeInfo_Type_SibType12_v1610, SIB_TypeInfo_Type_SibType13_v1610, SIB_TypeInfo_Type_SibType14_v1610, SIB_TypeInfo_Type_Spare3, SIB_TypeInfo_Type_Spare2, SIB_TypeInfo_Type_Spare1},
}

var sIBTypeInfoValueTagConstraints = per.Constrained(0, 31)

const (
	SIB_TypeInfo_AreaScope_True = 0
)

var sIBTypeInfoAreaScopeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB_TypeInfo_AreaScope_True},
}

type SIB_TypeInfo struct {
	Type      int64
	ValueTag  *int64
	AreaScope *int64
}

func (ie *SIB_TypeInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIBTypeInfoConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ValueTag != nil, ie.AreaScope != nil}); err != nil {
		return err
	}

	// 2. type: enumerated
	{
		if err := e.EncodeEnumerated(ie.Type, sIBTypeInfoTypeConstraints); err != nil {
			return err
		}
	}

	// 3. valueTag: integer
	{
		if ie.ValueTag != nil {
			if err := e.EncodeInteger(*ie.ValueTag, sIBTypeInfoValueTagConstraints); err != nil {
				return err
			}
		}
	}

	// 4. areaScope: enumerated
	{
		if ie.AreaScope != nil {
			if err := e.EncodeEnumerated(*ie.AreaScope, sIBTypeInfoAreaScopeConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SIB_TypeInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIBTypeInfoConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. type: enumerated
	{
		v0, err := d.DecodeEnumerated(sIBTypeInfoTypeConstraints)
		if err != nil {
			return err
		}
		ie.Type = v0
	}

	// 3. valueTag: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(sIBTypeInfoValueTagConstraints)
			if err != nil {
				return err
			}
			ie.ValueTag = &v
		}
	}

	// 4. areaScope: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sIBTypeInfoAreaScopeConstraints)
			if err != nil {
				return err
			}
			ie.AreaScope = &idx
		}
	}

	return nil
}
