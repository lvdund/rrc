// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetUplinkPerCC-v1850 (line 20678).

var featureSetUplinkPerCCV1850Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "twoPUSCH-CB-MultiDCI-STx2P-AdditionalTime-r18", Optional: true},
		{Name: "twoPUSCH-NonCB-MultiDCI-STx2P-AdditionalTime-r18", Optional: true},
	},
}

var featureSetUplinkPerCC_v1850TwoPUSCHCBMultiDCISTx2PAdditionalTimeR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "scs-60kHz-r18"},
		{Name: "scs-120kHz-r18"},
		{Name: "scs-480kHz-r18"},
		{Name: "scs-960kHz-r18"},
	},
}

const (
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_60kHz_r18  = 0
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_120kHz_r18 = 1
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_480kHz_r18 = 2
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_960kHz_r18 = 3
)

var featureSetUplinkPerCC_v1850TwoPUSCHNonCBMultiDCISTx2PAdditionalTimeR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "scs-60kHz-r18"},
		{Name: "scs-120kHz-r18"},
		{Name: "scs-480kHz-r18"},
		{Name: "scs-960kHz-r18"},
	},
}

const (
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_60kHz_r18  = 0
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_120kHz_r18 = 1
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_480kHz_r18 = 2
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_960kHz_r18 = 3
)

const (
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_60kHz_r18_Sym1  = 0
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_60kHz_r18_Sym4  = 1
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_60kHz_r18_Sym8  = 2
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_60kHz_r18_Sym16 = 3
)

var featureSetUplinkPerCCV1850TwoPUSCHCBMultiDCISTx2PAdditionalTimeR18Scs60kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_60kHz_r18_Sym1, FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_60kHz_r18_Sym4, FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_60kHz_r18_Sym8, FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_60kHz_r18_Sym16},
}

const (
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_120kHz_r18_Sym4  = 0
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_120kHz_r18_Sym8  = 1
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_120kHz_r18_Sym16 = 2
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_120kHz_r18_Sym32 = 3
)

var featureSetUplinkPerCCV1850TwoPUSCHCBMultiDCISTx2PAdditionalTimeR18Scs120kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_120kHz_r18_Sym4, FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_120kHz_r18_Sym8, FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_120kHz_r18_Sym16, FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_120kHz_r18_Sym32},
}

const (
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_480kHz_r18_Sym16  = 0
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_480kHz_r18_Sym32  = 1
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_480kHz_r18_Sym64  = 2
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_480kHz_r18_Sym128 = 3
)

var featureSetUplinkPerCCV1850TwoPUSCHCBMultiDCISTx2PAdditionalTimeR18Scs480kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_480kHz_r18_Sym16, FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_480kHz_r18_Sym32, FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_480kHz_r18_Sym64, FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_480kHz_r18_Sym128},
}

const (
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_960kHz_r18_Sym32  = 0
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_960kHz_r18_Sym64  = 1
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_960kHz_r18_Sym128 = 2
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_960kHz_r18_Sym256 = 3
)

var featureSetUplinkPerCCV1850TwoPUSCHCBMultiDCISTx2PAdditionalTimeR18Scs960kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_960kHz_r18_Sym32, FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_960kHz_r18_Sym64, FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_960kHz_r18_Sym128, FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_960kHz_r18_Sym256},
}

const (
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_60kHz_r18_Sym1  = 0
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_60kHz_r18_Sym4  = 1
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_60kHz_r18_Sym8  = 2
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_60kHz_r18_Sym16 = 3
)

var featureSetUplinkPerCCV1850TwoPUSCHNonCBMultiDCISTx2PAdditionalTimeR18Scs60kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_60kHz_r18_Sym1, FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_60kHz_r18_Sym4, FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_60kHz_r18_Sym8, FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_60kHz_r18_Sym16},
}

const (
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_120kHz_r18_Sym4  = 0
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_120kHz_r18_Sym8  = 1
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_120kHz_r18_Sym16 = 2
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_120kHz_r18_Sym32 = 3
)

var featureSetUplinkPerCCV1850TwoPUSCHNonCBMultiDCISTx2PAdditionalTimeR18Scs120kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_120kHz_r18_Sym4, FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_120kHz_r18_Sym8, FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_120kHz_r18_Sym16, FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_120kHz_r18_Sym32},
}

