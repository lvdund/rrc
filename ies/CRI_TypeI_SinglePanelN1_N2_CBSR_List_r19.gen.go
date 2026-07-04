// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CRI-TypeI-SinglePanelN1-N2-CBSR-List-r19 (line 6332).

var cRITypeISinglePanelN1N2CBSRListR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "two-one-r19"},
		{Name: "two-two-r19"},
		{Name: "four-one-r19"},
		{Name: "three-two-r19"},
		{Name: "six-one-r19"},
		{Name: "four-two-r19"},
		{Name: "eight-one-r19"},
		{Name: "four-three-r19"},
		{Name: "six-two-r19"},
		{Name: "twelve-one-r19"},
		{Name: "four-four-r19"},
		{Name: "eight-two-r19"},
		{Name: "sixteen-one-r19"},
	},
}

const (
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Two_One_r19     = 0
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Two_Two_r19     = 1
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_One_r19    = 2
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Three_Two_r19   = 3
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Six_One_r19     = 4
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_Two_r19    = 5
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Eight_One_r19   = 6
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_Three_r19  = 7
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Six_Two_r19     = 8
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Twelve_One_r19  = 9
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_Four_r19   = 10
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Eight_Two_r19   = 11
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Sixteen_One_r19 = 12
)

var cRITypeISinglePanelN1N2CBSRListR19TwoOneR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "no-cbsr-r19"},
		{Name: "cbsr-list-r19"},
	},
}

const (
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Two_One_r19_No_Cbsr_r19   = 0
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Two_One_r19_Cbsr_List_r19 = 1
)

var cRITypeISinglePanelN1N2CBSRListR19TwoOneR19CbsrListR19Constraints = per.SizeRange(1, 8)

var cRITypeISinglePanelN1N2CBSRListR19TwoTwoR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "no-cbsr-r19"},
		{Name: "cbsr-list-r19"},
	},
}

const (
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Two_Two_r19_No_Cbsr_r19   = 0
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Two_Two_r19_Cbsr_List_r19 = 1
)

var cRITypeISinglePanelN1N2CBSRListR19TwoTwoR19CbsrListR19Constraints = per.SizeRange(1, 8)

var cRITypeISinglePanelN1N2CBSRListR19FourOneR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "no-cbsr-r19"},
		{Name: "cbsr-list-r19"},
	},
}

const (
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_One_r19_No_Cbsr_r19   = 0
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_One_r19_Cbsr_List_r19 = 1
)

var cRITypeISinglePanelN1N2CBSRListR19FourOneR19CbsrListR19Constraints = per.SizeRange(1, 8)

var cRITypeISinglePanelN1N2CBSRListR19ThreeTwoR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "no-cbsr-r19"},
		{Name: "cbsr-list-r19"},
	},
}

const (
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Three_Two_r19_No_Cbsr_r19   = 0
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Three_Two_r19_Cbsr_List_r19 = 1
)

var cRITypeISinglePanelN1N2CBSRListR19ThreeTwoR19CbsrListR19Constraints = per.SizeRange(1, 8)

var cRITypeISinglePanelN1N2CBSRListR19SixOneR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "no-cbsr-r19"},
		{Name: "cbsr-list-r19"},
	},
}

const (
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Six_One_r19_No_Cbsr_r19   = 0
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Six_One_r19_Cbsr_List_r19 = 1
)

var cRITypeISinglePanelN1N2CBSRListR19SixOneR19CbsrListR19Constraints = per.SizeRange(1, 8)

var cRITypeISinglePanelN1N2CBSRListR19FourTwoR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "no-cbsr-r19"},
		{Name: "cbsr-list-r19"},
	},
}

const (
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_Two_r19_No_Cbsr_r19   = 0
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_Two_r19_Cbsr_List_r19 = 1
)

var cRITypeISinglePanelN1N2CBSRListR19FourTwoR19CbsrListR19Constraints = per.SizeRange(1, 8)

var cRITypeISinglePanelN1N2CBSRListR19EightOneR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "no-cbsr-r19"},
		{Name: "cbsr-list-r19"},
	},
}

const (
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Eight_One_r19_No_Cbsr_r19   = 0
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Eight_One_r19_Cbsr_List_r19 = 1
)

var cRITypeISinglePanelN1N2CBSRListR19EightOneR19CbsrListR19Constraints = per.SizeRange(1, 8)

var cRITypeISinglePanelN1N2CBSRListR19FourThreeR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "no-cbsr-r19"},
		{Name: "cbsr-list-r19"},
	},
}

