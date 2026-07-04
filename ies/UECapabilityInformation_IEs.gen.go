// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UECapabilityInformation-IEs (line 2888).

var uECapabilityInformationIEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ue-CapabilityRAT-ContainerList", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var uECapabilityInformationIEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type UECapabilityInformation_IEs struct {
	Ue_CapabilityRAT_ContainerList *UE_CapabilityRAT_ContainerList
	LateNonCriticalExtension       []byte
	NonCriticalExtension           *struct{}
}

func (ie *UECapabilityInformation_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uECapabilityInformationIEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ue_CapabilityRAT_ContainerList != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. ue-CapabilityRAT-ContainerList: ref
	{
		if ie.Ue_CapabilityRAT_ContainerList != nil {
			if err := ie.Ue_CapabilityRAT_ContainerList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, uECapabilityInformationIEsLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 4. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *UECapabilityInformation_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uECapabilityInformationIEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ue-CapabilityRAT-ContainerList: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ue_CapabilityRAT_ContainerList = new(UE_CapabilityRAT_ContainerList)
			if err := ie.Ue_CapabilityRAT_ContainerList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(uECapabilityInformationIEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 4. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(2) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
