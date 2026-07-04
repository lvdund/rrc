// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SupportedCSI-RS-Resource (line 19334).

var supportedCSIRSResourceConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberTxPortsPerResource"},
		{Name: "maxNumberResourcesPerBand"},
		{Name: "totalNumberTxPortsPerBand"},
	},
}

const (
	SupportedCSI_RS_Resource_MaxNumberTxPortsPerResource_P2  = 0
	SupportedCSI_RS_Resource_MaxNumberTxPortsPerResource_P4  = 1
	SupportedCSI_RS_Resource_MaxNumberTxPortsPerResource_P8  = 2
	SupportedCSI_RS_Resource_MaxNumberTxPortsPerResource_P12 = 3
	SupportedCSI_RS_Resource_MaxNumberTxPortsPerResource_P16 = 4
	SupportedCSI_RS_Resource_MaxNumberTxPortsPerResource_P24 = 5
	SupportedCSI_RS_Resource_MaxNumberTxPortsPerResource_P32 = 6
)

var supportedCSIRSResourceMaxNumberTxPortsPerResourceConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SupportedCSI_RS_Resource_MaxNumberTxPortsPerResource_P2, SupportedCSI_RS_Resource_MaxNumberTxPortsPerResource_P4, SupportedCSI_RS_Resource_MaxNumberTxPortsPerResource_P8, SupportedCSI_RS_Resource_MaxNumberTxPortsPerResource_P12, SupportedCSI_RS_Resource_MaxNumberTxPortsPerResource_P16, SupportedCSI_RS_Resource_MaxNumberTxPortsPerResource_P24, SupportedCSI_RS_Resource_MaxNumberTxPortsPerResource_P32},
}

var supportedCSIRSResourceMaxNumberResourcesPerBandConstraints = per.Constrained(1, 64)

var supportedCSIRSResourceTotalNumberTxPortsPerBandConstraints = per.Constrained(2, 256)

type SupportedCSI_RS_Resource struct {
	MaxNumberTxPortsPerResource int64
	MaxNumberResourcesPerBand   int64
	TotalNumberTxPortsPerBand   int64
}

func (ie *SupportedCSI_RS_Resource) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(supportedCSIRSResourceConstraints)
	_ = seq

	// 1. maxNumberTxPortsPerResource: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberTxPortsPerResource, supportedCSIRSResourceMaxNumberTxPortsPerResourceConstraints); err != nil {
			return err
		}
	}

	// 2. maxNumberResourcesPerBand: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberResourcesPerBand, supportedCSIRSResourceMaxNumberResourcesPerBandConstraints); err != nil {
			return err
		}
	}

	// 3. totalNumberTxPortsPerBand: integer
	{
		if err := e.EncodeInteger(ie.TotalNumberTxPortsPerBand, supportedCSIRSResourceTotalNumberTxPortsPerBandConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SupportedCSI_RS_Resource) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(supportedCSIRSResourceConstraints)
	_ = seq

	// 1. maxNumberTxPortsPerResource: enumerated
	{
		v0, err := d.DecodeEnumerated(supportedCSIRSResourceMaxNumberTxPortsPerResourceConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberTxPortsPerResource = v0
	}

	// 2. maxNumberResourcesPerBand: integer
	{
		v1, err := d.DecodeInteger(supportedCSIRSResourceMaxNumberResourcesPerBandConstraints)
		if err != nil {
			return err
		}
		ie.MaxNumberResourcesPerBand = v1
	}

	// 3. totalNumberTxPortsPerBand: integer
	{
		v2, err := d.DecodeInteger(supportedCSIRSResourceTotalNumberTxPortsPerBandConstraints)
		if err != nil {
			return err
		}
		ie.TotalNumberTxPortsPerBand = v2
	}

	return nil
}
