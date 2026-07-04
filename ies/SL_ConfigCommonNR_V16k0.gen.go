// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-ConfigCommonNR-v16k0 (line 4331).

var sLConfigCommonNRV16k0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-FreqInfoListExt-v16k0", Optional: true},
	},
}

var sLConfigCommonNRV16k0SlFreqInfoListExtV16k0Constraints = per.SizeRange(1, common.MaxNrofFreqSL_r16)

type SL_ConfigCommonNR_V16k0 struct {
	Sl_FreqInfoListExt_V16k0 []SL_FreqConfigCommonExt_V16k0
}

func (ie *SL_ConfigCommonNR_V16k0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLConfigCommonNRV16k0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_FreqInfoListExt_V16k0 != nil}); err != nil {
		return err
	}

	// 2. sl-FreqInfoListExt-v16k0: sequence-of
	{
		if ie.Sl_FreqInfoListExt_V16k0 != nil {
			seqOf := e.NewSequenceOfEncoder(sLConfigCommonNRV16k0SlFreqInfoListExtV16k0Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_FreqInfoListExt_V16k0))); err != nil {
				return err
			}
			for i := range ie.Sl_FreqInfoListExt_V16k0 {
				if err := ie.Sl_FreqInfoListExt_V16k0[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *SL_ConfigCommonNR_V16k0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLConfigCommonNRV16k0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-FreqInfoListExt-v16k0: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sLConfigCommonNRV16k0SlFreqInfoListExtV16k0Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_FreqInfoListExt_V16k0 = make([]SL_FreqConfigCommonExt_V16k0, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_FreqInfoListExt_V16k0[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
