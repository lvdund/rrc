// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCReconfigurationComplete-v1800-IEs (line 1185).

var rRCReconfigurationCompleteV1800IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "needForInterruptionInfoNR-r18", Optional: true},
		{Name: "flightPathInfoAvailable-r18", Optional: true},
		{Name: "selectedPSCellForCHO-WithSCG-r18", Optional: true},
		{Name: "selectedSK-Counter-r18", Optional: true},
		{Name: "measConfigReportAppLayerAvailable-r18", Optional: true},
		{Name: "appliedLTM-CandidateId-r18", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	RRCReconfigurationComplete_v1800_IEs_FlightPathInfoAvailable_r18_True = 0
)

var rRCReconfigurationCompleteV1800IEsFlightPathInfoAvailableR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCReconfigurationComplete_v1800_IEs_FlightPathInfoAvailable_r18_True},
}

const (
	RRCReconfigurationComplete_v1800_IEs_MeasConfigReportAppLayerAvailable_r18_True = 0
)

var rRCReconfigurationCompleteV1800IEsMeasConfigReportAppLayerAvailableR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCReconfigurationComplete_v1800_IEs_MeasConfigReportAppLayerAvailable_r18_True},
}

type RRCReconfigurationComplete_v1800_IEs struct {
	NeedForInterruptionInfoNR_r18         *NeedForInterruptionInfoNR_r18
	FlightPathInfoAvailable_r18           *int64
	SelectedPSCellForCHO_WithSCG_r18      *SelectedPSCellForCHO_WithSCG_r18
	SelectedSK_Counter_r18                *SK_Counter
	MeasConfigReportAppLayerAvailable_r18 *int64
	AppliedLTM_CandidateId_r18            *LTM_CandidateId_r18
	NonCriticalExtension                  *RRCReconfigurationComplete_v1900_IEs
}

func (ie *RRCReconfigurationComplete_v1800_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReconfigurationCompleteV1800IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.NeedForInterruptionInfoNR_r18 != nil, ie.FlightPathInfoAvailable_r18 != nil, ie.SelectedPSCellForCHO_WithSCG_r18 != nil, ie.SelectedSK_Counter_r18 != nil, ie.MeasConfigReportAppLayerAvailable_r18 != nil, ie.AppliedLTM_CandidateId_r18 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. needForInterruptionInfoNR-r18: ref
	{
		if ie.NeedForInterruptionInfoNR_r18 != nil {
			if err := ie.NeedForInterruptionInfoNR_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. flightPathInfoAvailable-r18: enumerated
	{
		if ie.FlightPathInfoAvailable_r18 != nil {
			if err := e.EncodeEnumerated(*ie.FlightPathInfoAvailable_r18, rRCReconfigurationCompleteV1800IEsFlightPathInfoAvailableR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. selectedPSCellForCHO-WithSCG-r18: ref
	{
		if ie.SelectedPSCellForCHO_WithSCG_r18 != nil {
			if err := ie.SelectedPSCellForCHO_WithSCG_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. selectedSK-Counter-r18: ref
	{
		if ie.SelectedSK_Counter_r18 != nil {
			if err := ie.SelectedSK_Counter_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. measConfigReportAppLayerAvailable-r18: enumerated
	{
		if ie.MeasConfigReportAppLayerAvailable_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MeasConfigReportAppLayerAvailable_r18, rRCReconfigurationCompleteV1800IEsMeasConfigReportAppLayerAvailableR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. appliedLTM-CandidateId-r18: ref
	{
		if ie.AppliedLTM_CandidateId_r18 != nil {
			if err := ie.AppliedLTM_CandidateId_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RRCReconfigurationComplete_v1800_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReconfigurationCompleteV1800IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. needForInterruptionInfoNR-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.NeedForInterruptionInfoNR_r18 = new(NeedForInterruptionInfoNR_r18)
			if err := ie.NeedForInterruptionInfoNR_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. flightPathInfoAvailable-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(rRCReconfigurationCompleteV1800IEsFlightPathInfoAvailableR18Constraints)
			if err != nil {
				return err
			}
			ie.FlightPathInfoAvailable_r18 = &idx
		}
	}

	// 4. selectedPSCellForCHO-WithSCG-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.SelectedPSCellForCHO_WithSCG_r18 = new(SelectedPSCellForCHO_WithSCG_r18)
			if err := ie.SelectedPSCellForCHO_WithSCG_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. selectedSK-Counter-r18: ref
	{
		if seq.IsComponentPresent(3) {
			ie.SelectedSK_Counter_r18 = new(SK_Counter)
			if err := ie.SelectedSK_Counter_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. measConfigReportAppLayerAvailable-r18: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(rRCReconfigurationCompleteV1800IEsMeasConfigReportAppLayerAvailableR18Constraints)
			if err != nil {
				return err
			}
			ie.MeasConfigReportAppLayerAvailable_r18 = &idx
		}
	}

	// 7. appliedLTM-CandidateId-r18: ref
	{
		if seq.IsComponentPresent(5) {
			ie.AppliedLTM_CandidateId_r18 = new(LTM_CandidateId_r18)
			if err := ie.AppliedLTM_CandidateId_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(6) {
			ie.NonCriticalExtension = new(RRCReconfigurationComplete_v1900_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
