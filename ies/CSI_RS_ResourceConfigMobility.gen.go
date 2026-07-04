// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CSI-RS-ResourceConfigMobility (line 7516).

var cSIRSResourceConfigMobilityConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "subcarrierSpacing"},
		{Name: "csi-RS-CellList-Mobility"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var cSIRSResourceConfigMobilityCsiRSCellListMobilityConstraints = per.SizeRange(1, common.MaxNrofCSI_RS_CellsRRM)

type CSI_RS_ResourceConfigMobility struct {
	SubcarrierSpacing        SubcarrierSpacing
	Csi_RS_CellList_Mobility []CSI_RS_CellMobility
	RefServCellIndex         *ServCellIndex
}

func (ie *CSI_RS_ResourceConfigMobility) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSIRSResourceConfigMobilityConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.RefServCellIndex != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. subcarrierSpacing: ref
	{
		if err := ie.SubcarrierSpacing.Encode(e); err != nil {
			return err
		}
	}

	// 3. csi-RS-CellList-Mobility: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(cSIRSResourceConfigMobilityCsiRSCellListMobilityConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.Csi_RS_CellList_Mobility))); err != nil {
			return err
		}
		for i := range ie.Csi_RS_CellList_Mobility {
			if err := ie.Csi_RS_CellList_Mobility[i].Encode(e); err != nil {
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
					{Name: "refServCellIndex", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.RefServCellIndex != nil}); err != nil {
				return err
			}

			if ie.RefServCellIndex != nil {
				if err := ie.RefServCellIndex.Encode(ex); err != nil {
					return err
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

func (ie *CSI_RS_ResourceConfigMobility) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSIRSResourceConfigMobilityConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. subcarrierSpacing: ref
	{
		if err := ie.SubcarrierSpacing.Decode(d); err != nil {
			return err
		}
	}

	// 3. csi-RS-CellList-Mobility: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(cSIRSResourceConfigMobilityCsiRSCellListMobilityConstraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Csi_RS_CellList_Mobility = make([]CSI_RS_CellMobility, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Csi_RS_CellList_Mobility[i].Decode(d); err != nil {
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
				{Name: "refServCellIndex", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.RefServCellIndex = new(ServCellIndex)
			if err := ie.RefServCellIndex.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
