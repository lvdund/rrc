// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DummyB (line 19823).

var dummyBConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberTxPortsPerResource"},
		{Name: "maxNumberResources"},
		{Name: "totalNumberTxPorts"},
		{Name: "supportedCodebookMode"},
		{Name: "maxNumberCSI-RS-PerResourceSet"},
	},
}

const (
	DummyB_MaxNumberTxPortsPerResource_P2  = 0
	DummyB_MaxNumberTxPortsPerResource_P4  = 1
	DummyB_MaxNumberTxPortsPerResource_P8  = 2
	DummyB_MaxNumberTxPortsPerResource_P12 = 3
	DummyB_MaxNumberTxPortsPerResource_P16 = 4
	DummyB_MaxNumberTxPortsPerResource_P24 = 5
	DummyB_MaxNumberTxPortsPerResource_P32 = 6
)

var dummyBMaxNumberTxPortsPerResourceConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DummyB_MaxNumberTxPortsPerResource_P2, DummyB_MaxNumberTxPortsPerResource_P4, DummyB_MaxNumberTxPortsPerResource_P8, DummyB_MaxNumberTxPortsPerResource_P12, DummyB_MaxNumberTxPortsPerResource_P16, DummyB_MaxNumberTxPortsPerResource_P24, DummyB_MaxNumberTxPortsPerResource_P32},
}

var dummyBMaxNumberResourcesConstraints = per.Constrained(1, 64)

var dummyBTotalNumberTxPortsConstraints = per.Constrained(2, 256)

const (
	DummyB_SupportedCodebookMode_Mode1         = 0
	DummyB_SupportedCodebookMode_Mode1AndMode2 = 1
)

var dummyBSupportedCodebookModeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DummyB_SupportedCodebookMode_Mode1, DummyB_SupportedCodebookMode_Mode1AndMode2},
}

var dummyBMaxNumberCSIRSPerResourceSetConstraints = per.Constrained(1, 8)

type DummyB struct {
	MaxNumberTxPortsPerResource    int64
	MaxNumberResources             int64
	TotalNumberTxPorts             int64
	SupportedCodebookMode          int64
	MaxNumberCSI_RS_PerResourceSet int64
}

func (ie *DummyB) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dummyBConstraints)
	_ = seq

	// 1. maxNumberTxPortsPerResource: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberTxPortsPerResource, dummyBMaxNumberTxPortsPerResourceConstraints); err != nil {
			return err
		}
	}

	// 2. maxNumberResources: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberResources, dummyBMaxNumberResourcesConstraints); err != nil {
			return err
		}
	}

	// 3. totalNumberTxPorts: integer
	{
		if err := e.EncodeInteger(ie.TotalNumberTxPorts, dummyBTotalNumberTxPortsConstraints); err != nil {
			return err
		}
	}

	// 4. supportedCodebookMode: enumerated
	{
		if err := e.EncodeEnumerated(ie.SupportedCodebookMode, dummyBSupportedCodebookModeConstraints); err != nil {
			return err
		}
	}

	// 5. maxNumberCSI-RS-PerResourceSet: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberCSI_RS_PerResourceSet, dummyBMaxNumberCSIRSPerResourceSetConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *DummyB) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dummyBConstraints)
	_ = seq

	// 1. maxNumberTxPortsPerResource: enumerated
	{
		v0, err := d.DecodeEnumerated(dummyBMaxNumberTxPortsPerResourceConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberTxPortsPerResource = v0
	}

	// 2. maxNumberResources: integer
	{
		v1, err := d.DecodeInteger(dummyBMaxNumberResourcesConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberResources = v1
	}

	// 3. totalNumberTxPorts: integer
	{
		v2, err := d.DecodeInteger(dummyBTotalNumberTxPortsConstraints)
		if err != nil {
			return err
		}
		ie.TotalNumberTxPorts = v2
	}

	// 4. supportedCodebookMode: enumerated
	{
		v3, err := d.DecodeEnumerated(dummyBSupportedCodebookModeConstraints)
		if err != nil {
			return err
		}
		ie.SupportedCodebookMode = v3
	}

	// 5. maxNumberCSI-RS-PerResourceSet: integer
	{
		v4, err := d.DecodeInteger(dummyBMaxNumberCSIRSPerResourceSetConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberCSI_RS_PerResourceSet = v4
	}

	return nil
}
