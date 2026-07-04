// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Q-OffsetRangeList (line 9525).

var qOffsetRangeListConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rsrpOffsetSSB", Optional: true},
		{Name: "rsrqOffsetSSB", Optional: true},
		{Name: "sinrOffsetSSB", Optional: true},
		{Name: "rsrpOffsetCSI-RS", Optional: true},
		{Name: "rsrqOffsetCSI-RS", Optional: true},
		{Name: "sinrOffsetCSI-RS", Optional: true},
	},
}

type Q_OffsetRangeList struct {
	RsrpOffsetSSB    *Q_OffsetRange
	RsrqOffsetSSB    *Q_OffsetRange
	SinrOffsetSSB    *Q_OffsetRange
	RsrpOffsetCSI_RS *Q_OffsetRange
	RsrqOffsetCSI_RS *Q_OffsetRange
	SinrOffsetCSI_RS *Q_OffsetRange
}

func (ie *Q_OffsetRangeList) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(qOffsetRangeListConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.RsrpOffsetSSB != nil && ie.RsrpOffsetSSB.Value != Q_OffsetRange_DB0, ie.RsrqOffsetSSB != nil && ie.RsrqOffsetSSB.Value != Q_OffsetRange_DB0, ie.SinrOffsetSSB != nil && ie.SinrOffsetSSB.Value != Q_OffsetRange_DB0, ie.RsrpOffsetCSI_RS != nil && ie.RsrpOffsetCSI_RS.Value != Q_OffsetRange_DB0, ie.RsrqOffsetCSI_RS != nil && ie.RsrqOffsetCSI_RS.Value != Q_OffsetRange_DB0, ie.SinrOffsetCSI_RS != nil && ie.SinrOffsetCSI_RS.Value != Q_OffsetRange_DB0}); err != nil {
		return err
	}

	// 2. rsrpOffsetSSB: ref
	{
		if ie.RsrpOffsetSSB != nil && ie.RsrpOffsetSSB.Value != Q_OffsetRange_DB0 {
			if err := ie.RsrpOffsetSSB.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. rsrqOffsetSSB: ref
	{
		if ie.RsrqOffsetSSB != nil && ie.RsrqOffsetSSB.Value != Q_OffsetRange_DB0 {
			if err := ie.RsrqOffsetSSB.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sinrOffsetSSB: ref
	{
		if ie.SinrOffsetSSB != nil && ie.SinrOffsetSSB.Value != Q_OffsetRange_DB0 {
			if err := ie.SinrOffsetSSB.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. rsrpOffsetCSI-RS: ref
	{
		if ie.RsrpOffsetCSI_RS != nil && ie.RsrpOffsetCSI_RS.Value != Q_OffsetRange_DB0 {
			if err := ie.RsrpOffsetCSI_RS.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. rsrqOffsetCSI-RS: ref
	{
		if ie.RsrqOffsetCSI_RS != nil && ie.RsrqOffsetCSI_RS.Value != Q_OffsetRange_DB0 {
			if err := ie.RsrqOffsetCSI_RS.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. sinrOffsetCSI-RS: ref
	{
		if ie.SinrOffsetCSI_RS != nil && ie.SinrOffsetCSI_RS.Value != Q_OffsetRange_DB0 {
			if err := ie.SinrOffsetCSI_RS.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *Q_OffsetRangeList) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(qOffsetRangeListConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. rsrpOffsetSSB: ref
	{
		if seq.IsComponentPresent(0) {
			ie.RsrpOffsetSSB = new(Q_OffsetRange)
			if err := ie.RsrpOffsetSSB.Decode(d); err != nil {
				return err
			}
		} else {
			ie.RsrpOffsetSSB = &Q_OffsetRange{Value: Q_OffsetRange_DB0}
		}
	}

	// 3. rsrqOffsetSSB: ref
	{
		if seq.IsComponentPresent(1) {
			ie.RsrqOffsetSSB = new(Q_OffsetRange)
			if err := ie.RsrqOffsetSSB.Decode(d); err != nil {
				return err
			}
		} else {
			ie.RsrqOffsetSSB = &Q_OffsetRange{Value: Q_OffsetRange_DB0}
		}
	}

	// 4. sinrOffsetSSB: ref
	{
		if seq.IsComponentPresent(2) {
			ie.SinrOffsetSSB = new(Q_OffsetRange)
			if err := ie.SinrOffsetSSB.Decode(d); err != nil {
				return err
			}
		} else {
			ie.SinrOffsetSSB = &Q_OffsetRange{Value: Q_OffsetRange_DB0}
		}
	}

	// 5. rsrpOffsetCSI-RS: ref
	{
		if seq.IsComponentPresent(3) {
			ie.RsrpOffsetCSI_RS = new(Q_OffsetRange)
			if err := ie.RsrpOffsetCSI_RS.Decode(d); err != nil {
				return err
			}
		} else {
			ie.RsrpOffsetCSI_RS = &Q_OffsetRange{Value: Q_OffsetRange_DB0}
		}
	}

	// 6. rsrqOffsetCSI-RS: ref
	{
		if seq.IsComponentPresent(4) {
			ie.RsrqOffsetCSI_RS = new(Q_OffsetRange)
			if err := ie.RsrqOffsetCSI_RS.Decode(d); err != nil {
				return err
			}
		} else {
			ie.RsrqOffsetCSI_RS = &Q_OffsetRange{Value: Q_OffsetRange_DB0}
		}
	}

	// 7. sinrOffsetCSI-RS: ref
	{
		if seq.IsComponentPresent(5) {
			ie.SinrOffsetCSI_RS = new(Q_OffsetRange)
			if err := ie.SinrOffsetCSI_RS.Decode(d); err != nil {
				return err
			}
		} else {
			ie.SinrOffsetCSI_RS = &Q_OffsetRange{Value: Q_OffsetRange_DB0}
		}
	}

	return nil
}
