// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersNR-v1740 (line 17596).

var cAParametersNRV1740Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "nack-OnlyFeedbackForSPS-Multicast-r17", Optional: true},
		{Name: "singlePUCCH-ConfigForMulticast-r17", Optional: true},
	},
}

const (
	CA_ParametersNR_v1740_Nack_OnlyFeedbackForSPS_Multicast_r17_Supported = 0
)

var cAParametersNRV1740NackOnlyFeedbackForSPSMulticastR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1740_Nack_OnlyFeedbackForSPS_Multicast_r17_Supported},
}

const (
	CA_ParametersNR_v1740_SinglePUCCH_ConfigForMulticast_r17_Supported = 0
)

var cAParametersNRV1740SinglePUCCHConfigForMulticastR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1740_SinglePUCCH_ConfigForMulticast_r17_Supported},
}

type CA_ParametersNR_v1740 struct {
	Nack_OnlyFeedbackForSPS_Multicast_r17 *int64
	SinglePUCCH_ConfigForMulticast_r17    *int64
}

func (ie *CA_ParametersNR_v1740) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRV1740Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Nack_OnlyFeedbackForSPS_Multicast_r17 != nil, ie.SinglePUCCH_ConfigForMulticast_r17 != nil}); err != nil {
		return err
	}

	// 2. nack-OnlyFeedbackForSPS-Multicast-r17: enumerated
	{
		if ie.Nack_OnlyFeedbackForSPS_Multicast_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Nack_OnlyFeedbackForSPS_Multicast_r17, cAParametersNRV1740NackOnlyFeedbackForSPSMulticastR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. singlePUCCH-ConfigForMulticast-r17: enumerated
	{
		if ie.SinglePUCCH_ConfigForMulticast_r17 != nil {
			if err := e.EncodeEnumerated(*ie.SinglePUCCH_ConfigForMulticast_r17, cAParametersNRV1740SinglePUCCHConfigForMulticastR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNR_v1740) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRV1740Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. nack-OnlyFeedbackForSPS-Multicast-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1740NackOnlyFeedbackForSPSMulticastR17Constraints)
			if err != nil {
				return err
			}
			ie.Nack_OnlyFeedbackForSPS_Multicast_r17 = &idx
		}
	}

	// 3. singlePUCCH-ConfigForMulticast-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1740SinglePUCCHConfigForMulticastR17Constraints)
			if err != nil {
				return err
			}
			ie.SinglePUCCH_ConfigForMulticast_r17 = &idx
		}
	}

	return nil
}
