// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SSB-ToMeasureAltitudeBased-r18 (line 9598).

var sSBToMeasureAltitudeBasedR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "altitudeRange-r18"},
		{Name: "ssb-ToMeasure-r18", Optional: true},
	},
}

var sSBToMeasureAltitudeBasedR18AltitudeRangeR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "altitudeMin-r18", Optional: true},
		{Name: "altitudeMax-r18", Optional: true},
		{Name: "altitudeHyst-r18", Optional: true},
	},
}

type SSB_ToMeasureAltitudeBased_r18 struct {
	AltitudeRange_r18 struct {
		AltitudeMin_r18  *Altitude_r18
		AltitudeMax_r18  *Altitude_r18
		AltitudeHyst_r18 *HysteresisAltitude_r18
	}
	Ssb_ToMeasure_r18 *SSB_ToMeasure
}

func (ie *SSB_ToMeasureAltitudeBased_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sSBToMeasureAltitudeBasedR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ssb_ToMeasure_r18 != nil}); err != nil {
		return err
	}

	// 2. altitudeRange-r18: sequence
	{
		{
			c := &ie.AltitudeRange_r18
			sSBToMeasureAltitudeBasedR18AltitudeRangeR18Seq := e.NewSequenceEncoder(sSBToMeasureAltitudeBasedR18AltitudeRangeR18Constraints)
			if err := sSBToMeasureAltitudeBasedR18AltitudeRangeR18Seq.EncodePreamble([]bool{c.AltitudeMin_r18 != nil, c.AltitudeMax_r18 != nil, c.AltitudeHyst_r18 != nil}); err != nil {
				return err
			}
			if c.AltitudeMin_r18 != nil {
				if err := c.AltitudeMin_r18.Encode(e); err != nil {
					return err
				}
			}
			if c.AltitudeMax_r18 != nil {
				if err := c.AltitudeMax_r18.Encode(e); err != nil {
					return err
				}
			}
			if c.AltitudeHyst_r18 != nil {
				if err := c.AltitudeHyst_r18.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. ssb-ToMeasure-r18: ref
	{
		if ie.Ssb_ToMeasure_r18 != nil {
			if err := ie.Ssb_ToMeasure_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SSB_ToMeasureAltitudeBased_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sSBToMeasureAltitudeBasedR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. altitudeRange-r18: sequence
	{
		{
			c := &ie.AltitudeRange_r18
			sSBToMeasureAltitudeBasedR18AltitudeRangeR18Seq := d.NewSequenceDecoder(sSBToMeasureAltitudeBasedR18AltitudeRangeR18Constraints)
			if err := sSBToMeasureAltitudeBasedR18AltitudeRangeR18Seq.DecodePreamble(); err != nil {
				return err
			}
			if sSBToMeasureAltitudeBasedR18AltitudeRangeR18Seq.IsComponentPresent(0) {
				c.AltitudeMin_r18 = new(Altitude_r18)
				if err := (*c.AltitudeMin_r18).Decode(d); err != nil {
					return err
				}
			}
			if sSBToMeasureAltitudeBasedR18AltitudeRangeR18Seq.IsComponentPresent(1) {
				c.AltitudeMax_r18 = new(Altitude_r18)
				if err := (*c.AltitudeMax_r18).Decode(d); err != nil {
					return err
				}
			}
			if sSBToMeasureAltitudeBasedR18AltitudeRangeR18Seq.IsComponentPresent(2) {
				c.AltitudeHyst_r18 = new(HysteresisAltitude_r18)
				if err := (*c.AltitudeHyst_r18).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. ssb-ToMeasure-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Ssb_ToMeasure_r18 = new(SSB_ToMeasure)
			if err := ie.Ssb_ToMeasure_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
