package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_UE_SelectedConfig_r16 struct {
	Sl_PSSCH_TxConfigList_r16     *SL_PSSCH_TxConfigList_r16                        `optional`
	Sl_ProbResourceKeep_r16       *SL_UE_SelectedConfig_r16_sl_ProbResourceKeep_r16 `optional`
	Sl_ReselectAfter_r16          *SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16    `optional`
	Sl_CBR_CommonTxConfigList_r16 *SL_CBR_CommonTxConfigList_r16                    `optional`
	Ul_PrioritizationThres_r16    *int64                                            `lb:1,ub:16,optional`
	Sl_PrioritizationThres_r16    *int64                                            `lb:1,ub:8,optional`
}

func (ie *SL_UE_SelectedConfig_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_PSSCH_TxConfigList_r16 != nil, ie.Sl_ProbResourceKeep_r16 != nil, ie.Sl_ReselectAfter_r16 != nil, ie.Sl_CBR_CommonTxConfigList_r16 != nil, ie.Ul_PrioritizationThres_r16 != nil, ie.Sl_PrioritizationThres_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_PSSCH_TxConfigList_r16 != nil {
		if err = ie.Sl_PSSCH_TxConfigList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PSSCH_TxConfigList_r16", err)
		}
	}
	if ie.Sl_ProbResourceKeep_r16 != nil {
		if err = ie.Sl_ProbResourceKeep_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_ProbResourceKeep_r16", err)
		}
	}
	if ie.Sl_ReselectAfter_r16 != nil {
		if err = ie.Sl_ReselectAfter_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_ReselectAfter_r16", err)
		}
	}
	if ie.Sl_CBR_CommonTxConfigList_r16 != nil {
		if err = ie.Sl_CBR_CommonTxConfigList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_CBR_CommonTxConfigList_r16", err)
		}
	}
	if ie.Ul_PrioritizationThres_r16 != nil {
		if err = w.WriteInteger(*ie.Ul_PrioritizationThres_r16, &aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode Ul_PrioritizationThres_r16", err)
		}
	}
	if ie.Sl_PrioritizationThres_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_PrioritizationThres_r16, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode Sl_PrioritizationThres_r16", err)
		}
	}
	return nil
}

func (ie *SL_UE_SelectedConfig_r16) Decode(r *aper.AperReader) error {
	var err error
	var Sl_PSSCH_TxConfigList_r16Present bool
	if Sl_PSSCH_TxConfigList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_ProbResourceKeep_r16Present bool
	if Sl_ProbResourceKeep_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_ReselectAfter_r16Present bool
	if Sl_ReselectAfter_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_CBR_CommonTxConfigList_r16Present bool
	if Sl_CBR_CommonTxConfigList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_PrioritizationThres_r16Present bool
	if Ul_PrioritizationThres_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PrioritizationThres_r16Present bool
	if Sl_PrioritizationThres_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_PSSCH_TxConfigList_r16Present {
		ie.Sl_PSSCH_TxConfigList_r16 = new(SL_PSSCH_TxConfigList_r16)
		if err = ie.Sl_PSSCH_TxConfigList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_PSSCH_TxConfigList_r16", err)
		}
	}
	if Sl_ProbResourceKeep_r16Present {
		ie.Sl_ProbResourceKeep_r16 = new(SL_UE_SelectedConfig_r16_sl_ProbResourceKeep_r16)
		if err = ie.Sl_ProbResourceKeep_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_ProbResourceKeep_r16", err)
		}
	}
	if Sl_ReselectAfter_r16Present {
		ie.Sl_ReselectAfter_r16 = new(SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16)
		if err = ie.Sl_ReselectAfter_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_ReselectAfter_r16", err)
		}
	}
	if Sl_CBR_CommonTxConfigList_r16Present {
		ie.Sl_CBR_CommonTxConfigList_r16 = new(SL_CBR_CommonTxConfigList_r16)
		if err = ie.Sl_CBR_CommonTxConfigList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_CBR_CommonTxConfigList_r16", err)
		}
	}
	if Ul_PrioritizationThres_r16Present {
		var tmp_int_Ul_PrioritizationThres_r16 int64
		if tmp_int_Ul_PrioritizationThres_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode Ul_PrioritizationThres_r16", err)
		}
		ie.Ul_PrioritizationThres_r16 = &tmp_int_Ul_PrioritizationThres_r16
	}
	if Sl_PrioritizationThres_r16Present {
		var tmp_int_Sl_PrioritizationThres_r16 int64
		if tmp_int_Sl_PrioritizationThres_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Sl_PrioritizationThres_r16", err)
		}
		ie.Sl_PrioritizationThres_r16 = &tmp_int_Sl_PrioritizationThres_r16
	}
	return nil
}
