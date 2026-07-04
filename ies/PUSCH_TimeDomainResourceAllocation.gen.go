// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUSCH-TimeDomainResourceAllocation (line 12703).

var pUSCHTimeDomainResourceAllocationConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "k2", Optional: true},
		{Name: "mappingType"},
		{Name: "startSymbolAndLength"},
	},
}

var pUSCHTimeDomainResourceAllocationK2Constraints = per.Constrained(0, 32)

const (
	PUSCH_TimeDomainResourceAllocation_MappingType_TypeA = 0
	PUSCH_TimeDomainResourceAllocation_MappingType_TypeB = 1
)

var pUSCHTimeDomainResourceAllocationMappingTypeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_TimeDomainResourceAllocation_MappingType_TypeA, PUSCH_TimeDomainResourceAllocation_MappingType_TypeB},
}

var pUSCHTimeDomainResourceAllocationStartSymbolAndLengthConstraints = per.Constrained(0, 127)

type PUSCH_TimeDomainResourceAllocation struct {
	K2                   *int64
	MappingType          int64
	StartSymbolAndLength int64
}

func (ie *PUSCH_TimeDomainResourceAllocation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUSCHTimeDomainResourceAllocationConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.K2 != nil}); err != nil {
		return err
	}

	// 2. k2: integer
	{
		if ie.K2 != nil {
			if err := e.EncodeInteger(*ie.K2, pUSCHTimeDomainResourceAllocationK2Constraints); err != nil {
				return err
			}
		}
	}

	// 3. mappingType: enumerated
	{
		if err := e.EncodeEnumerated(ie.MappingType, pUSCHTimeDomainResourceAllocationMappingTypeConstraints); err != nil {
			return err
		}
	}

	// 4. startSymbolAndLength: integer
	{
		if err := e.EncodeInteger(ie.StartSymbolAndLength, pUSCHTimeDomainResourceAllocationStartSymbolAndLengthConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PUSCH_TimeDomainResourceAllocation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUSCHTimeDomainResourceAllocationConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. k2: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(pUSCHTimeDomainResourceAllocationK2Constraints)
			if err != nil {
				return err
			}
			ie.K2 = &v
		}
	}

	// 3. mappingType: enumerated
	{
		v1, err := d.DecodeEnumerated(pUSCHTimeDomainResourceAllocationMappingTypeConstraints)
		if err != nil {
			return err
		}
		ie.MappingType = v1
	}

	// 4. startSymbolAndLength: integer
	{
		v2, err := d.DecodeInteger(pUSCHTimeDomainResourceAllocationStartSymbolAndLengthConstraints)
		if err != nil {
			return err
		}
		ie.StartSymbolAndLength = v2
	}

	return nil
}
