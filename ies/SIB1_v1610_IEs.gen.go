// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB1-v1610-IEs (line 2009).

var sIB1V1610IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "idleModeMeasurementsEUTRA-r16", Optional: true},
		{Name: "idleModeMeasurementsNR-r16", Optional: true},
		{Name: "posSI-SchedulingInfo-r16", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	SIB1_v1610_IEs_IdleModeMeasurementsEUTRA_r16_True = 0
)

var sIB1V1610IEsIdleModeMeasurementsEUTRAR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_v1610_IEs_IdleModeMeasurementsEUTRA_r16_True},
}

const (
	SIB1_v1610_IEs_IdleModeMeasurementsNR_r16_True = 0
)

var sIB1V1610IEsIdleModeMeasurementsNRR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_v1610_IEs_IdleModeMeasurementsNR_r16_True},
}

type SIB1_v1610_IEs struct {
	IdleModeMeasurementsEUTRA_r16 *int64
	IdleModeMeasurementsNR_r16    *int64
	PosSI_SchedulingInfo_r16      *PosSI_SchedulingInfo_r16
	NonCriticalExtension          *SIB1_v1630_IEs
}

func (ie *SIB1_v1610_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB1V1610IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.IdleModeMeasurementsEUTRA_r16 != nil, ie.IdleModeMeasurementsNR_r16 != nil, ie.PosSI_SchedulingInfo_r16 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. idleModeMeasurementsEUTRA-r16: enumerated
	{
		if ie.IdleModeMeasurementsEUTRA_r16 != nil {
			if err := e.EncodeEnumerated(*ie.IdleModeMeasurementsEUTRA_r16, sIB1V1610IEsIdleModeMeasurementsEUTRAR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. idleModeMeasurementsNR-r16: enumerated
	{
		if ie.IdleModeMeasurementsNR_r16 != nil {
			if err := e.EncodeEnumerated(*ie.IdleModeMeasurementsNR_r16, sIB1V1610IEsIdleModeMeasurementsNRR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. posSI-SchedulingInfo-r16: ref
	{
		if ie.PosSI_SchedulingInfo_r16 != nil {
			if err := ie.PosSI_SchedulingInfo_r16.Encode(e); err != nil {
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

func (ie *SIB1_v1610_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB1V1610IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. idleModeMeasurementsEUTRA-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sIB1V1610IEsIdleModeMeasurementsEUTRAR16Constraints)
			if err != nil {
				return err
			}
			ie.IdleModeMeasurementsEUTRA_r16 = &idx
		}
	}

	// 3. idleModeMeasurementsNR-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sIB1V1610IEsIdleModeMeasurementsNRR16Constraints)
			if err != nil {
				return err
			}
			ie.IdleModeMeasurementsNR_r16 = &idx
		}
	}

	// 4. posSI-SchedulingInfo-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.PosSI_SchedulingInfo_r16 = new(PosSI_SchedulingInfo_r16)
			if err := ie.PosSI_SchedulingInfo_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(3) {
			ie.NonCriticalExtension = new(SIB1_v1630_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
