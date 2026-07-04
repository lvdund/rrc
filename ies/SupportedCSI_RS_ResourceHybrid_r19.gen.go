// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SupportedCSI-RS-ResourceHybrid-r19 (line 19383).

var supportedCSIRSResourceHybridR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberTxPortsPerResource-r19"},
		{Name: "maxNumberResources-r19"},
		{Name: "totalNumberTxPorts-r19"},
	},
}

const (
	SupportedCSI_RS_ResourceHybrid_r19_MaxNumberTxPortsPerResource_r19_P2  = 0
	SupportedCSI_RS_ResourceHybrid_r19_MaxNumberTxPortsPerResource_r19_P4  = 1
	SupportedCSI_RS_ResourceHybrid_r19_MaxNumberTxPortsPerResource_r19_P8  = 2
	SupportedCSI_RS_ResourceHybrid_r19_MaxNumberTxPortsPerResource_r19_P12 = 3
	SupportedCSI_RS_ResourceHybrid_r19_MaxNumberTxPortsPerResource_r19_P16 = 4
	SupportedCSI_RS_ResourceHybrid_r19_MaxNumberTxPortsPerResource_r19_P24 = 5
	SupportedCSI_RS_ResourceHybrid_r19_MaxNumberTxPortsPerResource_r19_P32 = 6
)

var supportedCSIRSResourceHybridR19MaxNumberTxPortsPerResourceR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SupportedCSI_RS_ResourceHybrid_r19_MaxNumberTxPortsPerResource_r19_P2, SupportedCSI_RS_ResourceHybrid_r19_MaxNumberTxPortsPerResource_r19_P4, SupportedCSI_RS_ResourceHybrid_r19_MaxNumberTxPortsPerResource_r19_P8, SupportedCSI_RS_ResourceHybrid_r19_MaxNumberTxPortsPerResource_r19_P12, SupportedCSI_RS_ResourceHybrid_r19_MaxNumberTxPortsPerResource_r19_P16, SupportedCSI_RS_ResourceHybrid_r19_MaxNumberTxPortsPerResource_r19_P24, SupportedCSI_RS_ResourceHybrid_r19_MaxNumberTxPortsPerResource_r19_P32},
}

var supportedCSIRSResourceHybridR19MaxNumberResourcesR19Constraints = per.Constrained(1, 256)

var supportedCSI_RS_ResourceHybrid_r19TotalNumberTxPortsR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "totalNumberTxPorts1-r19"},
		{Name: "totalNumberTxPorts2-r19"},
	},
}

const (
	SupportedCSI_RS_ResourceHybrid_r19_TotalNumberTxPorts_r19_TotalNumberTxPorts1_r19 = 0
	SupportedCSI_RS_ResourceHybrid_r19_TotalNumberTxPorts_r19_TotalNumberTxPorts2_r19 = 1
)

const (
	SupportedCSI_RS_ResourceHybrid_r19_TotalNumberTxPorts_r19_TotalNumberTxPorts2_r19_N512  = 0
	SupportedCSI_RS_ResourceHybrid_r19_TotalNumberTxPorts_r19_TotalNumberTxPorts2_r19_N768  = 1
	SupportedCSI_RS_ResourceHybrid_r19_TotalNumberTxPorts_r19_TotalNumberTxPorts2_r19_N1024 = 2
)

var supportedCSIRSResourceHybridR19TotalNumberTxPortsR19TotalNumberTxPorts2R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SupportedCSI_RS_ResourceHybrid_r19_TotalNumberTxPorts_r19_TotalNumberTxPorts2_r19_N512, SupportedCSI_RS_ResourceHybrid_r19_TotalNumberTxPorts_r19_TotalNumberTxPorts2_r19_N768, SupportedCSI_RS_ResourceHybrid_r19_TotalNumberTxPorts_r19_TotalNumberTxPorts2_r19_N1024},
}

