package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_TxResourceReq_v1700 struct {
	Sl_DRX_InfoFromRxList_r17 []SL_DRX_ConfigUC_SemiStatic_r17              `lb:1,ub:maxNrofSL_RxInfoSet_r17,optional`
	Sl_DRX_Indication_r17     *SL_TxResourceReq_v1700_sl_DRX_Indication_r17 `optional`
}

func (ie *SL_TxResourceReq_v1700) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Sl_DRX_InfoFromRxList_r17) > 0, ie.Sl_DRX_Indication_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Sl_DRX_InfoFromRxList_r17) > 0 {
		tmp_Sl_DRX_InfoFromRxList_r17 := utils.NewSequence[*SL_DRX_ConfigUC_SemiStatic_r17]([]*SL_DRX_ConfigUC_SemiStatic_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofSL_RxInfoSet_r17}, false)
		for _, i := range ie.Sl_DRX_InfoFromRxList_r17 {
			tmp_Sl_DRX_InfoFromRxList_r17.Value = append(tmp_Sl_DRX_InfoFromRxList_r17.Value, &i)
		}
		if err = tmp_Sl_DRX_InfoFromRxList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_DRX_InfoFromRxList_r17", err)
		}
	}
	if ie.Sl_DRX_Indication_r17 != nil {
		if err = ie.Sl_DRX_Indication_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_DRX_Indication_r17", err)
		}
	}
	return nil
}

func (ie *SL_TxResourceReq_v1700) Decode(r *aper.AperReader) error {
	var err error
	var Sl_DRX_InfoFromRxList_r17Present bool
	if Sl_DRX_InfoFromRxList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_DRX_Indication_r17Present bool
	if Sl_DRX_Indication_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_DRX_InfoFromRxList_r17Present {
		tmp_Sl_DRX_InfoFromRxList_r17 := utils.NewSequence[*SL_DRX_ConfigUC_SemiStatic_r17]([]*SL_DRX_ConfigUC_SemiStatic_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofSL_RxInfoSet_r17}, false)
		fn_Sl_DRX_InfoFromRxList_r17 := func() *SL_DRX_ConfigUC_SemiStatic_r17 {
			return new(SL_DRX_ConfigUC_SemiStatic_r17)
		}
		if err = tmp_Sl_DRX_InfoFromRxList_r17.Decode(r, fn_Sl_DRX_InfoFromRxList_r17); err != nil {
			return utils.WrapError("Decode Sl_DRX_InfoFromRxList_r17", err)
		}
		ie.Sl_DRX_InfoFromRxList_r17 = []SL_DRX_ConfigUC_SemiStatic_r17{}
		for _, i := range tmp_Sl_DRX_InfoFromRxList_r17.Value {
			ie.Sl_DRX_InfoFromRxList_r17 = append(ie.Sl_DRX_InfoFromRxList_r17, *i)
		}
	}
	if Sl_DRX_Indication_r17Present {
		ie.Sl_DRX_Indication_r17 = new(SL_TxResourceReq_v1700_sl_DRX_Indication_r17)
		if err = ie.Sl_DRX_Indication_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_DRX_Indication_r17", err)
		}
	}
	return nil
}
