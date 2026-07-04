// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCReconfigurationComplete-v1610-IEs (line 1162).

var rRCReconfigurationCompleteV1610IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ue-MeasurementsAvailable-r16", Optional: true},
		{Name: "needForGapsInfoNR-r16", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type RRCReconfigurationComplete_v1610_IEs struct {
	Ue_MeasurementsAvailable_r16 *UE_MeasurementsAvailable_r16
	NeedForGapsInfoNR_r16        *NeedForGapsInfoNR_r16
	NonCriticalExtension         *RRCReconfigurationComplete_v1640_IEs
}

func (ie *RRCReconfigurationComplete_v1610_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReconfigurationCompleteV1610IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ue_MeasurementsAvailable_r16 != nil, ie.NeedForGapsInfoNR_r16 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. ue-MeasurementsAvailable-r16: ref
	{
		if ie.Ue_MeasurementsAvailable_r16 != nil {
			if err := ie.Ue_MeasurementsAvailable_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. needForGapsInfoNR-r16: ref
	{
		if ie.NeedForGapsInfoNR_r16 != nil {
			if err := ie.NeedForGapsInfoNR_r16.Encode(e); err != nil {
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

func (ie *RRCReconfigurationComplete_v1610_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReconfigurationCompleteV1610IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ue-MeasurementsAvailable-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ue_MeasurementsAvailable_r16 = new(UE_MeasurementsAvailable_r16)
			if err := ie.Ue_MeasurementsAvailable_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. needForGapsInfoNR-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NeedForGapsInfoNR_r16 = new(NeedForGapsInfoNR_r16)
			if err := ie.NeedForGapsInfoNR_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(2) {
			ie.NonCriticalExtension = new(RRCReconfigurationComplete_v1640_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
