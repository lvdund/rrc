// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: EUTRA-FreqNeighCellInfo (line 4172).

var eUTRAFreqNeighCellInfoConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "physCellId"},
		{Name: "dummy"},
		{Name: "q-RxLevMinOffsetCell", Optional: true},
		{Name: "q-QualMinOffsetCell", Optional: true},
	},
}

var eUTRAFreqNeighCellInfoQRxLevMinOffsetCellConstraints = per.Constrained(1, 8)

var eUTRAFreqNeighCellInfoQQualMinOffsetCellConstraints = per.Constrained(1, 8)

type EUTRA_FreqNeighCellInfo struct {
	PhysCellId           EUTRA_PhysCellId
	Dummy                EUTRA_Q_OffsetRange
	Q_RxLevMinOffsetCell *int64
	Q_QualMinOffsetCell  *int64
}

func (ie *EUTRA_FreqNeighCellInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(eUTRAFreqNeighCellInfoConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Q_RxLevMinOffsetCell != nil, ie.Q_QualMinOffsetCell != nil}); err != nil {
		return err
	}

	// 2. physCellId: ref
	{
		if err := ie.PhysCellId.Encode(e); err != nil {
			return err
		}
	}

	// 3. dummy: ref
	{
		if err := ie.Dummy.Encode(e); err != nil {
			return err
		}
	}

	// 4. q-RxLevMinOffsetCell: integer
	{
		if ie.Q_RxLevMinOffsetCell != nil {
			if err := e.EncodeInteger(*ie.Q_RxLevMinOffsetCell, eUTRAFreqNeighCellInfoQRxLevMinOffsetCellConstraints); err != nil {
				return err
			}
		}
	}

	// 5. q-QualMinOffsetCell: integer
	{
		if ie.Q_QualMinOffsetCell != nil {
			if err := e.EncodeInteger(*ie.Q_QualMinOffsetCell, eUTRAFreqNeighCellInfoQQualMinOffsetCellConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *EUTRA_FreqNeighCellInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(eUTRAFreqNeighCellInfoConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. physCellId: ref
	{
		if err := ie.PhysCellId.Decode(d); err != nil {
			return err
		}
	}

	// 3. dummy: ref
	{
		if err := ie.Dummy.Decode(d); err != nil {
			return err
		}
	}

	// 4. q-RxLevMinOffsetCell: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(eUTRAFreqNeighCellInfoQRxLevMinOffsetCellConstraints)
			if err != nil {
				return err
			}
			ie.Q_RxLevMinOffsetCell = &v
		}
	}

	// 5. q-QualMinOffsetCell: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(eUTRAFreqNeighCellInfoQQualMinOffsetCellConstraints)
			if err != nil {
				return err
			}
			ie.Q_QualMinOffsetCell = &v
		}
	}

	return nil
}
