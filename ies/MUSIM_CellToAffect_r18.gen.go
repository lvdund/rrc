// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MUSIM-CellToAffect-r18 (line 2643).

var mUSIMCellToAffectR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "musim-ServCellIndex-r18"},
		{Name: "musim-MIMO-Layers-DL-r18", Optional: true},
		{Name: "musim-MIMO-Layers-UL-r18", Optional: true},
		{Name: "musim-SupportedBandwidth-DL-r18", Optional: true},
		{Name: "musim-SupportedBandwidth-UL-r18", Optional: true},
	},
}

var mUSIMCellToAffectR18MusimMIMOLayersDLR18Constraints = per.Constrained(1, 8)

var mUSIMCellToAffectR18MusimMIMOLayersULR18Constraints = per.Constrained(1, 4)

type MUSIM_CellToAffect_r18 struct {
	Musim_ServCellIndex_r18         ServCellIndex
	Musim_MIMO_Layers_DL_r18        *int64
	Musim_MIMO_Layers_UL_r18        *int64
	Musim_SupportedBandwidth_DL_r18 *SupportedBandwidth_v1700
	Musim_SupportedBandwidth_UL_r18 *SupportedBandwidth_v1700
}

func (ie *MUSIM_CellToAffect_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mUSIMCellToAffectR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Musim_MIMO_Layers_DL_r18 != nil, ie.Musim_MIMO_Layers_UL_r18 != nil, ie.Musim_SupportedBandwidth_DL_r18 != nil, ie.Musim_SupportedBandwidth_UL_r18 != nil}); err != nil {
		return err
	}

	// 2. musim-ServCellIndex-r18: ref
	{
		if err := ie.Musim_ServCellIndex_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. musim-MIMO-Layers-DL-r18: integer
	{
		if ie.Musim_MIMO_Layers_DL_r18 != nil {
			if err := e.EncodeInteger(*ie.Musim_MIMO_Layers_DL_r18, mUSIMCellToAffectR18MusimMIMOLayersDLR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. musim-MIMO-Layers-UL-r18: integer
	{
		if ie.Musim_MIMO_Layers_UL_r18 != nil {
			if err := e.EncodeInteger(*ie.Musim_MIMO_Layers_UL_r18, mUSIMCellToAffectR18MusimMIMOLayersULR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. musim-SupportedBandwidth-DL-r18: ref
	{
		if ie.Musim_SupportedBandwidth_DL_r18 != nil {
			if err := ie.Musim_SupportedBandwidth_DL_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. musim-SupportedBandwidth-UL-r18: ref
	{
		if ie.Musim_SupportedBandwidth_UL_r18 != nil {
			if err := ie.Musim_SupportedBandwidth_UL_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MUSIM_CellToAffect_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mUSIMCellToAffectR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. musim-ServCellIndex-r18: ref
	{
		if err := ie.Musim_ServCellIndex_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. musim-MIMO-Layers-DL-r18: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(mUSIMCellToAffectR18MusimMIMOLayersDLR18Constraints)
			if err != nil {
				return err
			}
			ie.Musim_MIMO_Layers_DL_r18 = &v
		}
	}

	// 4. musim-MIMO-Layers-UL-r18: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(mUSIMCellToAffectR18MusimMIMOLayersULR18Constraints)
			if err != nil {
				return err
			}
			ie.Musim_MIMO_Layers_UL_r18 = &v
		}
	}

	// 5. musim-SupportedBandwidth-DL-r18: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Musim_SupportedBandwidth_DL_r18 = new(SupportedBandwidth_v1700)
			if err := ie.Musim_SupportedBandwidth_DL_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. musim-SupportedBandwidth-UL-r18: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Musim_SupportedBandwidth_UL_r18 = new(SupportedBandwidth_v1700)
			if err := ie.Musim_SupportedBandwidth_UL_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
