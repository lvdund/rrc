// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SupportedCSI-RS-ResourceExt-r19 (line 19347).

var supportedCSIRSResourceExtR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberResourcesPerBand-r19"},
		{Name: "totalNumberTxPortsPerBand-r19"},
	},
}

var supportedCSI_RS_ResourceExt_r19MaxNumberResourcesPerBandR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "maxNumberResourcesPerBand1-r19"},
		{Name: "maxNumberResourcesPerBand2-r19"},
	},
}

const (
	SupportedCSI_RS_ResourceExt_r19_MaxNumberResourcesPerBand_r19_MaxNumberResourcesPerBand1_r19 = 0
	SupportedCSI_RS_ResourceExt_r19_MaxNumberResourcesPerBand_r19_MaxNumberResourcesPerBand2_r19 = 1
)

var supportedCSI_RS_ResourceExt_r19TotalNumberTxPortsPerBandR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "totalNumberTxPortsPerBand1-r19"},
		{Name: "totalNumberTxPortsPerBand2-r19"},
	},
}

const (
	SupportedCSI_RS_ResourceExt_r19_TotalNumberTxPortsPerBand_r19_TotalNumberTxPortsPerBand1_r19 = 0
	SupportedCSI_RS_ResourceExt_r19_TotalNumberTxPortsPerBand_r19_TotalNumberTxPortsPerBand2_r19 = 1
)

const (
	SupportedCSI_RS_ResourceExt_r19_MaxNumberResourcesPerBand_r19_MaxNumberResourcesPerBand2_r19_N128 = 0
	SupportedCSI_RS_ResourceExt_r19_MaxNumberResourcesPerBand_r19_MaxNumberResourcesPerBand2_r19_N256 = 1
)

var supportedCSIRSResourceExtR19MaxNumberResourcesPerBandR19MaxNumberResourcesPerBand2R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SupportedCSI_RS_ResourceExt_r19_MaxNumberResourcesPerBand_r19_MaxNumberResourcesPerBand2_r19_N128, SupportedCSI_RS_ResourceExt_r19_MaxNumberResourcesPerBand_r19_MaxNumberResourcesPerBand2_r19_N256},
}

const (
	SupportedCSI_RS_ResourceExt_r19_TotalNumberTxPortsPerBand_r19_TotalNumberTxPortsPerBand2_r19_N512  = 0
	SupportedCSI_RS_ResourceExt_r19_TotalNumberTxPortsPerBand_r19_TotalNumberTxPortsPerBand2_r19_N768  = 1
	SupportedCSI_RS_ResourceExt_r19_TotalNumberTxPortsPerBand_r19_TotalNumberTxPortsPerBand2_r19_N1024 = 2
)

var supportedCSIRSResourceExtR19TotalNumberTxPortsPerBandR19TotalNumberTxPortsPerBand2R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SupportedCSI_RS_ResourceExt_r19_TotalNumberTxPortsPerBand_r19_TotalNumberTxPortsPerBand2_r19_N512, SupportedCSI_RS_ResourceExt_r19_TotalNumberTxPortsPerBand_r19_TotalNumberTxPortsPerBand2_r19_N768, SupportedCSI_RS_ResourceExt_r19_TotalNumberTxPortsPerBand_r19_TotalNumberTxPortsPerBand2_r19_N1024},
}

type SupportedCSI_RS_ResourceExt_r19 struct {
	MaxNumberResourcesPerBand_r19 struct {
		Choice                         int
		MaxNumberResourcesPerBand1_r19 *int64
		MaxNumberResourcesPerBand2_r19 *int64
	}
	TotalNumberTxPortsPerBand_r19 struct {
		Choice                         int
		TotalNumberTxPortsPerBand1_r19 *int64
		TotalNumberTxPortsPerBand2_r19 *int64
	}
}

