// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CellReselectionPriorities (line 1316).

var cellReselectionPrioritiesConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "freqPriorityListEUTRA", Optional: true},
		{Name: "freqPriorityListNR", Optional: true},
		{Name: "t320", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	CellReselectionPriorities_T320_Min5   = 0
	CellReselectionPriorities_T320_Min10  = 1
	CellReselectionPriorities_T320_Min20  = 2
	CellReselectionPriorities_T320_Min30  = 3
	CellReselectionPriorities_T320_Min60  = 4
	CellReselectionPriorities_T320_Min120 = 5
	CellReselectionPriorities_T320_Min180 = 6
	CellReselectionPriorities_T320_Spare1 = 7
)

var cellReselectionPrioritiesT320Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CellReselectionPriorities_T320_Min5, CellReselectionPriorities_T320_Min10, CellReselectionPriorities_T320_Min20, CellReselectionPriorities_T320_Min30, CellReselectionPriorities_T320_Min60, CellReselectionPriorities_T320_Min120, CellReselectionPriorities_T320_Min180, CellReselectionPriorities_T320_Spare1},
}

type CellReselectionPriorities struct {
	FreqPriorityListEUTRA                *FreqPriorityListEUTRA
	FreqPriorityListNR                   *FreqPriorityListNR
	T320                                 *int64
	FreqPriorityListDedicatedSlicing_r17 *FreqPriorityListDedicatedSlicing_r17
}

func (ie *CellReselectionPriorities) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cellReselectionPrioritiesConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.FreqPriorityListDedicatedSlicing_r17 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FreqPriorityListEUTRA != nil, ie.FreqPriorityListNR != nil, ie.T320 != nil}); err != nil {
		return err
	}

	// 3. freqPriorityListEUTRA: ref
	{
		if ie.FreqPriorityListEUTRA != nil {
			if err := ie.FreqPriorityListEUTRA.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. freqPriorityListNR: ref
	{
		if ie.FreqPriorityListNR != nil {
			if err := ie.FreqPriorityListNR.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. t320: enumerated
	{
		if ie.T320 != nil {
			if err := e.EncodeEnumerated(*ie.T320, cellReselectionPrioritiesT320Constraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "freqPriorityListDedicatedSlicing-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.FreqPriorityListDedicatedSlicing_r17 != nil}); err != nil {
				return err
			}

			if ie.FreqPriorityListDedicatedSlicing_r17 != nil {
				if err := ie.FreqPriorityListDedicatedSlicing_r17.Encode(ex); err != nil {
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

func (ie *CellReselectionPriorities) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cellReselectionPrioritiesConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. freqPriorityListEUTRA: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FreqPriorityListEUTRA = new(FreqPriorityListEUTRA)
			if err := ie.FreqPriorityListEUTRA.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. freqPriorityListNR: ref
	{
		if seq.IsComponentPresent(1) {
			ie.FreqPriorityListNR = new(FreqPriorityListNR)
			if err := ie.FreqPriorityListNR.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. t320: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(cellReselectionPrioritiesT320Constraints)
			if err != nil {
				return err
			}
			ie.T320 = &idx
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
				{Name: "freqPriorityListDedicatedSlicing-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.FreqPriorityListDedicatedSlicing_r17 = new(FreqPriorityListDedicatedSlicing_r17)
			if err := ie.FreqPriorityListDedicatedSlicing_r17.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
