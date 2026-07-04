// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CFRA-TwoStep-r16 (line 12958).

var cFRATwoStepR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "occasionsTwoStepRA-r16", Optional: true},
		{Name: "msgA-CFRA-PUSCH-r16"},
		{Name: "msgA-TransMax-r16", Optional: true},
		{Name: "resourcesTwoStep-r16"},
	},
}

const (
	CFRA_TwoStep_r16_MsgA_TransMax_r16_N1   = 0
	CFRA_TwoStep_r16_MsgA_TransMax_r16_N2   = 1
	CFRA_TwoStep_r16_MsgA_TransMax_r16_N4   = 2
	CFRA_TwoStep_r16_MsgA_TransMax_r16_N6   = 3
	CFRA_TwoStep_r16_MsgA_TransMax_r16_N8   = 4
	CFRA_TwoStep_r16_MsgA_TransMax_r16_N10  = 5
	CFRA_TwoStep_r16_MsgA_TransMax_r16_N20  = 6
	CFRA_TwoStep_r16_MsgA_TransMax_r16_N50  = 7
	CFRA_TwoStep_r16_MsgA_TransMax_r16_N100 = 8
	CFRA_TwoStep_r16_MsgA_TransMax_r16_N200 = 9
)

var cFRATwoStepR16MsgATransMaxR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CFRA_TwoStep_r16_MsgA_TransMax_r16_N1, CFRA_TwoStep_r16_MsgA_TransMax_r16_N2, CFRA_TwoStep_r16_MsgA_TransMax_r16_N4, CFRA_TwoStep_r16_MsgA_TransMax_r16_N6, CFRA_TwoStep_r16_MsgA_TransMax_r16_N8, CFRA_TwoStep_r16_MsgA_TransMax_r16_N10, CFRA_TwoStep_r16_MsgA_TransMax_r16_N20, CFRA_TwoStep_r16_MsgA_TransMax_r16_N50, CFRA_TwoStep_r16_MsgA_TransMax_r16_N100, CFRA_TwoStep_r16_MsgA_TransMax_r16_N200},
}

const (
	CFRA_TwoStep_r16_OccasionsTwoStepRA_r16_Ssb_PerRACH_OccasionTwoStepRA_r16_OneEighth = 0
	CFRA_TwoStep_r16_OccasionsTwoStepRA_r16_Ssb_PerRACH_OccasionTwoStepRA_r16_OneFourth = 1
	CFRA_TwoStep_r16_OccasionsTwoStepRA_r16_Ssb_PerRACH_OccasionTwoStepRA_r16_OneHalf   = 2
	CFRA_TwoStep_r16_OccasionsTwoStepRA_r16_Ssb_PerRACH_OccasionTwoStepRA_r16_One       = 3
	CFRA_TwoStep_r16_OccasionsTwoStepRA_r16_Ssb_PerRACH_OccasionTwoStepRA_r16_Two       = 4
	CFRA_TwoStep_r16_OccasionsTwoStepRA_r16_Ssb_PerRACH_OccasionTwoStepRA_r16_Four      = 5
	CFRA_TwoStep_r16_OccasionsTwoStepRA_r16_Ssb_PerRACH_OccasionTwoStepRA_r16_Eight     = 6
	CFRA_TwoStep_r16_OccasionsTwoStepRA_r16_Ssb_PerRACH_OccasionTwoStepRA_r16_Sixteen   = 7
)

var cFRATwoStepR16OccasionsTwoStepRAR16SsbPerRACHOccasionTwoStepRAR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CFRA_TwoStep_r16_OccasionsTwoStepRA_r16_Ssb_PerRACH_OccasionTwoStepRA_r16_OneEighth, CFRA_TwoStep_r16_OccasionsTwoStepRA_r16_Ssb_PerRACH_OccasionTwoStepRA_r16_OneFourth, CFRA_TwoStep_r16_OccasionsTwoStepRA_r16_Ssb_PerRACH_OccasionTwoStepRA_r16_OneHalf, CFRA_TwoStep_r16_OccasionsTwoStepRA_r16_Ssb_PerRACH_OccasionTwoStepRA_r16_One, CFRA_TwoStep_r16_OccasionsTwoStepRA_r16_Ssb_PerRACH_OccasionTwoStepRA_r16_Two, CFRA_TwoStep_r16_OccasionsTwoStepRA_r16_Ssb_PerRACH_OccasionTwoStepRA_r16_Four, CFRA_TwoStep_r16_OccasionsTwoStepRA_r16_Ssb_PerRACH_OccasionTwoStepRA_r16_Eight, CFRA_TwoStep_r16_OccasionsTwoStepRA_r16_Ssb_PerRACH_OccasionTwoStepRA_r16_Sixteen},
}

