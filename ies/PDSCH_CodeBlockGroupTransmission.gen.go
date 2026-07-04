// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDSCH-CodeBlockGroupTransmission (line 11499).

var pDSCHCodeBlockGroupTransmissionConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "maxCodeBlockGroupsPerTransportBlock"},
		{Name: "codeBlockGroupFlushIndicator"},
	},
}

const (
	PDSCH_CodeBlockGroupTransmission_MaxCodeBlockGroupsPerTransportBlock_N2 = 0
	PDSCH_CodeBlockGroupTransmission_MaxCodeBlockGroupsPerTransportBlock_N4 = 1
	PDSCH_CodeBlockGroupTransmission_MaxCodeBlockGroupsPerTransportBlock_N6 = 2
	PDSCH_CodeBlockGroupTransmission_MaxCodeBlockGroupsPerTransportBlock_N8 = 3
)

var pDSCHCodeBlockGroupTransmissionMaxCodeBlockGroupsPerTransportBlockConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_CodeBlockGroupTransmission_MaxCodeBlockGroupsPerTransportBlock_N2, PDSCH_CodeBlockGroupTransmission_MaxCodeBlockGroupsPerTransportBlock_N4, PDSCH_CodeBlockGroupTransmission_MaxCodeBlockGroupsPerTransportBlock_N6, PDSCH_CodeBlockGroupTransmission_MaxCodeBlockGroupsPerTransportBlock_N8},
}

type PDSCH_CodeBlockGroupTransmission struct {
	MaxCodeBlockGroupsPerTransportBlock int64
	CodeBlockGroupFlushIndicator        bool
}

func (ie *PDSCH_CodeBlockGroupTransmission) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDSCHCodeBlockGroupTransmissionConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. maxCodeBlockGroupsPerTransportBlock: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxCodeBlockGroupsPerTransportBlock, pDSCHCodeBlockGroupTransmissionMaxCodeBlockGroupsPerTransportBlockConstraints); err != nil {
			return err
		}
	}

	// 3. codeBlockGroupFlushIndicator: boolean
	{
		if err := e.EncodeBoolean(ie.CodeBlockGroupFlushIndicator); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PDSCH_CodeBlockGroupTransmission) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDSCHCodeBlockGroupTransmissionConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. maxCodeBlockGroupsPerTransportBlock: enumerated
	{
		v0, err := d.DecodeEnumerated(pDSCHCodeBlockGroupTransmissionMaxCodeBlockGroupsPerTransportBlockConstraints)
		if err != nil {
			return err
		}
		ie.MaxCodeBlockGroupsPerTransportBlock = v0
	}

	// 3. codeBlockGroupFlushIndicator: boolean
	{
		v1, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.CodeBlockGroupFlushIndicator = v1
	}

	return nil
}
