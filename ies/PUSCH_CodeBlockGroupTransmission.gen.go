// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUSCH-CodeBlockGroupTransmission (line 12689).

var pUSCHCodeBlockGroupTransmissionConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "maxCodeBlockGroupsPerTransportBlock"},
	},
}

const (
	PUSCH_CodeBlockGroupTransmission_MaxCodeBlockGroupsPerTransportBlock_N2 = 0
	PUSCH_CodeBlockGroupTransmission_MaxCodeBlockGroupsPerTransportBlock_N4 = 1
	PUSCH_CodeBlockGroupTransmission_MaxCodeBlockGroupsPerTransportBlock_N6 = 2
	PUSCH_CodeBlockGroupTransmission_MaxCodeBlockGroupsPerTransportBlock_N8 = 3
)

var pUSCHCodeBlockGroupTransmissionMaxCodeBlockGroupsPerTransportBlockConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_CodeBlockGroupTransmission_MaxCodeBlockGroupsPerTransportBlock_N2, PUSCH_CodeBlockGroupTransmission_MaxCodeBlockGroupsPerTransportBlock_N4, PUSCH_CodeBlockGroupTransmission_MaxCodeBlockGroupsPerTransportBlock_N6, PUSCH_CodeBlockGroupTransmission_MaxCodeBlockGroupsPerTransportBlock_N8},
}

type PUSCH_CodeBlockGroupTransmission struct {
	MaxCodeBlockGroupsPerTransportBlock int64
}

func (ie *PUSCH_CodeBlockGroupTransmission) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUSCHCodeBlockGroupTransmissionConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. maxCodeBlockGroupsPerTransportBlock: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxCodeBlockGroupsPerTransportBlock, pUSCHCodeBlockGroupTransmissionMaxCodeBlockGroupsPerTransportBlockConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PUSCH_CodeBlockGroupTransmission) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUSCHCodeBlockGroupTransmissionConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. maxCodeBlockGroupsPerTransportBlock: enumerated
	{
		v0, err := d.DecodeEnumerated(pUSCHCodeBlockGroupTransmissionMaxCodeBlockGroupsPerTransportBlockConstraints)
		if err != nil {
			return err
		}
		ie.MaxCodeBlockGroupsPerTransportBlock = v0
	}

	return nil
}
