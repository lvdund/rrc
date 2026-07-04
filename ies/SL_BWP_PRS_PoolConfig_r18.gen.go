// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-BWP-PRS-PoolConfig-r18 (line 26856).

var sLBWPPRSPoolConfigR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PRS-RxPool-r18", Optional: true},
		{Name: "sl-PRS-TxPoolSelectedNormal-r18", Optional: true},
		{Name: "sl-PRS-TxPoolScheduling-r18", Optional: true},
		{Name: "sl-PRS-TxPoolExceptional-r18", Optional: true},
	},
}

var sLBWPPRSPoolConfigR18SlPRSRxPoolR18Constraints = per.SizeRange(1, common.MaxNrofRXPool_r16)

type SL_BWP_PRS_PoolConfig_r18 struct {
	Sl_PRS_RxPool_r18               []SL_PRS_ResourcePool_r18
	Sl_PRS_TxPoolSelectedNormal_r18 *SL_PRS_TxPoolDedicated_r18
	Sl_PRS_TxPoolScheduling_r18     *SL_PRS_TxPoolDedicated_r18
	Sl_PRS_TxPoolExceptional_r18    *SL_PRS_ResourcePoolConfig_r18
}

func (ie *SL_BWP_PRS_PoolConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLBWPPRSPoolConfigR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_PRS_RxPool_r18 != nil, ie.Sl_PRS_TxPoolSelectedNormal_r18 != nil, ie.Sl_PRS_TxPoolScheduling_r18 != nil, ie.Sl_PRS_TxPoolExceptional_r18 != nil}); err != nil {
		return err
	}

	// 2. sl-PRS-RxPool-r18: sequence-of
	{
		if ie.Sl_PRS_RxPool_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLBWPPRSPoolConfigR18SlPRSRxPoolR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_PRS_RxPool_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_PRS_RxPool_r18 {
				if err := ie.Sl_PRS_RxPool_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-PRS-TxPoolSelectedNormal-r18: ref
	{
		if ie.Sl_PRS_TxPoolSelectedNormal_r18 != nil {
			if err := ie.Sl_PRS_TxPoolSelectedNormal_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sl-PRS-TxPoolScheduling-r18: ref
	{
		if ie.Sl_PRS_TxPoolScheduling_r18 != nil {
			if err := ie.Sl_PRS_TxPoolScheduling_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. sl-PRS-TxPoolExceptional-r18: ref
	{
		if ie.Sl_PRS_TxPoolExceptional_r18 != nil {
			if err := ie.Sl_PRS_TxPoolExceptional_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_BWP_PRS_PoolConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLBWPPRSPoolConfigR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-PRS-RxPool-r18: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sLBWPPRSPoolConfigR18SlPRSRxPoolR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_PRS_RxPool_r18 = make([]SL_PRS_ResourcePool_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_PRS_RxPool_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-PRS-TxPoolSelectedNormal-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_PRS_TxPoolSelectedNormal_r18 = new(SL_PRS_TxPoolDedicated_r18)
			if err := ie.Sl_PRS_TxPoolSelectedNormal_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sl-PRS-TxPoolScheduling-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_PRS_TxPoolScheduling_r18 = new(SL_PRS_TxPoolDedicated_r18)
			if err := ie.Sl_PRS_TxPoolScheduling_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. sl-PRS-TxPoolExceptional-r18: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Sl_PRS_TxPoolExceptional_r18 = new(SL_PRS_ResourcePoolConfig_r18)
			if err := ie.Sl_PRS_TxPoolExceptional_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
