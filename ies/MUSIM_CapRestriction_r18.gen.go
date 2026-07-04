// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MUSIM-CapRestriction-r18 (line 2626).

var mUSIMCapRestrictionR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "musim-Cell-SCG-ToRelease-r18", Optional: true},
		{Name: "musim-CellToAffectList-r18", Optional: true},
		{Name: "musim-AffectedBandsList-r18", Optional: true},
		{Name: "musim-AvoidedBandsList-r18", Optional: true},
		{Name: "musim-MaxCC-r18", Optional: true},
	},
}

type MUSIM_CapRestriction_r18 struct {
	Musim_Cell_SCG_ToRelease_r18 *MUSIM_Cell_SCG_ToRelease_r18
	Musim_CellToAffectList_r18   *MUSIM_CellToAffectList_r18
	Musim_AffectedBandsList_r18  *MUSIM_AffectedBandsList_r18
	Musim_AvoidedBandsList_r18   *MUSIM_AvoidedBandsList_r18
	Musim_MaxCC_r18              *MUSIM_MaxCC_r18
}

func (ie *MUSIM_CapRestriction_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mUSIMCapRestrictionR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Musim_Cell_SCG_ToRelease_r18 != nil, ie.Musim_CellToAffectList_r18 != nil, ie.Musim_AffectedBandsList_r18 != nil, ie.Musim_AvoidedBandsList_r18 != nil, ie.Musim_MaxCC_r18 != nil}); err != nil {
		return err
	}

	// 2. musim-Cell-SCG-ToRelease-r18: ref
	{
		if ie.Musim_Cell_SCG_ToRelease_r18 != nil {
			if err := ie.Musim_Cell_SCG_ToRelease_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. musim-CellToAffectList-r18: ref
	{
		if ie.Musim_CellToAffectList_r18 != nil {
			if err := ie.Musim_CellToAffectList_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. musim-AffectedBandsList-r18: ref
	{
		if ie.Musim_AffectedBandsList_r18 != nil {
			if err := ie.Musim_AffectedBandsList_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. musim-AvoidedBandsList-r18: ref
	{
		if ie.Musim_AvoidedBandsList_r18 != nil {
			if err := ie.Musim_AvoidedBandsList_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. musim-MaxCC-r18: ref
	{
		if ie.Musim_MaxCC_r18 != nil {
			if err := ie.Musim_MaxCC_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MUSIM_CapRestriction_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mUSIMCapRestrictionR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. musim-Cell-SCG-ToRelease-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Musim_Cell_SCG_ToRelease_r18 = new(MUSIM_Cell_SCG_ToRelease_r18)
			if err := ie.Musim_Cell_SCG_ToRelease_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. musim-CellToAffectList-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Musim_CellToAffectList_r18 = new(MUSIM_CellToAffectList_r18)
			if err := ie.Musim_CellToAffectList_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. musim-AffectedBandsList-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Musim_AffectedBandsList_r18 = new(MUSIM_AffectedBandsList_r18)
			if err := ie.Musim_AffectedBandsList_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. musim-AvoidedBandsList-r18: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Musim_AvoidedBandsList_r18 = new(MUSIM_AvoidedBandsList_r18)
			if err := ie.Musim_AvoidedBandsList_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. musim-MaxCC-r18: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Musim_MaxCC_r18 = new(MUSIM_MaxCC_r18)
			if err := ie.Musim_MaxCC_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
