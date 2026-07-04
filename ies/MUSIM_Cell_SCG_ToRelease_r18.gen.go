// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MUSIM-Cell-SCG-ToRelease-r18 (line 2634).

var mUSIMCellSCGToReleaseR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "musim-CellToRelease-r18", Optional: true},
		{Name: "scg-ReleasePreference-r18", Optional: true},
	},
}

const (
	MUSIM_Cell_SCG_ToRelease_r18_Scg_ReleasePreference_r18_True = 0
)

var mUSIMCellSCGToReleaseR18ScgReleasePreferenceR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MUSIM_Cell_SCG_ToRelease_r18_Scg_ReleasePreference_r18_True},
}

type MUSIM_Cell_SCG_ToRelease_r18 struct {
	Musim_CellToRelease_r18   *MUSIM_CellToRelease_r18
	Scg_ReleasePreference_r18 *int64
}

func (ie *MUSIM_Cell_SCG_ToRelease_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mUSIMCellSCGToReleaseR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Musim_CellToRelease_r18 != nil, ie.Scg_ReleasePreference_r18 != nil}); err != nil {
		return err
	}

	// 2. musim-CellToRelease-r18: ref
	{
		if ie.Musim_CellToRelease_r18 != nil {
			if err := ie.Musim_CellToRelease_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. scg-ReleasePreference-r18: enumerated
	{
		if ie.Scg_ReleasePreference_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Scg_ReleasePreference_r18, mUSIMCellSCGToReleaseR18ScgReleasePreferenceR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MUSIM_Cell_SCG_ToRelease_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mUSIMCellSCGToReleaseR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. musim-CellToRelease-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Musim_CellToRelease_r18 = new(MUSIM_CellToRelease_r18)
			if err := ie.Musim_CellToRelease_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. scg-ReleasePreference-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(mUSIMCellSCGToReleaseR18ScgReleasePreferenceR18Constraints)
			if err != nil {
				return err
			}
			ie.Scg_ReleasePreference_r18 = &idx
		}
	}

	return nil
}
