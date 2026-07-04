// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CFRA (line 12930).

var cFRAConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "occasions", Optional: true},
		{Name: "resources"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

var cFRAResourcesConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ssb"},
		{Name: "csirs"},
	},
}

const (
	CFRA_Resources_Ssb   = 0
	CFRA_Resources_Csirs = 1
)

var cFRATotalNumberOfRAPreamblesConstraints = per.Constrained(1, 63)

const (
	CFRA_Ext_Msg1_RepetitionNum_r18_N2     = 0
	CFRA_Ext_Msg1_RepetitionNum_r18_N4     = 1
	CFRA_Ext_Msg1_RepetitionNum_r18_N8     = 2
	CFRA_Ext_Msg1_RepetitionNum_r18_Spare1 = 3
)

var cFRAExtMsg1RepetitionNumR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CFRA_Ext_Msg1_RepetitionNum_r18_N2, CFRA_Ext_Msg1_RepetitionNum_r18_N4, CFRA_Ext_Msg1_RepetitionNum_r18_N8, CFRA_Ext_Msg1_RepetitionNum_r18_Spare1},
}

const (
	CFRA_Ext_Ra_OccasionType_r19_Sbfd = 0
)

var cFRAExtRaOccasionTypeR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CFRA_Ext_Ra_OccasionType_r19_Sbfd},
}

var cFRAOccasionsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rach-ConfigGeneric"},
		{Name: "ssb-perRACH-Occasion", Optional: true},
	},
}

const (
	CFRA_Occasions_Ssb_PerRACH_Occasion_OneEighth = 0
	CFRA_Occasions_Ssb_PerRACH_Occasion_OneFourth = 1
	CFRA_Occasions_Ssb_PerRACH_Occasion_OneHalf   = 2
	CFRA_Occasions_Ssb_PerRACH_Occasion_One       = 3
	CFRA_Occasions_Ssb_PerRACH_Occasion_Two       = 4
	CFRA_Occasions_Ssb_PerRACH_Occasion_Four      = 5
	CFRA_Occasions_Ssb_PerRACH_Occasion_Eight     = 6
	CFRA_Occasions_Ssb_PerRACH_Occasion_Sixteen   = 7
)

var cFRAOccasionsSsbPerRACHOccasionConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CFRA_Occasions_Ssb_PerRACH_Occasion_OneEighth, CFRA_Occasions_Ssb_PerRACH_Occasion_OneFourth, CFRA_Occasions_Ssb_PerRACH_Occasion_OneHalf, CFRA_Occasions_Ssb_PerRACH_Occasion_One, CFRA_Occasions_Ssb_PerRACH_Occasion_Two, CFRA_Occasions_Ssb_PerRACH_Occasion_Four, CFRA_Occasions_Ssb_PerRACH_Occasion_Eight, CFRA_Occasions_Ssb_PerRACH_Occasion_Sixteen},
}

var cFRAResourcesSsbSsbResourceListConstraints = per.SizeRange(1, common.MaxRA_SSB_Resources)

var cFRAResourcesCsirsCsirsResourceListConstraints = per.SizeRange(1, common.MaxRA_CSIRS_Resources)

type CFRA struct {
	Occasions *struct {
		Rach_ConfigGeneric   RACH_ConfigGeneric
		Ssb_PerRACH_Occasion *int64
	}
	Resources struct {
		Choice int
		Ssb    *struct {
			Ssb_ResourceList         []CFRA_SSB_Resource
			Ra_Ssb_OccasionMaskIndex int64
		}
		Csirs *struct {
			Csirs_ResourceList   []CFRA_CSIRS_Resource
			Rsrp_ThresholdCSI_RS RSRP_Range
		}
	}
	TotalNumberOfRA_Preambles *int64
	Msg1_RepetitionNum_r18    *int64
	Ra_OccasionType_r19       *int64
}

