// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UE-RadioPagingInfo-r17 (line 25949).

var uERadioPagingInfoR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "pei-SubgroupingSupportBandList-r17", Optional: true},
	},
}

var uERadioPagingInfoR17PeiSubgroupingSupportBandListR17Constraints = per.SizeRange(1, common.MaxBands)

type UE_RadioPagingInfo_r17 struct {
	Pei_SubgroupingSupportBandList_r17 []FreqBandIndicatorNR
}

func (ie *UE_RadioPagingInfo_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uERadioPagingInfoR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pei_SubgroupingSupportBandList_r17 != nil}); err != nil {
		return err
	}

	// 3. pei-SubgroupingSupportBandList-r17: sequence-of
	{
		if ie.Pei_SubgroupingSupportBandList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(uERadioPagingInfoR17PeiSubgroupingSupportBandListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Pei_SubgroupingSupportBandList_r17))); err != nil {
				return err
			}
			for i := range ie.Pei_SubgroupingSupportBandList_r17 {
				if err := ie.Pei_SubgroupingSupportBandList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *UE_RadioPagingInfo_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uERadioPagingInfoR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. pei-SubgroupingSupportBandList-r17: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(uERadioPagingInfoR17PeiSubgroupingSupportBandListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Pei_SubgroupingSupportBandList_r17 = make([]FreqBandIndicatorNR, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Pei_SubgroupingSupportBandList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
