package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SRS_TPC_PDCCH_Config struct {
	Srs_CC_SetIndexlist []SRS_CC_SetIndex `lb:1,ub:4,optional`
}

func (ie *SRS_TPC_PDCCH_Config) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Srs_CC_SetIndexlist) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Srs_CC_SetIndexlist) > 0 {
		tmp_Srs_CC_SetIndexlist := utils.NewSequence[*SRS_CC_SetIndex]([]*SRS_CC_SetIndex{}, aper.Constraint{Lb: 1, Ub: 4}, false)
		for _, i := range ie.Srs_CC_SetIndexlist {
			tmp_Srs_CC_SetIndexlist.Value = append(tmp_Srs_CC_SetIndexlist.Value, &i)
		}
		if err = tmp_Srs_CC_SetIndexlist.Encode(w); err != nil {
			return utils.WrapError("Encode Srs_CC_SetIndexlist", err)
		}
	}
	return nil
}

func (ie *SRS_TPC_PDCCH_Config) Decode(r *aper.AperReader) error {
	var err error
	var Srs_CC_SetIndexlistPresent bool
	if Srs_CC_SetIndexlistPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Srs_CC_SetIndexlistPresent {
		tmp_Srs_CC_SetIndexlist := utils.NewSequence[*SRS_CC_SetIndex]([]*SRS_CC_SetIndex{}, aper.Constraint{Lb: 1, Ub: 4}, false)
		fn_Srs_CC_SetIndexlist := func() *SRS_CC_SetIndex {
			return new(SRS_CC_SetIndex)
		}
		if err = tmp_Srs_CC_SetIndexlist.Decode(r, fn_Srs_CC_SetIndexlist); err != nil {
			return utils.WrapError("Decode Srs_CC_SetIndexlist", err)
		}
		ie.Srs_CC_SetIndexlist = []SRS_CC_SetIndex{}
		for _, i := range tmp_Srs_CC_SetIndexlist.Value {
			ie.Srs_CC_SetIndexlist = append(ie.Srs_CC_SetIndexlist, *i)
		}
	}
	return nil
}
