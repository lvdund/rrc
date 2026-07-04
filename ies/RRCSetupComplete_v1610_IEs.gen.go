// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCSetupComplete-v1610-IEs (line 1730).

var rRCSetupCompleteV1610IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "iab-NodeIndication-r16", Optional: true},
		{Name: "idleMeasAvailable-r16", Optional: true},
		{Name: "ue-MeasurementsAvailable-r16", Optional: true},
		{Name: "mobilityHistoryAvail-r16", Optional: true},
		{Name: "mobilityState-r16", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	RRCSetupComplete_v1610_IEs_Iab_NodeIndication_r16_True = 0
)

var rRCSetupCompleteV1610IEsIabNodeIndicationR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCSetupComplete_v1610_IEs_Iab_NodeIndication_r16_True},
}

const (
	RRCSetupComplete_v1610_IEs_IdleMeasAvailable_r16_True = 0
)

var rRCSetupCompleteV1610IEsIdleMeasAvailableR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCSetupComplete_v1610_IEs_IdleMeasAvailable_r16_True},
}

const (
	RRCSetupComplete_v1610_IEs_MobilityHistoryAvail_r16_True = 0
)

var rRCSetupCompleteV1610IEsMobilityHistoryAvailR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCSetupComplete_v1610_IEs_MobilityHistoryAvail_r16_True},
}

const (
	RRCSetupComplete_v1610_IEs_MobilityState_r16_Normal = 0
	RRCSetupComplete_v1610_IEs_MobilityState_r16_Medium = 1
	RRCSetupComplete_v1610_IEs_MobilityState_r16_High   = 2
	RRCSetupComplete_v1610_IEs_MobilityState_r16_Spare  = 3
)

var rRCSetupCompleteV1610IEsMobilityStateR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCSetupComplete_v1610_IEs_MobilityState_r16_Normal, RRCSetupComplete_v1610_IEs_MobilityState_r16_Medium, RRCSetupComplete_v1610_IEs_MobilityState_r16_High, RRCSetupComplete_v1610_IEs_MobilityState_r16_Spare},
}

type RRCSetupComplete_v1610_IEs struct {
	Iab_NodeIndication_r16       *int64
	IdleMeasAvailable_r16        *int64
	Ue_MeasurementsAvailable_r16 *UE_MeasurementsAvailable_r16
	MobilityHistoryAvail_r16     *int64
	MobilityState_r16            *int64
	NonCriticalExtension         *RRCSetupComplete_v1690_IEs
}

func (ie *RRCSetupComplete_v1610_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCSetupCompleteV1610IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Iab_NodeIndication_r16 != nil, ie.IdleMeasAvailable_r16 != nil, ie.Ue_MeasurementsAvailable_r16 != nil, ie.MobilityHistoryAvail_r16 != nil, ie.MobilityState_r16 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. iab-NodeIndication-r16: enumerated
	{
		if ie.Iab_NodeIndication_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Iab_NodeIndication_r16, rRCSetupCompleteV1610IEsIabNodeIndicationR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. idleMeasAvailable-r16: enumerated
	{
		if ie.IdleMeasAvailable_r16 != nil {
			if err := e.EncodeEnumerated(*ie.IdleMeasAvailable_r16, rRCSetupCompleteV1610IEsIdleMeasAvailableR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. ue-MeasurementsAvailable-r16: ref
	{
		if ie.Ue_MeasurementsAvailable_r16 != nil {
			if err := ie.Ue_MeasurementsAvailable_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. mobilityHistoryAvail-r16: enumerated
	{
		if ie.MobilityHistoryAvail_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MobilityHistoryAvail_r16, rRCSetupCompleteV1610IEsMobilityHistoryAvailR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. mobilityState-r16: enumerated
	{
		if ie.MobilityState_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MobilityState_r16, rRCSetupCompleteV1610IEsMobilityStateR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RRCSetupComplete_v1610_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCSetupCompleteV1610IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. iab-NodeIndication-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(rRCSetupCompleteV1610IEsIabNodeIndicationR16Constraints)
			if err != nil {
				return err
			}
			ie.Iab_NodeIndication_r16 = &idx
		}
	}

	// 3. idleMeasAvailable-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(rRCSetupCompleteV1610IEsIdleMeasAvailableR16Constraints)
			if err != nil {
				return err
			}
			ie.IdleMeasAvailable_r16 = &idx
		}
	}

	// 4. ue-MeasurementsAvailable-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Ue_MeasurementsAvailable_r16 = new(UE_MeasurementsAvailable_r16)
			if err := ie.Ue_MeasurementsAvailable_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. mobilityHistoryAvail-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(rRCSetupCompleteV1610IEsMobilityHistoryAvailR16Constraints)
			if err != nil {
				return err
			}
			ie.MobilityHistoryAvail_r16 = &idx
		}
	}

	// 6. mobilityState-r16: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(rRCSetupCompleteV1610IEsMobilityStateR16Constraints)
			if err != nil {
				return err
			}
			ie.MobilityState_r16 = &idx
		}
	}

	// 7. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(5) {
			ie.NonCriticalExtension = new(RRCSetupComplete_v1690_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
