// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UECapabilityEnquiry-v1560-IEs (line 2862).

var uECapabilityEnquiryV1560IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "capabilityRequestFilterCommon", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type UECapabilityEnquiry_v1560_IEs struct {
	CapabilityRequestFilterCommon *UE_CapabilityRequestFilterCommon
	NonCriticalExtension          *UECapabilityEnquiry_v1610_IEs
}

func (ie *UECapabilityEnquiry_v1560_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uECapabilityEnquiryV1560IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CapabilityRequestFilterCommon != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. capabilityRequestFilterCommon: ref
	{
		if ie.CapabilityRequestFilterCommon != nil {
			if err := ie.CapabilityRequestFilterCommon.Encode(e); err != nil {
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

func (ie *UECapabilityEnquiry_v1560_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uECapabilityEnquiryV1560IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. capabilityRequestFilterCommon: ref
	{
		if seq.IsComponentPresent(0) {
			ie.CapabilityRequestFilterCommon = new(UE_CapabilityRequestFilterCommon)
			if err := ie.CapabilityRequestFilterCommon.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(UECapabilityEnquiry_v1610_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
