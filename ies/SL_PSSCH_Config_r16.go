package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_PSSCH_Config_r16 struct {
	Sl_PSSCH_DMRS_TimePatternList_r16 []int64                             `lb:1,ub:3,e_lb:2,e_ub:4,optional`
	Sl_BetaOffsets2ndSCI_r16          []SL_BetaOffsets_r16                `lb:4,ub:4,optional`
	Sl_Scaling_r16                    *SL_PSSCH_Config_r16_sl_Scaling_r16 `optional`
}

func (ie *SL_PSSCH_Config_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Sl_PSSCH_DMRS_TimePatternList_r16) > 0, len(ie.Sl_BetaOffsets2ndSCI_r16) > 0, ie.Sl_Scaling_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Sl_PSSCH_DMRS_TimePatternList_r16) > 0 {
		tmp_Sl_PSSCH_DMRS_TimePatternList_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: 3}, false)
		for _, i := range ie.Sl_PSSCH_DMRS_TimePatternList_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 2, Ub: 4}, false)
			tmp_Sl_PSSCH_DMRS_TimePatternList_r16.Value = append(tmp_Sl_PSSCH_DMRS_TimePatternList_r16.Value, &tmp_ie)
		}
		if err = tmp_Sl_PSSCH_DMRS_TimePatternList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PSSCH_DMRS_TimePatternList_r16", err)
		}
	}
	if len(ie.Sl_BetaOffsets2ndSCI_r16) > 0 {
		tmp_Sl_BetaOffsets2ndSCI_r16 := utils.NewSequence[*SL_BetaOffsets_r16]([]*SL_BetaOffsets_r16{}, aper.Constraint{Lb: 4, Ub: 4}, false)
		for _, i := range ie.Sl_BetaOffsets2ndSCI_r16 {
			tmp_Sl_BetaOffsets2ndSCI_r16.Value = append(tmp_Sl_BetaOffsets2ndSCI_r16.Value, &i)
		}
		if err = tmp_Sl_BetaOffsets2ndSCI_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_BetaOffsets2ndSCI_r16", err)
		}
	}
	if ie.Sl_Scaling_r16 != nil {
		if err = ie.Sl_Scaling_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_Scaling_r16", err)
		}
	}
	return nil
}

func (ie *SL_PSSCH_Config_r16) Decode(r *aper.AperReader) error {
	var err error
	var Sl_PSSCH_DMRS_TimePatternList_r16Present bool
	if Sl_PSSCH_DMRS_TimePatternList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_BetaOffsets2ndSCI_r16Present bool
	if Sl_BetaOffsets2ndSCI_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_Scaling_r16Present bool
	if Sl_Scaling_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_PSSCH_DMRS_TimePatternList_r16Present {
		tmp_Sl_PSSCH_DMRS_TimePatternList_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: 3}, false)
		fn_Sl_PSSCH_DMRS_TimePatternList_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 2, Ub: 4}, false)
			return &ie
		}
		if err = tmp_Sl_PSSCH_DMRS_TimePatternList_r16.Decode(r, fn_Sl_PSSCH_DMRS_TimePatternList_r16); err != nil {
			return utils.WrapError("Decode Sl_PSSCH_DMRS_TimePatternList_r16", err)
		}
		ie.Sl_PSSCH_DMRS_TimePatternList_r16 = []int64{}
		for _, i := range tmp_Sl_PSSCH_DMRS_TimePatternList_r16.Value {
			ie.Sl_PSSCH_DMRS_TimePatternList_r16 = append(ie.Sl_PSSCH_DMRS_TimePatternList_r16, int64(i.Value))
		}
	}
	if Sl_BetaOffsets2ndSCI_r16Present {
		tmp_Sl_BetaOffsets2ndSCI_r16 := utils.NewSequence[*SL_BetaOffsets_r16]([]*SL_BetaOffsets_r16{}, aper.Constraint{Lb: 4, Ub: 4}, false)
		fn_Sl_BetaOffsets2ndSCI_r16 := func() *SL_BetaOffsets_r16 {
			return new(SL_BetaOffsets_r16)
		}
		if err = tmp_Sl_BetaOffsets2ndSCI_r16.Decode(r, fn_Sl_BetaOffsets2ndSCI_r16); err != nil {
			return utils.WrapError("Decode Sl_BetaOffsets2ndSCI_r16", err)
		}
		ie.Sl_BetaOffsets2ndSCI_r16 = []SL_BetaOffsets_r16{}
		for _, i := range tmp_Sl_BetaOffsets2ndSCI_r16.Value {
			ie.Sl_BetaOffsets2ndSCI_r16 = append(ie.Sl_BetaOffsets2ndSCI_r16, *i)
		}
	}
	if Sl_Scaling_r16Present {
		ie.Sl_Scaling_r16 = new(SL_PSSCH_Config_r16_sl_Scaling_r16)
		if err = ie.Sl_Scaling_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_Scaling_r16", err)
		}
	}
	return nil
}
