// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UECapabilityEnquiry-v17b0-IEs (line 2872).

var uECapabilityEnquiryV17b0IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rrc-MaxCapaSegAllowed-r17", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var uECapabilityEnquiryV17b0IEsRrcMaxCapaSegAllowedR17Constraints = per.Constrained(2, 16)

type UECapabilityEnquiry_V17b0_IEs struct {
	Rrc_MaxCapaSegAllowed_r17 *int64
	NonCriticalExtension      *struct{}
}

func (ie *UECapabilityEnquiry_V17b0_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uECapabilityEnquiryV17b0IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Rrc_MaxCapaSegAllowed_r17 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. rrc-MaxCapaSegAllowed-r17: integer
	{
		if ie.Rrc_MaxCapaSegAllowed_r17 != nil {
			if err := e.EncodeInteger(*ie.Rrc_MaxCapaSegAllowed_r17, uECapabilityEnquiryV17b0IEsRrcMaxCapaSegAllowedR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *UECapabilityEnquiry_V17b0_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uECapabilityEnquiryV17b0IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. rrc-MaxCapaSegAllowed-r17: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(uECapabilityEnquiryV17b0IEsRrcMaxCapaSegAllowedR17Constraints)
			if err != nil {
				return err
			}
			ie.Rrc_MaxCapaSegAllowed_r17 = &v
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
