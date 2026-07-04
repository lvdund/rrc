// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: RA-PrioritizationSliceInfo-r17 (line 13099).

var rAPrioritizationSliceInfoR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nsag-ID-List-r17"},
		{Name: "ra-Prioritization-r17"},
	},
}

var rAPrioritizationSliceInfoR17NsagIDListR17Constraints = per.SizeRange(1, common.MaxSliceInfo_r17)

type RA_PrioritizationSliceInfo_r17 struct {
	Nsag_ID_List_r17      []NSAG_ID_r17
	Ra_Prioritization_r17 RA_Prioritization
}

func (ie *RA_PrioritizationSliceInfo_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rAPrioritizationSliceInfoR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. nsag-ID-List-r17: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(rAPrioritizationSliceInfoR17NsagIDListR17Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Nsag_ID_List_r17))); err != nil {
			return err
		}
		for i := range ie.Nsag_ID_List_r17 {
			if err := ie.Nsag_ID_List_r17[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. ra-Prioritization-r17: ref
	{
		if err := ie.Ra_Prioritization_r17.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *RA_PrioritizationSliceInfo_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rAPrioritizationSliceInfoR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. nsag-ID-List-r17: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(rAPrioritizationSliceInfoR17NsagIDListR17Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Nsag_ID_List_r17 = make([]NSAG_ID_r17, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Nsag_ID_List_r17[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. ra-Prioritization-r17: ref
	{
		if err := ie.Ra_Prioritization_r17.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
