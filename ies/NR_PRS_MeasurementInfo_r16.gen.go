// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NR-PRS-MeasurementInfo-r16 (line 8577).

var nRPRSMeasurementInfoR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dl-PRS-PointA-r16"},
		{Name: "nr-MeasPRS-RepetitionAndOffset-r16"},
		{Name: "nr-MeasPRS-length-r16"},
	},
}

var nR_PRS_MeasurementInfo_r16NrMeasPRSRepetitionAndOffsetR16Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ms20-r16"},
		{Name: "ms40-r16"},
		{Name: "ms80-r16"},
		{Name: "ms160-r16"},
	},
}

const (
	NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_RepetitionAndOffset_r16_Ms20_r16  = 0
	NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_RepetitionAndOffset_r16_Ms40_r16  = 1
	NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_RepetitionAndOffset_r16_Ms80_r16  = 2
	NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_RepetitionAndOffset_r16_Ms160_r16 = 3
)

const (
	NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_Length_r16_Ms1dot5 = 0
	NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_Length_r16_Ms3     = 1
	NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_Length_r16_Ms3dot5 = 2
	NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_Length_r16_Ms4     = 3
	NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_Length_r16_Ms5dot5 = 4
	NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_Length_r16_Ms6     = 5
	NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_Length_r16_Ms10    = 6
	NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_Length_r16_Ms20    = 7
)

var nRPRSMeasurementInfoR16NrMeasPRSLengthR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_Length_r16_Ms1dot5, NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_Length_r16_Ms3, NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_Length_r16_Ms3dot5, NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_Length_r16_Ms4, NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_Length_r16_Ms5dot5, NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_Length_r16_Ms6, NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_Length_r16_Ms10, NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_Length_r16_Ms20},
}

type NR_PRS_MeasurementInfo_r16 struct {
	Dl_PRS_PointA_r16                  ARFCN_ValueNR
	Nr_MeasPRS_RepetitionAndOffset_r16 struct {
		Choice    int
		Ms20_r16  *int64
		Ms40_r16  *int64
		Ms80_r16  *int64
		Ms160_r16 *int64
	}
	Nr_MeasPRS_Length_r16 int64
}

func (ie *NR_PRS_MeasurementInfo_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRPRSMeasurementInfoR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. dl-PRS-PointA-r16: ref
	{
		if err := ie.Dl_PRS_PointA_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. nr-MeasPRS-RepetitionAndOffset-r16: choice
	{
		choiceEnc := e.NewChoiceEncoder(nR_PRS_MeasurementInfo_r16NrMeasPRSRepetitionAndOffsetR16Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Nr_MeasPRS_RepetitionAndOffset_r16.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Nr_MeasPRS_RepetitionAndOffset_r16.Choice {
		case NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_RepetitionAndOffset_r16_Ms20_r16:
			if err := e.EncodeInteger((*ie.Nr_MeasPRS_RepetitionAndOffset_r16.Ms20_r16), per.Constrained(0, 19)); err != nil {
				return err
			}
		case NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_RepetitionAndOffset_r16_Ms40_r16:
			if err := e.EncodeInteger((*ie.Nr_MeasPRS_RepetitionAndOffset_r16.Ms40_r16), per.Constrained(0, 39)); err != nil {
				return err
			}
		case NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_RepetitionAndOffset_r16_Ms80_r16:
			if err := e.EncodeInteger((*ie.Nr_MeasPRS_RepetitionAndOffset_r16.Ms80_r16), per.Constrained(0, 79)); err != nil {
				return err
			}
		case NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_RepetitionAndOffset_r16_Ms160_r16:
			if err := e.EncodeInteger((*ie.Nr_MeasPRS_RepetitionAndOffset_r16.Ms160_r16), per.Constrained(0, 159)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Nr_MeasPRS_RepetitionAndOffset_r16.Choice), Constraint: "index out of range"}
		}
	}

	// 4. nr-MeasPRS-length-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Nr_MeasPRS_Length_r16, nRPRSMeasurementInfoR16NrMeasPRSLengthR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *NR_PRS_MeasurementInfo_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRPRSMeasurementInfoR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. dl-PRS-PointA-r16: ref
	{
		if err := ie.Dl_PRS_PointA_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. nr-MeasPRS-RepetitionAndOffset-r16: choice
	{
		choiceDec := d.NewChoiceDecoder(nR_PRS_MeasurementInfo_r16NrMeasPRSRepetitionAndOffsetR16Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Nr_MeasPRS_RepetitionAndOffset_r16.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_RepetitionAndOffset_r16_Ms20_r16:
			ie.Nr_MeasPRS_RepetitionAndOffset_r16.Ms20_r16 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 19))
			if err != nil {
				return err
			}
			(*ie.Nr_MeasPRS_RepetitionAndOffset_r16.Ms20_r16) = v
		case NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_RepetitionAndOffset_r16_Ms40_r16:
			ie.Nr_MeasPRS_RepetitionAndOffset_r16.Ms40_r16 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 39))
			if err != nil {
				return err
			}
			(*ie.Nr_MeasPRS_RepetitionAndOffset_r16.Ms40_r16) = v
		case NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_RepetitionAndOffset_r16_Ms80_r16:
			ie.Nr_MeasPRS_RepetitionAndOffset_r16.Ms80_r16 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 79))
			if err != nil {
				return err
			}
			(*ie.Nr_MeasPRS_RepetitionAndOffset_r16.Ms80_r16) = v
		case NR_PRS_MeasurementInfo_r16_Nr_MeasPRS_RepetitionAndOffset_r16_Ms160_r16:
			ie.Nr_MeasPRS_RepetitionAndOffset_r16.Ms160_r16 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 159))
			if err != nil {
				return err
			}
			(*ie.Nr_MeasPRS_RepetitionAndOffset_r16.Ms160_r16) = v
		}
	}

	// 4. nr-MeasPRS-length-r16: enumerated
	{
		v2, err := d.DecodeEnumerated(nRPRSMeasurementInfoR16NrMeasPRSLengthR16Constraints)
		if err != nil {
			return err
		}
		ie.Nr_MeasPRS_Length_r16 = v2
	}

	return nil
}
