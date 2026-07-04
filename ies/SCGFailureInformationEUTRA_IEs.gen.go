// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SCGFailureInformationEUTRA-IEs (line 1897).

var sCGFailureInformationEUTRAIEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "failureReportSCG-EUTRA", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type SCGFailureInformationEUTRA_IEs struct {
	FailureReportSCG_EUTRA *FailureReportSCG_EUTRA
	NonCriticalExtension   *SCGFailureInformationEUTRA_v1590_IEs
}

func (ie *SCGFailureInformationEUTRA_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sCGFailureInformationEUTRAIEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FailureReportSCG_EUTRA != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. failureReportSCG-EUTRA: ref
	{
		if ie.FailureReportSCG_EUTRA != nil {
			if err := ie.FailureReportSCG_EUTRA.Encode(e); err != nil {
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

func (ie *SCGFailureInformationEUTRA_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sCGFailureInformationEUTRAIEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. failureReportSCG-EUTRA: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FailureReportSCG_EUTRA = new(FailureReportSCG_EUTRA)
			if err := ie.FailureReportSCG_EUTRA.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(SCGFailureInformationEUTRA_v1590_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