type SupportedCSI_RS_ResourceHybrid_r19 struct {
	MaxNumberTxPortsPerResource_r19 int64
	MaxNumberResources_r19          int64
	TotalNumberTxPorts_r19          struct {
		Choice                  int
		TotalNumberTxPorts1_r19 *int64
		TotalNumberTxPorts2_r19 *int64
	}
}

func (ie *SupportedCSI_RS_ResourceHybrid_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(supportedCSIRSResourceHybridR19Constraints)
	_ = seq

	// 1. maxNumberTxPortsPerResource-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberTxPortsPerResource_r19, supportedCSIRSResourceHybridR19MaxNumberTxPortsPerResourceR19Constraints); err != nil {
			return err
		}
	}

	// 2. maxNumberResources-r19: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberResources_r19, supportedCSIRSResourceHybridR19MaxNumberResourcesR19Constraints); err != nil {
			return err
		}
	}

	// 3. totalNumberTxPorts-r19: choice
	{
		choiceEnc := e.NewChoiceEncoder(supportedCSI_RS_ResourceHybrid_r19TotalNumberTxPortsR19Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.TotalNumberTxPorts_r19.Choice), false, nil); err != nil {
			return err
		}
		switch ie.TotalNumberTxPorts_r19.Choice {
		case SupportedCSI_RS_ResourceHybrid_r19_TotalNumberTxPorts_r19_TotalNumberTxPorts1_r19:
			if err := e.EncodeInteger((*ie.TotalNumberTxPorts_r19.TotalNumberTxPorts1_r19), per.Constrained(64, 256)); err != nil {
				return err
			}
		case SupportedCSI_RS_ResourceHybrid_r19_TotalNumberTxPorts_r19_TotalNumberTxPorts2_r19:
			if err := e.EncodeEnumerated((*ie.TotalNumberTxPorts_r19.TotalNumberTxPorts2_r19), supportedCSIRSResourceHybridR19TotalNumberTxPortsR19TotalNumberTxPorts2R19Constraints); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.TotalNumberTxPorts_r19.Choice), Constraint: "index out of range"}
		}
	}

	return nil
}

func (ie *SupportedCSI_RS_ResourceHybrid_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(supportedCSIRSResourceHybridR19Constraints)
	_ = seq

	// 1. maxNumberTxPortsPerResource-r19: enumerated
	{
		v0, err := d.DecodeEnumerated(supportedCSIRSResourceHybridR19MaxNumberTxPortsPerResourceR19Constraints)
		if err != nil {
			return err
		}
		ie.MaxNumberTxPortsPerResource_r19 = v0
	}

	// 2. maxNumberResources-r19: integer
	{
		v1, err := d.DecodeInteger(supportedCSIRSResourceHybridR19MaxNumberResourcesR19Constraints)
		if err != nil {
			return err
		}
		ie.MaxNumberResources_r19 = v1
	}

	// 3. totalNumberTxPorts-r19: choice
	{
		choiceDec := d.NewChoiceDecoder(supportedCSI_RS_ResourceHybrid_r19TotalNumberTxPortsR19Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.TotalNumberTxPorts_r19.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SupportedCSI_RS_ResourceHybrid_r19_TotalNumberTxPorts_r19_TotalNumberTxPorts1_r19:
			ie.TotalNumberTxPorts_r19.TotalNumberTxPorts1_r19 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(64, 256))
			if err != nil {
				return err
			}
			(*ie.TotalNumberTxPorts_r19.TotalNumberTxPorts1_r19) = v
		case SupportedCSI_RS_ResourceHybrid_r19_TotalNumberTxPorts_r19_TotalNumberTxPorts2_r19:
			ie.TotalNumberTxPorts_r19.TotalNumberTxPorts2_r19 = new(int64)
			v, err := d.DecodeEnumerated(supportedCSIRSResourceHybridR19TotalNumberTxPortsR19TotalNumberTxPorts2R19Constraints)
			if err != nil {
				return err
			}
			(*ie.TotalNumberTxPorts_r19.TotalNumberTxPorts2_r19) = v
		}
	}

	return nil
}
