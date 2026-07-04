// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-BWP-DiscPoolConfig-r17 (line 26807).

var sLBWPDiscPoolConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-DiscRxPool-r17", Optional: true},
		{Name: "sl-DiscTxPoolSelected-r17", Optional: true},
		{Name: "sl-DiscTxPoolScheduling-r17", Optional: true},
	},
}

var sLBWPDiscPoolConfigR17SlDiscRxPoolR17Constraints = per.SizeRange(1, common.MaxNrofRXPool_r16)

type SL_BWP_DiscPoolConfig_r17 struct {
	Sl_DiscRxPool_r17           []SL_ResourcePool_r16
	Sl_DiscTxPoolSelected_r17   *SL_TxPoolDedicated_r16
	Sl_DiscTxPoolScheduling_r17 *SL_TxPoolDedicated_r16
}

func (ie *SL_BWP_DiscPoolConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLBWPDiscPoolConfigR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_DiscRxPool_r17 != nil, ie.Sl_DiscTxPoolSelected_r17 != nil, ie.Sl_DiscTxPoolScheduling_r17 != nil}); err != nil {
		return err
	}

	// 2. sl-DiscRxPool-r17: sequence-of
	{
		if ie.Sl_DiscRxPool_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(sLBWPDiscPoolConfigR17SlDiscRxPoolR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_DiscRxPool_r17))); err != nil {
				return err
			}
			for i := range ie.Sl_DiscRxPool_r17 {
				if err := ie.Sl_DiscRxPool_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-DiscTxPoolSelected-r17: ref
	{
		if ie.Sl_DiscTxPoolSelected_r17 != nil {
			if err := ie.Sl_DiscTxPoolSelected_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sl-DiscTxPoolScheduling-r17: ref
	{
		if ie.Sl_DiscTxPoolScheduling_r17 != nil {
			if err := ie.Sl_DiscTxPoolScheduling_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_BWP_DiscPoolConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLBWPDiscPoolConfigR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-DiscRxPool-r17: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sLBWPDiscPoolConfigR17SlDiscRxPoolR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_DiscRxPool_r17 = make([]SL_ResourcePool_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_DiscRxPool_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-DiscTxPoolSelected-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_DiscTxPoolSelected_r17 = new(SL_TxPoolDedicated_r16)
			if err := ie.Sl_DiscTxPoolSelected_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sl-DiscTxPoolScheduling-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_DiscTxPoolScheduling_r17 = new(SL_TxPoolDedicated_r16)
			if err := ie.Sl_DiscTxPoolScheduling_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
