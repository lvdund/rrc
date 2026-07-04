// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDSCH-ConfigDCI-1-3-r18 (line 11449).

var pDSCHConfigDCI13R18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "resourceAllocationDCI-1-3-r18", Optional: true},
		{Name: "rbg-SizeDCI-1-3-r18", Optional: true},
		{Name: "resourceAllocationType1GranularityDCI-1-3-r18", Optional: true},
		{Name: "numberOfBitsForRV-DCI-1-3-r18", Optional: true},
		{Name: "harq-ProcessNumberSizeDCI-1-3-r18", Optional: true},
	},
}

const (
	PDSCH_ConfigDCI_1_3_r18_ResourceAllocationDCI_1_3_r18_ResourceAllocationType0 = 0
	PDSCH_ConfigDCI_1_3_r18_ResourceAllocationDCI_1_3_r18_ResourceAllocationType1 = 1
	PDSCH_ConfigDCI_1_3_r18_ResourceAllocationDCI_1_3_r18_DynamicSwitch           = 2
)

var pDSCHConfigDCI13R18ResourceAllocationDCI13R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_ConfigDCI_1_3_r18_ResourceAllocationDCI_1_3_r18_ResourceAllocationType0, PDSCH_ConfigDCI_1_3_r18_ResourceAllocationDCI_1_3_r18_ResourceAllocationType1, PDSCH_ConfigDCI_1_3_r18_ResourceAllocationDCI_1_3_r18_DynamicSwitch},
}

const (
	PDSCH_ConfigDCI_1_3_r18_Rbg_SizeDCI_1_3_r18_Config1 = 0
	PDSCH_ConfigDCI_1_3_r18_Rbg_SizeDCI_1_3_r18_Config2 = 1
	PDSCH_ConfigDCI_1_3_r18_Rbg_SizeDCI_1_3_r18_Config3 = 2
	PDSCH_ConfigDCI_1_3_r18_Rbg_SizeDCI_1_3_r18_Spare1  = 3
)

var pDSCHConfigDCI13R18RbgSizeDCI13R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_ConfigDCI_1_3_r18_Rbg_SizeDCI_1_3_r18_Config1, PDSCH_ConfigDCI_1_3_r18_Rbg_SizeDCI_1_3_r18_Config2, PDSCH_ConfigDCI_1_3_r18_Rbg_SizeDCI_1_3_r18_Config3, PDSCH_ConfigDCI_1_3_r18_Rbg_SizeDCI_1_3_r18_Spare1},
}

const (
	PDSCH_ConfigDCI_1_3_r18_ResourceAllocationType1GranularityDCI_1_3_r18_N2  = 0
	PDSCH_ConfigDCI_1_3_r18_ResourceAllocationType1GranularityDCI_1_3_r18_N4  = 1
	PDSCH_ConfigDCI_1_3_r18_ResourceAllocationType1GranularityDCI_1_3_r18_N8  = 2
	PDSCH_ConfigDCI_1_3_r18_ResourceAllocationType1GranularityDCI_1_3_r18_N16 = 3
)

var pDSCHConfigDCI13R18ResourceAllocationType1GranularityDCI13R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_ConfigDCI_1_3_r18_ResourceAllocationType1GranularityDCI_1_3_r18_N2, PDSCH_ConfigDCI_1_3_r18_ResourceAllocationType1GranularityDCI_1_3_r18_N4, PDSCH_ConfigDCI_1_3_r18_ResourceAllocationType1GranularityDCI_1_3_r18_N8, PDSCH_ConfigDCI_1_3_r18_ResourceAllocationType1GranularityDCI_1_3_r18_N16},
}

var pDSCHConfigDCI13R18NumberOfBitsForRVDCI13R18Constraints = per.Constrained(0, 2)

var pDSCHConfigDCI13R18HarqProcessNumberSizeDCI13R18Constraints = per.Constrained(0, 5)

