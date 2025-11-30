package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_DRX_ConfigUC_r17_sl_drx_HARQ_RTT_Timer2_r17_Enum_sl0    aper.Enumerated = 0
	SL_DRX_ConfigUC_r17_sl_drx_HARQ_RTT_Timer2_r17_Enum_sl1    aper.Enumerated = 1
	SL_DRX_ConfigUC_r17_sl_drx_HARQ_RTT_Timer2_r17_Enum_sl2    aper.Enumerated = 2
	SL_DRX_ConfigUC_r17_sl_drx_HARQ_RTT_Timer2_r17_Enum_sl4    aper.Enumerated = 3
	SL_DRX_ConfigUC_r17_sl_drx_HARQ_RTT_Timer2_r17_Enum_spare4 aper.Enumerated = 4
	SL_DRX_ConfigUC_r17_sl_drx_HARQ_RTT_Timer2_r17_Enum_spare3 aper.Enumerated = 5
	SL_DRX_ConfigUC_r17_sl_drx_HARQ_RTT_Timer2_r17_Enum_spare2 aper.Enumerated = 6
	SL_DRX_ConfigUC_r17_sl_drx_HARQ_RTT_Timer2_r17_Enum_spare1 aper.Enumerated = 7
)

type SL_DRX_ConfigUC_r17_sl_drx_HARQ_RTT_Timer2_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SL_DRX_ConfigUC_r17_sl_drx_HARQ_RTT_Timer2_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SL_DRX_ConfigUC_r17_sl_drx_HARQ_RTT_Timer2_r17", err)
	}
	return nil
}

func (ie *SL_DRX_ConfigUC_r17_sl_drx_HARQ_RTT_Timer2_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SL_DRX_ConfigUC_r17_sl_drx_HARQ_RTT_Timer2_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
