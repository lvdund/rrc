// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DummyD (line 19840).

var dummyDConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberTxPortsPerResource"},
		{Name: "maxNumberResources"},
		{Name: "totalNumberTxPorts"},
		{Name: "parameterLx"},
		{Name: "amplitudeScalingType"},
		{Name: "amplitudeSubsetRestriction", Optional: true},
		{Name: "maxNumberCSI-RS-PerResourceSet"},
	},
}

const (
	DummyD_MaxNumberTxPortsPerResource_P4  = 0
	DummyD_MaxNumberTxPortsPerResource_P8  = 1
	DummyD_MaxNumberTxPortsPerResource_P12 = 2
	DummyD_MaxNumberTxPortsPerResource_P16 = 3
	DummyD_MaxNumberTxPortsPerResource_P24 = 4
	DummyD_MaxNumberTxPortsPerResource_P32 = 5
)

var dummyDMaxNumberTxPortsPerResourceConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DummyD_MaxNumberTxPortsPerResource_P4, DummyD_MaxNumberTxPortsPerResource_P8, DummyD_MaxNumberTxPortsPerResource_P12, DummyD_MaxNumberTxPortsPerResource_P16, DummyD_MaxNumberTxPortsPerResource_P24, DummyD_MaxNumberTxPortsPerResource_P32},
}

var dummyDMaxNumberResourcesConstraints = per.Constrained(1, 64)

var dummyDTotalNumberTxPortsConstraints = per.Constrained(2, 256)

var dummyDParameterLxConstraints = per.Constrained(2, 4)

const (
	DummyD_AmplitudeScalingType_Wideband           = 0
	DummyD_AmplitudeScalingType_WidebandAndSubband = 1
)

var dummyDAmplitudeScalingTypeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DummyD_AmplitudeScalingType_Wideband, DummyD_AmplitudeScalingType_WidebandAndSubband},
}

const (
	DummyD_AmplitudeSubsetRestriction_Supported = 0
)

var dummyDAmplitudeSubsetRestrictionConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DummyD_AmplitudeSubsetRestriction_Supported},
}

var dummyDMaxNumberCSIRSPerResourceSetConstraints = per.Constrained(1, 8)

type DummyD struct {
	MaxNumberTxPortsPerResource    int64
	MaxNumberResources             int64
	TotalNumberTxPorts             int64
	ParameterLx                    int64
	AmplitudeScalingType           int64
	AmplitudeSubsetRestriction     *int64
	MaxNumberCSI_RS_PerResourceSet int64
}

func (ie *DummyD) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dummyDConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AmplitudeSubsetRestriction != nil}); err != nil {
		return err
	}

	// 2. maxNumberTxPortsPerResource: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberTxPortsPerResource, dummyDMaxNumberTxPortsPerResourceConstraints); err != nil {
			return err
		}
	}

	// 3. maxNumberResources: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberResources, dummyDMaxNumberResourcesConstraints); err != nil {
			return err
		}
	}

	// 4. totalNumberTxPorts: integer
	{
		if err := e.EncodeInteger(ie.TotalNumberTxPorts, dummyDTotalNumberTxPortsConstraints); err != nil {
			return err
		}
	}

	// 5. parameterLx: integer
	{
		if err := e.EncodeInteger(ie.ParameterLx, dummyDParameterLxConstraints); err != nil {
			return err
		}
	}

	// 6. amplitudeScalingType: enumerated
	{
		if err := e.EncodeEnumerated(ie.AmplitudeScalingType, dummyDAmplitudeScalingTypeConstraints); err != nil {
			return err
		}
	}

	// 7. amplitudeSubsetRestriction: enumerated
	{
		if ie.AmplitudeSubsetRestriction != nil {
			if err := e.EncodeEnumerated(*ie.AmplitudeSubsetRestriction, dummyDAmplitudeSubsetRestrictionConstraints); err != nil {
				return err
			}
		}
	}

	// 8. maxNumberCSI-RS-PerResourceSet: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberCSI_RS_PerResourceSet, dummyDMaxNumberCSIRSPerResourceSetConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *DummyD) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dummyDConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. maxNumberTxPortsPerResource: enumerated
	{
		v0, err := d.DecodeEnumerated(dummyDMaxNumberTxPortsPerResourceConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberTxPortsPerResource = v0
	}

	// 3. maxNumberResources: integer
	{
		v1, err := d.DecodeInteger(dummyDMaxNumberResourcesConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberResources = v1
	}

	// 4. totalNumberTxPorts: integer
	{
		v2, err := d.DecodeInteger(dummyDTotalNumberTxPortsConstraints)
		if err != nil {
			return err
		}
		ie.TotalNumberTxPorts = v2
	}

	// 5. parameterLx: integer
	{
		v3, err := d.DecodeInteger(dummyDParameterLxConstraints)
		if err != nil {
			return err
		}
		ie.ParameterLx = v3
	}

	// 6. amplitudeScalingType: enumerated
	{
		v4, err := d.DecodeEnumerated(dummyDAmplitudeScalingTypeConstraints)
		if err != nil {
			return err
		}
		ie.AmplitudeScalingType = v4
	}

	// 7. amplitudeSubsetRestriction: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(dummyDAmplitudeSubsetRestrictionConstraints)
			if err != nil {
				return err
			}
			ie.AmplitudeSubsetRestriction = &idx
		}
	}

	// 8. maxNumberCSI-RS-PerResourceSet: integer
	{
		v6, err := d.DecodeInteger(dummyDMaxNumberCSIRSPerResourceSetConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberCSI_RS_PerResourceSet = v6
	}

	return nil
}
