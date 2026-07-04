// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCResume-v1610-IEs (line 1552).

var rRCResumeV1610IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "idleModeMeasurementReq-r16", Optional: true},
		{Name: "restoreMCG-SCells-r16", Optional: true},
		{Name: "restoreSCG-r16", Optional: true},
		{Name: "mrdc-SecondaryCellGroup-r16", Optional: true},
		{Name: "needForGapsConfigNR-r16", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	RRCResume_v1610_IEs_IdleModeMeasurementReq_r16_True = 0
)

var rRCResumeV1610IEsIdleModeMeasurementReqR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCResume_v1610_IEs_IdleModeMeasurementReq_r16_True},
}

const (
	RRCResume_v1610_IEs_RestoreMCG_SCells_r16_True = 0
)

var rRCResumeV1610IEsRestoreMCGSCellsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCResume_v1610_IEs_RestoreMCG_SCells_r16_True},
}

const (
	RRCResume_v1610_IEs_RestoreSCG_r16_True = 0
)

var rRCResumeV1610IEsRestoreSCGR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCResume_v1610_IEs_RestoreSCG_r16_True},
}

var rRCResume_v1610_IEsMrdcSecondaryCellGroupR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nr-SCG-r16"},
		{Name: "eutra-SCG-r16"},
	},
}

const (
	RRCResume_v1610_IEs_Mrdc_SecondaryCellGroup_r16_Nr_SCG_r16    = 0
	RRCResume_v1610_IEs_Mrdc_SecondaryCellGroup_r16_Eutra_SCG_r16 = 1
)

var rRCResume_v1610_IEsNeedForGapsConfigNRR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCResume_v1610_IEs_NeedForGapsConfigNR_r16_Release = 0
	RRCResume_v1610_IEs_NeedForGapsConfigNR_r16_Setup   = 1
)

type RRCResume_v1610_IEs struct {
	IdleModeMeasurementReq_r16  *int64
	RestoreMCG_SCells_r16       *int64
	RestoreSCG_r16              *int64
	Mrdc_SecondaryCellGroup_r16 *struct {
		Choice        int
		Nr_SCG_r16    *[]byte
		Eutra_SCG_r16 *[]byte
	}
	NeedForGapsConfigNR_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *NeedForGapsConfigNR_r16
	}
	NonCriticalExtension *RRCResume_v1700_IEs
}

