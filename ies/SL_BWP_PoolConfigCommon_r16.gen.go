// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-BWP-PoolConfigCommon-r16 (line 26847).

var sLBWPPoolConfigCommonR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-RxPool-r16", Optional: true},
		{Name: "sl-TxPoolSelectedNormal-r16", Optional: true},
		{Name: "sl-TxPoolExceptional-r16", Optional: true},
	},
}

var sLBWPPoolConfigCommonR16SlRxPoolR16Constraints = per.SizeRange(1, common.MaxNrofRXPool_r16)

var sLBWPPoolConfigCommonR16SlTxPoolSelectedNormalR16Constraints = per.SizeRange(1, common.MaxNrofTXPool_r16)

type SL_BWP_PoolConfigCommon_r16 struct {
	Sl_RxPool_r16               []SL_ResourcePool_r16
	Sl_TxPoolSelectedNormal_r16 []SL_ResourcePoolConfig_r16
	Sl_TxPoolExceptional_r16    *SL_ResourcePoolConfig_r16
}

func (ie *SL_BWP_PoolConfigCommon_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLBWPPoolConfigCommonR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_RxPool_r16 != nil, ie.Sl_TxPoolSelectedNormal_r16 != nil, ie.Sl_TxPoolExceptional_r16 != nil}); err != nil {
		return err
	}

	// 2. sl-RxPool-r16: sequence-of
	{
		if ie.Sl_RxPool_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLBWPPoolConfigCommonR16SlRxPoolR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_RxPool_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_RxPool_r16 {
				if err := ie.Sl_RxPool_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-TxPoolSelectedNormal-r16: sequence-of
	{
		if ie.Sl_TxPoolSelectedNormal_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(sLBWPPoolConfigCommonR16SlTxPoolSelectedNormalR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_TxPoolSelectedNormal_r16))); err != nil {
				return err
			}
			for i := range ie.Sl_TxPoolSelectedNormal_r16 {
				if err := ie.Sl_TxPoolSelectedNormal_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. sl-TxPoolExceptional-r16: ref
	{
		if ie.Sl_TxPoolExceptional_r16 != nil {
			if err := ie.Sl_TxPoolExceptional_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_BWP_PoolConfigCommon_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLBWPPoolConfigCommonR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-RxPool-r16: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sLBWPPoolConfigCommonR16SlRxPoolR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_RxPool_r16 = make([]SL_ResourcePool_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_RxPool_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-TxPoolSelectedNormal-r16: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sLBWPPoolConfigCommonR16SlTxPoolSelectedNormalR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_TxPoolSelectedNormal_r16 = make([]SL_ResourcePoolConfig_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_TxPoolSelectedNormal_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. sl-TxPoolExceptional-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_TxPoolExceptional_r16 = new(SL_ResourcePoolConfig_r16)
			if err := ie.Sl_TxPoolExceptional_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