const (
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_Three_r19_No_Cbsr_r19   = 0
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_Three_r19_Cbsr_List_r19 = 1
)

var cRITypeISinglePanelN1N2CBSRListR19FourThreeR19CbsrListR19Constraints = per.SizeRange(1, 8)

var cRITypeISinglePanelN1N2CBSRListR19SixTwoR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "no-cbsr-r19"},
		{Name: "cbsr-list-r19"},
	},
}

const (
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Six_Two_r19_No_Cbsr_r19   = 0
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Six_Two_r19_Cbsr_List_r19 = 1
)

var cRITypeISinglePanelN1N2CBSRListR19SixTwoR19CbsrListR19Constraints = per.SizeRange(1, 8)

var cRITypeISinglePanelN1N2CBSRListR19TwelveOneR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "no-cbsr-r19"},
		{Name: "cbsr-list-r19"},
	},
}

const (
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Twelve_One_r19_No_Cbsr_r19   = 0
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Twelve_One_r19_Cbsr_List_r19 = 1
)

var cRITypeISinglePanelN1N2CBSRListR19TwelveOneR19CbsrListR19Constraints = per.SizeRange(1, 8)

var cRITypeISinglePanelN1N2CBSRListR19FourFourR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "no-cbsr-r19"},
		{Name: "cbsr-list-r19"},
	},
}

const (
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_Four_r19_No_Cbsr_r19   = 0
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_Four_r19_Cbsr_List_r19 = 1
)

var cRITypeISinglePanelN1N2CBSRListR19FourFourR19CbsrListR19Constraints = per.SizeRange(1, 8)

var cRITypeISinglePanelN1N2CBSRListR19EightTwoR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "no-cbsr-r19"},
		{Name: "cbsr-list-r19"},
	},
}

const (
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Eight_Two_r19_No_Cbsr_r19   = 0
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Eight_Two_r19_Cbsr_List_r19 = 1
)

var cRITypeISinglePanelN1N2CBSRListR19EightTwoR19CbsrListR19Constraints = per.SizeRange(1, 8)

var cRITypeISinglePanelN1N2CBSRListR19SixteenOneR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "no-cbsr-r19"},
		{Name: "cbsr-list-r19"},
	},
}

const (
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Sixteen_One_r19_No_Cbsr_r19   = 0
	CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Sixteen_One_r19_Cbsr_List_r19 = 1
)

var cRITypeISinglePanelN1N2CBSRListR19SixteenOneR19CbsrListR19Constraints = per.SizeRange(1, 8)

type CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19 struct {
	Choice      int
	Two_One_r19 *struct {
		Choice        int
		No_Cbsr_r19   *struct{}
		Cbsr_List_r19 *[]per.BitString
	}
	Two_Two_r19 *struct {
		Choice        int
		No_Cbsr_r19   *struct{}
		Cbsr_List_r19 *[]per.BitString
	}
	Four_One_r19 *struct {
		Choice        int
		No_Cbsr_r19   *struct{}
		Cbsr_List_r19 *[]per.BitString
	}
	Three_Two_r19 *struct {
		Choice        int
		No_Cbsr_r19   *struct{}
		Cbsr_List_r19 *[]per.BitString
	}
	Six_One_r19 *struct {
		Choice        int
		No_Cbsr_r19   *struct{}
		Cbsr_List_r19 *[]per.BitString
	}
	Four_Two_r19 *struct {
		Choice        int
		No_Cbsr_r19   *struct{}
		Cbsr_List_r19 *[]per.BitString
	}
	Eight_One_r19 *struct {
		Choice        int
		No_Cbsr_r19   *struct{}
		Cbsr_List_r19 *[]per.BitString
	}
	Four_Three_r19 *struct {
		Choice        int
		No_Cbsr_r19   *struct{}
		Cbsr_List_r19 *[]per.BitString
	}
	Six_Two_r19 *struct {
		Choice        int
		No_Cbsr_r19   *struct{}
		Cbsr_List_r19 *[]per.BitString
	}
	Twelve_One_r19 *struct {
		Choice        int
		No_Cbsr_r19   *struct{}
		Cbsr_List_r19 *[]per.BitString
	}
	Four_Four_r19 *struct {
		Choice        int
		No_Cbsr_r19   *struct{}
		Cbsr_List_r19 *[]per.BitString
	}
	Eight_Two_r19 *struct {
		Choice        int
		No_Cbsr_r19   *struct{}
		Cbsr_List_r19 *[]per.BitString
	}
	Sixteen_One_r19 *struct {
		Choice        int
		No_Cbsr_r19   *struct{}
		Cbsr_List_r19 *[]per.BitString
	}
}

