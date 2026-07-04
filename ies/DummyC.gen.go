// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DummyC (line 19831).

var dummyCConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberTxPortsPerResource"},
		{Name: "maxNumberResources"},
		{Name: "totalNumberTxPorts"},
		{Name: "supportedCodebookMode"},
		{Name: "supportedNumberPanels"},
		{Name: "maxNumberCSI-RS-PerResourceSet"},
	},
}

const (
	DummyC_MaxNumberTxPortsPerResource_P8  = 0
	DummyC_MaxNumberTxPortsPerResource_P16 = 1
	DummyC_MaxNumberTxPortsPerResource_P32 = 2
)

var dummyCMaxNumberTxPortsPerResourceConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DummyC_MaxNumberTxPortsPerResource_P8, DummyC_MaxNumberTxPortsPerResource_P16, DummyC_MaxNumberTxPortsPerResource_P32},
}

var dummyCMaxNumberResourcesConstraints = per.Constrained(1, 64)

var dummyCTotalNumberTxPortsConstraints = per.Constrained(2, 256)

const (
	DummyC_SupportedCodebookMode_Mode1 = 0
	DummyC_SupportedCodebookMode_Mode2 = 1
	DummyC_SupportedCodebookMode_Both  = 2
)

var dummyCSupportedCodebookModeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DummyC_SupportedCodebookMode_Mode1, DummyC_SupportedCodebookMode_Mode2, DummyC_SupportedCodebookMode_Both},
}

const (
	DummyC_SupportedNumberPanels_N2 = 0
	DummyC_SupportedNumberPanels_N4 = 1
)

var dummyCSupportedNumberPanelsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DummyC_SupportedNumberPanels_N2, DummyC_SupportedNumberPanels_N4},
}

var dummyCMaxNumberCSIRSPerResourceSetConstraints = per.Constrained(1, 8)

type DummyC struct {
	MaxNumberTxPortsPerResource    int64
	MaxNumberResources             int64
	TotalNumberTxPorts             int64
	SupportedCodebookMode          int64
	SupportedNumberPanels          int64
	MaxNumberCSI_RS_PerResourceSet int64
}

func (ie *DummyC) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dummyCConstraints)
	_ = seq

	// 1. maxNumberTxPortsPerResource: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberTxPortsPerResource, dummyCMaxNumberTxPortsPerResourceConstraints); err != nil {
			return err
		}
	}

	// 2. maxNumberResources: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberResources, dummyCMaxNumberResourcesConstraints); err != nil {
			return err
		}
	}

	// 3. totalNumberTxPorts: integer
	{
		if err := e.EncodeInteger(ie.TotalNumberTxPorts, dummyCTotalNumberTxPortsConstraints); err != nil {
			return err
		}
	}

	// 4. supportedCodebookMode: enumerated
	{
		if err := e.EncodeEnumerated(ie.SupportedCodebookMode, dummyCSupportedCodebookModeConstraints); err != nil {
			return err
		}
	}

	// 5. supportedNumberPanels: enumerated
	{
		if err := e.EncodeEnumerated(ie.SupportedNumberPanels, dummyCSupportedNumberPanelsConstraints); err != nil {
			return err
		}
	}

	// 6. maxNumberCSI-RS-PerResourceSet: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberCSI_RS_PerResourceSet, dummyCMaxNumberCSIRSPerResourceSetConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *DummyC) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dummyCConstraints)
	_ = seq

	// 1. maxNumberTxPortsPerResource: enumerated
	{
		v0, err := d.DecodeEnumerated(dummyCMaxNumberTxPortsPerResourceConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberTxPortsPerResource = v0
	}

	// 2. maxNumberResources: integer
	{
		v1, err := d.DecodeInteger(dummyCMaxNumberResourcesConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberResources = v1
	}

	// 3. totalNumberTxPorts: integer
	{
		v2, err := d.DecodeInteger(dummyCTotalNumberTxPortsConstraints)
		if err != nil {
			return err
		}
		ie.TotalNumberTxPorts = v2
	}

	// 4. supportedCodebookMode: enumerated
	{
		v3, err := d.DecodeEnumerated(dummyCSupportedCodebookModeConstraints)
		if err != nil {
			return err
		}
		ie.SupportedCodebookMode = v3
	}

	// 5. supportedNumberPanels: enumerated
	{
		v4, err := d.DecodeEnumerated(dummyCSupportedNumberPanelsConstraints)
		if err != nil {
			return err
		}
		ie.SupportedNumberPanels = v4
	}

	// 6. maxNumberCSI-RS-PerResourceSet: integer
	{
		v5, err := d.DecodeInteger(dummyCMaxNumberCSIRSPerResourceSetConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberCSI_RS_PerResourceSet = v5
	}

	return nil
}
