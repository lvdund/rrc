// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-PRS-TxPoolDedicated-r18 (line 26863).

var sLPRSTxPoolDedicatedR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PRS-PoolToReleaseList-r18", Optional: true},
		{Name: "sl-PRS-PoolToAddModList-r18", Optional: true},
	},
}

var sLPRSTxPoolDedicatedR18SlPRSPoolToReleaseListR18Constraints = per.SizeRange(1, common.MaxNrofSL_PRS_TxPool_r18)

var sLPRSTxPoolDedicatedR18SlPRSPoolToAddModListR18Constraints = per.SizeRange(1, common.MaxNrofSL_PRS_TxPool_r18)

type SL_PRS_TxPoolDedicated_r18 struct {
	Sl_PRS_PoolToReleaseList_r18 []SL_PRS_ResourcePoolID_r18
	Sl_PRS_PoolToAddModList_r18  []SL_PRS_ResourcePoolConfig_r18
}

func (ie *SL_PRS_TxPoolDedicated_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPRSTxPoolDedicatedR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_PRS_PoolToReleaseList_r18 != nil, ie.Sl_PRS_PoolToAddModList_r18 != nil}); err != nil {
		return err
	}

	// 2. sl-PRS-PoolToReleaseList-r18: sequence-of
	{
		if ie.Sl_PRS_PoolToReleaseList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLPRSTxPoolDedicatedR18SlPRSPoolToReleaseListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_PRS_PoolToReleaseList_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_PRS_PoolToReleaseList_r18 {
				if err := ie.Sl_PRS_PoolToReleaseList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-PRS-PoolToAddModList-r18: sequence-of
	{
		if ie.Sl_PRS_PoolToAddModList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLPRSTxPoolDedicatedR18SlPRSPoolToAddModListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_PRS_PoolToAddModList_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_PRS_PoolToAddModList_r18 {
				if err := ie.Sl_PRS_PoolToAddModList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SL_PRS_TxPoolDedicated_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPRSTxPoolDedicatedR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-PRS-PoolToReleaseList-r18: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sLPRSTxPoolDedicatedR18SlPRSPoolToReleaseListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_PRS_PoolToReleaseList_r18 = make([]SL_PRS_ResourcePoolID_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_PRS_PoolToReleaseList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-PRS-PoolToAddModList-r18: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sLPRSTxPoolDedicatedR18SlPRSPoolToAddModListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_PRS_PoolToAddModList_r18 = make([]SL_PRS_ResourcePoolConfig_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_PRS_PoolToAddModList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
