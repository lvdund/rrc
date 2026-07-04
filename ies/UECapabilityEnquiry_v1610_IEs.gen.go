// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UECapabilityEnquiry-v1610-IEs (line 2867).

var uECapabilityEnquiryV1610IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rrc-SegAllowed-r16", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	UECapabilityEnquiry_v1610_IEs_Rrc_SegAllowed_r16_Enabled = 0
)

var uECapabilityEnquiryV1610IEsRrcSegAllowedR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UECapabilityEnquiry_v1610_IEs_Rrc_SegAllowed_r16_Enabled},
}

type UECapabilityEnquiry_v1610_IEs struct {
	Rrc_SegAllowed_r16   *int64
	NonCriticalExtension *UECapabilityEnquiry_V17b0_IEs
}

func (ie *UECapabilityEnquiry_v1610_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uECapabilityEnquiryV1610IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Rrc_SegAllowed_r16 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. rrc-SegAllowed-r16: enumerated
	{
		if ie.Rrc_SegAllowed_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Rrc_SegAllowed_r16, uECapabilityEnquiryV1610IEsRrcSegAllowedR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UECapabilityEnquiry_v1610_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uECapabilityEnquiryV1610IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. rrc-SegAllowed-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(uECapabilityEnquiryV1610IEsRrcSegAllowedR16Constraints)
			if err != nil {
				return err
			}
			ie.Rrc_SegAllowed_r16 = &idx
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(UECapabilityEnquiry_V17b0_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
