// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MUSIM-MaxCC-r18 (line 2671).

var mUSIMMaxCCR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "musim-MaxCC-TotalDL-r18", Optional: true},
		{Name: "musim-MaxCC-TotalUL-r18", Optional: true},
		{Name: "musim-MaxCC-FR1-DL-r18", Optional: true},
		{Name: "musim-MaxCC-FR1-UL-r18", Optional: true},
		{Name: "musim-MaxCC-FR2-1-DL-r18", Optional: true},
		{Name: "musim-MaxCC-FR2-1-UL-r18", Optional: true},
		{Name: "musim-MaxCC-FR2-2-DL-r18", Optional: true},
		{Name: "musim-MaxCC-FR2-2-UL-r18", Optional: true},
	},
}

var mUSIMMaxCCR18MusimMaxCCTotalDLR18Constraints = per.Constrained(1, 32)

var mUSIMMaxCCR18MusimMaxCCTotalULR18Constraints = per.Constrained(1, 32)

var mUSIMMaxCCR18MusimMaxCCFR1DLR18Constraints = per.Constrained(1, 32)

var mUSIMMaxCCR18MusimMaxCCFR1ULR18Constraints = per.Constrained(1, 32)

var mUSIMMaxCCR18MusimMaxCCFR21DLR18Constraints = per.Constrained(1, 32)

var mUSIMMaxCCR18MusimMaxCCFR21ULR18Constraints = per.Constrained(1, 32)

var mUSIMMaxCCR18MusimMaxCCFR22DLR18Constraints = per.Constrained(1, 32)

var mUSIMMaxCCR18MusimMaxCCFR22ULR18Constraints = per.Constrained(1, 32)

type MUSIM_MaxCC_r18 struct {
	Musim_MaxCC_TotalDL_r18  *int64
	Musim_MaxCC_TotalUL_r18  *int64
	Musim_MaxCC_FR1_DL_r18   *int64
	Musim_MaxCC_FR1_UL_r18   *int64
	Musim_MaxCC_FR2_1_DL_r18 *int64
	Musim_MaxCC_FR2_1_UL_r18 *int64
	Musim_MaxCC_FR2_2_DL_r18 *int64
	Musim_MaxCC_FR2_2_UL_r18 *int64
}

