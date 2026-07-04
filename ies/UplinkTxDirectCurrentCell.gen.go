// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UplinkTxDirectCurrentCell (line 16372).

var uplinkTxDirectCurrentCellConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "servCellIndex"},
		{Name: "uplinkDirectCurrentBWP"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var uplinkTxDirectCurrentCellUplinkDirectCurrentBWPConstraints = per.SizeRange(1, common.MaxNrofBWPs)

var uplinkTxDirectCurrentCellExtUplinkDirectCurrentBWPSULConstraints = per.SizeRange(1, common.MaxNrofBWPs)

type UplinkTxDirectCurrentCell struct {
	ServCellIndex              ServCellIndex
	UplinkDirectCurrentBWP     []UplinkTxDirectCurrentBWP
	UplinkDirectCurrentBWP_SUL []UplinkTxDirectCurrentBWP
}

func (ie *UplinkTxDirectCurrentCell) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uplinkTxDirectCurrentCellConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.UplinkDirectCurrentBWP_SUL != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. servCellIndex: ref
	{
		if err := ie.ServCellIndex.Encode(e); err != nil {
			return err
		}
	}

	// 3. uplinkDirectCurrentBWP: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(uplinkTxDirectCurrentCellUplinkDirectCurrentBWPConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.UplinkDirectCurrentBWP))); err != nil {
			return err
		}
		for i := range ie.UplinkDirectCurrentBWP {
			if err := ie.UplinkDirectCurrentBWP[i].Encode(e); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "uplinkDirectCurrentBWP-SUL", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.UplinkDirectCurrentBWP_SUL != nil}); err != nil {
				return err
			}

			if ie.UplinkDirectCurrentBWP_SUL != nil {
				seqOf := ex.NewSequenceOfEncoder(uplinkTxDirectCurrentCellExtUplinkDirectCurrentBWPSULConstraints)
				if err := seqOf.EncodeLength(int64(len(ie.UplinkDirectCurrentBWP_SUL))); err != nil {
					return err
				}
				for i := range ie.UplinkDirectCurrentBWP_SUL {
					if err := ie.UplinkDirectCurrentBWP_SUL[i].Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *UplinkTxDirectCurrentCell) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uplinkTxDirectCurrentCellConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. servCellIndex: ref
	{
		if err := ie.ServCellIndex.Decode(d); err != nil {
			return err
		}
	}

	// 3. uplinkDirectCurrentBWP: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(uplinkTxDirectCurrentCellUplinkDirectCurrentBWPConstraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.UplinkDirectCurrentBWP = make([]UplinkTxDirectCurrentBWP, n)
		for i := int64(0); i < n; i++ {
			if err := ie.UplinkDirectCurrentBWP[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "uplinkDirectCurrentBWP-SUL", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(uplinkTxDirectCurrentCellExtUplinkDirectCurrentBWPSULConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.UplinkDirectCurrentBWP_SUL = make([]UplinkTxDirectCurrentBWP, n)
			for i := int64(0); i < n; i++ {
				if err := ie.UplinkDirectCurrentBWP_SUL[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
