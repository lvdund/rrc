// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SliceInfo-r17 (line 8383).

var sliceInfoR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "nsag-IdentityInfo-r17"},
		{Name: "nsag-CellReselectionPriority-r17", Optional: true},
		{Name: "nsag-CellReselectionSubPriority-r17", Optional: true},
		{Name: "sliceCellListNR-r17", Optional: true},
	},
}

var sliceInfo_r17SliceCellListNRR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sliceAllowedCellListNR-r17"},
		{Name: "sliceExcludedCellListNR-r17"},
	},
}

const (
	SliceInfo_r17_SliceCellListNR_r17_SliceAllowedCellListNR_r17  = 0
	SliceInfo_r17_SliceCellListNR_r17_SliceExcludedCellListNR_r17 = 1
)

type SliceInfo_r17 struct {
	Nsag_IdentityInfo_r17               NSAG_IdentityInfo_r17
	Nsag_CellReselectionPriority_r17    *CellReselectionPriority
	Nsag_CellReselectionSubPriority_r17 *CellReselectionSubPriority
	SliceCellListNR_r17                 *struct {
		Choice                      int
		SliceAllowedCellListNR_r17  *SliceCellListNR_r17
		SliceExcludedCellListNR_r17 *SliceCellListNR_r17
	}
}

func (ie *SliceInfo_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sliceInfoR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Nsag_CellReselectionPriority_r17 != nil, ie.Nsag_CellReselectionSubPriority_r17 != nil, ie.SliceCellListNR_r17 != nil}); err != nil {
		return err
	}

	// 2. nsag-IdentityInfo-r17: ref
	{
		if err := ie.Nsag_IdentityInfo_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. nsag-CellReselectionPriority-r17: ref
	{
		if ie.Nsag_CellReselectionPriority_r17 != nil {
			if err := ie.Nsag_CellReselectionPriority_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. nsag-CellReselectionSubPriority-r17: ref
	{
		if ie.Nsag_CellReselectionSubPriority_r17 != nil {
			if err := ie.Nsag_CellReselectionSubPriority_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. sliceCellListNR-r17: choice
	{
		if ie.SliceCellListNR_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(sliceInfo_r17SliceCellListNRR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.SliceCellListNR_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.SliceCellListNR_r17).Choice {
			case SliceInfo_r17_SliceCellListNR_r17_SliceAllowedCellListNR_r17:
				if err := (*ie.SliceCellListNR_r17).SliceAllowedCellListNR_r17.Encode(e); err != nil {
					return err
				}
			case SliceInfo_r17_SliceCellListNR_r17_SliceExcludedCellListNR_r17:
				if err := (*ie.SliceCellListNR_r17).SliceExcludedCellListNR_r17.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.SliceCellListNR_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *SliceInfo_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sliceInfoR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. nsag-IdentityInfo-r17: ref
	{
		if err := ie.Nsag_IdentityInfo_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. nsag-CellReselectionPriority-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Nsag_CellReselectionPriority_r17 = new(CellReselectionPriority)
			if err := ie.Nsag_CellReselectionPriority_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. nsag-CellReselectionSubPriority-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Nsag_CellReselectionSubPriority_r17 = new(CellReselectionSubPriority)
			if err := ie.Nsag_CellReselectionSubPriority_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. sliceCellListNR-r17: choice
	{
		if seq.IsComponentPresent(3) {
			ie.SliceCellListNR_r17 = &struct {
				Choice                      int
				SliceAllowedCellListNR_r17  *SliceCellListNR_r17
				SliceExcludedCellListNR_r17 *SliceCellListNR_r17
			}{}
			choiceDec := d.NewChoiceDecoder(sliceInfo_r17SliceCellListNRR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.SliceCellListNR_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SliceInfo_r17_SliceCellListNR_r17_SliceAllowedCellListNR_r17:
				(*ie.SliceCellListNR_r17).SliceAllowedCellListNR_r17 = new(SliceCellListNR_r17)
				if err := (*ie.SliceCellListNR_r17).SliceAllowedCellListNR_r17.Decode(d); err != nil {
					return err
				}
			case SliceInfo_r17_SliceCellListNR_r17_SliceExcludedCellListNR_r17:
				(*ie.SliceCellListNR_r17).SliceExcludedCellListNR_r17 = new(SliceCellListNR_r17)
				if err := (*ie.SliceCellListNR_r17).SliceExcludedCellListNR_r17.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