func (ie *MUSIM_MaxCC_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mUSIMMaxCCR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Musim_MaxCC_TotalDL_r18 != nil, ie.Musim_MaxCC_TotalUL_r18 != nil, ie.Musim_MaxCC_FR1_DL_r18 != nil, ie.Musim_MaxCC_FR1_UL_r18 != nil, ie.Musim_MaxCC_FR2_1_DL_r18 != nil, ie.Musim_MaxCC_FR2_1_UL_r18 != nil, ie.Musim_MaxCC_FR2_2_DL_r18 != nil, ie.Musim_MaxCC_FR2_2_UL_r18 != nil}); err != nil {
		return err
	}

	// 2. musim-MaxCC-TotalDL-r18: integer
	{
		if ie.Musim_MaxCC_TotalDL_r18 != nil {
			if err := e.EncodeInteger(*ie.Musim_MaxCC_TotalDL_r18, mUSIMMaxCCR18MusimMaxCCTotalDLR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. musim-MaxCC-TotalUL-r18: integer
	{
		if ie.Musim_MaxCC_TotalUL_r18 != nil {
			if err := e.EncodeInteger(*ie.Musim_MaxCC_TotalUL_r18, mUSIMMaxCCR18MusimMaxCCTotalULR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. musim-MaxCC-FR1-DL-r18: integer
	{
		if ie.Musim_MaxCC_FR1_DL_r18 != nil {
			if err := e.EncodeInteger(*ie.Musim_MaxCC_FR1_DL_r18, mUSIMMaxCCR18MusimMaxCCFR1DLR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. musim-MaxCC-FR1-UL-r18: integer
	{
		if ie.Musim_MaxCC_FR1_UL_r18 != nil {
			if err := e.EncodeInteger(*ie.Musim_MaxCC_FR1_UL_r18, mUSIMMaxCCR18MusimMaxCCFR1ULR18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. musim-MaxCC-FR2-1-DL-r18: integer
	{
		if ie.Musim_MaxCC_FR2_1_DL_r18 != nil {
			if err := e.EncodeInteger(*ie.Musim_MaxCC_FR2_1_DL_r18, mUSIMMaxCCR18MusimMaxCCFR21DLR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. musim-MaxCC-FR2-1-UL-r18: integer
	{
		if ie.Musim_MaxCC_FR2_1_UL_r18 != nil {
			if err := e.EncodeInteger(*ie.Musim_MaxCC_FR2_1_UL_r18, mUSIMMaxCCR18MusimMaxCCFR21ULR18Constraints); err != nil {
				return err
			}
		}
	}

	// 8. musim-MaxCC-FR2-2-DL-r18: integer
	{
		if ie.Musim_MaxCC_FR2_2_DL_r18 != nil {
			if err := e.EncodeInteger(*ie.Musim_MaxCC_FR2_2_DL_r18, mUSIMMaxCCR18MusimMaxCCFR22DLR18Constraints); err != nil {
				return err
			}
		}
	}

	// 9. musim-MaxCC-FR2-2-UL-r18: integer
	{
		if ie.Musim_MaxCC_FR2_2_UL_r18 != nil {
			if err := e.EncodeInteger(*ie.Musim_MaxCC_FR2_2_UL_r18, mUSIMMaxCCR18MusimMaxCCFR22ULR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MUSIM_MaxCC_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mUSIMMaxCCR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. musim-MaxCC-TotalDL-r18: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(mUSIMMaxCCR18MusimMaxCCTotalDLR18Constraints)
			if err != nil {
				return err
			}
			ie.Musim_MaxCC_TotalDL_r18 = &v
		}
	}

	// 3. musim-MaxCC-TotalUL-r18: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(mUSIMMaxCCR18MusimMaxCCTotalULR18Constraints)
			if err != nil {
				return err
			}
			ie.Musim_MaxCC_TotalUL_r18 = &v
		}
	}

	// 4. musim-MaxCC-FR1-DL-r18: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(mUSIMMaxCCR18MusimMaxCCFR1DLR18Constraints)
			if err != nil {
				return err
			}
			ie.Musim_MaxCC_FR1_DL_r18 = &v
		}
	}

	// 5. musim-MaxCC-FR1-UL-r18: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(mUSIMMaxCCR18MusimMaxCCFR1ULR18Constraints)
			if err != nil {
				return err
			}
			ie.Musim_MaxCC_FR1_UL_r18 = &v
		}
	}

	// 6. musim-MaxCC-FR2-1-DL-r18: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(mUSIMMaxCCR18MusimMaxCCFR21DLR18Constraints)
			if err != nil {
				return err
			}
			ie.Musim_MaxCC_FR2_1_DL_r18 = &v
		}
	}

	// 7. musim-MaxCC-FR2-1-UL-r18: integer
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeInteger(mUSIMMaxCCR18MusimMaxCCFR21ULR18Constraints)
			if err != nil {
				return err
			}
			ie.Musim_MaxCC_FR2_1_UL_r18 = &v
		}
	}

	// 8. musim-MaxCC-FR2-2-DL-r18: integer
	{
		if seq.IsComponentPresent(6) {
			v, err := d.DecodeInteger(mUSIMMaxCCR18MusimMaxCCFR22DLR18Constraints)
			if err != nil {
				return err
			}
			ie.Musim_MaxCC_FR2_2_DL_r18 = &v
		}
	}

	// 9. musim-MaxCC-FR2-2-UL-r18: integer
	{
		if seq.IsComponentPresent(7) {
			v, err := d.DecodeInteger(mUSIMMaxCCR18MusimMaxCCFR22ULR18Constraints)
			if err != nil {
				return err
			}
			ie.Musim_MaxCC_FR2_2_UL_r18 = &v
		}
	}

	return nil
}
