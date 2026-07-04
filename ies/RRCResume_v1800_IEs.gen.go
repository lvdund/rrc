// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCResume-v1800-IEs (line 1574).

var rRCResumeV1800IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "needForInterruptionConfigNR-r18", Optional: true},
		{Name: "reselectionMeasurementReq-r18", Optional: true},
		{Name: "validatedMeasurementsReq-r18", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	RRCResume_v1800_IEs_NeedForInterruptionConfigNR_r18_Disabled = 0
	RRCResume_v1800_IEs_NeedForInterruptionConfigNR_r18_Enabled  = 1
)

var rRCResumeV1800IEsNeedForInterruptionConfigNRR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCResume_v1800_IEs_NeedForInterruptionConfigNR_r18_Disabled, RRCResume_v1800_IEs_NeedForInterruptionConfigNR_r18_Enabled},
}

const (
	RRCResume_v1800_IEs_ReselectionMeasurementReq_r18_True = 0
)

var rRCResumeV1800IEsReselectionMeasurementReqR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCResume_v1800_IEs_ReselectionMeasurementReq_r18_True},
}

const (
	RRCResume_v1800_IEs_ValidatedMeasurementsReq_r18_True = 0
)

var rRCResumeV1800IEsValidatedMeasurementsReqR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCResume_v1800_IEs_ValidatedMeasurementsReq_r18_True},
}

type RRCResume_v1800_IEs struct {
	NeedForInterruptionConfigNR_r18 *int64
	ReselectionMeasurementReq_r18   *int64
	ValidatedMeasurementsReq_r18    *int64
	NonCriticalExtension            *RRCResume_v1910_IEs
}

func (ie *RRCResume_v1800_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCResumeV1800IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.NeedForInterruptionConfigNR_r18 != nil, ie.ReselectionMeasurementReq_r18 != nil, ie.ValidatedMeasurementsReq_r18 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. needForInterruptionConfigNR-r18: enumerated
	{
		if ie.NeedForInterruptionConfigNR_r18 != nil {
			if err := e.EncodeEnumerated(*ie.NeedForInterruptionConfigNR_r18, rRCResumeV1800IEsNeedForInterruptionConfigNRR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. reselectionMeasurementReq-r18: enumerated
	{
		if ie.ReselectionMeasurementReq_r18 != nil {
			if err := e.EncodeEnumerated(*ie.ReselectionMeasurementReq_r18, rRCResumeV1800IEsReselectionMeasurementReqR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. validatedMeasurementsReq-r18: enumerated
	{
		if ie.ValidatedMeasurementsReq_r18 != nil {
			if err := e.EncodeEnumerated(*ie.ValidatedMeasurementsReq_r18, rRCResumeV1800IEsValidatedMeasurementsReqR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RRCResume_v1800_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCResumeV1800IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. needForInterruptionConfigNR-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(rRCResumeV1800IEsNeedForInterruptionConfigNRR18Constraints)
			if err != nil {
				return err
			}
			ie.NeedForInterruptionConfigNR_r18 = &idx
		}
	}

	// 3. reselectionMeasurementReq-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(rRCResumeV1800IEsReselectionMeasurementReqR18Constraints)
			if err != nil {
				return err
			}
			ie.ReselectionMeasurementReq_r18 = &idx
		}
	}

	// 4. validatedMeasurementsReq-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(rRCResumeV1800IEsValidatedMeasurementsReqR18Constraints)
			if err != nil {
				return err
			}
			ie.ValidatedMeasurementsReq_r18 = &idx
		}
	}

	// 5. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(3) {
			ie.NonCriticalExtension = new(RRCResume_v1910_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
