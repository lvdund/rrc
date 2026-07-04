// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: Phy-ParametersMRDC (line 23258).

var phyParametersMRDCConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "naics-Capability-List", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var phyParametersMRDCNaicsCapabilityListConstraints = per.SizeRange(1, common.MaxNrofNAICS_Entries)

const (
	Phy_ParametersMRDC_Ext_Tdd_PCellUL_TX_AllUL_Subframe_r16_Supported = 0
)

var phyParametersMRDCExtTddPCellULTXAllULSubframeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersMRDC_Ext_Tdd_PCellUL_TX_AllUL_Subframe_r16_Supported},
}

const (
	Phy_ParametersMRDC_Ext_Fdd_PCellUL_TX_AllUL_Subframe_r16_Supported = 0
)

var phyParametersMRDCExtFddPCellULTXAllULSubframeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersMRDC_Ext_Fdd_PCellUL_TX_AllUL_Subframe_r16_Supported},
}

type Phy_ParametersMRDC struct {
	Naics_Capability_List             []NAICS_Capability_Entry
	SpCellPlacement                   *CarrierAggregationVariant
	Tdd_PCellUL_TX_AllUL_Subframe_r16 *int64
	Fdd_PCellUL_TX_AllUL_Subframe_r16 *int64
}

func (ie *Phy_ParametersMRDC) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(phyParametersMRDCConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.SpCellPlacement != nil
	hasExtGroup1 := ie.Tdd_PCellUL_TX_AllUL_Subframe_r16 != nil || ie.Fdd_PCellUL_TX_AllUL_Subframe_r16 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Naics_Capability_List != nil}); err != nil {
		return err
	}

	// 3. naics-Capability-List: sequence-of
	{
		if ie.Naics_Capability_List != nil {
			seqOf := e.NewSequenceOfEncoder(phyParametersMRDCNaicsCapabilityListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Naics_Capability_List))); err != nil {
				return err
			}
			for i := range ie.Naics_Capability_List {
				if err := ie.Naics_Capability_List[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "spCellPlacement", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SpCellPlacement != nil}); err != nil {
				return err
			}

			if ie.SpCellPlacement != nil {
				if err := ie.SpCellPlacement.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "tdd-PCellUL-TX-AllUL-Subframe-r16", Optional: true},
					{Name: "fdd-PCellUL-TX-AllUL-Subframe-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Tdd_PCellUL_TX_AllUL_Subframe_r16 != nil, ie.Fdd_PCellUL_TX_AllUL_Subframe_r16 != nil}); err != nil {
				return err
			}

			if ie.Tdd_PCellUL_TX_AllUL_Subframe_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Tdd_PCellUL_TX_AllUL_Subframe_r16, phyParametersMRDCExtTddPCellULTXAllULSubframeR16Constraints); err != nil {
					return err
				}
			}

			if ie.Fdd_PCellUL_TX_AllUL_Subframe_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Fdd_PCellUL_TX_AllUL_Subframe_r16, phyParametersMRDCExtFddPCellULTXAllULSubframeR16Constraints); err != nil {
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

func (ie *Phy_ParametersMRDC) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(phyParametersMRDCConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. naics-Capability-List: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(phyParametersMRDCNaicsCapabilityListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Naics_Capability_List = make([]NAICS_Capability_Entry, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Naics_Capability_List[i].Decode(d); err != nil {
					return err
				}
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
				{Name: "spCellPlacement", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SpCellPlacement = new(CarrierAggregationVariant)
			if err := ie.SpCellPlacement.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "tdd-PCellUL-TX-AllUL-Subframe-r16", Optional: true},
				{Name: "fdd-PCellUL-TX-AllUL-Subframe-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(phyParametersMRDCExtTddPCellULTXAllULSubframeR16Constraints)
			if err != nil {
				return err
			}
			ie.Tdd_PCellUL_TX_AllUL_Subframe_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(phyParametersMRDCExtFddPCellULTXAllULSubframeR16Constraints)
			if err != nil {
				return err
			}
			ie.Fdd_PCellUL_TX_AllUL_Subframe_r16 = &v
		}
	}

	return nil
}
