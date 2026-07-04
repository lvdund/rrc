// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FailureInformation-IEs (line 404).

var failureInformationIEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "failureInfoRLC-Bearer", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var failureInformationIEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type FailureInformation_IEs struct {
	FailureInfoRLC_Bearer    *FailureInfoRLC_Bearer
	LateNonCriticalExtension []byte
	NonCriticalExtension     *FailureInformation_v1610_IEs
}

func (ie *FailureInformation_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(failureInformationIEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FailureInfoRLC_Bearer != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. failureInfoRLC-Bearer: ref
	{
		if ie.FailureInfoRLC_Bearer != nil {
			if err := ie.FailureInfoRLC_Bearer.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, failureInformationIEsLateNonCriticalExtensionConstraints); err != nil {
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

func (ie *FailureInformation_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(failureInformationIEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. failureInfoRLC-Bearer: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FailureInfoRLC_Bearer = new(FailureInfoRLC_Bearer)
			if err := ie.FailureInfoRLC_Bearer.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(failureInformationIEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(2) {
			ie.NonCriticalExtension = new(FailureInformation_v1610_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
