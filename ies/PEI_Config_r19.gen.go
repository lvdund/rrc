// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PEI-Config-r19 (line 8079).

var pEIConfigR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "po-NumPerPEI-r19"},
		{Name: "payloadSizeDCI-2-7-r19"},
		{Name: "pei-FrameOffset-r19"},
	},
}

const (
	PEI_Config_r19_Po_NumPerPEI_r19_Po1 = 0
	PEI_Config_r19_Po_NumPerPEI_r19_Po2 = 1
	PEI_Config_r19_Po_NumPerPEI_r19_Po4 = 2
	PEI_Config_r19_Po_NumPerPEI_r19_Po8 = 3
)

var pEIConfigR19PoNumPerPEIR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PEI_Config_r19_Po_NumPerPEI_r19_Po1, PEI_Config_r19_Po_NumPerPEI_r19_Po2, PEI_Config_r19_Po_NumPerPEI_r19_Po4, PEI_Config_r19_Po_NumPerPEI_r19_Po8},
}

var pEIConfigR19PayloadSizeDCI27R19Constraints = per.Constrained(1, common.MaxDCI_2_7_Size_r17)

var pEIConfigR19PeiFrameOffsetR19Constraints = per.Constrained(0, 32)

type PEI_Config_r19 struct {
	Po_NumPerPEI_r19       int64
	PayloadSizeDCI_2_7_r19 int64
	Pei_FrameOffset_r19    int64
}

func (ie *PEI_Config_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pEIConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. po-NumPerPEI-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.Po_NumPerPEI_r19, pEIConfigR19PoNumPerPEIR19Constraints); err != nil {
			return err
		}
	}

	// 3. payloadSizeDCI-2-7-r19: integer
	{
		if err := e.EncodeInteger(ie.PayloadSizeDCI_2_7_r19, pEIConfigR19PayloadSizeDCI27R19Constraints); err != nil {
			return err
		}
	}

	// 4. pei-FrameOffset-r19: integer
	{
		if err := e.EncodeInteger(ie.Pei_FrameOffset_r19, pEIConfigR19PeiFrameOffsetR19Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PEI_Config_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pEIConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. po-NumPerPEI-r19: enumerated
	{
		v0, err := d.DecodeEnumerated(pEIConfigR19PoNumPerPEIR19Constraints)
		if err != nil {
			return err
		}
		ie.Po_NumPerPEI_r19 = v0
	}

	// 3. payloadSizeDCI-2-7-r19: integer
	{
		v1, err := d.DecodeInteger(pEIConfigR19PayloadSizeDCI27R19Constraints)
		if err != nil {
			return err
		}
		ie.PayloadSizeDCI_2_7_r19 = v1
	}

	// 4. pei-FrameOffset-r19: integer
	{
		v2, err := d.DecodeInteger(pEIConfigR19PeiFrameOffsetR19Constraints)
		if err != nil {
			return err
		}
		ie.Pei_FrameOffset_r19 = v2
	}

	return nil
}
