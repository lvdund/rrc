// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IntraFreqNeighCellInfo (line 3900).

var intraFreqNeighCellInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "physCellId"},
		{Name: "q-OffsetCell"},
		{Name: "q-RxLevMinOffsetCell", Optional: true},
		{Name: "q-RxLevMinOffsetCellSUL", Optional: true},
		{Name: "q-QualMinOffsetCell", Optional: true},
	},
}

var intraFreqNeighCellInfoQRxLevMinOffsetCellConstraints = per.Constrained(1, 8)

var intraFreqNeighCellInfoQRxLevMinOffsetCellSULConstraints = per.Constrained(1, 8)

var intraFreqNeighCellInfoQQualMinOffsetCellConstraints = per.Constrained(1, 8)

type IntraFreqNeighCellInfo struct {
	PhysCellId              PhysCellId
	Q_OffsetCell            Q_OffsetRange
	Q_RxLevMinOffsetCell    *int64
	Q_RxLevMinOffsetCellSUL *int64
	Q_QualMinOffsetCell     *int64
}

func (ie *IntraFreqNeighCellInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(intraFreqNeighCellInfoConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Q_RxLevMinOffsetCell != nil, ie.Q_RxLevMinOffsetCellSUL != nil, ie.Q_QualMinOffsetCell != nil}); err != nil {
		return err
	}

	// 3. physCellId: ref
	{
		if err := ie.PhysCellId.Encode(e); err != nil {
			return err
		}
	}

	// 4. q-OffsetCell: ref
	{
		if err := ie.Q_OffsetCell.Encode(e); err != nil {
			return err
		}
	}

	// 5. q-RxLevMinOffsetCell: integer
	{
		if ie.Q_RxLevMinOffsetCell != nil {
			if err := e.EncodeInteger(*ie.Q_RxLevMinOffsetCell, intraFreqNeighCellInfoQRxLevMinOffsetCellConstraints); err != nil {
				return err
			}
		}
	}

	// 6. q-RxLevMinOffsetCellSUL: integer
	{
		if ie.Q_RxLevMinOffsetCellSUL != nil {
			if err := e.EncodeInteger(*ie.Q_RxLevMinOffsetCellSUL, intraFreqNeighCellInfoQRxLevMinOffsetCellSULConstraints); err != nil {
				return err
			}
		}
	}

	// 7. q-QualMinOffsetCell: integer
	{
		if ie.Q_QualMinOffsetCell != nil {
			if err := e.EncodeInteger(*ie.Q_QualMinOffsetCell, intraFreqNeighCellInfoQQualMinOffsetCellConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *IntraFreqNeighCellInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(intraFreqNeighCellInfoConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. physCellId: ref
	{
		if err := ie.PhysCellId.Decode(d); err != nil {
			return err
		}
	}

	// 4. q-OffsetCell: ref
	{
		if err := ie.Q_OffsetCell.Decode(d); err != nil {
			return err
		}
	}

	// 5. q-RxLevMinOffsetCell: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(intraFreqNeighCellInfoQRxLevMinOffsetCellConstraints)
			if err != nil {
				return err
			}
			ie.Q_RxLevMinOffsetCell = &v
		}
	}

	// 6. q-RxLevMinOffsetCellSUL: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(intraFreqNeighCellInfoQRxLevMinOffsetCellSULConstraints)
			if err != nil {
				return err
			}
			ie.Q_RxLevMinOffsetCellSUL = &v
		}
	}

	// 7. q-QualMinOffsetCell: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(intraFreqNeighCellInfoQQualMinOffsetCellConstraints)
			if err != nil {
				return err
			}
			ie.Q_QualMinOffsetCell = &v
		}
	}

	return nil
}