const (
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_480kHz_r18_Sym16  = 0
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_480kHz_r18_Sym32  = 1
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_480kHz_r18_Sym64  = 2
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_480kHz_r18_Sym128 = 3
)

var featureSetUplinkPerCCV1850TwoPUSCHNonCBMultiDCISTx2PAdditionalTimeR18Scs480kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_480kHz_r18_Sym16, FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_480kHz_r18_Sym32, FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_480kHz_r18_Sym64, FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_480kHz_r18_Sym128},
}

const (
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_960kHz_r18_Sym32  = 0
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_960kHz_r18_Sym64  = 1
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_960kHz_r18_Sym128 = 2
	FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_960kHz_r18_Sym256 = 3
)

var featureSetUplinkPerCCV1850TwoPUSCHNonCBMultiDCISTx2PAdditionalTimeR18Scs960kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_960kHz_r18_Sym32, FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_960kHz_r18_Sym64, FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_960kHz_r18_Sym128, FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_960kHz_r18_Sym256},
}

type FeatureSetUplinkPerCC_v1850 struct {
	TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18 *struct {
		Choice         int
		Scs_60kHz_r18  *int64
		Scs_120kHz_r18 *int64
		Scs_480kHz_r18 *int64
		Scs_960kHz_r18 *int64
	}
	TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18 *struct {
		Choice         int
		Scs_60kHz_r18  *int64
		Scs_120kHz_r18 *int64
		Scs_480kHz_r18 *int64
		Scs_960kHz_r18 *int64
	}
}

