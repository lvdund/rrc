// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB21-r17 (line 4591).

var sIB21R17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "mbs-FSAI-IntraFreq-r17", Optional: true},
		{Name: "mbs-FSAI-InterFreqList-r17", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
}

var sIB21R17LateNonCriticalExtensionConstraints = per.SizeConstraints{}

type SIB21_r17 struct {
	Mbs_FSAI_IntraFreq_r17     *MBS_FSAI_List_r17
	Mbs_FSAI_InterFreqList_r17 *MBS_FSAI_InterFreqList_r17
	LateNonCriticalExtension   []byte
}

func (ie *SIB21_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB21R17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mbs_FSAI_IntraFreq_r17 != nil, ie.Mbs_FSAI_InterFreqList_r17 != nil, ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. mbs-FSAI-IntraFreq-r17: ref
	{
		if ie.Mbs_FSAI_IntraFreq_r17 != nil {
			if err := ie.Mbs_FSAI_IntraFreq_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. mbs-FSAI-InterFreqList-r17: ref
	{
		if ie.Mbs_FSAI_InterFreqList_r17 != nil {
			if err := ie.Mbs_FSAI_InterFreqList_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB21R17LateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SIB21_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB21R17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. mbs-FSAI-IntraFreq-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Mbs_FSAI_IntraFreq_r17 = new(MBS_FSAI_List_r17)
			if err := ie.Mbs_FSAI_IntraFreq_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. mbs-FSAI-InterFreqList-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Mbs_FSAI_InterFreqList_r17 = new(MBS_FSAI_InterFreqList_r17)
			if err := ie.Mbs_FSAI_InterFreqList_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeOctetString(sIB21R17LateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	return nil
}
