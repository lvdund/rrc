// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-BWP-DiscPoolConfigCommon-r17 (line 26816).

var sLBWPDiscPoolConfigCommonR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-DiscRxPool-r17", Optional: true},
		{Name: "sl-DiscTxPoolSelected-r17", Optional: true},
	},
}

var sLBWPDiscPoolConfigCommonR17SlDiscRxPoolR17Constraints = per.SizeRange(1, common.MaxNrofRXPool_r16)

var sLBWPDiscPoolConfigCommonR17SlDiscTxPoolSelectedR17Constraints = per.SizeRange(1, common.MaxNrofTXPool_r16)

type SL_BWP_DiscPoolConfigCommon_r17 struct {
	Sl_DiscRxPool_r17         []SL_ResourcePool_r16
	Sl_DiscTxPoolSelected_r17 []SL_ResourcePoolConfig_r16
}

func (ie *SL_BWP_DiscPoolConfigCommon_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLBWPDiscPoolConfigCommonR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_DiscRxPool_r17 != nil, ie.Sl_DiscTxPoolSelected_r17 != nil}); err != nil {
		return err
	}

	// 3. sl-DiscRxPool-r17: sequence-of
	{
		if ie.Sl_DiscRxPool_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(sLBWPDiscPoolConfigCommonR17SlDiscRxPoolR17Constraints)
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

	// 4. sl-DiscTxPoolSelected-r17: sequence-of
	{
		if ie.Sl_DiscTxPoolSelected_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(sLBWPDiscPoolConfigCommonR17SlDiscTxPoolSelectedR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_DiscTxPoolSelected_r17))); err != nil {
				return err
			}
			for i := range ie.Sl_DiscTxPoolSelected_r17 {
				if err := ie.Sl_DiscTxPoolSelected_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SL_BWP_DiscPoolConfigCommon_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLBWPDiscPoolConfigCommonR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-DiscRxPool-r17: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sLBWPDiscPoolConfigCommonR17SlDiscRxPoolR17Constraints)
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

	// 4. sl-DiscTxPoolSelected-r17: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sLBWPDiscPoolConfigCommonR17SlDiscTxPoolSelectedR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_DiscTxPoolSelected_r17 = make([]SL_ResourcePoolConfig_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_DiscTxPoolSelected_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