func (ie *FeatureSetUplinkPerCC_v1850) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetUplinkPerCCV1850Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18 != nil, ie.TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18 != nil}); err != nil {
		return err
	}

	// 2. twoPUSCH-CB-MultiDCI-STx2P-AdditionalTime-r18: choice
	{
		if ie.TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(featureSetUplinkPerCC_v1850TwoPUSCHCBMultiDCISTx2PAdditionalTimeR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18).Choice {
			case FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_60kHz_r18:
				if err := e.EncodeEnumerated((*(*ie.TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18).Scs_60kHz_r18), featureSetUplinkPerCCV1850TwoPUSCHCBMultiDCISTx2PAdditionalTimeR18Scs60kHzR18Constraints); err != nil {
					return err
				}
			case FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_120kHz_r18:
				if err := e.EncodeEnumerated((*(*ie.TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18).Scs_120kHz_r18), featureSetUplinkPerCCV1850TwoPUSCHCBMultiDCISTx2PAdditionalTimeR18Scs120kHzR18Constraints); err != nil {
					return err
				}
			case FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_480kHz_r18:
				if err := e.EncodeEnumerated((*(*ie.TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18).Scs_480kHz_r18), featureSetUplinkPerCCV1850TwoPUSCHCBMultiDCISTx2PAdditionalTimeR18Scs480kHzR18Constraints); err != nil {
					return err
				}
			case FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_960kHz_r18:
				if err := e.EncodeEnumerated((*(*ie.TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18).Scs_960kHz_r18), featureSetUplinkPerCCV1850TwoPUSCHCBMultiDCISTx2PAdditionalTimeR18Scs960kHzR18Constraints); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 3. twoPUSCH-NonCB-MultiDCI-STx2P-AdditionalTime-r18: choice
	{
		if ie.TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(featureSetUplinkPerCC_v1850TwoPUSCHNonCBMultiDCISTx2PAdditionalTimeR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18).Choice {
			case FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_60kHz_r18:
				if err := e.EncodeEnumerated((*(*ie.TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18).Scs_60kHz_r18), featureSetUplinkPerCCV1850TwoPUSCHNonCBMultiDCISTx2PAdditionalTimeR18Scs60kHzR18Constraints); err != nil {
					return err
				}
			case FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_120kHz_r18:
				if err := e.EncodeEnumerated((*(*ie.TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18).Scs_120kHz_r18), featureSetUplinkPerCCV1850TwoPUSCHNonCBMultiDCISTx2PAdditionalTimeR18Scs120kHzR18Constraints); err != nil {
					return err
				}
			case FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_480kHz_r18:
				if err := e.EncodeEnumerated((*(*ie.TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18).Scs_480kHz_r18), featureSetUplinkPerCCV1850TwoPUSCHNonCBMultiDCISTx2PAdditionalTimeR18Scs480kHzR18Constraints); err != nil {
					return err
				}
			case FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_960kHz_r18:
				if err := e.EncodeEnumerated((*(*ie.TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18).Scs_960kHz_r18), featureSetUplinkPerCCV1850TwoPUSCHNonCBMultiDCISTx2PAdditionalTimeR18Scs960kHzR18Constraints); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *FeatureSetUplinkPerCC_v1850) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetUplinkPerCCV1850Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. twoPUSCH-CB-MultiDCI-STx2P-AdditionalTime-r18: choice
	{
		if seq.IsComponentPresent(0) {
			ie.TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18 = &struct {
				Choice         int
				Scs_60kHz_r18  *int64
				Scs_120kHz_r18 *int64
				Scs_480kHz_r18 *int64
				Scs_960kHz_r18 *int64
			}{}
			choiceDec := d.NewChoiceDecoder(featureSetUplinkPerCC_v1850TwoPUSCHCBMultiDCISTx2PAdditionalTimeR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_60kHz_r18:
				(*ie.TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18).Scs_60kHz_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1850TwoPUSCHCBMultiDCISTx2PAdditionalTimeR18Scs60kHzR18Constraints)
				if err != nil {
					return err
				}
				(*(*ie.TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18).Scs_60kHz_r18) = v
			case FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_120kHz_r18:
				(*ie.TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18).Scs_120kHz_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1850TwoPUSCHCBMultiDCISTx2PAdditionalTimeR18Scs120kHzR18Constraints)
				if err != nil {
					return err
				}
				(*(*ie.TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18).Scs_120kHz_r18) = v
			case FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_480kHz_r18:
				(*ie.TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18).Scs_480kHz_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1850TwoPUSCHCBMultiDCISTx2PAdditionalTimeR18Scs480kHzR18Constraints)
				if err != nil {
					return err
				}
				(*(*ie.TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18).Scs_480kHz_r18) = v
			case FeatureSetUplinkPerCC_v1850_TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18_Scs_960kHz_r18:
				(*ie.TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18).Scs_960kHz_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1850TwoPUSCHCBMultiDCISTx2PAdditionalTimeR18Scs960kHzR18Constraints)
				if err != nil {
					return err
				}
				(*(*ie.TwoPUSCH_CB_MultiDCI_STx2P_AdditionalTime_r18).Scs_960kHz_r18) = v
			}
		}
	}

	// 3. twoPUSCH-NonCB-MultiDCI-STx2P-AdditionalTime-r18: choice
	{
		if seq.IsComponentPresent(1) {
			ie.TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18 = &struct {
				Choice         int
				Scs_60kHz_r18  *int64
				Scs_120kHz_r18 *int64
				Scs_480kHz_r18 *int64
				Scs_960kHz_r18 *int64
			}{}
			choiceDec := d.NewChoiceDecoder(featureSetUplinkPerCC_v1850TwoPUSCHNonCBMultiDCISTx2PAdditionalTimeR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_60kHz_r18:
				(*ie.TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18).Scs_60kHz_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1850TwoPUSCHNonCBMultiDCISTx2PAdditionalTimeR18Scs60kHzR18Constraints)
				if err != nil {
					return err
				}
				(*(*ie.TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18).Scs_60kHz_r18) = v
			case FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_120kHz_r18:
				(*ie.TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18).Scs_120kHz_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1850TwoPUSCHNonCBMultiDCISTx2PAdditionalTimeR18Scs120kHzR18Constraints)
				if err != nil {
					return err
				}
				(*(*ie.TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18).Scs_120kHz_r18) = v
			case FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_480kHz_r18:
				(*ie.TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18).Scs_480kHz_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1850TwoPUSCHNonCBMultiDCISTx2PAdditionalTimeR18Scs480kHzR18Constraints)
				if err != nil {
					return err
				}
				(*(*ie.TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18).Scs_480kHz_r18) = v
			case FeatureSetUplinkPerCC_v1850_TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18_Scs_960kHz_r18:
				(*ie.TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18).Scs_960kHz_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1850TwoPUSCHNonCBMultiDCISTx2PAdditionalTimeR18Scs960kHzR18Constraints)
				if err != nil {
					return err
				}
				(*(*ie.TwoPUSCH_NonCB_MultiDCI_STx2P_AdditionalTime_r18).Scs_960kHz_r18) = v
			}
		}
	}

	return nil
}
