// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MCGFailureInformation-r16-IEs (line 699).

var mCGFailureInformationR16IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "failureReportMCG-r16", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var mCGFailureInformationR16IEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type MCGFailureInformation_r16_IEs struct {
	FailureReportMCG_r16     *FailureReportMCG_r16
	LateNonCriticalExtension []byte
	NonCriticalExtension     *struct{}
}

func (ie *MCGFailureInformation_r16_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mCGFailureInformationR16IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FailureReportMCG_r16 != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. failureReportMCG-r16: ref
	{
		if ie.FailureReportMCG_r16 != nil {
			if err := ie.FailureReportMCG_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, mCGFailureInformationR16IEsLateNonCriticalExtensionConstraints); err != nil {
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

func (ie *MCGFailureInformation_r16_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mCGFailureInformationR16IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. failureReportMCG-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FailureReportMCG_r16 = new(FailureReportMCG_r16)
			if err := ie.FailureReportMCG_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(mCGFailureInformationR16IEsLateNonCriticalExtensionConstraints)
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
