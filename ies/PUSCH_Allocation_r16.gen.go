// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUSCH-Allocation-r16 (line 12717).

var pUSCHAllocationR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "mappingType-r16", Optional: true},
		{Name: "startSymbolAndLength-r16", Optional: true},
		{Name: "startSymbol-r16", Optional: true},
		{Name: "length-r16", Optional: true},
		{Name: "numberOfRepetitions-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

const (
	PUSCH_Allocation_r16_MappingType_r16_TypeA = 0
	PUSCH_Allocation_r16_MappingType_r16_TypeB = 1
)

var pUSCHAllocationR16MappingTypeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Allocation_r16_MappingType_r16_TypeA, PUSCH_Allocation_r16_MappingType_r16_TypeB},
}

var pUSCHAllocationR16StartSymbolAndLengthR16Constraints = per.Constrained(0, 127)

var pUSCHAllocationR16StartSymbolR16Constraints = per.Constrained(0, 13)

var pUSCHAllocationR16LengthR16Constraints = per.Constrained(1, 14)

const (
	PUSCH_Allocation_r16_NumberOfRepetitions_r16_N1  = 0
	PUSCH_Allocation_r16_NumberOfRepetitions_r16_N2  = 1
	PUSCH_Allocation_r16_NumberOfRepetitions_r16_N3  = 2
	PUSCH_Allocation_r16_NumberOfRepetitions_r16_N4  = 3
	PUSCH_Allocation_r16_NumberOfRepetitions_r16_N7  = 4
	PUSCH_Allocation_r16_NumberOfRepetitions_r16_N8  = 5
	PUSCH_Allocation_r16_NumberOfRepetitions_r16_N12 = 6
	PUSCH_Allocation_r16_NumberOfRepetitions_r16_N16 = 7
)

var pUSCHAllocationR16NumberOfRepetitionsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Allocation_r16_NumberOfRepetitions_r16_N1, PUSCH_Allocation_r16_NumberOfRepetitions_r16_N2, PUSCH_Allocation_r16_NumberOfRepetitions_r16_N3, PUSCH_Allocation_r16_NumberOfRepetitions_r16_N4, PUSCH_Allocation_r16_NumberOfRepetitions_r16_N7, PUSCH_Allocation_r16_NumberOfRepetitions_r16_N8, PUSCH_Allocation_r16_NumberOfRepetitions_r16_N12, PUSCH_Allocation_r16_NumberOfRepetitions_r16_N16},
}

const (
	PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_N1     = 0
	PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_N2     = 1
	PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_N3     = 2
	PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_N4     = 3
	PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_N7     = 4
	PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_N8     = 5
	PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_N12    = 6
	PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_N16    = 7
	PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_N20    = 8
	PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_N24    = 9
	PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_N28    = 10
	PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_N32    = 11
	PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_Spare4 = 12
	PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_Spare3 = 13
	PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_Spare2 = 14
	PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_Spare1 = 15
)

var pUSCHAllocationR16ExtNumberOfRepetitionsExtR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_N1, PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_N2, PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_N3, PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_N4, PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_N7, PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_N8, PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_N12, PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_N16, PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_N20, PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_N24, PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_N28, PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_N32, PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_Spare4, PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_Spare3, PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_Spare2, PUSCH_Allocation_r16_Ext_NumberOfRepetitionsExt_r17_Spare1},
}

const (
	PUSCH_Allocation_r16_Ext_NumberOfSlotsTBoMS_r17_N1     = 0
	PUSCH_Allocation_r16_Ext_NumberOfSlotsTBoMS_r17_N2     = 1
	PUSCH_Allocation_r16_Ext_NumberOfSlotsTBoMS_r17_N4     = 2
	PUSCH_Allocation_r16_Ext_NumberOfSlotsTBoMS_r17_N8     = 3
	PUSCH_Allocation_r16_Ext_NumberOfSlotsTBoMS_r17_Spare4 = 4
	PUSCH_Allocation_r16_Ext_NumberOfSlotsTBoMS_r17_Spare3 = 5
	PUSCH_Allocation_r16_Ext_NumberOfSlotsTBoMS_r17_Spare2 = 6
	PUSCH_Allocation_r16_Ext_NumberOfSlotsTBoMS_r17_Spare1 = 7
)

