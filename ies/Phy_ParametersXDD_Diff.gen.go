// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Phy-ParametersXDD-Diff (line 23079).

var phyParametersXDDDiffConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dynamicSFI", Optional: true},
		{Name: "twoPUCCH-F0-2-ConsecSymbols", Optional: true},
		{Name: "twoDifferentTPC-Loop-PUSCH", Optional: true},
		{Name: "twoDifferentTPC-Loop-PUCCH", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	Phy_ParametersXDD_Diff_DynamicSFI_Supported = 0
)

var phyParametersXDDDiffDynamicSFIConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersXDD_Diff_DynamicSFI_Supported},
}

const (
	Phy_ParametersXDD_Diff_TwoPUCCH_F0_2_ConsecSymbols_Supported = 0
)

var phyParametersXDDDiffTwoPUCCHF02ConsecSymbolsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersXDD_Diff_TwoPUCCH_F0_2_ConsecSymbols_Supported},
}

const (
	Phy_ParametersXDD_Diff_TwoDifferentTPC_Loop_PUSCH_Supported = 0
)

var phyParametersXDDDiffTwoDifferentTPCLoopPUSCHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersXDD_Diff_TwoDifferentTPC_Loop_PUSCH_Supported},
}

const (
	Phy_ParametersXDD_Diff_TwoDifferentTPC_Loop_PUCCH_Supported = 0
)

var phyParametersXDDDiffTwoDifferentTPCLoopPUCCHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersXDD_Diff_TwoDifferentTPC_Loop_PUCCH_Supported},
}

const (
	Phy_ParametersXDD_Diff_Ext_Dl_SchedulingOffset_PDSCH_TypeA_Supported = 0
)

var phyParametersXDDDiffExtDlSchedulingOffsetPDSCHTypeAConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersXDD_Diff_Ext_Dl_SchedulingOffset_PDSCH_TypeA_Supported},
}

const (
	Phy_ParametersXDD_Diff_Ext_Dl_SchedulingOffset_PDSCH_TypeB_Supported = 0
)

var phyParametersXDDDiffExtDlSchedulingOffsetPDSCHTypeBConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersXDD_Diff_Ext_Dl_SchedulingOffset_PDSCH_TypeB_Supported},
}

const (
	Phy_ParametersXDD_Diff_Ext_Ul_SchedulingOffset_Supported = 0
)

var phyParametersXDDDiffExtUlSchedulingOffsetConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersXDD_Diff_Ext_Ul_SchedulingOffset_Supported},
}

type Phy_ParametersXDD_Diff struct {
	DynamicSFI                      *int64
	TwoPUCCH_F0_2_ConsecSymbols     *int64
	TwoDifferentTPC_Loop_PUSCH      *int64
	TwoDifferentTPC_Loop_PUCCH      *int64
	Dl_SchedulingOffset_PDSCH_TypeA *int64
	Dl_SchedulingOffset_PDSCH_TypeB *int64
	Ul_SchedulingOffset             *int64
}

