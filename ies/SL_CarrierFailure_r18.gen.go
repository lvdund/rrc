// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-CarrierFailure-r18 (line 2307).

var sLCarrierFailureR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-DestinationIdentity-r18"},
		{Name: "sl-CarrierFailure-r18"},
	},
}

var sLCarrierFailureR18SlCarrierFailureR18Constraints = per.SizeRange(1, common.MaxNrofFreqSL_r16)

type SL_CarrierFailure_r18 struct {
	Sl_DestinationIdentity_r18 SL_DestinationIdentity_r16
	Sl_CarrierFailure_r18      []int64
}

func (ie *SL_CarrierFailure_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLCarrierFailureR18Constraints)
	_ = seq

	// 1. sl-DestinationIdentity-r18: ref
	{
		if err := ie.Sl_DestinationIdentity_r18.Encode(e); err != nil {
			return err
		}
	}

	// 2. sl-CarrierFailure-r18: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(sLCarrierFailureR18SlCarrierFailureR18Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Sl_CarrierFailure_r18))); err != nil {
			return err
		}
		for i := range ie.Sl_CarrierFailure_r18 {
			if err := e.EncodeInteger(ie.Sl_CarrierFailure_r18[i], per.Constrained(1, common.MaxNrofFreqSL_r16)); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_CarrierFailure_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLCarrierFailureR18Constraints)
	_ = seq

	// 1. sl-DestinationIdentity-r18: ref
	{
		if err := ie.Sl_DestinationIdentity_r18.Decode(d); err != nil {
			return err
		}
	}

	// 2. sl-CarrierFailure-r18: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(sLCarrierFailureR18SlCarrierFailureR18Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Sl_CarrierFailure_r18 = make([]int64, n)
		for i := int64(0); i < n; i++ {
			v, err := d.DecodeInteger(per.Constrained(1, common.MaxNrofFreqSL_r16))
			if err != nil {
				return err
			}
			ie.Sl_CarrierFailure_r18[i] = v
		}
	}

	return nil
}