var pUSCHAllocationR16ExtNumberOfSlotsTBoMSR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Allocation_r16_Ext_NumberOfSlotsTBoMS_r17_N1, PUSCH_Allocation_r16_Ext_NumberOfSlotsTBoMS_r17_N2, PUSCH_Allocation_r16_Ext_NumberOfSlotsTBoMS_r17_N4, PUSCH_Allocation_r16_Ext_NumberOfSlotsTBoMS_r17_N8, PUSCH_Allocation_r16_Ext_NumberOfSlotsTBoMS_r17_Spare4, PUSCH_Allocation_r16_Ext_NumberOfSlotsTBoMS_r17_Spare3, PUSCH_Allocation_r16_Ext_NumberOfSlotsTBoMS_r17_Spare2, PUSCH_Allocation_r16_Ext_NumberOfSlotsTBoMS_r17_Spare1},
}

var pUSCHAllocationR16ExtendedK2R17Constraints = per.Constrained(0, 128)

const (
	PUSCH_Allocation_r16_Ext_Ul_MutingIndicator_r19_Enabled = 0
)

var pUSCHAllocationR16ExtUlMutingIndicatorR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Allocation_r16_Ext_Ul_MutingIndicator_r19_Enabled},
}

const (
	PUSCH_Allocation_r16_Ext_Occ_Length_r19_N2 = 0
	PUSCH_Allocation_r16_Ext_Occ_Length_r19_N4 = 1
)

var pUSCHAllocationR16ExtOccLengthR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Allocation_r16_Ext_Occ_Length_r19_N2, PUSCH_Allocation_r16_Ext_Occ_Length_r19_N4},
}

type PUSCH_Allocation_r16 struct {
	MappingType_r16            *int64
	StartSymbolAndLength_r16   *int64
	StartSymbol_r16            *int64
	Length_r16                 *int64
	NumberOfRepetitions_r16    *int64
	NumberOfRepetitionsExt_r17 *int64
	NumberOfSlotsTBoMS_r17     *int64
	ExtendedK2_r17             *int64
	Ul_MutingIndicator_r19     *int64
	Occ_Length_r19             *int64
}