func (ie *SupportedCSI_RS_ResourceExt_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(supportedCSIRSResourceExtR19Constraints)
	_ = seq

	// 1. maxNumberResourcesPerBand-r19: choice
	{
		choiceEnc := e.NewChoiceEncoder(supportedCSI_RS_ResourceExt_r19MaxNumberResourcesPerBandR19Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.MaxNumberResourcesPerBand_r19.Choice), false, nil); err != nil {
			return err
		}
		switch ie.MaxNumberResourcesPerBand_r19.Choice {
		case SupportedCSI_RS_ResourceExt_r19_MaxNumberResourcesPerBand_r19_MaxNumberResourcesPerBand1_r19:
			if err := e.EncodeInteger((*ie.MaxNumberResourcesPerBand_r19.MaxNumberResourcesPerBand1_r19), per.Constrained(1, 64)); err != nil {
				return err
			}
		case SupportedCSI_RS_ResourceExt_r19_MaxNumberResourcesPerBand_r19_MaxNumberResourcesPerBand2_r19:
			if err := e.EncodeEnumerated((*ie.MaxNumberResourcesPerBand_r19.MaxNumberResourcesPerBand2_r19), supportedCSIRSResourceExtR19MaxNumberResourcesPerBandR19MaxNumberResourcesPerBand2R19Constraints); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.MaxNumberResourcesPerBand_r19.Choice), Constraint: "index out of range"}
		}
	}

	// 2. totalNumberTxPortsPerBand-r19: choice
	{
		choiceEnc := e.NewChoiceEncoder(supportedCSI_RS_ResourceExt_r19TotalNumberTxPortsPerBandR19Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.TotalNumberTxPortsPerBand_r19.Choice), false, nil); err != nil {
			return err
		}
		switch ie.TotalNumberTxPortsPerBand_r19.Choice {
		case SupportedCSI_RS_ResourceExt_r19_TotalNumberTxPortsPerBand_r19_TotalNumberTxPortsPerBand1_r19:
			if err := e.EncodeInteger((*ie.TotalNumberTxPortsPerBand_r19.TotalNumberTxPortsPerBand1_r19), per.Constrained(64, 256)); err != nil {
				return err
			}
		case SupportedCSI_RS_ResourceExt_r19_TotalNumberTxPortsPerBand_r19_TotalNumberTxPortsPerBand2_r19:
			if err := e.EncodeEnumerated((*ie.TotalNumberTxPortsPerBand_r19.TotalNumberTxPortsPerBand2_r19), supportedCSIRSResourceExtR19TotalNumberTxPortsPerBandR19TotalNumberTxPortsPerBand2R19Constraints); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.TotalNumberTxPortsPerBand_r19.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *SupportedCSI_RS_ResourceExt_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(supportedCSIRSResourceExtR19Constraints)
	_ = seq

	// 1. maxNumberResourcesPerBand-r19: choice
	{
		choiceDec := d.NewChoiceDecoder(supportedCSI_RS_ResourceExt_r19MaxNumberResourcesPerBandR19Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.MaxNumberResourcesPerBand_r19.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SupportedCSI_RS_ResourceExt_r19_MaxNumberResourcesPerBand_r19_MaxNumberResourcesPerBand1_r19:
			ie.MaxNumberResourcesPerBand_r19.MaxNumberResourcesPerBand1_r19 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(1, 64))
			if err != nil {
				return err
			}
			(*ie.MaxNumberResourcesPerBand_r19.MaxNumberResourcesPerBand1_r19) = v
		case SupportedCSI_RS_ResourceExt_r19_MaxNumberResourcesPerBand_r19_MaxNumberResourcesPerBand2_r19:
			ie.MaxNumberResourcesPerBand_r19.MaxNumberResourcesPerBand2_r19 = new(int64)
			v, err := d.DecodeEnumerated(supportedCSIRSResourceExtR19MaxNumberResourcesPerBandR19MaxNumberResourcesPerBand2R19Constraints)
			if err != nil {
				return err
			}
			(*ie.MaxNumberResourcesPerBand_r19.MaxNumberResourcesPerBand2_r19) = v
		}
	}

	// 2. totalNumberTxPortsPerBand-r19: choice
	{
		choiceDec := d.NewChoiceDecoder(supportedCSI_RS_ResourceExt_r19TotalNumberTxPortsPerBandR19Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.TotalNumberTxPortsPerBand_r19.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SupportedCSI_RS_ResourceExt_r19_TotalNumberTxPortsPerBand_r19_TotalNumberTxPortsPerBand1_r19:
			ie.TotalNumberTxPortsPerBand_r19.TotalNumberTxPortsPerBand1_r19 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(64, 256))
			if err != nil {
				return err
			}
			(*ie.TotalNumberTxPortsPerBand_r19.TotalNumberTxPortsPerBand1_r19) = v
		case SupportedCSI_RS_ResourceExt_r19_TotalNumberTxPortsPerBand_r19_TotalNumberTxPortsPerBand2_r19:
			ie.TotalNumberTxPortsPerBand_r19.TotalNumberTxPortsPerBand2_r19 = new(int64)
			v, err := d.DecodeEnumerated(supportedCSIRSResourceExtR19TotalNumberTxPortsPerBandR19TotalNumberTxPortsPerBand2R19Constraints)
			if err != nil {
				return err
			}
			(*ie.TotalNumberTxPortsPerBand_r19.TotalNumberTxPortsPerBand2_r19) = v
		}
	}

	return nil
}
