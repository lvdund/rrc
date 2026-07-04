// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CodebookParametersetype2CJT-r18 (line 18936).

var codebookParametersetype2CJTR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "eType2CJT-r18"},
		{Name: "eType2CJT-FD-IO-r18", Optional: true},
		{Name: "eType2CJT-FD-FO-r18", Optional: true},
		{Name: "eType2CJT-R2-r18", Optional: true},
		{Name: "eType2CJT-PV-Beta-r18", Optional: true},
		{Name: "eType2CJT-2NN1N2-r18", Optional: true},
		{Name: "eType2CJT-Rank3Rank4-r18", Optional: true},
		{Name: "eType2CJT-L6-r18", Optional: true},
		{Name: "eType2CJT-NN-r18", Optional: true},
		{Name: "eType2CJT-NL-SD-r18", Optional: true},
		{Name: "eType2CJT-Unequal-r18", Optional: true},
	},
}

var codebookParametersetype2CJTR18EType2CJTFDIOR18Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	CodebookParametersetype2CJT_r18_EType2CJT_FD_FO_r18_Supported = 0
)

var codebookParametersetype2CJTR18EType2CJTFDFOR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersetype2CJT_r18_EType2CJT_FD_FO_r18_Supported},
}

var codebookParametersetype2CJTR18EType2CJTR2R18Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	CodebookParametersetype2CJT_r18_EType2CJT_PV_Beta_r18_Supported = 0
)

var codebookParametersetype2CJTR18EType2CJTPVBetaR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersetype2CJT_r18_EType2CJT_PV_Beta_r18_Supported},
}

const (
	CodebookParametersetype2CJT_r18_EType2CJT_2NN1N2_r18_N64  = 0
	CodebookParametersetype2CJT_r18_EType2CJT_2NN1N2_r18_N96  = 1
	CodebookParametersetype2CJT_r18_EType2CJT_2NN1N2_r18_N128 = 2
)

var codebookParametersetype2CJTR18EType2CJT2NN1N2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersetype2CJT_r18_EType2CJT_2NN1N2_r18_N64, CodebookParametersetype2CJT_r18_EType2CJT_2NN1N2_r18_N96, CodebookParametersetype2CJT_r18_EType2CJT_2NN1N2_r18_N128},
}

const (
	CodebookParametersetype2CJT_r18_EType2CJT_Rank3Rank4_r18_Supported = 0
)

var codebookParametersetype2CJTR18EType2CJTRank3Rank4R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersetype2CJT_r18_EType2CJT_Rank3Rank4_r18_Supported},
}

const (
	CodebookParametersetype2CJT_r18_EType2CJT_L6_r18_Supported = 0
)

var codebookParametersetype2CJTR18EType2CJTL6R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersetype2CJT_r18_EType2CJT_L6_r18_Supported},
}

const (
	CodebookParametersetype2CJT_r18_EType2CJT_NN_r18_Supported = 0
)

var codebookParametersetype2CJTR18EType2CJTNNR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersetype2CJT_r18_EType2CJT_NN_r18_Supported},
}

const (
	CodebookParametersetype2CJT_r18_EType2CJT_NL_SD_r18_N2 = 0
	CodebookParametersetype2CJT_r18_EType2CJT_NL_SD_r18_N4 = 1
)

var codebookParametersetype2CJTR18EType2CJTNLSDR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersetype2CJT_r18_EType2CJT_NL_SD_r18_N2, CodebookParametersetype2CJT_r18_EType2CJT_NL_SD_r18_N4},
}

const (
	CodebookParametersetype2CJT_r18_EType2CJT_Unequal_r18_Supported = 0
)

var codebookParametersetype2CJTR18EType2CJTUnequalR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersetype2CJT_r18_EType2CJT_Unequal_r18_Supported},
}

var codebookParametersetype2CJTR18EType2CJTR18SupportedCSIRSResourceListR18Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	CodebookParametersetype2CJT_r18_EType2CJT_r18_Scalingfactor_r18_N1     = 0
	CodebookParametersetype2CJT_r18_EType2CJT_r18_Scalingfactor_r18_N1dot5 = 1
	CodebookParametersetype2CJT_r18_EType2CJT_r18_Scalingfactor_r18_N2     = 2
)

var codebookParametersetype2CJTR18EType2CJTR18ScalingfactorR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersetype2CJT_r18_EType2CJT_r18_Scalingfactor_r18_N1, CodebookParametersetype2CJT_r18_EType2CJT_r18_Scalingfactor_r18_N1dot5, CodebookParametersetype2CJT_r18_EType2CJT_r18_Scalingfactor_r18_N2},
}

