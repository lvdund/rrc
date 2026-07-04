// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PEI-Config-r17 (line 7926).

var pEIConfigR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "po-NumPerPEI-r17"},
		{Name: "payloadSizeDCI-2-7-r17"},
		{Name: "pei-FrameOffset-r17"},
		{Name: "subgroupConfig-r17"},
		{Name: "lastUsedCellOnly-r17", Optional: true},
	},
}

const (
	PEI_Config_r17_Po_NumPerPEI_r17_Po1 = 0
	PEI_Config_r17_Po_NumPerPEI_r17_Po2 = 1
	PEI_Config_r17_Po_NumPerPEI_r17_Po4 = 2
	PEI_Config_r17_Po_NumPerPEI_r17_Po8 = 3
)

var pEIConfigR17PoNumPerPEIR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PEI_Config_r17_Po_NumPerPEI_r17_Po1, PEI_Config_r17_Po_NumPerPEI_r17_Po2, PEI_Config_r17_Po_NumPerPEI_r17_Po4, PEI_Config_r17_Po_NumPerPEI_r17_Po8},
}

var pEIConfigR17PayloadSizeDCI27R17Constraints = per.Constrained(1, common.MaxDCI_2_7_Size_r17)

var pEIConfigR17PeiFrameOffsetR17Constraints = per.Constrained(0, 16)

const (
	PEI_Config_r17_LastUsedCellOnly_r17_True = 0
)

var pEIConfigR17LastUsedCellOnlyR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PEI_Config_r17_LastUsedCellOnly_r17_True},
}

type PEI_Config_r17 struct {
	Po_NumPerPEI_r17       int64
	PayloadSizeDCI_2_7_r17 int64
	Pei_FrameOffset_r17    int64
	SubgroupConfig_r17     SubgroupConfig_r17
	LastUsedCellOnly_r17   *int64
}

func (ie *PEI_Config_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pEIConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LastUsedCellOnly_r17 != nil}); err != nil {
		return err
	}

	// 3. po-NumPerPEI-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Po_NumPerPEI_r17, pEIConfigR17PoNumPerPEIR17Constraints); err != nil {
			return err
		}
	}

	// 4. payloadSizeDCI-2-7-r17: integer
	{
		if err := e.EncodeInteger(ie.PayloadSizeDCI_2_7_r17, pEIConfigR17PayloadSizeDCI27R17Constraints); err != nil {
			return err
		}
	}

	// 5. pei-FrameOffset-r17: integer
	{
		if err := e.EncodeInteger(ie.Pei_FrameOffset_r17, pEIConfigR17PeiFrameOffsetR17Constraints); err != nil {
			return err
		}
	}

	// 6. subgroupConfig-r17: ref
	{
		if err := ie.SubgroupConfig_r17.Encode(e); err != nil {
			return err
		}
	}

	// 7. lastUsedCellOnly-r17: enumerated
	{
		if ie.LastUsedCellOnly_r17 != nil {
			if err := e.EncodeEnumerated(*ie.LastUsedCellOnly_r17, pEIConfigR17LastUsedCellOnlyR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PEI_Config_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pEIConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. po-NumPerPEI-r17: enumerated
	{
		v0, err := d.DecodeEnumerated(pEIConfigR17PoNumPerPEIR17Constraints)
		if err != nil {
			return err
		}
		ie.Po_NumPerPEI_r17 = v0
	}

	// 4. payloadSizeDCI-2-7-r17: integer
	{
		v1, err := d.DecodeInteger(pEIConfigR17PayloadSizeDCI27R17Constraints)
		if err != nil {
			return err
		}
		ie.PayloadSizeDCI_2_7_r17 = v1
	}

	// 5. pei-FrameOffset-r17: integer
	{
		v2, err := d.DecodeInteger(pEIConfigR17PeiFrameOffsetR17Constraints)
		if err != nil {
			return err
		}
		ie.Pei_FrameOffset_r17 = v2
	}

	// 6. subgroupConfig-r17: ref
	{
		if err := ie.SubgroupConfig_r17.Decode(d); err != nil {
			return err
		}
	}

	// 7. lastUsedCellOnly-r17: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(pEIConfigR17LastUsedCellOnlyR17Constraints)
			if err != nil {
				return err
			}
			ie.LastUsedCellOnly_r17 = &idx
		}
	}

	return nil
}
