// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDSCH-TimeDomainResourceAllocation-r16 (line 11523).

var pDSCHTimeDomainResourceAllocationR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "k0-r16", Optional: true},
		{Name: "mappingType-r16"},
		{Name: "startSymbolAndLength-r16"},
		{Name: "repetitionNumber-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var pDSCHTimeDomainResourceAllocationR16K0R16Constraints = per.Constrained(0, 32)

const (
	PDSCH_TimeDomainResourceAllocation_r16_MappingType_r16_TypeA = 0
	PDSCH_TimeDomainResourceAllocation_r16_MappingType_r16_TypeB = 1
)

var pDSCHTimeDomainResourceAllocationR16MappingTypeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_TimeDomainResourceAllocation_r16_MappingType_r16_TypeA, PDSCH_TimeDomainResourceAllocation_r16_MappingType_r16_TypeB},
}

var pDSCHTimeDomainResourceAllocationR16StartSymbolAndLengthR16Constraints = per.Constrained(0, 127)

const (
	PDSCH_TimeDomainResourceAllocation_r16_RepetitionNumber_r16_N2  = 0
	PDSCH_TimeDomainResourceAllocation_r16_RepetitionNumber_r16_N3  = 1
	PDSCH_TimeDomainResourceAllocation_r16_RepetitionNumber_r16_N4  = 2
	PDSCH_TimeDomainResourceAllocation_r16_RepetitionNumber_r16_N5  = 3
	PDSCH_TimeDomainResourceAllocation_r16_RepetitionNumber_r16_N6  = 4
	PDSCH_TimeDomainResourceAllocation_r16_RepetitionNumber_r16_N7  = 5
	PDSCH_TimeDomainResourceAllocation_r16_RepetitionNumber_r16_N8  = 6
	PDSCH_TimeDomainResourceAllocation_r16_RepetitionNumber_r16_N16 = 7
)

var pDSCHTimeDomainResourceAllocationR16RepetitionNumberR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_TimeDomainResourceAllocation_r16_RepetitionNumber_r16_N2, PDSCH_TimeDomainResourceAllocation_r16_RepetitionNumber_r16_N3, PDSCH_TimeDomainResourceAllocation_r16_RepetitionNumber_r16_N4, PDSCH_TimeDomainResourceAllocation_r16_RepetitionNumber_r16_N5, PDSCH_TimeDomainResourceAllocation_r16_RepetitionNumber_r16_N6, PDSCH_TimeDomainResourceAllocation_r16_RepetitionNumber_r16_N7, PDSCH_TimeDomainResourceAllocation_r16_RepetitionNumber_r16_N8, PDSCH_TimeDomainResourceAllocation_r16_RepetitionNumber_r16_N16},
}

var pDSCHTimeDomainResourceAllocationR16K0V1710Constraints = per.Constrained(33, 128)

const (
	PDSCH_TimeDomainResourceAllocation_r16_Ext_RepetitionNumber_v1730_N2  = 0
	PDSCH_TimeDomainResourceAllocation_r16_Ext_RepetitionNumber_v1730_N3  = 1
	PDSCH_TimeDomainResourceAllocation_r16_Ext_RepetitionNumber_v1730_N4  = 2
	PDSCH_TimeDomainResourceAllocation_r16_Ext_RepetitionNumber_v1730_N5  = 3
	PDSCH_TimeDomainResourceAllocation_r16_Ext_RepetitionNumber_v1730_N6  = 4
	PDSCH_TimeDomainResourceAllocation_r16_Ext_RepetitionNumber_v1730_N7  = 5
	PDSCH_TimeDomainResourceAllocation_r16_Ext_RepetitionNumber_v1730_N8  = 6
	PDSCH_TimeDomainResourceAllocation_r16_Ext_RepetitionNumber_v1730_N16 = 7
)