type CodebookParametersetype2CJT_r18 struct {
	EType2CJT_r18 struct {
		SupportedCSI_RS_ResourceList_r18     []int64
		Scalingfactor_r18                    int64
		MaxNumberNZP_CSI_RS_MultiTRP_CJT_r18 int64
	}
	EType2CJT_FD_IO_r18      []int64
	EType2CJT_FD_FO_r18      *int64
	EType2CJT_R2_r18         []int64
	EType2CJT_PV_Beta_r18    *int64
	EType2CJT_2NN1N2_r18     *int64
	EType2CJT_Rank3Rank4_r18 *int64
	EType2CJT_L6_r18         *int64
	EType2CJT_NN_r18         *int64
	EType2CJT_NL_SD_r18      *int64
	EType2CJT_Unequal_r18    *int64
}

func (ie *CodebookParametersetype2CJT_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(codebookParametersetype2CJTR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.EType2CJT_FD_IO_r18 != nil, ie.EType2CJT_FD_FO_r18 != nil, ie.EType2CJT_R2_r18 != nil, ie.EType2CJT_PV_Beta_r18 != nil, ie.EType2CJT_2NN1N2_r18 != nil, ie.EType2CJT_Rank3Rank4_r18 != nil, ie.EType2CJT_L6_r18 != nil, ie.EType2CJT_NN_r18 != nil, ie.EType2CJT_NL_SD_r18 != nil, ie.EType2CJT_Unequal_r18 != nil}); err != nil {
		return err
	}

	// 2. eType2CJT-r18: sequence
	{
		{
			c := &ie.EType2CJT_r18
			{
				seqOf := e.NewSequenceOfEncoder(codebookParametersetype2CJTR18EType2CJTR18SupportedCSIRSResourceListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceList_r18))); err != nil {
					return err
				}
				for i := range c.SupportedCSI_RS_ResourceList_r18 {
					if err := e.EncodeInteger(c.SupportedCSI_RS_ResourceList_r18[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeEnumerated(c.Scalingfactor_r18, codebookParametersetype2CJTR18EType2CJTR18ScalingfactorR18Constraints); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.MaxNumberNZP_CSI_RS_MultiTRP_CJT_r18, per.Constrained(2, 4)); err != nil {
				return err
			}
		}
	}

	// 3. eType2CJT-FD-IO-r18: sequence-of
	{
		if ie.EType2CJT_FD_IO_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookParametersetype2CJTR18EType2CJTFDIOR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.EType2CJT_FD_IO_r18))); err != nil {
				return err
			}
			for i := range ie.EType2CJT_FD_IO_r18 {
				if err := e.EncodeInteger(ie.EType2CJT_FD_IO_r18[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 4. eType2CJT-FD-FO-r18: enumerated
	{
		if ie.EType2CJT_FD_FO_r18 != nil {
			if err := e.EncodeEnumerated(*ie.EType2CJT_FD_FO_r18, codebookParametersetype2CJTR18EType2CJTFDFOR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. eType2CJT-R2-r18: sequence-of
	{
		if ie.EType2CJT_R2_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(codebookParametersetype2CJTR18EType2CJTR2R18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.EType2CJT_R2_r18))); err != nil {
				return err
			}
			for i := range ie.EType2CJT_R2_r18 {
				if err := e.EncodeInteger(ie.EType2CJT_R2_r18[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
					return err
				}
			}
		}
	}

	// 6. eType2CJT-PV-Beta-r18: enumerated
	{
		if ie.EType2CJT_PV_Beta_r18 != nil {
			if err := e.EncodeEnumerated(*ie.EType2CJT_PV_Beta_r18, codebookParametersetype2CJTR18EType2CJTPVBetaR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. eType2CJT-2NN1N2-r18: enumerated
	{
		if ie.EType2CJT_2NN1N2_r18 != nil {
			if err := e.EncodeEnumerated(*ie.EType2CJT_2NN1N2_r18, codebookParametersetype2CJTR18EType2CJT2NN1N2R18Constraints); err != nil {
				return err
			}
		}
	}

	// 8. eType2CJT-Rank3Rank4-r18: enumerated
	{
		if ie.EType2CJT_Rank3Rank4_r18 != nil {
			if err := e.EncodeEnumerated(*ie.EType2CJT_Rank3Rank4_r18, codebookParametersetype2CJTR18EType2CJTRank3Rank4R18Constraints); err != nil {
				return err
			}
		}
	}

	// 9. eType2CJT-L6-r18: enumerated
	{
		if ie.EType2CJT_L6_r18 != nil {
			if err := e.EncodeEnumerated(*ie.EType2CJT_L6_r18, codebookParametersetype2CJTR18EType2CJTL6R18Constraints); err != nil {
				return err
			}
		}
	}

	// 10. eType2CJT-NN-r18: enumerated
	{
		if ie.EType2CJT_NN_r18 != nil {
			if err := e.EncodeEnumerated(*ie.EType2CJT_NN_r18, codebookParametersetype2CJTR18EType2CJTNNR18Constraints); err != nil {
				return err
			}
		}
	}

	// 11. eType2CJT-NL-SD-r18: enumerated
	{
		if ie.EType2CJT_NL_SD_r18 != nil {
			if err := e.EncodeEnumerated(*ie.EType2CJT_NL_SD_r18, codebookParametersetype2CJTR18EType2CJTNLSDR18Constraints); err != nil {
				return err
			}
		}
	}

	// 12. eType2CJT-Unequal-r18: enumerated
	{
		if ie.EType2CJT_Unequal_r18 != nil {
			if err := e.EncodeEnumerated(*ie.EType2CJT_Unequal_r18, codebookParametersetype2CJTR18EType2CJTUnequalR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CodebookParametersetype2CJT_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(codebookParametersetype2CJTR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. eType2CJT-r18: sequence
	{
		{
			c := &ie.EType2CJT_r18
			{
				seqOf := d.NewSequenceOfDecoder(codebookParametersetype2CJTR18EType2CJTR18SupportedCSIRSResourceListR18Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceList_r18 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceList_r18[i] = v
				}
			}
			{
				v, err := d.DecodeEnumerated(codebookParametersetype2CJTR18EType2CJTR18ScalingfactorR18Constraints)
				if err != nil {
					return err
				}
				c.Scalingfactor_r18 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(2, 4))
				if err != nil {
					return err
				}
				c.MaxNumberNZP_CSI_RS_MultiTRP_CJT_r18 = v
			}
		}
	}

	// 3. eType2CJT-FD-IO-r18: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(codebookParametersetype2CJTR18EType2CJTFDIOR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.EType2CJT_FD_IO_r18 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.EType2CJT_FD_IO_r18[i] = v
			}
		}
	}

	// 4. eType2CJT-FD-FO-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(codebookParametersetype2CJTR18EType2CJTFDFOR18Constraints)
			if err != nil {
				return err
			}
			ie.EType2CJT_FD_FO_r18 = &idx
		}
	}

	// 5. eType2CJT-R2-r18: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(codebookParametersetype2CJTR18EType2CJTR2R18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.EType2CJT_R2_r18 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
				if err != nil {
					return err
				}
				ie.EType2CJT_R2_r18[i] = v
			}
		}
	}

	// 6. eType2CJT-PV-Beta-r18: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(codebookParametersetype2CJTR18EType2CJTPVBetaR18Constraints)
			if err != nil {
				return err
			}
			ie.EType2CJT_PV_Beta_r18 = &idx
		}
	}

	// 7. eType2CJT-2NN1N2-r18: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(codebookParametersetype2CJTR18EType2CJT2NN1N2R18Constraints)
			if err != nil {
				return err
			}
			ie.EType2CJT_2NN1N2_r18 = &idx
		}
	}

	// 8. eType2CJT-Rank3Rank4-r18: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(codebookParametersetype2CJTR18EType2CJTRank3Rank4R18Constraints)
			if err != nil {
				return err
			}
			ie.EType2CJT_Rank3Rank4_r18 = &idx
		}
	}

	// 9. eType2CJT-L6-r18: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(codebookParametersetype2CJTR18EType2CJTL6R18Constraints)
			if err != nil {
				return err
			}
			ie.EType2CJT_L6_r18 = &idx
		}
	}

	// 10. eType2CJT-NN-r18: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(codebookParametersetype2CJTR18EType2CJTNNR18Constraints)
			if err != nil {
				return err
			}
			ie.EType2CJT_NN_r18 = &idx
		}
	}

	// 11. eType2CJT-NL-SD-r18: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(codebookParametersetype2CJTR18EType2CJTNLSDR18Constraints)
			if err != nil {
				return err
			}
			ie.EType2CJT_NL_SD_r18 = &idx
		}
	}

	// 12. eType2CJT-Unequal-r18: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(codebookParametersetype2CJTR18EType2CJTUnequalR18Constraints)
			if err != nil {
				return err
			}
			ie.EType2CJT_Unequal_r18 = &idx
		}
	}

	return nil
}
