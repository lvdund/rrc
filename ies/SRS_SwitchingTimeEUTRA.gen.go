// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRS-SwitchingTimeEUTRA (line 25392).

var sRSSwitchingTimeEUTRAConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "switchingTimeDL", Optional: true},
		{Name: "switchingTimeUL", Optional: true},
	},
}

const (
	SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N0     = 0
	SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N0dot5 = 1
	SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N1     = 2
	SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N1dot5 = 3
	SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N2     = 4
	SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N2dot5 = 5
	SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N3     = 6
	SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N3dot5 = 7
	SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N4     = 8
	SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N4dot5 = 9
	SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N5     = 10
	SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N5dot5 = 11
	SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N6     = 12
	SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N6dot5 = 13
	SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N7     = 14
)

var sRSSwitchingTimeEUTRASwitchingTimeDLConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N0, SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N0dot5, SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N1, SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N1dot5, SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N2, SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N2dot5, SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N3, SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N3dot5, SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N4, SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N4dot5, SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N5, SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N5dot5, SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N6, SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N6dot5, SRS_SwitchingTimeEUTRA_SwitchingTimeDL_N7},
}

const (
	SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N0     = 0
	SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N0dot5 = 1
	SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N1     = 2
	SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N1dot5 = 3
	SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N2     = 4
	SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N2dot5 = 5
	SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N3     = 6
	SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N3dot5 = 7
	SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N4     = 8
	SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N4dot5 = 9
	SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N5     = 10
	SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N5dot5 = 11
	SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N6     = 12
	SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N6dot5 = 13
	SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N7     = 14
)

var sRSSwitchingTimeEUTRASwitchingTimeULConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N0, SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N0dot5, SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N1, SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N1dot5, SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N2, SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N2dot5, SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N3, SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N3dot5, SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N4, SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N4dot5, SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N5, SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N5dot5, SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N6, SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N6dot5, SRS_SwitchingTimeEUTRA_SwitchingTimeUL_N7},
}

type SRS_SwitchingTimeEUTRA struct {
	SwitchingTimeDL *int64
	SwitchingTimeUL *int64
}

func (ie *SRS_SwitchingTimeEUTRA) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSSwitchingTimeEUTRAConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SwitchingTimeDL != nil, ie.SwitchingTimeUL != nil}); err != nil {
		return err
	}

	// 2. switchingTimeDL: enumerated
	{
		if ie.SwitchingTimeDL != nil {
			if err := e.EncodeEnumerated(*ie.SwitchingTimeDL, sRSSwitchingTimeEUTRASwitchingTimeDLConstraints); err != nil {
				return err
			}
		}
	}

	// 3. switchingTimeUL: enumerated
	{
		if ie.SwitchingTimeUL != nil {
			if err := e.EncodeEnumerated(*ie.SwitchingTimeUL, sRSSwitchingTimeEUTRASwitchingTimeULConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SRS_SwitchingTimeEUTRA) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSSwitchingTimeEUTRAConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. switchingTimeDL: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sRSSwitchingTimeEUTRASwitchingTimeDLConstraints)
			if err != nil {
				return err
			}
			ie.SwitchingTimeDL = &idx
		}
	}

	// 3. switchingTimeUL: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sRSSwitchingTimeEUTRASwitchingTimeULConstraints)
			if err != nil {
				return err
			}
			ie.SwitchingTimeUL = &idx
		}
	}

	return nil
}