var pDSCHTimeDomainResourceAllocationR16ExtRepetitionNumberV1730Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_TimeDomainResourceAllocation_r16_Ext_RepetitionNumber_v1730_N2, PDSCH_TimeDomainResourceAllocation_r16_Ext_RepetitionNumber_v1730_N3, PDSCH_TimeDomainResourceAllocation_r16_Ext_RepetitionNumber_v1730_N4, PDSCH_TimeDomainResourceAllocation_r16_Ext_RepetitionNumber_v1730_N5, PDSCH_TimeDomainResourceAllocation_r16_Ext_RepetitionNumber_v1730_N6, PDSCH_TimeDomainResourceAllocation_r16_Ext_RepetitionNumber_v1730_N7, PDSCH_TimeDomainResourceAllocation_r16_Ext_RepetitionNumber_v1730_N8, PDSCH_TimeDomainResourceAllocation_r16_Ext_RepetitionNumber_v1730_N16},
}

type PDSCH_TimeDomainResourceAllocation_r16 struct {
	K0_r16                   *int64
	MappingType_r16          int64
	StartSymbolAndLength_r16 int64
	RepetitionNumber_r16     *int64
	K0_v1710                 *int64
	RepetitionNumber_v1730   *int64
}

func (ie *PDSCH_TimeDomainResourceAllocation_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDSCHTimeDomainResourceAllocationR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.K0_v1710 != nil
	hasExtGroup1 := ie.RepetitionNumber_v1730 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.K0_r16 != nil, ie.RepetitionNumber_r16 != nil}); err != nil {
		return err
	}

	// 3. k0-r16: integer
	{
		if ie.K0_r16 != nil {
			if err := e.EncodeInteger(*ie.K0_r16, pDSCHTimeDomainResourceAllocationR16K0R16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. mappingType-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.MappingType_r16, pDSCHTimeDomainResourceAllocationR16MappingTypeR16Constraints); err != nil {
			return err
		}
	}

	// 5. startSymbolAndLength-r16: integer
	{
		if err := e.EncodeInteger(ie.StartSymbolAndLength_r16, pDSCHTimeDomainResourceAllocationR16StartSymbolAndLengthR16Constraints); err != nil {
			return err
		}
	}

	// 6. repetitionNumber-r16: enumerated
	{
		if ie.RepetitionNumber_r16 != nil {
			if err := e.EncodeEnumerated(*ie.RepetitionNumber_r16, pDSCHTimeDomainResourceAllocationR16RepetitionNumberR16Constraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "k0-v1710", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.K0_v1710 != nil}); err != nil {
				return err
			}

			if ie.K0_v1710 != nil {
				if err := ex.EncodeInteger(*ie.K0_v1710, pDSCHTimeDomainResourceAllocationR16K0V1710Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "repetitionNumber-v1730", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.RepetitionNumber_v1730 != nil}); err != nil {
				return err
			}

			if ie.RepetitionNumber_v1730 != nil {
				if err := ex.EncodeEnumerated(*ie.RepetitionNumber_v1730, pDSCHTimeDomainResourceAllocationR16ExtRepetitionNumberV1730Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PDSCH_TimeDomainResourceAllocation_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDSCHTimeDomainResourceAllocationR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. k0-r16: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(pDSCHTimeDomainResourceAllocationR16K0R16Constraints)
			if err != nil {
				return err
			}
			ie.K0_r16 = &v
		}
	}

	// 4. mappingType-r16: enumerated
	{
		v1, err := d.DecodeEnumerated(pDSCHTimeDomainResourceAllocationR16MappingTypeR16Constraints)
		if err != nil {
			return err
		}
		ie.MappingType_r16 = v1
	}

	// 5. startSymbolAndLength-r16: integer
	{
		v2, err := d.DecodeInteger(pDSCHTimeDomainResourceAllocationR16StartSymbolAndLengthR16Constraints)
		if err != nil {
			return err
		}
		ie.StartSymbolAndLength_r16 = v2
	}

	// 6. repetitionNumber-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(pDSCHTimeDomainResourceAllocationR16RepetitionNumberR16Constraints)
			if err != nil {
				return err
			}
			ie.RepetitionNumber_r16 = &idx
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "k0-v1710", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(pDSCHTimeDomainResourceAllocationR16K0V1710Constraints)
			if err != nil {
				return err
			}
			ie.K0_v1710 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "repetitionNumber-v1730", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(pDSCHTimeDomainResourceAllocationR16ExtRepetitionNumberV1730Constraints)
			if err != nil {
				return err
			}
			ie.RepetitionNumber_v1730 = &v
		}
	}

	return nil
}