func (ie *CFRA) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cFRAConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.TotalNumberOfRA_Preambles != nil
	hasExtGroup1 := ie.Msg1_RepetitionNum_r18 != nil
	hasExtGroup2 := ie.Ra_OccasionType_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Occasions != nil}); err != nil {
		return err
	}

	// 3. occasions: sequence
	{
		if ie.Occasions != nil {
			c := ie.Occasions
			cFRAOccasionsSeq := e.NewSequenceEncoder(cFRAOccasionsConstraints)
			if err := cFRAOccasionsSeq.EncodePreamble([]bool{c.Ssb_PerRACH_Occasion != nil}); err != nil {
				return err
			}
			if err := c.Rach_ConfigGeneric.Encode(e); err != nil {
				return err
			}
			if c.Ssb_PerRACH_Occasion != nil {
				if err := e.EncodeEnumerated((*c.Ssb_PerRACH_Occasion), cFRAOccasionsSsbPerRACHOccasionConstraints); err != nil {
					return err
				}
			}
		}
	}

	// 4. resources: choice
	{
		choiceEnc := e.NewChoiceEncoder(cFRAResourcesConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Resources.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Resources.Choice {
		case CFRA_Resources_Ssb:
			c := (*ie.Resources.Ssb)
			{
				seqOf := e.NewSequenceOfEncoder(cFRAResourcesSsbSsbResourceListConstraints)
				if err := seqOf.EncodeLength(int64(len(c.Ssb_ResourceList))); err != nil {
					return err
				}
				for i := range c.Ssb_ResourceList {
					if err := c.Ssb_ResourceList[i].Encode(e); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeInteger(c.Ra_Ssb_OccasionMaskIndex, per.Constrained(0, 15)); err != nil {
				return err
			}
		case CFRA_Resources_Csirs:
			c := (*ie.Resources.Csirs)
			{
				seqOf := e.NewSequenceOfEncoder(cFRAResourcesCsirsCsirsResourceListConstraints)
				if err := seqOf.EncodeLength(int64(len(c.Csirs_ResourceList))); err != nil {
					return err
				}
				for i := range c.Csirs_ResourceList {
					if err := c.Csirs_ResourceList[i].Encode(e); err != nil {
						return err
					}
				}
			}
			if err := c.Rsrp_ThresholdCSI_RS.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Resources.Choice), Constraint: "index out of range"}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "totalNumberOfRA-Preambles", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.TotalNumberOfRA_Preambles != nil}); err != nil {
				return err
			}

			if ie.TotalNumberOfRA_Preambles != nil {
				if err := ex.EncodeInteger(*ie.TotalNumberOfRA_Preambles, cFRATotalNumberOfRAPreamblesConstraints); err != nil {
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
					{Name: "msg1-RepetitionNum-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Msg1_RepetitionNum_r18 != nil}); err != nil {
				return err
			}

			if ie.Msg1_RepetitionNum_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Msg1_RepetitionNum_r18, cFRAExtMsg1RepetitionNumR18Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 2:
		if hasExtGroup2 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "ra-OccasionType-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ra_OccasionType_r19 != nil}); err != nil {
				return err
			}

			if ie.Ra_OccasionType_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Ra_OccasionType_r19, cFRAExtRaOccasionTypeR19Constraints); err != nil {
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

func (ie *CFRA) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cFRAConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. occasions: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.Occasions = &struct {
				Rach_ConfigGeneric   RACH_ConfigGeneric
				Ssb_PerRACH_Occasion *int64
			}{}
			c := ie.Occasions
			cFRAOccasionsSeq := d.NewSequenceDecoder(cFRAOccasionsConstraints)
			if err := cFRAOccasionsSeq.DecodePreamble(); err != nil {
				return err
			}
			{
				if err := c.Rach_ConfigGeneric.Decode(d); err != nil {
					return err
				}
			}
			if cFRAOccasionsSeq.IsComponentPresent(1) {
				c.Ssb_PerRACH_Occasion = new(int64)
				v, err := d.DecodeEnumerated(cFRAOccasionsSsbPerRACHOccasionConstraints)
				if err != nil {
					return err
				}
				(*c.Ssb_PerRACH_Occasion) = v
			}
		}
	}

	// 4. resources: choice
	{
		choiceDec := d.NewChoiceDecoder(cFRAResourcesConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Resources.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CFRA_Resources_Ssb:
			ie.Resources.Ssb = &struct {
				Ssb_ResourceList         []CFRA_SSB_Resource
				Ra_Ssb_OccasionMaskIndex int64
			}{}
			c := (*ie.Resources.Ssb)
			{
				seqOf := d.NewSequenceOfDecoder(cFRAResourcesSsbSsbResourceListConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.Ssb_ResourceList = make([]CFRA_SSB_Resource, n)
				for i := int64(0); i < n; i++ {
					if err := c.Ssb_ResourceList[i].Decode(d); err != nil {
						return err
					}
				}
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, 15))
				if err != nil {
					return err
				}
				c.Ra_Ssb_OccasionMaskIndex = v
			}
		case CFRA_Resources_Csirs:
			ie.Resources.Csirs = &struct {
				Csirs_ResourceList   []CFRA_CSIRS_Resource
				Rsrp_ThresholdCSI_RS RSRP_Range
			}{}
			c := (*ie.Resources.Csirs)
			{
				seqOf := d.NewSequenceOfDecoder(cFRAResourcesCsirsCsirsResourceListConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.Csirs_ResourceList = make([]CFRA_CSIRS_Resource, n)
				for i := int64(0); i < n; i++ {
					if err := c.Csirs_ResourceList[i].Decode(d); err != nil {
						return err
					}
				}
			}
			{
				if err := c.Rsrp_ThresholdCSI_RS.Decode(d); err != nil {
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
				{Name: "totalNumberOfRA-Preambles", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(cFRATotalNumberOfRAPreamblesConstraints)
			if err != nil {
				return err
			}
			ie.TotalNumberOfRA_Preambles = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "msg1-RepetitionNum-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(cFRAExtMsg1RepetitionNumR18Constraints)
			if err != nil {
				return err
			}
			ie.Msg1_RepetitionNum_r18 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "ra-OccasionType-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(cFRAExtRaOccasionTypeR19Constraints)
			if err != nil {
				return err
			}
			ie.Ra_OccasionType_r19 = &v
		}
	}

	return nil
}
