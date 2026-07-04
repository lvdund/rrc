// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DummyE (line 19850).

var dummyEConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberTxPortsPerResource"},
		{Name: "maxNumberResources"},
		{Name: "totalNumberTxPorts"},
		{Name: "parameterLx"},
		{Name: "amplitudeScalingType"},
		{Name: "maxNumberCSI-RS-PerResourceSet"},
	},
}

const (
	DummyE_MaxNumberTxPortsPerResource_P4  = 0
	DummyE_MaxNumberTxPortsPerResource_P8  = 1
	DummyE_MaxNumberTxPortsPerResource_P12 = 2
	DummyE_MaxNumberTxPortsPerResource_P16 = 3
	DummyE_MaxNumberTxPortsPerResource_P24 = 4
	DummyE_MaxNumberTxPortsPerResource_P32 = 5
)

var dummyEMaxNumberTxPortsPerResourceConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DummyE_MaxNumberTxPortsPerResource_P4, DummyE_MaxNumberTxPortsPerResource_P8, DummyE_MaxNumberTxPortsPerResource_P12, DummyE_MaxNumberTxPortsPerResource_P16, DummyE_MaxNumberTxPortsPerResource_P24, DummyE_MaxNumberTxPortsPerResource_P32},
}

var dummyEMaxNumberResourcesConstraints = per.Constrained(1, 64)

var dummyETotalNumberTxPortsConstraints = per.Constrained(2, 256)

var dummyEParameterLxConstraints = per.Constrained(2, 4)

const (
	DummyE_AmplitudeScalingType_Wideband           = 0
	DummyE_AmplitudeScalingType_WidebandAndSubband = 1
)

var dummyEAmplitudeScalingTypeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DummyE_AmplitudeScalingType_Wideband, DummyE_AmplitudeScalingType_WidebandAndSubband},
}

var dummyEMaxNumberCSIRSPerResourceSetConstraints = per.Constrained(1, 8)

type DummyE struct {
	MaxNumberTxPortsPerResource    int64
	MaxNumberResources             int64
	TotalNumberTxPorts             int64
	ParameterLx                    int64
	AmplitudeScalingType           int64
	MaxNumberCSI_RS_PerResourceSet int64
}

func (ie *DummyE) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dummyEConstraints)
	_ = seq

	// 1. maxNumberTxPortsPerResource: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberTxPortsPerResource, dummyEMaxNumberTxPortsPerResourceConstraints); err != nil {
			return err
		}
	}

	// 2. maxNumberResources: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberResources, dummyEMaxNumberResourcesConstraints); err != nil {
			return err
		}
	}

	// 3. totalNumberTxPorts: integer
	{
		if err := e.EncodeInteger(ie.TotalNumberTxPorts, dummyETotalNumberTxPortsConstraints); err != nil {
			return err
		}
	}

	// 4. parameterLx: integer
	{
		if err := e.EncodeInteger(ie.ParameterLx, dummyEParameterLxConstraints); err != nil {
			return err
		}
	}

	// 5. amplitudeScalingType: enumerated
	{
		if err := e.EncodeEnumerated(ie.AmplitudeScalingType, dummyEAmplitudeScalingTypeConstraints); err != nil {
			return err
		}
	}

	// 6. maxNumberCSI-RS-PerResourceSet: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberCSI_RS_PerResourceSet, dummyEMaxNumberCSIRSPerResourceSetConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *DummyE) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dummyEConstraints)
	_ = seq

	// 1. maxNumberTxPortsPerResource: enumerated
	{
		v0, err := d.DecodeEnumerated(dummyEMaxNumberTxPortsPerResourceConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberTxPortsPerResource = v0
	}

	// 2. maxNumberResources: integer
	{
		v1, err := d.DecodeInteger(dummyEMaxNumberResourcesConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberResources = v1
	}

	// 3. totalNumberTxPorts: integer
	{
		v2, err := d.DecodeInteger(dummyETotalNumberTxPortsConstraints)
		if err != nil {
			return err
		}
		ie.TotalNumberTxPorts = v2
	}

	// 4. parameterLx: integer
	{
		v3, err := d.DecodeInteger(dummyEParameterLxConstraints)
		if err != nil {
			return err
		}
		ie.ParameterLx = v3
	}

	// 5. amplitudeScalingType: enumerated
	{
		v4, err := d.DecodeEnumerated(dummyEAmplitudeScalingTypeConstraints)
		if err != nil {
			return err
		}
		ie.AmplitudeScalingType = v4
	}

	// 6. maxNumberCSI-RS-PerResourceSet: integer
	{
		v5, err := d.DecodeInteger(dummyEMaxNumberCSIRSPerResourceSetConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberCSI_RS_PerResourceSet = v5
	}

	return nil
}
