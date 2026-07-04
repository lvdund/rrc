// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUCCH-TPC-CommandConfig (line 12323).

var pUCCHTPCCommandConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "tpc-IndexPCell", Optional: true},
		{Name: "tpc-IndexPUCCH-SCell", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var pUCCHTPCCommandConfigTpcIndexPCellConstraints = per.Constrained(1, 15)

var pUCCHTPCCommandConfigTpcIndexPUCCHSCellConstraints = per.Constrained(1, 15)

var pUCCHTPCCommandConfigTpcIndexPUCCHSSCellR17Constraints = per.Constrained(1, 15)

var pUCCHTPCCommandConfigTpcIndexPUCCHSScellSecondaryPUCCHgroupR17Constraints = per.Constrained(1, 15)

type PUCCH_TPC_CommandConfig struct {
	Tpc_IndexPCell                               *int64
	Tpc_IndexPUCCH_SCell                         *int64
	Tpc_IndexPUCCH_SSCell_r17                    *int64
	Tpc_IndexPUCCH_SScellSecondaryPUCCHgroup_r17 *int64
}

func (ie *PUCCH_TPC_CommandConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUCCHTPCCommandConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Tpc_IndexPUCCH_SSCell_r17 != nil || ie.Tpc_IndexPUCCH_SScellSecondaryPUCCHgroup_r17 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Tpc_IndexPCell != nil, ie.Tpc_IndexPUCCH_SCell != nil}); err != nil {
		return err
	}

	// 3. tpc-IndexPCell: integer
	{
		if ie.Tpc_IndexPCell != nil {
			if err := e.EncodeInteger(*ie.Tpc_IndexPCell, pUCCHTPCCommandConfigTpcIndexPCellConstraints); err != nil {
				return err
			}
		}
	}

	// 4. tpc-IndexPUCCH-SCell: integer
	{
		if ie.Tpc_IndexPUCCH_SCell != nil {
			if err := e.EncodeInteger(*ie.Tpc_IndexPUCCH_SCell, pUCCHTPCCommandConfigTpcIndexPUCCHSCellConstraints); err != nil {
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
					{Name: "tpc-IndexPUCCH-sSCell-r17", Optional: true},
					{Name: "tpc-IndexPUCCH-sScellSecondaryPUCCHgroup-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Tpc_IndexPUCCH_SSCell_r17 != nil, ie.Tpc_IndexPUCCH_SScellSecondaryPUCCHgroup_r17 != nil}); err != nil {
				return err
			}

			if ie.Tpc_IndexPUCCH_SSCell_r17 != nil {
				if err := ex.EncodeInteger(*ie.Tpc_IndexPUCCH_SSCell_r17, pUCCHTPCCommandConfigTpcIndexPUCCHSSCellR17Constraints); err != nil {
					return err
				}
			}

			if ie.Tpc_IndexPUCCH_SScellSecondaryPUCCHgroup_r17 != nil {
				if err := ex.EncodeInteger(*ie.Tpc_IndexPUCCH_SScellSecondaryPUCCHgroup_r17, pUCCHTPCCommandConfigTpcIndexPUCCHSScellSecondaryPUCCHgroupR17Constraints); err != nil {
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

func (ie *PUCCH_TPC_CommandConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUCCHTPCCommandConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. tpc-IndexPCell: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(pUCCHTPCCommandConfigTpcIndexPCellConstraints)
			if err != nil {
				return err
			}
			ie.Tpc_IndexPCell = &v
		}
	}

	// 4. tpc-IndexPUCCH-SCell: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(pUCCHTPCCommandConfigTpcIndexPUCCHSCellConstraints)
			if err != nil {
				return err
			}
			ie.Tpc_IndexPUCCH_SCell = &v
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
				{Name: "tpc-IndexPUCCH-sSCell-r17", Optional: true},
				{Name: "tpc-IndexPUCCH-sScellSecondaryPUCCHgroup-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(pUCCHTPCCommandConfigTpcIndexPUCCHSSCellR17Constraints)
			if err != nil {
				return err
			}
			ie.Tpc_IndexPUCCH_SSCell_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeInteger(pUCCHTPCCommandConfigTpcIndexPUCCHSScellSecondaryPUCCHgroupR17Constraints)
			if err != nil {
				return err
			}
			ie.Tpc_IndexPUCCH_SScellSecondaryPUCCHgroup_r17 = &v
		}
	}

	return nil
}
