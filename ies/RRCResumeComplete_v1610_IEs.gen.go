// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCResumeComplete-v1610-IEs (line 1605).

var rRCResumeCompleteV1610IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "idleMeasAvailable-r16", Optional: true},
		{Name: "measResultIdleEUTRA-r16", Optional: true},
		{Name: "measResultIdleNR-r16", Optional: true},
		{Name: "scg-Response-r16", Optional: true},
		{Name: "ue-MeasurementsAvailable-r16", Optional: true},
		{Name: "mobilityHistoryAvail-r16", Optional: true},
		{Name: "mobilityState-r16", Optional: true},
		{Name: "needForGapsInfoNR-r16", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	RRCResumeComplete_v1610_IEs_IdleMeasAvailable_r16_True = 0
)

var rRCResumeCompleteV1610IEsIdleMeasAvailableR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCResumeComplete_v1610_IEs_IdleMeasAvailable_r16_True},
}

var rRCResumeComplete_v1610_IEsScgResponseR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nr-SCG-Response"},
		{Name: "eutra-SCG-Response"},
	},
}

const (
	RRCResumeComplete_v1610_IEs_Scg_Response_r16_Nr_SCG_Response    = 0
	RRCResumeComplete_v1610_IEs_Scg_Response_r16_Eutra_SCG_Response = 1
)

const (
	RRCResumeComplete_v1610_IEs_MobilityHistoryAvail_r16_True = 0
)

var rRCResumeCompleteV1610IEsMobilityHistoryAvailR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCResumeComplete_v1610_IEs_MobilityHistoryAvail_r16_True},
}

const (
	RRCResumeComplete_v1610_IEs_MobilityState_r16_Normal = 0
	RRCResumeComplete_v1610_IEs_MobilityState_r16_Medium = 1
	RRCResumeComplete_v1610_IEs_MobilityState_r16_High   = 2
	RRCResumeComplete_v1610_IEs_MobilityState_r16_Spare  = 3
)

var rRCResumeCompleteV1610IEsMobilityStateR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCResumeComplete_v1610_IEs_MobilityState_r16_Normal, RRCResumeComplete_v1610_IEs_MobilityState_r16_Medium, RRCResumeComplete_v1610_IEs_MobilityState_r16_High, RRCResumeComplete_v1610_IEs_MobilityState_r16_Spare},
}

type RRCResumeComplete_v1610_IEs struct {
	IdleMeasAvailable_r16   *int64
	MeasResultIdleEUTRA_r16 *MeasResultIdleEUTRA_r16
	MeasResultIdleNR_r16    *MeasResultIdleNR_r16
	Scg_Response_r16        *struct {
		Choice             int
		Nr_SCG_Response    *[]byte
		Eutra_SCG_Response *[]byte
	}
	Ue_MeasurementsAvailable_r16 *UE_MeasurementsAvailable_r16
	MobilityHistoryAvail_r16     *int64
	MobilityState_r16            *int64
	NeedForGapsInfoNR_r16        *NeedForGapsInfoNR_r16
	NonCriticalExtension         *RRCResumeComplete_v1640_IEs
}

