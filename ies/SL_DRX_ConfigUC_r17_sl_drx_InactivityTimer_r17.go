package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_ms0    aper.Enumerated = 0
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_ms1    aper.Enumerated = 1
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_ms2    aper.Enumerated = 2
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_ms3    aper.Enumerated = 3
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_ms4    aper.Enumerated = 4
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_ms5    aper.Enumerated = 5
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_ms6    aper.Enumerated = 6
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_ms8    aper.Enumerated = 7
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_ms10   aper.Enumerated = 8
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_ms20   aper.Enumerated = 9
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_ms30   aper.Enumerated = 10
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_ms40   aper.Enumerated = 11
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_ms50   aper.Enumerated = 12
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_ms60   aper.Enumerated = 13
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_ms80   aper.Enumerated = 14
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_ms100  aper.Enumerated = 15
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_ms200  aper.Enumerated = 16
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_ms300  aper.Enumerated = 17
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_ms500  aper.Enumerated = 18
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_ms750  aper.Enumerated = 19
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_ms1280 aper.Enumerated = 20
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_ms1920 aper.Enumerated = 21
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_ms2560 aper.Enumerated = 22
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_spare9 aper.Enumerated = 23
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_spare8 aper.Enumerated = 24
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_spare7 aper.Enumerated = 25
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_spare6 aper.Enumerated = 26
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_spare5 aper.Enumerated = 27
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_spare4 aper.Enumerated = 28
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_spare3 aper.Enumerated = 29
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_spare2 aper.Enumerated = 30
	SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17_Enum_spare1 aper.Enumerated = 31
)

type SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17 struct {
	Value aper.Enumerated `lb:0,ub:31,madatory`
}

func (ie *SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Encode SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17", err)
	}
	return nil
}

func (ie *SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Decode SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
