// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SupportedCSI-RS-ResourceAggregate-r19 (line 19358).

var supportedCSIRSResourceAggregateR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberTxPortsPerAggregatedResource-r19"},
		{Name: "maxNumberAggregatedResources-r19"},
		{Name: "totalNumberTxPorts-r19"},
	},
}

const (
	SupportedCSI_RS_ResourceAggregate_r19_MaxNumberTxPortsPerAggregatedResource_r19_P48  = 0
	SupportedCSI_RS_ResourceAggregate_r19_MaxNumberTxPortsPerAggregatedResource_r19_P64  = 1
	SupportedCSI_RS_ResourceAggregate_r19_MaxNumberTxPortsPerAggregatedResource_r19_P128 = 2
)

var supportedCSIRSResourceAggregateR19MaxNumberTxPortsPerAggregatedResourceR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SupportedCSI_RS_ResourceAggregate_r19_MaxNumberTxPortsPerAggregatedResource_r19_P48, SupportedCSI_RS_ResourceAggregate_r19_MaxNumberTxPortsPerAggregatedResource_r19_P64, SupportedCSI_RS_ResourceAggregate_r19_MaxNumberTxPortsPerAggregatedResource_r19_P128},
}

var supportedCSI_RS_ResourceAggregate_r19MaxNumberAggregatedResourcesR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "maxNumberAggregatedResources1-r19"},
		{Name: "maxNumberAggregatedResources2-r19"},
	},
}

const (
	SupportedCSI_RS_ResourceAggregate_r19_MaxNumberAggregatedResources_r19_MaxNumberAggregatedResources1_r19 = 0
	SupportedCSI_RS_ResourceAggregate_r19_MaxNumberAggregatedResources_r19_MaxNumberAggregatedResources2_r19 = 1
)

var supportedCSI_RS_ResourceAggregate_r19TotalNumberTxPortsR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "totalNumberTxPorts1-r19"},
		{Name: "totalNumberTxPorts2-r19"},
	},
}

const (
	SupportedCSI_RS_ResourceAggregate_r19_TotalNumberTxPorts_r19_TotalNumberTxPorts1_r19 = 0
	SupportedCSI_RS_ResourceAggregate_r19_TotalNumberTxPorts_r19_TotalNumberTxPorts2_r19 = 1
)

const (
	SupportedCSI_RS_ResourceAggregate_r19_MaxNumberAggregatedResources_r19_MaxNumberAggregatedResources2_r19_N128 = 0
	SupportedCSI_RS_ResourceAggregate_r19_MaxNumberAggregatedResources_r19_MaxNumberAggregatedResources2_r19_N256 = 1
)

var supportedCSIRSResourceAggregateR19MaxNumberAggregatedResourcesR19MaxNumberAggregatedResources2R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SupportedCSI_RS_ResourceAggregate_r19_MaxNumberAggregatedResources_r19_MaxNumberAggregatedResources2_r19_N128, SupportedCSI_RS_ResourceAggregate_r19_MaxNumberAggregatedResources_r19_MaxNumberAggregatedResources2_r19_N256},
}

const (
	SupportedCSI_RS_ResourceAggregate_r19_TotalNumberTxPorts_r19_TotalNumberTxPorts2_r19_N512  = 0
	SupportedCSI_RS_ResourceAggregate_r19_TotalNumberTxPorts_r19_TotalNumberTxPorts2_r19_N768  = 1
	SupportedCSI_RS_ResourceAggregate_r19_TotalNumberTxPorts_r19_TotalNumberTxPorts2_r19_N1024 = 2
)

var supportedCSIRSResourceAggregateR19TotalNumberTxPortsR19TotalNumberTxPorts2R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SupportedCSI_RS_ResourceAggregate_r19_TotalNumberTxPorts_r19_TotalNumberTxPorts2_r19_N512, SupportedCSI_RS_ResourceAggregate_r19_TotalNumberTxPorts_r19_TotalNumberTxPorts2_r19_N768, SupportedCSI_RS_ResourceAggregate_r19_TotalNumberTxPorts_r19_TotalNumberTxPorts2_r19_N1024},
}

type SupportedCSI_RS_ResourceAggregate_r19 struct {
	MaxNumberTxPortsPerAggregatedResource_r19 int64
	MaxNumberAggregatedResources_r19          struct {
		Choice                            int
		MaxNumberAggregatedResources1_r19 *int64
		MaxNumberAggregatedResources2_r19 *int64
	}
	TotalNumberTxPorts_r19 struct {
		Choice                  int
		TotalNumberTxPorts1_r19 *int64
		TotalNumberTxPorts2_r19 *int64
	}
}

