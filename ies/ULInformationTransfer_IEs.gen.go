// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ULInformationTransfer-IEs (line 3640).

var uLInformationTransferIEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dedicatedNAS-Message", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var uLInformationTransferIEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type ULInformationTransfer_IEs struct {
	DedicatedNAS_Message     *DedicatedNAS_Message
	LateNonCriticalExtension []byte
	NonCriticalExtension     *ULInformationTransfer_v1700_IEs
}

func (ie *ULInformationTransfer_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uLInformationTransferIEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DedicatedNAS_Message != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. dedicatedNAS-Message: ref
	{
		if ie.DedicatedNAS_Message != nil {
			if err := ie.DedicatedNAS_Message.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, uLInformationTransferIEsLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *ULInformationTransfer_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uLInformationTransferIEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. dedicatedNAS-Message: ref
	{
		if seq.IsComponentPresent(0) {
			ie.DedicatedNAS_Message = new(DedicatedNAS_Message)
			if err := ie.DedicatedNAS_Message.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(uLInformationTransferIEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(2) {
			ie.NonCriticalExtension = new(ULInformationTransfer_v1700_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
