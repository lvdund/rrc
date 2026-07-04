// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDSCH-TimeDomainResourceAllocation (line 11515).

var pDSCHTimeDomainResourceAllocationConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "k0", Optional: true},
		{Name: "mappingType"},
		{Name: "startSymbolAndLength"},
	},
}

var pDSCHTimeDomainResourceAllocationK0Constraints = per.Constrained(0, 32)

const (
	PDSCH_TimeDomainResourceAllocation_MappingType_TypeA = 0
	PDSCH_TimeDomainResourceAllocation_MappingType_TypeB = 1
)

var pDSCHTimeDomainResourceAllocationMappingTypeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_TimeDomainResourceAllocation_MappingType_TypeA, PDSCH_TimeDomainResourceAllocation_MappingType_TypeB},
}

var pDSCHTimeDomainResourceAllocationStartSymbolAndLengthConstraints = per.Constrained(0, 127)

type PDSCH_TimeDomainResourceAllocation struct {
	K0                   *int64
	MappingType          int64
	StartSymbolAndLength int64
}

func (ie *PDSCH_TimeDomainResourceAllocation) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDSCHTimeDomainResourceAllocationConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.K0 != nil}); err != nil {
		return err
	}

	// 2. k0: integer
	{
		if ie.K0 != nil {
			if err := e.EncodeInteger(*ie.K0, pDSCHTimeDomainResourceAllocationK0Constraints); err != nil {
				return err
			}
		}
	}

	// 3. mappingType: enumerated
	{
		if err := e.EncodeEnumerated(ie.MappingType, pDSCHTimeDomainResourceAllocationMappingTypeConstraints); err != nil {
			return err
		}
	}

	// 4. startSymbolAndLength: integer
	{
		if err := e.EncodeInteger(ie.StartSymbolAndLength, pDSCHTimeDomainResourceAllocationStartSymbolAndLengthConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PDSCH_TimeDomainResourceAllocation) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDSCHTimeDomainResourceAllocationConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. k0: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(pDSCHTimeDomainResourceAllocationK0Constraints)
			if err != nil {
				return err
			}
			ie.K0 = &v
		}
	}

	// 3. mappingType: enumerated
	{
		v1, err := d.DecodeEnumerated(pDSCHTimeDomainResourceAllocationMappingTypeConstraints)
		if err != nil {
			return err
		}
		ie.MappingType = v1
	}

	// 4. startSymbolAndLength: integer
	{
		v2, err := d.DecodeInteger(pDSCHTimeDomainResourceAllocationStartSymbolAndLengthConstraints)
		if err != nil {
			return err
		}
		ie.StartSymbolAndLength = v2
	}

	return nil
}