func (ie *SupportedCSI_RS_ResourceAggregate_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(supportedCSIRSResourceAggregateR19Constraints)
	_ = seq

	// 1. maxNumberTxPortsPerAggregatedResource-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberTxPortsPerAggregatedResource_r19, supportedCSIRSResourceAggregateR19MaxNumberTxPortsPerAggregatedResourceR19Constraints); err != nil {
			return err
		}
	}

	// 2. maxNumberAggregatedResources-r19: choice
	{
		choiceEnc := e.NewChoiceEncoder(supportedCSI_RS_ResourceAggregate_r19MaxNumberAggregatedResourcesR19Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.MaxNumberAggregatedResources_r19.Choice), false, nil); err != nil {
			return err
		}
		switch ie.MaxNumberAggregatedResources_r19.Choice {
		case SupportedCSI_RS_ResourceAggregate_r19_MaxNumberAggregatedResources_r19_MaxNumberAggregatedResources1_r19:
			if err := e.EncodeInteger((*ie.MaxNumberAggregatedResources_r19.MaxNumberAggregatedResources1_r19), per.Constrained(1, 64)); err != nil {
				return err
			}
		case SupportedCSI_RS_ResourceAggregate_r19_MaxNumberAggregatedResources_r19_MaxNumberAggregatedResources2_r19:
			if err := e.EncodeEnumerated((*ie.MaxNumberAggregatedResources_r19.MaxNumberAggregatedResources2_r19), supportedCSIRSResourceAggregateR19MaxNumberAggregatedResourcesR19MaxNumberAggregatedResources2R19Constraints); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.MaxNumberAggregatedResources_r19.Choice), Constraint: "index out of range"}
		}
	}

	// 3. totalNumberTxPorts-r19: choice
	{
		choiceEnc := e.NewChoiceEncoder(supportedCSI_RS_ResourceAggregate_r19TotalNumberTxPortsR19Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.TotalNumberTxPorts_r19.Choice), false, nil); err != nil {
			return err
		}
		switch ie.TotalNumberTxPorts_r19.Choice {
		case SupportedCSI_RS_ResourceAggregate_r19_TotalNumberTxPorts_r19_TotalNumberTxPorts1_r19:
			if err := e.EncodeInteger((*ie.TotalNumberTxPorts_r19.TotalNumberTxPorts1_r19), per.Constrained(48, 256)); err != nil {
				return err
			}
		case SupportedCSI_RS_ResourceAggregate_r19_TotalNumberTxPorts_r19_TotalNumberTxPorts2_r19:
			if err := e.EncodeEnumerated((*ie.TotalNumberTxPorts_r19.TotalNumberTxPorts2_r19), supportedCSIRSResourceAggregateR19TotalNumberTxPortsR19TotalNumberTxPorts2R19Constraints); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.TotalNumberTxPorts_r19.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *SupportedCSI_RS_ResourceAggregate_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(supportedCSIRSResourceAggregateR19Constraints)
	_ = seq

	// 1. maxNumberTxPortsPerAggregatedResource-r19: enumerated
	{
		v0, err := d.DecodeEnumerated(supportedCSIRSResourceAggregateR19MaxNumberTxPortsPerAggregatedResourceR19Constraints)
		if err != nil {
			return err
		}
		ie.MaxNumberTxPortsPerAggregatedResource_r19 = v0
	}

	// 2. maxNumberAggregatedResources-r19: choice
	{
		choiceDec := d.NewChoiceDecoder(supportedCSI_RS_ResourceAggregate_r19MaxNumberAggregatedResourcesR19Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.MaxNumberAggregatedResources_r19.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SupportedCSI_RS_ResourceAggregate_r19_MaxNumberAggregatedResources_r19_MaxNumberAggregatedResources1_r19:
			ie.MaxNumberAggregatedResources_r19.MaxNumberAggregatedResources1_r19 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(1, 64))
			if err != nil {
				return err
			}
			(*ie.MaxNumberAggregatedResources_r19.MaxNumberAggregatedResources1_r19) = v
		case SupportedCSI_RS_ResourceAggregate_r19_MaxNumberAggregatedResources_r19_MaxNumberAggregatedResources2_r19:
			ie.MaxNumberAggregatedResources_r19.MaxNumberAggregatedResources2_r19 = new(int64)
			v, err := d.DecodeEnumerated(supportedCSIRSResourceAggregateR19MaxNumberAggregatedResourcesR19MaxNumberAggregatedResources2R19Constraints)
			if err != nil {
				return err
			}
			(*ie.MaxNumberAggregatedResources_r19.MaxNumberAggregatedResources2_r19) = v
		}
	}

	// 3. totalNumberTxPorts-r19: choice
	{
		choiceDec := d.NewChoiceDecoder(supportedCSI_RS_ResourceAggregate_r19TotalNumberTxPortsR19Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.TotalNumberTxPorts_r19.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SupportedCSI_RS_ResourceAggregate_r19_TotalNumberTxPorts_r19_TotalNumberTxPorts1_r19:
			ie.TotalNumberTxPorts_r19.TotalNumberTxPorts1_r19 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(48, 256))
			if err != nil {
				return err
			}
			(*ie.TotalNumberTxPorts_r19.TotalNumberTxPorts1_r19) = v
		case SupportedCSI_RS_ResourceAggregate_r19_TotalNumberTxPorts_r19_TotalNumberTxPorts2_r19:
			ie.TotalNumberTxPorts_r19.TotalNumberTxPorts2_r19 = new(int64)
			v, err := d.DecodeEnumerated(supportedCSIRSResourceAggregateR19TotalNumberTxPortsR19TotalNumberTxPorts2R19Constraints)
			if err != nil {
				return err
			}
			(*ie.TotalNumberTxPorts_r19.TotalNumberTxPorts2_r19) = v
		}
	}

	return nil
}