type PDSCH_ConfigDCI_1_3_r18 struct {
	ResourceAllocationDCI_1_3_r18                 *int64
	Rbg_SizeDCI_1_3_r18                           *int64
	ResourceAllocationType1GranularityDCI_1_3_r18 *int64
	NumberOfBitsForRV_DCI_1_3_r18                 *int64
	Harq_ProcessNumberSizeDCI_1_3_r18             *int64
}

func (ie *PDSCH_ConfigDCI_1_3_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDSCHConfigDCI13R18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ResourceAllocationDCI_1_3_r18 != nil, ie.Rbg_SizeDCI_1_3_r18 != nil, ie.ResourceAllocationType1GranularityDCI_1_3_r18 != nil, ie.NumberOfBitsForRV_DCI_1_3_r18 != nil, ie.Harq_ProcessNumberSizeDCI_1_3_r18 != nil}); err != nil {
		return err
	}

	// 2. resourceAllocationDCI-1-3-r18: enumerated
	{
		if ie.ResourceAllocationDCI_1_3_r18 != nil {
			if err := e.EncodeEnumerated(*ie.ResourceAllocationDCI_1_3_r18, pDSCHConfigDCI13R18ResourceAllocationDCI13R18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. rbg-SizeDCI-1-3-r18: enumerated
	{
		if ie.Rbg_SizeDCI_1_3_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Rbg_SizeDCI_1_3_r18, pDSCHConfigDCI13R18RbgSizeDCI13R18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. resourceAllocationType1GranularityDCI-1-3-r18: enumerated
	{
		if ie.ResourceAllocationType1GranularityDCI_1_3_r18 != nil {
			if err := e.EncodeEnumerated(*ie.ResourceAllocationType1GranularityDCI_1_3_r18, pDSCHConfigDCI13R18ResourceAllocationType1GranularityDCI13R18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. numberOfBitsForRV-DCI-1-3-r18: integer
	{
		if ie.NumberOfBitsForRV_DCI_1_3_r18 != nil {
			if err := e.EncodeInteger(*ie.NumberOfBitsForRV_DCI_1_3_r18, pDSCHConfigDCI13R18NumberOfBitsForRVDCI13R18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. harq-ProcessNumberSizeDCI-1-3-r18: integer
	{
		if ie.Harq_ProcessNumberSizeDCI_1_3_r18 != nil {
			if err := e.EncodeInteger(*ie.Harq_ProcessNumberSizeDCI_1_3_r18, pDSCHConfigDCI13R18HarqProcessNumberSizeDCI13R18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PDSCH_ConfigDCI_1_3_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDSCHConfigDCI13R18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. resourceAllocationDCI-1-3-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(pDSCHConfigDCI13R18ResourceAllocationDCI13R18Constraints)
			if err != nil {
				return err
			}
			ie.ResourceAllocationDCI_1_3_r18 = &idx
		}
	}

	// 3. rbg-SizeDCI-1-3-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(pDSCHConfigDCI13R18RbgSizeDCI13R18Constraints)
			if err != nil {
				return err
			}
			ie.Rbg_SizeDCI_1_3_r18 = &idx
		}
	}

	// 4. resourceAllocationType1GranularityDCI-1-3-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(pDSCHConfigDCI13R18ResourceAllocationType1GranularityDCI13R18Constraints)
			if err != nil {
				return err
			}
			ie.ResourceAllocationType1GranularityDCI_1_3_r18 = &idx
		}
	}

	// 5. numberOfBitsForRV-DCI-1-3-r18: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(pDSCHConfigDCI13R18NumberOfBitsForRVDCI13R18Constraints)
			if err != nil {
				return err
			}
			ie.NumberOfBitsForRV_DCI_1_3_r18 = &v
		}
	}

	// 6. harq-ProcessNumberSizeDCI-1-3-r18: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(pDSCHConfigDCI13R18HarqProcessNumberSizeDCI13R18Constraints)
			if err != nil {
				return err
			}
			ie.Harq_ProcessNumberSizeDCI_1_3_r18 = &v
		}
	}

	return nil
}
