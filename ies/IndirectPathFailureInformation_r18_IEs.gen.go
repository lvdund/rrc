// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IndirectPathFailureInformation-r18-IEs (line 500).

var indirectPathFailureInformationR18IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "failureReportIndirectPath-r18", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var indirectPathFailureInformationR18IEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type IndirectPathFailureInformation_r18_IEs struct {
	FailureReportIndirectPath_r18 *FailureReportIndirectPath_r18
	LateNonCriticalExtension      []byte
	NonCriticalExtension          *struct{}
}

func (ie *IndirectPathFailureInformation_r18_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(indirectPathFailureInformationR18IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FailureReportIndirectPath_r18 != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. failureReportIndirectPath-r18: ref
	{
		if ie.FailureReportIndirectPath_r18 != nil {
			if err := ie.FailureReportIndirectPath_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, indirectPathFailureInformationR18IEsLateNonCriticalExtensionConstraints); err != nil {
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

func (ie *IndirectPathFailureInformation_r18_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(indirectPathFailureInformationR18IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. failureReportIndirectPath-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FailureReportIndirectPath_r18 = new(FailureReportIndirectPath_r18)
			if err := ie.FailureReportIndirectPath_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(indirectPathFailureInformationR18IEsLateNonCriticalExtensionConstraints)
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