func (ie *RRCResumeComplete_v1610_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCResumeCompleteV1610IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.IdleMeasAvailable_r16 != nil, ie.MeasResultIdleEUTRA_r16 != nil, ie.MeasResultIdleNR_r16 != nil, ie.Scg_Response_r16 != nil, ie.Ue_MeasurementsAvailable_r16 != nil, ie.MobilityHistoryAvail_r16 != nil, ie.MobilityState_r16 != nil, ie.NeedForGapsInfoNR_r16 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. idleMeasAvailable-r16: enumerated
	{
		if ie.IdleMeasAvailable_r16 != nil {
			if err := e.EncodeEnumerated(*ie.IdleMeasAvailable_r16, rRCResumeCompleteV1610IEsIdleMeasAvailableR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. measResultIdleEUTRA-r16: ref
	{
		if ie.MeasResultIdleEUTRA_r16 != nil {
			if err := ie.MeasResultIdleEUTRA_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. measResultIdleNR-r16: ref
	{
		if ie.MeasResultIdleNR_r16 != nil {
			if err := ie.MeasResultIdleNR_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. scg-Response-r16: choice
	{
		if ie.Scg_Response_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCResumeComplete_v1610_IEsScgResponseR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Scg_Response_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Scg_Response_r16).Choice {
			case RRCResumeComplete_v1610_IEs_Scg_Response_r16_Nr_SCG_Response:
				if err := e.EncodeOctetString((*(*ie.Scg_Response_r16).Nr_SCG_Response), per.SizeConstraints{}); err != nil {
					return err
				}
			case RRCResumeComplete_v1610_IEs_Scg_Response_r16_Eutra_SCG_Response:
				if err := e.EncodeOctetString((*(*ie.Scg_Response_r16).Eutra_SCG_Response), per.SizeConstraints{}); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Scg_Response_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 6. ue-MeasurementsAvailable-r16: ref
	{
		if ie.Ue_MeasurementsAvailable_r16 != nil {
			if err := ie.Ue_MeasurementsAvailable_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. mobilityHistoryAvail-r16: enumerated
	{
		if ie.MobilityHistoryAvail_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MobilityHistoryAvail_r16, rRCResumeCompleteV1610IEsMobilityHistoryAvailR16Constraints); err != nil {
				return err
			}
		}
	}

	// 8. mobilityState-r16: enumerated
	{
		if ie.MobilityState_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MobilityState_r16, rRCResumeCompleteV1610IEsMobilityStateR16Constraints); err != nil {
				return err
			}
		}
	}

	// 9. needForGapsInfoNR-r16: ref
	{
		if ie.NeedForGapsInfoNR_r16 != nil {
			if err := ie.NeedForGapsInfoNR_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 10. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RRCResumeComplete_v1610_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCResumeCompleteV1610IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. idleMeasAvailable-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(rRCResumeCompleteV1610IEsIdleMeasAvailableR16Constraints)
			if err != nil {
				return err
			}
			ie.IdleMeasAvailable_r16 = &idx
		}
	}

	// 3. measResultIdleEUTRA-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.MeasResultIdleEUTRA_r16 = new(MeasResultIdleEUTRA_r16)
			if err := ie.MeasResultIdleEUTRA_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. measResultIdleNR-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.MeasResultIdleNR_r16 = new(MeasResultIdleNR_r16)
			if err := ie.MeasResultIdleNR_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. scg-Response-r16: choice
	{
		if seq.IsComponentPresent(3) {
			ie.Scg_Response_r16 = &struct {
				Choice             int
				Nr_SCG_Response    *[]byte
				Eutra_SCG_Response *[]byte
			}{}
			choiceDec := d.NewChoiceDecoder(rRCResumeComplete_v1610_IEsScgResponseR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Scg_Response_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCResumeComplete_v1610_IEs_Scg_Response_r16_Nr_SCG_Response:
				(*ie.Scg_Response_r16).Nr_SCG_Response = new([]byte)
				v, err := d.DecodeOctetString(per.SizeConstraints{})
				if err != nil {
					return err
				}
				(*(*ie.Scg_Response_r16).Nr_SCG_Response) = v
			case RRCResumeComplete_v1610_IEs_Scg_Response_r16_Eutra_SCG_Response:
				(*ie.Scg_Response_r16).Eutra_SCG_Response = new([]byte)
				v, err := d.DecodeOctetString(per.SizeConstraints{})
				if err != nil {
					return err
				}
				(*(*ie.Scg_Response_r16).Eutra_SCG_Response) = v
			}
		}
	}

	// 6. ue-MeasurementsAvailable-r16: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Ue_MeasurementsAvailable_r16 = new(UE_MeasurementsAvailable_r16)
			if err := ie.Ue_MeasurementsAvailable_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. mobilityHistoryAvail-r16: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(rRCResumeCompleteV1610IEsMobilityHistoryAvailR16Constraints)
			if err != nil {
				return err
			}
			ie.MobilityHistoryAvail_r16 = &idx
		}
	}

	// 8. mobilityState-r16: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(rRCResumeCompleteV1610IEsMobilityStateR16Constraints)
			if err != nil {
				return err
			}
			ie.MobilityState_r16 = &idx
		}
	}

	// 9. needForGapsInfoNR-r16: ref
	{
		if seq.IsComponentPresent(7) {
			ie.NeedForGapsInfoNR_r16 = new(NeedForGapsInfoNR_r16)
			if err := ie.NeedForGapsInfoNR_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 10. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(8) {
			ie.NonCriticalExtension = new(RRCResumeComplete_v1640_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
