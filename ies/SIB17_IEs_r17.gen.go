// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SIB17-IEs-r17 (line 4413).

var sIB17IEsR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "trs-ResourceSetConfig-r17"},
		{Name: "validityDuration-r17", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
}

var sIB17IEsR17TrsResourceSetConfigR17Constraints = per.SizeRange(1, common.MaxNrofTRS_ResourceSets_r17)

const (
	SIB17_IEs_r17_ValidityDuration_r17_T1       = 0
	SIB17_IEs_r17_ValidityDuration_r17_T2       = 1
	SIB17_IEs_r17_ValidityDuration_r17_T4       = 2
	SIB17_IEs_r17_ValidityDuration_r17_T8       = 3
	SIB17_IEs_r17_ValidityDuration_r17_T16      = 4
	SIB17_IEs_r17_ValidityDuration_r17_T32      = 5
	SIB17_IEs_r17_ValidityDuration_r17_T64      = 6
	SIB17_IEs_r17_ValidityDuration_r17_T128     = 7
	SIB17_IEs_r17_ValidityDuration_r17_T256     = 8
	SIB17_IEs_r17_ValidityDuration_r17_T512     = 9
	SIB17_IEs_r17_ValidityDuration_r17_Infinity = 10
	SIB17_IEs_r17_ValidityDuration_r17_Spare5   = 11
	SIB17_IEs_r17_ValidityDuration_r17_Spare4   = 12
	SIB17_IEs_r17_ValidityDuration_r17_Spare3   = 13
	SIB17_IEs_r17_ValidityDuration_r17_Spare2   = 14
	SIB17_IEs_r17_ValidityDuration_r17_Spare1   = 15
)

var sIB17IEsR17ValidityDurationR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB17_IEs_r17_ValidityDuration_r17_T1, SIB17_IEs_r17_ValidityDuration_r17_T2, SIB17_IEs_r17_ValidityDuration_r17_T4, SIB17_IEs_r17_ValidityDuration_r17_T8, SIB17_IEs_r17_ValidityDuration_r17_T16, SIB17_IEs_r17_ValidityDuration_r17_T32, SIB17_IEs_r17_ValidityDuration_r17_T64, SIB17_IEs_r17_ValidityDuration_r17_T128, SIB17_IEs_r17_ValidityDuration_r17_T256, SIB17_IEs_r17_ValidityDuration_r17_T512, SIB17_IEs_r17_ValidityDuration_r17_Infinity, SIB17_IEs_r17_ValidityDuration_r17_Spare5, SIB17_IEs_r17_ValidityDuration_r17_Spare4, SIB17_IEs_r17_ValidityDuration_r17_Spare3, SIB17_IEs_r17_ValidityDuration_r17_Spare2, SIB17_IEs_r17_ValidityDuration_r17_Spare1},
}

var sIB17IEsR17LateNonCriticalExtensionConstraints = per.SizeConstraints{}

type SIB17_IEs_r17 struct {
	Trs_ResourceSetConfig_r17 []TRS_ResourceSet_r17
	ValidityDuration_r17      *int64
	LateNonCriticalExtension  []byte
}

func (ie *SIB17_IEs_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB17IEsR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ValidityDuration_r17 != nil, ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. trs-ResourceSetConfig-r17: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(sIB17IEsR17TrsResourceSetConfigR17Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Trs_ResourceSetConfig_r17))); err != nil {
			return err
		}
		for i := range ie.Trs_ResourceSetConfig_r17 {
			if err := ie.Trs_ResourceSetConfig_r17[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. validityDuration-r17: enumerated
	{
		if ie.ValidityDuration_r17 != nil {
			if err := e.EncodeEnumerated(*ie.ValidityDuration_r17, sIB17IEsR17ValidityDurationR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB17IEsR17LateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SIB17_IEs_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB17IEsR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. trs-ResourceSetConfig-r17: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(sIB17IEsR17TrsResourceSetConfigR17Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Trs_ResourceSetConfig_r17 = make([]TRS_ResourceSet_r17, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Trs_ResourceSetConfig_r17[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. validityDuration-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sIB17IEsR17ValidityDurationR17Constraints)
			if err != nil {
				return err
			}
			ie.ValidityDuration_r17 = &idx
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeOctetString(sIB17IEsR17LateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	return nil
}
