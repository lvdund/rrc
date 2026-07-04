// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MRDC-SecondaryCellGroupConfig (line 1071).

var mRDCSecondaryCellGroupConfigConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mrdc-ReleaseAndAdd", Optional: true},
		{Name: "mrdc-SecondaryCellGroup"},
	},
}

const (
	MRDC_SecondaryCellGroupConfig_Mrdc_ReleaseAndAdd_True = 0
)

var mRDCSecondaryCellGroupConfigMrdcReleaseAndAddConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_SecondaryCellGroupConfig_Mrdc_ReleaseAndAdd_True},
}

var mRDC_SecondaryCellGroupConfigMrdcSecondaryCellGroupConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "nr-SCG"},
		{Name: "eutra-SCG"},
	},
}

const (
	MRDC_SecondaryCellGroupConfig_Mrdc_SecondaryCellGroup_Nr_SCG    = 0
	MRDC_SecondaryCellGroupConfig_Mrdc_SecondaryCellGroup_Eutra_SCG = 1
)

type MRDC_SecondaryCellGroupConfig struct {
	Mrdc_ReleaseAndAdd      *int64
	Mrdc_SecondaryCellGroup struct {
		Choice    int
		Nr_SCG    *[]byte
		Eutra_SCG *[]byte
	}
}

func (ie *MRDC_SecondaryCellGroupConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mRDCSecondaryCellGroupConfigConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mrdc_ReleaseAndAdd != nil}); err != nil {
		return err
	}

	// 2. mrdc-ReleaseAndAdd: enumerated
	{
		if ie.Mrdc_ReleaseAndAdd != nil {
			if err := e.EncodeEnumerated(*ie.Mrdc_ReleaseAndAdd, mRDCSecondaryCellGroupConfigMrdcReleaseAndAddConstraints); err != nil {
				return err
			}
		}
	}

	// 3. mrdc-SecondaryCellGroup: choice
	{
		choiceEnc := e.NewChoiceEncoder(mRDC_SecondaryCellGroupConfigMrdcSecondaryCellGroupConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Mrdc_SecondaryCellGroup.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Mrdc_SecondaryCellGroup.Choice {
		case MRDC_SecondaryCellGroupConfig_Mrdc_SecondaryCellGroup_Nr_SCG:
			if err := e.EncodeOctetString((*ie.Mrdc_SecondaryCellGroup.Nr_SCG), per.SizeConstraints{}); err != nil {
				return err
			}
		case MRDC_SecondaryCellGroupConfig_Mrdc_SecondaryCellGroup_Eutra_SCG:
			if err := e.EncodeOctetString((*ie.Mrdc_SecondaryCellGroup.Eutra_SCG), per.SizeConstraints{}); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Mrdc_SecondaryCellGroup.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *MRDC_SecondaryCellGroupConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mRDCSecondaryCellGroupConfigConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. mrdc-ReleaseAndAdd: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(mRDCSecondaryCellGroupConfigMrdcReleaseAndAddConstraints)
			if err != nil {
				return err
			}
			ie.Mrdc_ReleaseAndAdd = &idx
		}
	}

	// 3. mrdc-SecondaryCellGroup: choice
	{
		choiceDec := d.NewChoiceDecoder(mRDC_SecondaryCellGroupConfigMrdcSecondaryCellGroupConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Mrdc_SecondaryCellGroup.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case MRDC_SecondaryCellGroupConfig_Mrdc_SecondaryCellGroup_Nr_SCG:
			ie.Mrdc_SecondaryCellGroup.Nr_SCG = new([]byte)
			v, err := d.DecodeOctetString(per.SizeConstraints{})
			if err != nil {
				return err
			}
			(*ie.Mrdc_SecondaryCellGroup.Nr_SCG) = v
		case MRDC_SecondaryCellGroupConfig_Mrdc_SecondaryCellGroup_Eutra_SCG:
			ie.Mrdc_SecondaryCellGroup.Eutra_SCG = new([]byte)
			v, err := d.DecodeOctetString(per.SizeConstraints{})
			if err != nil {
				return err
			}
			(*ie.Mrdc_SecondaryCellGroup.Eutra_SCG) = v
		}
	}

	return nil
}
