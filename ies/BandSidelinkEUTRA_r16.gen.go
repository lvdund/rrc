// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandSidelinkEUTRA-r16 (line 25092).

var bandSidelinkEUTRAR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "freqBandSidelinkEUTRA-r16"},
		{Name: "gnb-ScheduledMode3SidelinkEUTRA-r16", Optional: true},
		{Name: "gnb-ScheduledMode4SidelinkEUTRA-r16", Optional: true},
	},
}

const (
	BandSidelinkEUTRA_r16_Gnb_ScheduledMode4SidelinkEUTRA_r16_Supported = 0
)

var bandSidelinkEUTRAR16GnbScheduledMode4SidelinkEUTRAR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelinkEUTRA_r16_Gnb_ScheduledMode4SidelinkEUTRA_r16_Supported},
}

const (
	BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms0       = 0
	BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms0dot25  = 1
	BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms0dot5   = 2
	BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms0dot625 = 3
	BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms0dot75  = 4
	BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms1       = 5
	BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms1dot25  = 6
	BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms1dot5   = 7
	BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms1dot75  = 8
	BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms2       = 9
	BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms2dot5   = 10
	BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms3       = 11
	BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms4       = 12
	BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms5       = 13
	BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms6       = 14
	BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms8       = 15
	BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms10      = 16
	BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms20      = 17
)

var bandSidelinkEUTRAR16GnbScheduledMode3SidelinkEUTRAR16GnbScheduledMode3DelaySidelinkEUTRAR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms0, BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms0dot25, BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms0dot5, BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms0dot625, BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms0dot75, BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms1, BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms1dot25, BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms1dot5, BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms1dot75, BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms2, BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms2dot5, BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms3, BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms4, BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms5, BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms6, BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms8, BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms10, BandSidelinkEUTRA_r16_Gnb_ScheduledMode3SidelinkEUTRA_r16_Gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Ms20},
}

type BandSidelinkEUTRA_r16 struct {
	FreqBandSidelinkEUTRA_r16           FreqBandIndicatorEUTRA
	Gnb_ScheduledMode3SidelinkEUTRA_r16 *struct{ Gnb_ScheduledMode3DelaySidelinkEUTRA_r16 int64 }
	Gnb_ScheduledMode4SidelinkEUTRA_r16 *int64
}

func (ie *BandSidelinkEUTRA_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandSidelinkEUTRAR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Gnb_ScheduledMode3SidelinkEUTRA_r16 != nil, ie.Gnb_ScheduledMode4SidelinkEUTRA_r16 != nil}); err != nil {
		return err
	}

	// 2. freqBandSidelinkEUTRA-r16: ref
	{
		if err := ie.FreqBandSidelinkEUTRA_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. gnb-ScheduledMode3SidelinkEUTRA-r16: sequence
	{
		if ie.Gnb_ScheduledMode3SidelinkEUTRA_r16 != nil {
			c := ie.Gnb_ScheduledMode3SidelinkEUTRA_r16
			if err := e.EncodeEnumerated(c.Gnb_ScheduledMode3DelaySidelinkEUTRA_r16, bandSidelinkEUTRAR16GnbScheduledMode3SidelinkEUTRAR16GnbScheduledMode3DelaySidelinkEUTRAR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. gnb-ScheduledMode4SidelinkEUTRA-r16: enumerated
	{
		if ie.Gnb_ScheduledMode4SidelinkEUTRA_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Gnb_ScheduledMode4SidelinkEUTRA_r16, bandSidelinkEUTRAR16GnbScheduledMode4SidelinkEUTRAR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandSidelinkEUTRA_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandSidelinkEUTRAR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. freqBandSidelinkEUTRA-r16: ref
	{
		if err := ie.FreqBandSidelinkEUTRA_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. gnb-ScheduledMode3SidelinkEUTRA-r16: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.Gnb_ScheduledMode3SidelinkEUTRA_r16 = &struct{ Gnb_ScheduledMode3DelaySidelinkEUTRA_r16 int64 }{}
			c := ie.Gnb_ScheduledMode3SidelinkEUTRA_r16
			{
				v, err := d.DecodeEnumerated(bandSidelinkEUTRAR16GnbScheduledMode3SidelinkEUTRAR16GnbScheduledMode3DelaySidelinkEUTRAR16Constraints)
				if err != nil {
					return err
				}
				c.Gnb_ScheduledMode3DelaySidelinkEUTRA_r16 = v
			}
		}
	}

	// 4. gnb-ScheduledMode4SidelinkEUTRA-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(bandSidelinkEUTRAR16GnbScheduledMode4SidelinkEUTRAR16Constraints)
			if err != nil {
				return err
			}
			ie.Gnb_ScheduledMode4SidelinkEUTRA_r16 = &idx
		}
	}

	return nil
}
