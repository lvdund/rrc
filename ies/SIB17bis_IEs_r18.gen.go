// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SIB17bis-IEs-r18 (line 4453).

var sIB17bisIEsR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "trs-ResourceSetConfig-r18", Optional: true},
		{Name: "validityDuration-r18", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
}

var sIB17bisIEsR18TrsResourceSetConfigR18Constraints = per.SizeRange(1, common.MaxNrofTRS_ResourceSets_r17)

const (
	SIB17bis_IEs_r18_ValidityDuration_r18_T1       = 0
	SIB17bis_IEs_r18_ValidityDuration_r18_T2       = 1
	SIB17bis_IEs_r18_ValidityDuration_r18_T4       = 2
	SIB17bis_IEs_r18_ValidityDuration_r18_T8       = 3
	SIB17bis_IEs_r18_ValidityDuration_r18_T16      = 4
	SIB17bis_IEs_r18_ValidityDuration_r18_T32      = 5
	SIB17bis_IEs_r18_ValidityDuration_r18_T64      = 6
	SIB17bis_IEs_r18_ValidityDuration_r18_T128     = 7
	SIB17bis_IEs_r18_ValidityDuration_r18_T256     = 8
	SIB17bis_IEs_r18_ValidityDuration_r18_T512     = 9
	SIB17bis_IEs_r18_ValidityDuration_r18_Infinity = 10
	SIB17bis_IEs_r18_ValidityDuration_r18_Spare5   = 11
	SIB17bis_IEs_r18_ValidityDuration_r18_Spare4   = 12
	SIB17bis_IEs_r18_ValidityDuration_r18_Spare3   = 13
	SIB17bis_IEs_r18_ValidityDuration_r18_Spare2   = 14
	SIB17bis_IEs_r18_ValidityDuration_r18_Spare1   = 15
)

var sIB17bisIEsR18ValidityDurationR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB17bis_IEs_r18_ValidityDuration_r18_T1, SIB17bis_IEs_r18_ValidityDuration_r18_T2, SIB17bis_IEs_r18_ValidityDuration_r18_T4, SIB17bis_IEs_r18_ValidityDuration_r18_T8, SIB17bis_IEs_r18_ValidityDuration_r18_T16, SIB17bis_IEs_r18_ValidityDuration_r18_T32, SIB17bis_IEs_r18_ValidityDuration_r18_T64, SIB17bis_IEs_r18_ValidityDuration_r18_T128, SIB17bis_IEs_r18_ValidityDuration_r18_T256, SIB17bis_IEs_r18_ValidityDuration_r18_T512, SIB17bis_IEs_r18_ValidityDuration_r18_Infinity, SIB17bis_IEs_r18_ValidityDuration_r18_Spare5, SIB17bis_IEs_r18_ValidityDuration_r18_Spare4, SIB17bis_IEs_r18_ValidityDuration_r18_Spare3, SIB17bis_IEs_r18_ValidityDuration_r18_Spare2, SIB17bis_IEs_r18_ValidityDuration_r18_Spare1},
}

var sIB17bisIEsR18LateNonCriticalExtensionConstraints = per.SizeConstraints{}

type SIB17bis_IEs_r18 struct {
	Trs_ResourceSetConfig_r18 []TRS_ResourceSet_r18
	ValidityDuration_r18      *int64
	LateNonCriticalExtension  []byte
}

func (ie *SIB17bis_IEs_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB17bisIEsR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Trs_ResourceSetConfig_r18 != nil, ie.ValidityDuration_r18 != nil, ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. trs-ResourceSetConfig-r18: sequence-of
	{
		if ie.Trs_ResourceSetConfig_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sIB17bisIEsR18TrsResourceSetConfigR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Trs_ResourceSetConfig_r18))); err != nil {
				return err
			}
			for i := range ie.Trs_ResourceSetConfig_r18 {
				if err := ie.Trs_ResourceSetConfig_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. validityDuration-r18: enumerated
	{
		if ie.ValidityDuration_r18 != nil {
			if err := e.EncodeEnumerated(*ie.ValidityDuration_r18, sIB17bisIEsR18ValidityDurationR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB17bisIEsR18LateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SIB17bis_IEs_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB17bisIEsR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. trs-ResourceSetConfig-r18: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sIB17bisIEsR18TrsResourceSetConfigR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Trs_ResourceSetConfig_r18 = make([]TRS_ResourceSet_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Trs_ResourceSetConfig_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. validityDuration-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sIB17bisIEsR18ValidityDurationR18Constraints)
			if err != nil {
				return err
			}
			ie.ValidityDuration_r18 = &idx
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeOctetString(sIB17bisIEsR18LateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	return nil
}