func (ie *PUSCH_Allocation_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUSCHAllocationR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.NumberOfRepetitionsExt_r17 != nil || ie.NumberOfSlotsTBoMS_r17 != nil || ie.ExtendedK2_r17 != nil
	hasExtGroup1 := ie.Ul_MutingIndicator_r19 != nil || ie.Occ_Length_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MappingType_r16 != nil, ie.StartSymbolAndLength_r16 != nil, ie.StartSymbol_r16 != nil, ie.Length_r16 != nil, ie.NumberOfRepetitions_r16 != nil}); err != nil {
		return err
	}

	// 3. mappingType-r16: enumerated
	{
		if ie.MappingType_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MappingType_r16, pUSCHAllocationR16MappingTypeR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. startSymbolAndLength-r16: integer
	{
		if ie.StartSymbolAndLength_r16 != nil {
			if err := e.EncodeInteger(*ie.StartSymbolAndLength_r16, pUSCHAllocationR16StartSymbolAndLengthR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. startSymbol-r16: integer
	{
		if ie.StartSymbol_r16 != nil {
			if err := e.EncodeInteger(*ie.StartSymbol_r16, pUSCHAllocationR16StartSymbolR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. length-r16: integer
	{
		if ie.Length_r16 != nil {
			if err := e.EncodeInteger(*ie.Length_r16, pUSCHAllocationR16LengthR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. numberOfRepetitions-r16: enumerated
	{
		if ie.NumberOfRepetitions_r16 != nil {
			if err := e.EncodeEnumerated(*ie.NumberOfRepetitions_r16, pUSCHAllocationR16NumberOfRepetitionsR16Constraints); err != nil {
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
					{Name: "numberOfRepetitionsExt-r17", Optional: true},
					{Name: "numberOfSlotsTBoMS-r17", Optional: true},
					{Name: "extendedK2-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.NumberOfRepetitionsExt_r17 != nil, ie.NumberOfSlotsTBoMS_r17 != nil, ie.ExtendedK2_r17 != nil}); err != nil {
				return err
			}

			if ie.NumberOfRepetitionsExt_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.NumberOfRepetitionsExt_r17, pUSCHAllocationR16ExtNumberOfRepetitionsExtR17Constraints); err != nil {
					return err
				}
			}

			if ie.NumberOfSlotsTBoMS_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.NumberOfSlotsTBoMS_r17, pUSCHAllocationR16ExtNumberOfSlotsTBoMSR17Constraints); err != nil {
					return err
				}
			}

			if ie.ExtendedK2_r17 != nil {
				if err := ex.EncodeInteger(*ie.ExtendedK2_r17, pUSCHAllocationR16ExtendedK2R17Constraints); err != nil {
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
					{Name: "ul-MutingIndicator-r19", Optional: true},
					{Name: "occ-Length-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ul_MutingIndicator_r19 != nil, ie.Occ_Length_r19 != nil}); err != nil {
				return err
			}

			if ie.Ul_MutingIndicator_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Ul_MutingIndicator_r19, pUSCHAllocationR16ExtUlMutingIndicatorR19Constraints); err != nil {
					return err
				}
			}

			if ie.Occ_Length_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Occ_Length_r19, pUSCHAllocationR16ExtOccLengthR19Constraints); err != nil {
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

func (ie *PUSCH_Allocation_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUSCHAllocationR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. mappingType-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(pUSCHAllocationR16MappingTypeR16Constraints)
			if err != nil {
				return err
			}
			ie.MappingType_r16 = &idx
		}
	}

	// 4. startSymbolAndLength-r16: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(pUSCHAllocationR16StartSymbolAndLengthR16Constraints)
			if err != nil {
				return err
			}
			ie.StartSymbolAndLength_r16 = &v
		}
	}

	// 5. startSymbol-r16: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(pUSCHAllocationR16StartSymbolR16Constraints)
			if err != nil {
				return err
			}
			ie.StartSymbol_r16 = &v
		}
	}

	// 6. length-r16: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(pUSCHAllocationR16LengthR16Constraints)
			if err != nil {
				return err
			}
			ie.Length_r16 = &v
		}
	}

	// 7. numberOfRepetitions-r16: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(pUSCHAllocationR16NumberOfRepetitionsR16Constraints)
			if err != nil {
				return err
			}
			ie.NumberOfRepetitions_r16 = &idx
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
				{Name: "numberOfRepetitionsExt-r17", Optional: true},
				{Name: "numberOfSlotsTBoMS-r17", Optional: true},
				{Name: "extendedK2-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(pUSCHAllocationR16ExtNumberOfRepetitionsExtR17Constraints)
			if err != nil {
				return err
			}
			ie.NumberOfRepetitionsExt_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(pUSCHAllocationR16ExtNumberOfSlotsTBoMSR17Constraints)
			if err != nil {
				return err
			}
			ie.NumberOfSlotsTBoMS_r17 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeInteger(pUSCHAllocationR16ExtendedK2R17Constraints)
			if err != nil {
				return err
			}
			ie.ExtendedK2_r17 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "ul-MutingIndicator-r19", Optional: true},
				{Name: "occ-Length-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(pUSCHAllocationR16ExtUlMutingIndicatorR19Constraints)
			if err != nil {
				return err
			}
			ie.Ul_MutingIndicator_r19 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(pUSCHAllocationR16ExtOccLengthR19Constraints)
			if err != nil {
				return err
			}
			ie.Occ_Length_r19 = &v
		}
	}

	return nil
}
