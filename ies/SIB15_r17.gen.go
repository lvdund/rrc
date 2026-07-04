// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SIB15-r17 (line 4381).

var sIB15R17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "commonPLMNsWithDisasterCondition-r17", Optional: true},
		{Name: "applicableDisasterInfoList-r17", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
}

var sIB15R17CommonPLMNsWithDisasterConditionR17Constraints = per.SizeRange(1, common.MaxPLMN)

var sIB15R17ApplicableDisasterInfoListR17Constraints = per.SizeRange(1, common.MaxPLMN)

var sIB15R17LateNonCriticalExtensionConstraints = per.SizeConstraints{}

type SIB15_r17 struct {
	CommonPLMNsWithDisasterCondition_r17 []PLMN_Identity
	ApplicableDisasterInfoList_r17       []ApplicableDisasterInfo_r17
	LateNonCriticalExtension             []byte
}

func (ie *SIB15_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB15R17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CommonPLMNsWithDisasterCondition_r17 != nil, ie.ApplicableDisasterInfoList_r17 != nil, ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. commonPLMNsWithDisasterCondition-r17: sequence-of
	{
		if ie.CommonPLMNsWithDisasterCondition_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(sIB15R17CommonPLMNsWithDisasterConditionR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.CommonPLMNsWithDisasterCondition_r17))); err != nil {
				return err
			}
			for i := range ie.CommonPLMNsWithDisasterCondition_r17 {
				if err := ie.CommonPLMNsWithDisasterCondition_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. applicableDisasterInfoList-r17: sequence-of
	{
		if ie.ApplicableDisasterInfoList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(sIB15R17ApplicableDisasterInfoListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.ApplicableDisasterInfoList_r17))); err != nil {
				return err
			}
			for i := range ie.ApplicableDisasterInfoList_r17 {
				if err := ie.ApplicableDisasterInfoList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB15R17LateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SIB15_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB15R17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. commonPLMNsWithDisasterCondition-r17: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sIB15R17CommonPLMNsWithDisasterConditionR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.CommonPLMNsWithDisasterCondition_r17 = make([]PLMN_Identity, n)
			for i := int64(0); i < n; i++ {
				if err := ie.CommonPLMNsWithDisasterCondition_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. applicableDisasterInfoList-r17: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sIB15R17ApplicableDisasterInfoListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.ApplicableDisasterInfoList_r17 = make([]ApplicableDisasterInfo_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.ApplicableDisasterInfoList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeOctetString(sIB15R17LateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	return nil
}
