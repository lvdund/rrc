// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRS-SwitchingTimeNR (line 25384).

var sRSSwitchingTimeNRConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "switchingTimeDL", Optional: true},
		{Name: "switchingTimeUL", Optional: true},
	},
}

const (
	SRS_SwitchingTimeNR_SwitchingTimeDL_N0us   = 0
	SRS_SwitchingTimeNR_SwitchingTimeDL_N30us  = 1
	SRS_SwitchingTimeNR_SwitchingTimeDL_N100us = 2
	SRS_SwitchingTimeNR_SwitchingTimeDL_N140us = 3
	SRS_SwitchingTimeNR_SwitchingTimeDL_N200us = 4
	SRS_SwitchingTimeNR_SwitchingTimeDL_N300us = 5
	SRS_SwitchingTimeNR_SwitchingTimeDL_N500us = 6
	SRS_SwitchingTimeNR_SwitchingTimeDL_N900us = 7
)

var sRSSwitchingTimeNRSwitchingTimeDLConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_SwitchingTimeNR_SwitchingTimeDL_N0us, SRS_SwitchingTimeNR_SwitchingTimeDL_N30us, SRS_SwitchingTimeNR_SwitchingTimeDL_N100us, SRS_SwitchingTimeNR_SwitchingTimeDL_N140us, SRS_SwitchingTimeNR_SwitchingTimeDL_N200us, SRS_SwitchingTimeNR_SwitchingTimeDL_N300us, SRS_SwitchingTimeNR_SwitchingTimeDL_N500us, SRS_SwitchingTimeNR_SwitchingTimeDL_N900us},
}

const (
	SRS_SwitchingTimeNR_SwitchingTimeUL_N0us   = 0
	SRS_SwitchingTimeNR_SwitchingTimeUL_N30us  = 1
	SRS_SwitchingTimeNR_SwitchingTimeUL_N100us = 2
	SRS_SwitchingTimeNR_SwitchingTimeUL_N140us = 3
	SRS_SwitchingTimeNR_SwitchingTimeUL_N200us = 4
	SRS_SwitchingTimeNR_SwitchingTimeUL_N300us = 5
	SRS_SwitchingTimeNR_SwitchingTimeUL_N500us = 6
	SRS_SwitchingTimeNR_SwitchingTimeUL_N900us = 7
)

var sRSSwitchingTimeNRSwitchingTimeULConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRS_SwitchingTimeNR_SwitchingTimeUL_N0us, SRS_SwitchingTimeNR_SwitchingTimeUL_N30us, SRS_SwitchingTimeNR_SwitchingTimeUL_N100us, SRS_SwitchingTimeNR_SwitchingTimeUL_N140us, SRS_SwitchingTimeNR_SwitchingTimeUL_N200us, SRS_SwitchingTimeNR_SwitchingTimeUL_N300us, SRS_SwitchingTimeNR_SwitchingTimeUL_N500us, SRS_SwitchingTimeNR_SwitchingTimeUL_N900us},
}

type SRS_SwitchingTimeNR struct {
	SwitchingTimeDL *int64
	SwitchingTimeUL *int64
}

func (ie *SRS_SwitchingTimeNR) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSSwitchingTimeNRConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SwitchingTimeDL != nil, ie.SwitchingTimeUL != nil}); err != nil {
		return err
	}

	// 2. switchingTimeDL: enumerated
	{
		if ie.SwitchingTimeDL != nil {
			if err := e.EncodeEnumerated(*ie.SwitchingTimeDL, sRSSwitchingTimeNRSwitchingTimeDLConstraints); err != nil {
				return err
			}
		}
	}

	// 3. switchingTimeUL: enumerated
	{
		if ie.SwitchingTimeUL != nil {
			if err := e.EncodeEnumerated(*ie.SwitchingTimeUL, sRSSwitchingTimeNRSwitchingTimeULConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SRS_SwitchingTimeNR) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSSwitchingTimeNRConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. switchingTimeDL: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sRSSwitchingTimeNRSwitchingTimeDLConstraints)
			if err != nil {
				return err
			}
			ie.SwitchingTimeDL = &idx
		}
	}

	// 3. switchingTimeUL: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sRSSwitchingTimeNRSwitchingTimeULConstraints)
			if err != nil {
				return err
			}
			ie.SwitchingTimeUL = &idx
		}
	}

	return nil
}