func (ie *RRCResume_v1610_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCResumeV1610IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.IdleModeMeasurementReq_r16 != nil, ie.RestoreMCG_SCells_r16 != nil, ie.RestoreSCG_r16 != nil, ie.Mrdc_SecondaryCellGroup_r16 != nil, ie.NeedForGapsConfigNR_r16 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. idleModeMeasurementReq-r16: enumerated
	{
		if ie.IdleModeMeasurementReq_r16 != nil {
			if err := e.EncodeEnumerated(*ie.IdleModeMeasurementReq_r16, rRCResumeV1610IEsIdleModeMeasurementReqR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. restoreMCG-SCells-r16: enumerated
	{
		if ie.RestoreMCG_SCells_r16 != nil {
			if err := e.EncodeEnumerated(*ie.RestoreMCG_SCells_r16, rRCResumeV1610IEsRestoreMCGSCellsR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. restoreSCG-r16: enumerated
	{
		if ie.RestoreSCG_r16 != nil {
			if err := e.EncodeEnumerated(*ie.RestoreSCG_r16, rRCResumeV1610IEsRestoreSCGR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. mrdc-SecondaryCellGroup-r16: choice
	{
		if ie.Mrdc_SecondaryCellGroup_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCResume_v1610_IEsMrdcSecondaryCellGroupR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Mrdc_SecondaryCellGroup_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Mrdc_SecondaryCellGroup_r16).Choice {
			case RRCResume_v1610_IEs_Mrdc_SecondaryCellGroup_r16_Nr_SCG_r16:
				if err := e.EncodeOctetString((*(*ie.Mrdc_SecondaryCellGroup_r16).Nr_SCG_r16), per.SizeConstraints{}); err != nil {
					return err
				}
			case RRCResume_v1610_IEs_Mrdc_SecondaryCellGroup_r16_Eutra_SCG_r16:
				if err := e.EncodeOctetString((*(*ie.Mrdc_SecondaryCellGroup_r16).Eutra_SCG_r16), per.SizeConstraints{}); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Mrdc_SecondaryCellGroup_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 6. needForGapsConfigNR-r16: choice
	{
		if ie.NeedForGapsConfigNR_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCResume_v1610_IEsNeedForGapsConfigNRR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.NeedForGapsConfigNR_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.NeedForGapsConfigNR_r16).Choice {
			case RRCResume_v1610_IEs_NeedForGapsConfigNR_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCResume_v1610_IEs_NeedForGapsConfigNR_r16_Setup:
				if err := (*ie.NeedForGapsConfigNR_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.NeedForGapsConfigNR_r16).Choice), Constraint: "index out of range"}
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

func (ie *RRCResume_v1610_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCResumeV1610IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. idleModeMeasurementReq-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(rRCResumeV1610IEsIdleModeMeasurementReqR16Constraints)
			if err != nil {
				return err
			}
			ie.IdleModeMeasurementReq_r16 = &idx
		}
	}

	// 3. restoreMCG-SCells-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(rRCResumeV1610IEsRestoreMCGSCellsR16Constraints)
			if err != nil {
				return err
			}
			ie.RestoreMCG_SCells_r16 = &idx
		}
	}

	// 4. restoreSCG-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(rRCResumeV1610IEsRestoreSCGR16Constraints)
			if err != nil {
				return err
			}
			ie.RestoreSCG_r16 = &idx
		}
	}

	// 5. mrdc-SecondaryCellGroup-r16: choice
	{
		if seq.IsComponentPresent(3) {
			ie.Mrdc_SecondaryCellGroup_r16 = &struct {
				Choice        int
				Nr_SCG_r16    *[]byte
				Eutra_SCG_r16 *[]byte
			}{}
			choiceDec := d.NewChoiceDecoder(rRCResume_v1610_IEsMrdcSecondaryCellGroupR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Mrdc_SecondaryCellGroup_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCResume_v1610_IEs_Mrdc_SecondaryCellGroup_r16_Nr_SCG_r16:
				(*ie.Mrdc_SecondaryCellGroup_r16).Nr_SCG_r16 = new([]byte)
				v, err := d.DecodeOctetString(per.SizeConstraints{})
				if err != nil {
					return err
				}
				(*(*ie.Mrdc_SecondaryCellGroup_r16).Nr_SCG_r16) = v
			case RRCResume_v1610_IEs_Mrdc_SecondaryCellGroup_r16_Eutra_SCG_r16:
				(*ie.Mrdc_SecondaryCellGroup_r16).Eutra_SCG_r16 = new([]byte)
				v, err := d.DecodeOctetString(per.SizeConstraints{})
				if err != nil {
					return err
				}
				(*(*ie.Mrdc_SecondaryCellGroup_r16).Eutra_SCG_r16) = v
			}
		}
	}

	// 6. needForGapsConfigNR-r16: choice
	{
		if seq.IsComponentPresent(4) {
			ie.NeedForGapsConfigNR_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *NeedForGapsConfigNR_r16
			}{}
			choiceDec := d.NewChoiceDecoder(rRCResume_v1610_IEsNeedForGapsConfigNRR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.NeedForGapsConfigNR_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCResume_v1610_IEs_NeedForGapsConfigNR_r16_Release:
				(*ie.NeedForGapsConfigNR_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCResume_v1610_IEs_NeedForGapsConfigNR_r16_Setup:
				(*ie.NeedForGapsConfigNR_r16).Setup = new(NeedForGapsConfigNR_r16)
				if err := (*ie.NeedForGapsConfigNR_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(5) {
			ie.NonCriticalExtension = new(RRCResume_v1700_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
