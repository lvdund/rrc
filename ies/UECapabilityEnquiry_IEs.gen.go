// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UECapabilityEnquiry-IEs (line 2856).

var uECapabilityEnquiryIEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ue-CapabilityRAT-RequestList"},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "ue-CapabilityEnquiryExt", Optional: true},
	},
}

var uECapabilityEnquiryIEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

var uECapabilityEnquiryIEsUeCapabilityEnquiryExtConstraints = per.SizeConstraints{}

type UECapabilityEnquiry_IEs struct {
	Ue_CapabilityRAT_RequestList UE_CapabilityRAT_RequestList
	LateNonCriticalExtension     []byte
	Ue_CapabilityEnquiryExt      []byte
}

func (ie *UECapabilityEnquiry_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uECapabilityEnquiryIEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LateNonCriticalExtension != nil, ie.Ue_CapabilityEnquiryExt != nil}); err != nil {
		return err
	}

	// 2. ue-CapabilityRAT-RequestList: ref
	{
		if err := ie.Ue_CapabilityRAT_RequestList.Encode(e); err != nil {
			return err
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, uECapabilityEnquiryIEsLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 4. ue-CapabilityEnquiryExt: octet-string
	{
		if ie.Ue_CapabilityEnquiryExt != nil {
			if err := e.EncodeOctetString(ie.Ue_CapabilityEnquiryExt, uECapabilityEnquiryIEsUeCapabilityEnquiryExtConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UECapabilityEnquiry_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uECapabilityEnquiryIEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ue-CapabilityRAT-RequestList: ref
	{
		if err := ie.Ue_CapabilityRAT_RequestList.Decode(d); err != nil {
			return err
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(uECapabilityEnquiryIEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 4. ue-CapabilityEnquiryExt: octet-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeOctetString(uECapabilityEnquiryIEsUeCapabilityEnquiryExtConstraints)
			if err != nil {
				return err
			}
			ie.Ue_CapabilityEnquiryExt = v
		}
	}

	return nil
}