func (ie *CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(cRITypeISinglePanelN1N2CBSRListR19Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Two_One_r19:
		choiceEnc := e.NewChoiceEncoder(cRITypeISinglePanelN1N2CBSRListR19TwoOneR19Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Two_One_r19).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Two_One_r19).Choice {
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Two_One_r19_No_Cbsr_r19:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Two_One_r19_Cbsr_List_r19:
			seqOf := e.NewSequenceOfEncoder(cRITypeISinglePanelN1N2CBSRListR19TwoOneR19CbsrListR19Constraints)
			if err := seqOf.EncodeLength(int64(len((*(*ie.Two_One_r19).Cbsr_List_r19)))); err != nil {
				return err
			}
			for i := range *(*ie.Two_One_r19).Cbsr_List_r19 {
				if err := e.EncodeBitString((*(*ie.Two_One_r19).Cbsr_List_r19)[i], per.FixedSize(8)); err != nil {
					return err
				}
			}
		}
	case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Two_Two_r19:
		choiceEnc := e.NewChoiceEncoder(cRITypeISinglePanelN1N2CBSRListR19TwoTwoR19Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Two_Two_r19).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Two_Two_r19).Choice {
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Two_Two_r19_No_Cbsr_r19:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Two_Two_r19_Cbsr_List_r19:
			seqOf := e.NewSequenceOfEncoder(cRITypeISinglePanelN1N2CBSRListR19TwoTwoR19CbsrListR19Constraints)
			if err := seqOf.EncodeLength(int64(len((*(*ie.Two_Two_r19).Cbsr_List_r19)))); err != nil {
				return err
			}
			for i := range *(*ie.Two_Two_r19).Cbsr_List_r19 {
				if err := e.EncodeBitString((*(*ie.Two_Two_r19).Cbsr_List_r19)[i], per.FixedSize(64)); err != nil {
					return err
				}
			}
		}
	case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_One_r19:
		choiceEnc := e.NewChoiceEncoder(cRITypeISinglePanelN1N2CBSRListR19FourOneR19Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Four_One_r19).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Four_One_r19).Choice {
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_One_r19_No_Cbsr_r19:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_One_r19_Cbsr_List_r19:
			seqOf := e.NewSequenceOfEncoder(cRITypeISinglePanelN1N2CBSRListR19FourOneR19CbsrListR19Constraints)
			if err := seqOf.EncodeLength(int64(len((*(*ie.Four_One_r19).Cbsr_List_r19)))); err != nil {
				return err
			}
			for i := range *(*ie.Four_One_r19).Cbsr_List_r19 {
				if err := e.EncodeBitString((*(*ie.Four_One_r19).Cbsr_List_r19)[i], per.FixedSize(16)); err != nil {
					return err
				}
			}
		}
	case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Three_Two_r19:
		choiceEnc := e.NewChoiceEncoder(cRITypeISinglePanelN1N2CBSRListR19ThreeTwoR19Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Three_Two_r19).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Three_Two_r19).Choice {
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Three_Two_r19_No_Cbsr_r19:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Three_Two_r19_Cbsr_List_r19:
			seqOf := e.NewSequenceOfEncoder(cRITypeISinglePanelN1N2CBSRListR19ThreeTwoR19CbsrListR19Constraints)
			if err := seqOf.EncodeLength(int64(len((*(*ie.Three_Two_r19).Cbsr_List_r19)))); err != nil {
				return err
			}
			for i := range *(*ie.Three_Two_r19).Cbsr_List_r19 {
				if err := e.EncodeBitString((*(*ie.Three_Two_r19).Cbsr_List_r19)[i], per.FixedSize(96)); err != nil {
					return err
				}
			}
		}
	case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Six_One_r19:
		choiceEnc := e.NewChoiceEncoder(cRITypeISinglePanelN1N2CBSRListR19SixOneR19Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Six_One_r19).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Six_One_r19).Choice {
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Six_One_r19_No_Cbsr_r19:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Six_One_r19_Cbsr_List_r19:
			seqOf := e.NewSequenceOfEncoder(cRITypeISinglePanelN1N2CBSRListR19SixOneR19CbsrListR19Constraints)
			if err := seqOf.EncodeLength(int64(len((*(*ie.Six_One_r19).Cbsr_List_r19)))); err != nil {
				return err
			}
			for i := range *(*ie.Six_One_r19).Cbsr_List_r19 {
				if err := e.EncodeBitString((*(*ie.Six_One_r19).Cbsr_List_r19)[i], per.FixedSize(24)); err != nil {
					return err
				}
			}
		}
	case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_Two_r19:
		choiceEnc := e.NewChoiceEncoder(cRITypeISinglePanelN1N2CBSRListR19FourTwoR19Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Four_Two_r19).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Four_Two_r19).Choice {
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_Two_r19_No_Cbsr_r19:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_Two_r19_Cbsr_List_r19:
			seqOf := e.NewSequenceOfEncoder(cRITypeISinglePanelN1N2CBSRListR19FourTwoR19CbsrListR19Constraints)
			if err := seqOf.EncodeLength(int64(len((*(*ie.Four_Two_r19).Cbsr_List_r19)))); err != nil {
				return err
			}
			for i := range *(*ie.Four_Two_r19).Cbsr_List_r19 {
				if err := e.EncodeBitString((*(*ie.Four_Two_r19).Cbsr_List_r19)[i], per.FixedSize(128)); err != nil {
					return err
				}
			}
		}
	case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Eight_One_r19:
		choiceEnc := e.NewChoiceEncoder(cRITypeISinglePanelN1N2CBSRListR19EightOneR19Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Eight_One_r19).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Eight_One_r19).Choice {
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Eight_One_r19_No_Cbsr_r19:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Eight_One_r19_Cbsr_List_r19:
			seqOf := e.NewSequenceOfEncoder(cRITypeISinglePanelN1N2CBSRListR19EightOneR19CbsrListR19Constraints)
			if err := seqOf.EncodeLength(int64(len((*(*ie.Eight_One_r19).Cbsr_List_r19)))); err != nil {
				return err
			}
			for i := range *(*ie.Eight_One_r19).Cbsr_List_r19 {
				if err := e.EncodeBitString((*(*ie.Eight_One_r19).Cbsr_List_r19)[i], per.FixedSize(32)); err != nil {
					return err
				}
			}
		}
	case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_Three_r19:
		choiceEnc := e.NewChoiceEncoder(cRITypeISinglePanelN1N2CBSRListR19FourThreeR19Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Four_Three_r19).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Four_Three_r19).Choice {
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_Three_r19_No_Cbsr_r19:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_Three_r19_Cbsr_List_r19:
			seqOf := e.NewSequenceOfEncoder(cRITypeISinglePanelN1N2CBSRListR19FourThreeR19CbsrListR19Constraints)
			if err := seqOf.EncodeLength(int64(len((*(*ie.Four_Three_r19).Cbsr_List_r19)))); err != nil {
				return err
			}
			for i := range *(*ie.Four_Three_r19).Cbsr_List_r19 {
				if err := e.EncodeBitString((*(*ie.Four_Three_r19).Cbsr_List_r19)[i], per.FixedSize(192)); err != nil {
					return err
				}
			}
		}
	case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Six_Two_r19:
		choiceEnc := e.NewChoiceEncoder(cRITypeISinglePanelN1N2CBSRListR19SixTwoR19Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Six_Two_r19).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Six_Two_r19).Choice {
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Six_Two_r19_No_Cbsr_r19:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Six_Two_r19_Cbsr_List_r19:
			seqOf := e.NewSequenceOfEncoder(cRITypeISinglePanelN1N2CBSRListR19SixTwoR19CbsrListR19Constraints)
			if err := seqOf.EncodeLength(int64(len((*(*ie.Six_Two_r19).Cbsr_List_r19)))); err != nil {
				return err
			}
			for i := range *(*ie.Six_Two_r19).Cbsr_List_r19 {
				if err := e.EncodeBitString((*(*ie.Six_Two_r19).Cbsr_List_r19)[i], per.FixedSize(192)); err != nil {
					return err
				}
			}
		}
	case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Twelve_One_r19:
		choiceEnc := e.NewChoiceEncoder(cRITypeISinglePanelN1N2CBSRListR19TwelveOneR19Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Twelve_One_r19).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Twelve_One_r19).Choice {
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Twelve_One_r19_No_Cbsr_r19:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Twelve_One_r19_Cbsr_List_r19:
			seqOf := e.NewSequenceOfEncoder(cRITypeISinglePanelN1N2CBSRListR19TwelveOneR19CbsrListR19Constraints)
			if err := seqOf.EncodeLength(int64(len((*(*ie.Twelve_One_r19).Cbsr_List_r19)))); err != nil {
				return err
			}
			for i := range *(*ie.Twelve_One_r19).Cbsr_List_r19 {
				if err := e.EncodeBitString((*(*ie.Twelve_One_r19).Cbsr_List_r19)[i], per.FixedSize(48)); err != nil {
					return err
				}
			}
		}
	case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_Four_r19:
		choiceEnc := e.NewChoiceEncoder(cRITypeISinglePanelN1N2CBSRListR19FourFourR19Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Four_Four_r19).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Four_Four_r19).Choice {
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_Four_r19_No_Cbsr_r19:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_Four_r19_Cbsr_List_r19:
			seqOf := e.NewSequenceOfEncoder(cRITypeISinglePanelN1N2CBSRListR19FourFourR19CbsrListR19Constraints)
			if err := seqOf.EncodeLength(int64(len((*(*ie.Four_Four_r19).Cbsr_List_r19)))); err != nil {
				return err
			}
			for i := range *(*ie.Four_Four_r19).Cbsr_List_r19 {
				if err := e.EncodeBitString((*(*ie.Four_Four_r19).Cbsr_List_r19)[i], per.FixedSize(256)); err != nil {
					return err
				}
			}
		}
	case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Eight_Two_r19:
		choiceEnc := e.NewChoiceEncoder(cRITypeISinglePanelN1N2CBSRListR19EightTwoR19Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Eight_Two_r19).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Eight_Two_r19).Choice {
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Eight_Two_r19_No_Cbsr_r19:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Eight_Two_r19_Cbsr_List_r19:
			seqOf := e.NewSequenceOfEncoder(cRITypeISinglePanelN1N2CBSRListR19EightTwoR19CbsrListR19Constraints)
			if err := seqOf.EncodeLength(int64(len((*(*ie.Eight_Two_r19).Cbsr_List_r19)))); err != nil {
				return err
			}
			for i := range *(*ie.Eight_Two_r19).Cbsr_List_r19 {
				if err := e.EncodeBitString((*(*ie.Eight_Two_r19).Cbsr_List_r19)[i], per.FixedSize(256)); err != nil {
					return err
				}
			}
		}
	case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Sixteen_One_r19:
		choiceEnc := e.NewChoiceEncoder(cRITypeISinglePanelN1N2CBSRListR19SixteenOneR19Constraints)
		if err := choiceEnc.EncodeChoice(int64((*ie.Sixteen_One_r19).Choice), false, nil); err != nil {
			return err
		}
		switch (*ie.Sixteen_One_r19).Choice {
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Sixteen_One_r19_No_Cbsr_r19:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Sixteen_One_r19_Cbsr_List_r19:
			seqOf := e.NewSequenceOfEncoder(cRITypeISinglePanelN1N2CBSRListR19SixteenOneR19CbsrListR19Constraints)
			if err := seqOf.EncodeLength(int64(len((*(*ie.Sixteen_One_r19).Cbsr_List_r19)))); err != nil {
				return err
			}
			for i := range *(*ie.Sixteen_One_r19).Cbsr_List_r19 {
				if err := e.EncodeBitString((*(*ie.Sixteen_One_r19).Cbsr_List_r19)[i], per.FixedSize(64)); err != nil {
					return err
				}
			}
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "CRI-TypeI-SinglePanelN1-N2-CBSR-List-r19",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(cRITypeISinglePanelN1N2CBSRListR19Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "CRI-TypeI-SinglePanelN1-N2-CBSR-List-r19", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Two_One_r19:
		ie.Two_One_r19 = &struct {
			Choice        int
			No_Cbsr_r19   *struct{}
			Cbsr_List_r19 *[]per.BitString
		}{}
		choiceDec := d.NewChoiceDecoder(cRITypeISinglePanelN1N2CBSRListR19TwoOneR19Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Two_One_r19).Choice = int(index)
		switch index {
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Two_One_r19_No_Cbsr_r19:
			(*ie.Two_One_r19).No_Cbsr_r19 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Two_One_r19_Cbsr_List_r19:
			(*ie.Two_One_r19).Cbsr_List_r19 = new([]per.BitString)
			seqOf := d.NewSequenceOfDecoder(cRITypeISinglePanelN1N2CBSRListR19TwoOneR19CbsrListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			(*(*ie.Two_One_r19).Cbsr_List_r19) = make([]per.BitString, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeBitString(per.FixedSize(8))
				if err != nil {
					return err
				}
				(*(*ie.Two_One_r19).Cbsr_List_r19)[i] = v
			}
		}
	case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Two_Two_r19:
		ie.Two_Two_r19 = &struct {
			Choice        int
			No_Cbsr_r19   *struct{}
			Cbsr_List_r19 *[]per.BitString
		}{}
		choiceDec := d.NewChoiceDecoder(cRITypeISinglePanelN1N2CBSRListR19TwoTwoR19Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Two_Two_r19).Choice = int(index)
		switch index {
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Two_Two_r19_No_Cbsr_r19:
			(*ie.Two_Two_r19).No_Cbsr_r19 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Two_Two_r19_Cbsr_List_r19:
			(*ie.Two_Two_r19).Cbsr_List_r19 = new([]per.BitString)
			seqOf := d.NewSequenceOfDecoder(cRITypeISinglePanelN1N2CBSRListR19TwoTwoR19CbsrListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			(*(*ie.Two_Two_r19).Cbsr_List_r19) = make([]per.BitString, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeBitString(per.FixedSize(64))
				if err != nil {
					return err
				}
				(*(*ie.Two_Two_r19).Cbsr_List_r19)[i] = v
			}
		}
	case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_One_r19:
		ie.Four_One_r19 = &struct {
			Choice        int
			No_Cbsr_r19   *struct{}
			Cbsr_List_r19 *[]per.BitString
		}{}
		choiceDec := d.NewChoiceDecoder(cRITypeISinglePanelN1N2CBSRListR19FourOneR19Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Four_One_r19).Choice = int(index)
		switch index {
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_One_r19_No_Cbsr_r19:
			(*ie.Four_One_r19).No_Cbsr_r19 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_One_r19_Cbsr_List_r19:
			(*ie.Four_One_r19).Cbsr_List_r19 = new([]per.BitString)
			seqOf := d.NewSequenceOfDecoder(cRITypeISinglePanelN1N2CBSRListR19FourOneR19CbsrListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			(*(*ie.Four_One_r19).Cbsr_List_r19) = make([]per.BitString, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeBitString(per.FixedSize(16))
				if err != nil {
					return err
				}
				(*(*ie.Four_One_r19).Cbsr_List_r19)[i] = v
			}
		}
	case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Three_Two_r19:
		ie.Three_Two_r19 = &struct {
			Choice        int
			No_Cbsr_r19   *struct{}
			Cbsr_List_r19 *[]per.BitString
		}{}
		choiceDec := d.NewChoiceDecoder(cRITypeISinglePanelN1N2CBSRListR19ThreeTwoR19Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Three_Two_r19).Choice = int(index)
		switch index {
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Three_Two_r19_No_Cbsr_r19:
			(*ie.Three_Two_r19).No_Cbsr_r19 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Three_Two_r19_Cbsr_List_r19:
			(*ie.Three_Two_r19).Cbsr_List_r19 = new([]per.BitString)
			seqOf := d.NewSequenceOfDecoder(cRITypeISinglePanelN1N2CBSRListR19ThreeTwoR19CbsrListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			(*(*ie.Three_Two_r19).Cbsr_List_r19) = make([]per.BitString, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeBitString(per.FixedSize(96))
				if err != nil {
					return err
				}
				(*(*ie.Three_Two_r19).Cbsr_List_r19)[i] = v
			}
		}
	case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Six_One_r19:
		ie.Six_One_r19 = &struct {
			Choice        int
			No_Cbsr_r19   *struct{}
			Cbsr_List_r19 *[]per.BitString
		}{}
		choiceDec := d.NewChoiceDecoder(cRITypeISinglePanelN1N2CBSRListR19SixOneR19Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Six_One_r19).Choice = int(index)
		switch index {
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Six_One_r19_No_Cbsr_r19:
			(*ie.Six_One_r19).No_Cbsr_r19 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Six_One_r19_Cbsr_List_r19:
			(*ie.Six_One_r19).Cbsr_List_r19 = new([]per.BitString)
			seqOf := d.NewSequenceOfDecoder(cRITypeISinglePanelN1N2CBSRListR19SixOneR19CbsrListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			(*(*ie.Six_One_r19).Cbsr_List_r19) = make([]per.BitString, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeBitString(per.FixedSize(24))
				if err != nil {
					return err
				}
				(*(*ie.Six_One_r19).Cbsr_List_r19)[i] = v
			}
		}
	case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_Two_r19:
		ie.Four_Two_r19 = &struct {
			Choice        int
			No_Cbsr_r19   *struct{}
			Cbsr_List_r19 *[]per.BitString
		}{}
		choiceDec := d.NewChoiceDecoder(cRITypeISinglePanelN1N2CBSRListR19FourTwoR19Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Four_Two_r19).Choice = int(index)
		switch index {
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_Two_r19_No_Cbsr_r19:
			(*ie.Four_Two_r19).No_Cbsr_r19 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_Two_r19_Cbsr_List_r19:
			(*ie.Four_Two_r19).Cbsr_List_r19 = new([]per.BitString)
			seqOf := d.NewSequenceOfDecoder(cRITypeISinglePanelN1N2CBSRListR19FourTwoR19CbsrListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			(*(*ie.Four_Two_r19).Cbsr_List_r19) = make([]per.BitString, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeBitString(per.FixedSize(128))
				if err != nil {
					return err
				}
				(*(*ie.Four_Two_r19).Cbsr_List_r19)[i] = v
			}
		}
	case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Eight_One_r19:
		ie.Eight_One_r19 = &struct {
			Choice        int
			No_Cbsr_r19   *struct{}
			Cbsr_List_r19 *[]per.BitString
		}{}
		choiceDec := d.NewChoiceDecoder(cRITypeISinglePanelN1N2CBSRListR19EightOneR19Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Eight_One_r19).Choice = int(index)
		switch index {
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Eight_One_r19_No_Cbsr_r19:
			(*ie.Eight_One_r19).No_Cbsr_r19 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Eight_One_r19_Cbsr_List_r19:
			(*ie.Eight_One_r19).Cbsr_List_r19 = new([]per.BitString)
			seqOf := d.NewSequenceOfDecoder(cRITypeISinglePanelN1N2CBSRListR19EightOneR19CbsrListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			(*(*ie.Eight_One_r19).Cbsr_List_r19) = make([]per.BitString, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeBitString(per.FixedSize(32))
				if err != nil {
					return err
				}
				(*(*ie.Eight_One_r19).Cbsr_List_r19)[i] = v
			}
		}
	case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_Three_r19:
		ie.Four_Three_r19 = &struct {
			Choice        int
			No_Cbsr_r19   *struct{}
			Cbsr_List_r19 *[]per.BitString
		}{}
		choiceDec := d.NewChoiceDecoder(cRITypeISinglePanelN1N2CBSRListR19FourThreeR19Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Four_Three_r19).Choice = int(index)
		switch index {
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_Three_r19_No_Cbsr_r19:
			(*ie.Four_Three_r19).No_Cbsr_r19 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_Three_r19_Cbsr_List_r19:
			(*ie.Four_Three_r19).Cbsr_List_r19 = new([]per.BitString)
			seqOf := d.NewSequenceOfDecoder(cRITypeISinglePanelN1N2CBSRListR19FourThreeR19CbsrListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			(*(*ie.Four_Three_r19).Cbsr_List_r19) = make([]per.BitString, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeBitString(per.FixedSize(192))
				if err != nil {
					return err
				}
				(*(*ie.Four_Three_r19).Cbsr_List_r19)[i] = v
			}
		}
	case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Six_Two_r19:
		ie.Six_Two_r19 = &struct {
			Choice        int
			No_Cbsr_r19   *struct{}
			Cbsr_List_r19 *[]per.BitString
		}{}
		choiceDec := d.NewChoiceDecoder(cRITypeISinglePanelN1N2CBSRListR19SixTwoR19Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Six_Two_r19).Choice = int(index)
		switch index {
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Six_Two_r19_No_Cbsr_r19:
			(*ie.Six_Two_r19).No_Cbsr_r19 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Six_Two_r19_Cbsr_List_r19:
			(*ie.Six_Two_r19).Cbsr_List_r19 = new([]per.BitString)
			seqOf := d.NewSequenceOfDecoder(cRITypeISinglePanelN1N2CBSRListR19SixTwoR19CbsrListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			(*(*ie.Six_Two_r19).Cbsr_List_r19) = make([]per.BitString, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeBitString(per.FixedSize(192))
				if err != nil {
					return err
				}
				(*(*ie.Six_Two_r19).Cbsr_List_r19)[i] = v
			}
		}
	case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Twelve_One_r19:
		ie.Twelve_One_r19 = &struct {
			Choice        int
			No_Cbsr_r19   *struct{}
			Cbsr_List_r19 *[]per.BitString
		}{}
		choiceDec := d.NewChoiceDecoder(cRITypeISinglePanelN1N2CBSRListR19TwelveOneR19Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Twelve_One_r19).Choice = int(index)
		switch index {
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Twelve_One_r19_No_Cbsr_r19:
			(*ie.Twelve_One_r19).No_Cbsr_r19 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Twelve_One_r19_Cbsr_List_r19:
			(*ie.Twelve_One_r19).Cbsr_List_r19 = new([]per.BitString)
			seqOf := d.NewSequenceOfDecoder(cRITypeISinglePanelN1N2CBSRListR19TwelveOneR19CbsrListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			(*(*ie.Twelve_One_r19).Cbsr_List_r19) = make([]per.BitString, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeBitString(per.FixedSize(48))
				if err != nil {
					return err
				}
				(*(*ie.Twelve_One_r19).Cbsr_List_r19)[i] = v
			}
		}
	case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_Four_r19:
		ie.Four_Four_r19 = &struct {
			Choice        int
			No_Cbsr_r19   *struct{}
			Cbsr_List_r19 *[]per.BitString
		}{}
		choiceDec := d.NewChoiceDecoder(cRITypeISinglePanelN1N2CBSRListR19FourFourR19Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Four_Four_r19).Choice = int(index)
		switch index {
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_Four_r19_No_Cbsr_r19:
			(*ie.Four_Four_r19).No_Cbsr_r19 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Four_Four_r19_Cbsr_List_r19:
			(*ie.Four_Four_r19).Cbsr_List_r19 = new([]per.BitString)
			seqOf := d.NewSequenceOfDecoder(cRITypeISinglePanelN1N2CBSRListR19FourFourR19CbsrListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			(*(*ie.Four_Four_r19).Cbsr_List_r19) = make([]per.BitString, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeBitString(per.FixedSize(256))
				if err != nil {
					return err
				}
				(*(*ie.Four_Four_r19).Cbsr_List_r19)[i] = v
			}
		}
	case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Eight_Two_r19:
		ie.Eight_Two_r19 = &struct {
			Choice        int
			No_Cbsr_r19   *struct{}
			Cbsr_List_r19 *[]per.BitString
		}{}
		choiceDec := d.NewChoiceDecoder(cRITypeISinglePanelN1N2CBSRListR19EightTwoR19Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Eight_Two_r19).Choice = int(index)
		switch index {
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Eight_Two_r19_No_Cbsr_r19:
			(*ie.Eight_Two_r19).No_Cbsr_r19 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Eight_Two_r19_Cbsr_List_r19:
			(*ie.Eight_Two_r19).Cbsr_List_r19 = new([]per.BitString)
			seqOf := d.NewSequenceOfDecoder(cRITypeISinglePanelN1N2CBSRListR19EightTwoR19CbsrListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			(*(*ie.Eight_Two_r19).Cbsr_List_r19) = make([]per.BitString, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeBitString(per.FixedSize(256))
				if err != nil {
					return err
				}
				(*(*ie.Eight_Two_r19).Cbsr_List_r19)[i] = v
			}
		}
	case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Sixteen_One_r19:
		ie.Sixteen_One_r19 = &struct {
			Choice        int
			No_Cbsr_r19   *struct{}
			Cbsr_List_r19 *[]per.BitString
		}{}
		choiceDec := d.NewChoiceDecoder(cRITypeISinglePanelN1N2CBSRListR19SixteenOneR19Constraints)
		index, _, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		(*ie.Sixteen_One_r19).Choice = int(index)
		switch index {
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Sixteen_One_r19_No_Cbsr_r19:
			(*ie.Sixteen_One_r19).No_Cbsr_r19 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case CRI_TypeI_SinglePanelN1_N2_CBSR_List_r19_Sixteen_One_r19_Cbsr_List_r19:
			(*ie.Sixteen_One_r19).Cbsr_List_r19 = new([]per.BitString)
			seqOf := d.NewSequenceOfDecoder(cRITypeISinglePanelN1N2CBSRListR19SixteenOneR19CbsrListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			(*(*ie.Sixteen_One_r19).Cbsr_List_r19) = make([]per.BitString, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeBitString(per.FixedSize(64))
				if err != nil {
					return err
				}
				(*(*ie.Sixteen_One_r19).Cbsr_List_r19)[i] = v
			}
		}
	default:
		return &per.DecodeError{TypeName: "CRI-TypeI-SinglePanelN1-N2-CBSR-List-r19", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
