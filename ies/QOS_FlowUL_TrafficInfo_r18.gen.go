// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: QOS-FlowUL-TrafficInfo-r18 (line 2768).

var qOSFlowULTrafficInfoR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "qfi-r18"},
		{Name: "jitterRange-r18", Optional: true},
		{Name: "burstArrivalTime-r18", Optional: true},
		{Name: "trafficPeriodicity-r18", Optional: true},
		{Name: "pdu-SetIdentification-r18", Optional: true},
		{Name: "psi-Identification-r18", Optional: true},
	},
}

var qOS_FlowUL_TrafficInfo_r18BurstArrivalTimeR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "referenceTime"},
		{Name: "referenceSFN-AndSlot"},
	},
}

const (
	QOS_FlowUL_TrafficInfo_r18_BurstArrivalTime_r18_ReferenceTime        = 0
	QOS_FlowUL_TrafficInfo_r18_BurstArrivalTime_r18_ReferenceSFN_AndSlot = 1
)

var qOSFlowULTrafficInfoR18TrafficPeriodicityR18Constraints = per.Constrained(1, 640000)

type QOS_FlowUL_TrafficInfo_r18 struct {
	Qfi_r18         QFI
	JitterRange_r18 *struct {
		LowerBound_r18 JitterBound_r18
		UpperBound_r18 JitterBound_r18
	}
	BurstArrivalTime_r18 *struct {
		Choice               int
		ReferenceTime        *ReferenceTime_r16
		ReferenceSFN_AndSlot *ReferenceSFN_AndSlot_r18
	}
	TrafficPeriodicity_r18    *int64
	Pdu_SetIdentification_r18 *bool
	Psi_Identification_r18    *bool
}

func (ie *QOS_FlowUL_TrafficInfo_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(qOSFlowULTrafficInfoR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.JitterRange_r18 != nil, ie.BurstArrivalTime_r18 != nil, ie.TrafficPeriodicity_r18 != nil, ie.Pdu_SetIdentification_r18 != nil, ie.Psi_Identification_r18 != nil}); err != nil {
		return err
	}

	// 3. qfi-r18: ref
	{
		if err := ie.Qfi_r18.Encode(e); err != nil {
			return err
		}
	}

	// 4. jitterRange-r18: sequence
	{
		if ie.JitterRange_r18 != nil {
			c := ie.JitterRange_r18
			if err := c.LowerBound_r18.Encode(e); err != nil {
				return err
			}
			if err := c.UpperBound_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. burstArrivalTime-r18: choice
	{
		if ie.BurstArrivalTime_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(qOS_FlowUL_TrafficInfo_r18BurstArrivalTimeR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.BurstArrivalTime_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.BurstArrivalTime_r18).Choice {
			case QOS_FlowUL_TrafficInfo_r18_BurstArrivalTime_r18_ReferenceTime:
				if err := (*ie.BurstArrivalTime_r18).ReferenceTime.Encode(e); err != nil {
					return err
				}
			case QOS_FlowUL_TrafficInfo_r18_BurstArrivalTime_r18_ReferenceSFN_AndSlot:
				if err := (*ie.BurstArrivalTime_r18).ReferenceSFN_AndSlot.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.BurstArrivalTime_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 6. trafficPeriodicity-r18: integer
	{
		if ie.TrafficPeriodicity_r18 != nil {
			if err := e.EncodeInteger(*ie.TrafficPeriodicity_r18, qOSFlowULTrafficInfoR18TrafficPeriodicityR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. pdu-SetIdentification-r18: boolean
	{
		if ie.Pdu_SetIdentification_r18 != nil {
			if err := e.EncodeBoolean(*ie.Pdu_SetIdentification_r18); err != nil {
				return err
			}
		}
	}

	// 8. psi-Identification-r18: boolean
	{
		if ie.Psi_Identification_r18 != nil {
			if err := e.EncodeBoolean(*ie.Psi_Identification_r18); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *QOS_FlowUL_TrafficInfo_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(qOSFlowULTrafficInfoR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. qfi-r18: ref
	{
		if err := ie.Qfi_r18.Decode(d); err != nil {
			return err
		}
	}

	// 4. jitterRange-r18: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.JitterRange_r18 = &struct {
				LowerBound_r18 JitterBound_r18
				UpperBound_r18 JitterBound_r18
			}{}
			c := ie.JitterRange_r18
			{
				if err := c.LowerBound_r18.Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.UpperBound_r18.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. burstArrivalTime-r18: choice
	{
		if seq.IsComponentPresent(2) {
			ie.BurstArrivalTime_r18 = &struct {
				Choice               int
				ReferenceTime        *ReferenceTime_r16
				ReferenceSFN_AndSlot *ReferenceSFN_AndSlot_r18
			}{}
			choiceDec := d.NewChoiceDecoder(qOS_FlowUL_TrafficInfo_r18BurstArrivalTimeR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.BurstArrivalTime_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case QOS_FlowUL_TrafficInfo_r18_BurstArrivalTime_r18_ReferenceTime:
				(*ie.BurstArrivalTime_r18).ReferenceTime = new(ReferenceTime_r16)
				if err := (*ie.BurstArrivalTime_r18).ReferenceTime.Decode(d); err != nil {
					return err
				}
			case QOS_FlowUL_TrafficInfo_r18_BurstArrivalTime_r18_ReferenceSFN_AndSlot:
				(*ie.BurstArrivalTime_r18).ReferenceSFN_AndSlot = new(ReferenceSFN_AndSlot_r18)
				if err := (*ie.BurstArrivalTime_r18).ReferenceSFN_AndSlot.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. trafficPeriodicity-r18: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(qOSFlowULTrafficInfoR18TrafficPeriodicityR18Constraints)
			if err != nil {
				return err
			}
			ie.TrafficPeriodicity_r18 = &v
		}
	}

	// 7. pdu-SetIdentification-r18: boolean
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.Pdu_SetIdentification_r18 = &v
		}
	}

	// 8. psi-Identification-r18: boolean
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.Psi_Identification_r18 = &v
		}
	}

	return nil
}
