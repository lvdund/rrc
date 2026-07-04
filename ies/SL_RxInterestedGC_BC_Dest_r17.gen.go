// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-RxInterestedGC-BC-Dest-r17 (line 2196).

var sLRxInterestedGCBCDestR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-RxInterestedQoS-InfoList-r17"},
		{Name: "sl-DestinationIdentity-r16"},
	},
}

var sLRxInterestedGCBCDestR17SlRxInterestedQoSInfoListR17Constraints = per.SizeRange(1, common.MaxNrofSL_QFIsPerDest_r16)

type SL_RxInterestedGC_BC_Dest_r17 struct {
	Sl_RxInterestedQoS_InfoList_r17 []SL_QoS_Info_r16
	Sl_DestinationIdentity_r16      SL_DestinationIdentity_r16
}

func (ie *SL_RxInterestedGC_BC_Dest_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLRxInterestedGCBCDestR17Constraints)
	_ = seq

	// 1. sl-RxInterestedQoS-InfoList-r17: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(sLRxInterestedGCBCDestR17SlRxInterestedQoSInfoListR17Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Sl_RxInterestedQoS_InfoList_r17))); err != nil {
			return err
		}
		for i := range ie.Sl_RxInterestedQoS_InfoList_r17 {
			if err := ie.Sl_RxInterestedQoS_InfoList_r17[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 2. sl-DestinationIdentity-r16: ref
	{
		if err := ie.Sl_DestinationIdentity_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_RxInterestedGC_BC_Dest_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLRxInterestedGCBCDestR17Constraints)
	_ = seq

	// 1. sl-RxInterestedQoS-InfoList-r17: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(sLRxInterestedGCBCDestR17SlRxInterestedQoSInfoListR17Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Sl_RxInterestedQoS_InfoList_r17 = make([]SL_QoS_Info_r16, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Sl_RxInterestedQoS_InfoList_r17[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 2. sl-DestinationIdentity-r16: ref
	{
		if err := ie.Sl_DestinationIdentity_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