func (ie *Phy_ParametersXDD_Diff) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(phyParametersXDDDiffConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Dl_SchedulingOffset_PDSCH_TypeA != nil || ie.Dl_SchedulingOffset_PDSCH_TypeB != nil || ie.Ul_SchedulingOffset != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DynamicSFI != nil, ie.TwoPUCCH_F0_2_ConsecSymbols != nil, ie.TwoDifferentTPC_Loop_PUSCH != nil, ie.TwoDifferentTPC_Loop_PUCCH != nil}); err != nil {
		return err
	}

	// 3. dynamicSFI: enumerated
	{
		if ie.DynamicSFI != nil {
			if err := e.EncodeEnumerated(*ie.DynamicSFI, phyParametersXDDDiffDynamicSFIConstraints); err != nil {
				return err
			}
		}
	}

	// 4. twoPUCCH-F0-2-ConsecSymbols: enumerated
	{
		if ie.TwoPUCCH_F0_2_ConsecSymbols != nil {
			if err := e.EncodeEnumerated(*ie.TwoPUCCH_F0_2_ConsecSymbols, phyParametersXDDDiffTwoPUCCHF02ConsecSymbolsConstraints); err != nil {
				return err
			}
		}
	}

	// 5. twoDifferentTPC-Loop-PUSCH: enumerated
	{
		if ie.TwoDifferentTPC_Loop_PUSCH != nil {
			if err := e.EncodeEnumerated(*ie.TwoDifferentTPC_Loop_PUSCH, phyParametersXDDDiffTwoDifferentTPCLoopPUSCHConstraints); err != nil {
				return err
			}
		}
	}

	// 6. twoDifferentTPC-Loop-PUCCH: enumerated
	{
		if ie.TwoDifferentTPC_Loop_PUCCH != nil {
			if err := e.EncodeEnumerated(*ie.TwoDifferentTPC_Loop_PUCCH, phyParametersXDDDiffTwoDifferentTPCLoopPUCCHConstraints); err != nil {
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
					{Name: "dl-SchedulingOffset-PDSCH-TypeA", Optional: true},
					{Name: "dl-SchedulingOffset-PDSCH-TypeB", Optional: true},
					{Name: "ul-SchedulingOffset", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Dl_SchedulingOffset_PDSCH_TypeA != nil, ie.Dl_SchedulingOffset_PDSCH_TypeB != nil, ie.Ul_SchedulingOffset != nil}); err != nil {
				return err
			}

			if ie.Dl_SchedulingOffset_PDSCH_TypeA != nil {
				if err := ex.EncodeEnumerated(*ie.Dl_SchedulingOffset_PDSCH_TypeA, phyParametersXDDDiffExtDlSchedulingOffsetPDSCHTypeAConstraints); err != nil {
					return err
				}
			}

			if ie.Dl_SchedulingOffset_PDSCH_TypeB != nil {
				if err := ex.EncodeEnumerated(*ie.Dl_SchedulingOffset_PDSCH_TypeB, phyParametersXDDDiffExtDlSchedulingOffsetPDSCHTypeBConstraints); err != nil {
					return err
				}
			}

			if ie.Ul_SchedulingOffset != nil {
				if err := ex.EncodeEnumerated(*ie.Ul_SchedulingOffset, phyParametersXDDDiffExtUlSchedulingOffsetConstraints); err != nil {
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

func (ie *Phy_ParametersXDD_Diff) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(phyParametersXDDDiffConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. dynamicSFI: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(phyParametersXDDDiffDynamicSFIConstraints)
			if err != nil {
				return err
			}
			ie.DynamicSFI = &idx
		}
	}

	// 4. twoPUCCH-F0-2-ConsecSymbols: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(phyParametersXDDDiffTwoPUCCHF02ConsecSymbolsConstraints)
			if err != nil {
				return err
			}
			ie.TwoPUCCH_F0_2_ConsecSymbols = &idx
		}
	}

	// 5. twoDifferentTPC-Loop-PUSCH: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(phyParametersXDDDiffTwoDifferentTPCLoopPUSCHConstraints)
			if err != nil {
				return err
			}
			ie.TwoDifferentTPC_Loop_PUSCH = &idx
		}
	}

	// 6. twoDifferentTPC-Loop-PUCCH: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(phyParametersXDDDiffTwoDifferentTPCLoopPUCCHConstraints)
			if err != nil {
				return err
			}
			ie.TwoDifferentTPC_Loop_PUCCH = &idx
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
				{Name: "dl-SchedulingOffset-PDSCH-TypeA", Optional: true},
				{Name: "dl-SchedulingOffset-PDSCH-TypeB", Optional: true},
				{Name: "ul-SchedulingOffset", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(phyParametersXDDDiffExtDlSchedulingOffsetPDSCHTypeAConstraints)
			if err != nil {
				return err
			}
			ie.Dl_SchedulingOffset_PDSCH_TypeA = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(phyParametersXDDDiffExtDlSchedulingOffsetPDSCHTypeBConstraints)
			if err != nil {
				return err
			}
			ie.Dl_SchedulingOffset_PDSCH_TypeB = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(phyParametersXDDDiffExtUlSchedulingOffsetConstraints)
			if err != nil {
				return err
			}
			ie.Ul_SchedulingOffset = &v
		}
	}

	return nil
}
