// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SCGFailureInformation-IEs (line 1838).

var sCGFailureInformationIEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "failureReportSCG", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type SCGFailureInformation_IEs struct {
	FailureReportSCG     *FailureReportSCG
	NonCriticalExtension *SCGFailureInformation_v1590_IEs
}

func (ie *SCGFailureInformation_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sCGFailureInformationIEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FailureReportSCG != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. failureReportSCG: ref
	{
		if ie.FailureReportSCG != nil {
			if err := ie.FailureReportSCG.Encode(e); err != nil {
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

func (ie *SCGFailureInformation_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sCGFailureInformationIEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. failureReportSCG: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FailureReportSCG = new(FailureReportSCG)
			if err := ie.FailureReportSCG.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(SCGFailureInformation_v1590_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