var cFRATwoStepR16ResourcesTwoStepR16SsbResourceListConstraints = per.SizeRange(1, common.MaxRA_SSB_Resources)

type CFRA_TwoStep_r16 struct {
	OccasionsTwoStepRA_r16 *struct {
		Rach_ConfigGenericTwoStepRA_r16   RACH_ConfigGenericTwoStepRA_r16
		Ssb_PerRACH_OccasionTwoStepRA_r16 int64
	}
	MsgA_CFRA_PUSCH_r16  MsgA_PUSCH_Resource_r16
	MsgA_TransMax_r16    *int64
	ResourcesTwoStep_r16 struct {
		Ssb_ResourceList         []CFRA_SSB_Resource
		Ra_Ssb_OccasionMaskIndex int64
	}
}

func (ie *CFRA_TwoStep_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cFRATwoStepR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.OccasionsTwoStepRA_r16 != nil, ie.MsgA_TransMax_r16 != nil}); err != nil {
		return err
	}

	// 3. occasionsTwoStepRA-r16: sequence
	{
		if ie.OccasionsTwoStepRA_r16 != nil {
			c := ie.OccasionsTwoStepRA_r16
			if err := c.Rach_ConfigGenericTwoStepRA_r16.Encode(e); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.Ssb_PerRACH_OccasionTwoStepRA_r16, cFRATwoStepR16OccasionsTwoStepRAR16SsbPerRACHOccasionTwoStepRAR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. msgA-CFRA-PUSCH-r16: ref
	{
		if err := ie.MsgA_CFRA_PUSCH_r16.Encode(e); err != nil {
			return err
		}
	}

	// 5. msgA-TransMax-r16: enumerated
	{
		if ie.MsgA_TransMax_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MsgA_TransMax_r16, cFRATwoStepR16MsgATransMaxR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. resourcesTwoStep-r16: sequence
	{
		{
			c := &ie.ResourcesTwoStep_r16
			{
				seqOf := e.NewSequenceOfEncoder(cFRATwoStepR16ResourcesTwoStepR16SsbResourceListConstraints)
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
		}
	}

	return nil
}

func (ie *CFRA_TwoStep_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cFRATwoStepR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. occasionsTwoStepRA-r16: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.OccasionsTwoStepRA_r16 = &struct {
				Rach_ConfigGenericTwoStepRA_r16   RACH_ConfigGenericTwoStepRA_r16
				Ssb_PerRACH_OccasionTwoStepRA_r16 int64
			}{}
			c := ie.OccasionsTwoStepRA_r16
			{
				if err := c.Rach_ConfigGenericTwoStepRA_r16.Decode(d); err != nil {
					return err
				}
			}
			{
				v, err := d.DecodeEnumerated(cFRATwoStepR16OccasionsTwoStepRAR16SsbPerRACHOccasionTwoStepRAR16Constraints)
				if err != nil {
					return err
				}
				c.Ssb_PerRACH_OccasionTwoStepRA_r16 = v
			}
		}
	}

	// 4. msgA-CFRA-PUSCH-r16: ref
	{
		if err := ie.MsgA_CFRA_PUSCH_r16.Decode(d); err != nil {
			return err
		}
	}

	// 5. msgA-TransMax-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(cFRATwoStepR16MsgATransMaxR16Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_TransMax_r16 = &idx
		}
	}

	// 6. resourcesTwoStep-r16: sequence
	{
		{
			c := &ie.ResourcesTwoStep_r16
			{
				seqOf := d.NewSequenceOfDecoder(cFRATwoStepR16ResourcesTwoStepR16SsbResourceListConstraints)
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
		}
	}

	return nil
}
