// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCReestablishmentComplete-v1610-IEs (line 926).

var rRCReestablishmentCompleteV1610IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ue-MeasurementsAvailable-r16", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type RRCReestablishmentComplete_v1610_IEs struct {
	Ue_MeasurementsAvailable_r16 *UE_MeasurementsAvailable_r16
	NonCriticalExtension         *RRCReestablishmentComplete_v1800_IEs
}

func (ie *RRCReestablishmentComplete_v1610_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReestablishmentCompleteV1610IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ue_MeasurementsAvailable_r16 != nil, ie.NonCriticalExtension != nil}); err != nil {
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

func (ie *RRCReestablishmentComplete_v1610_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReestablishmentCompleteV1610IEsConstraints)

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

	// 3. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = new(RRCReestablishmentComplete_v1800_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
