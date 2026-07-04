// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-TxPoolDedicated-r16 (line 26832).

var sLTxPoolDedicatedR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PoolToReleaseList-r16", Optional: true},
		{Name: "sl-PoolToAddModList-r16", Optional: true},
	},
}

var sLTxPoolDedicatedR16SlPoolToReleaseListR16Constraints = per.SizeRange(1, common.MaxNrofTXPool_r16)

var sLTxPoolDedicatedR16SlPoolToAddModListR16Constraints = per.SizeRange(1, common.MaxNrofTXPool_r16)

type SL_TxPoolDedicated_r16 struct {
	Sl_PoolToReleaseList_r16 []SL_ResourcePoolID_r16
	Sl_PoolToAddModList_r16  []SL_ResourcePoolConfig_r16
}

func (ie *SL_TxPoolDedicated_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLTxPoolDedicatedR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_PoolToReleaseList_r16 != nil, ie.Sl_PoolToAddModList_r16 != nil}); err != nil {
		return err
	}

	// 2. sl-PoolToReleaseList-r16: sequence-of
	{
		if ie.Sl_PoolToReleaseList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLTxPoolDedicatedR16SlPoolToReleaseListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_PoolToReleaseList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_PoolToReleaseList_r16 {
				if err := ie.Sl_PoolToReleaseList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-PoolToAddModList-r16: sequence-of
	{
		if ie.Sl_PoolToAddModList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLTxPoolDedicatedR16SlPoolToAddModListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_PoolToAddModList_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_PoolToAddModList_r16 {
				if err := ie.Sl_PoolToAddModList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SL_TxPoolDedicated_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLTxPoolDedicatedR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-PoolToReleaseList-r16: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sLTxPoolDedicatedR16SlPoolToReleaseListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_PoolToReleaseList_r16 = make([]SL_ResourcePoolID_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_PoolToReleaseList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-PoolToAddModList-r16: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sLTxPoolDedicatedR16SlPoolToAddModListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_PoolToAddModList_r16 = make([]SL_ResourcePoolConfig_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_PoolToAddModList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
