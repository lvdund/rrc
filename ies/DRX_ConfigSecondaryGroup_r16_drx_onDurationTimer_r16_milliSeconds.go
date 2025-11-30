package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_ms1    aper.Enumerated = 0
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_ms2    aper.Enumerated = 1
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_ms3    aper.Enumerated = 2
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_ms4    aper.Enumerated = 3
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_ms5    aper.Enumerated = 4
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_ms6    aper.Enumerated = 5
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_ms8    aper.Enumerated = 6
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_ms10   aper.Enumerated = 7
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_ms20   aper.Enumerated = 8
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_ms30   aper.Enumerated = 9
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_ms40   aper.Enumerated = 10
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_ms50   aper.Enumerated = 11
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_ms60   aper.Enumerated = 12
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_ms80   aper.Enumerated = 13
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_ms100  aper.Enumerated = 14
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_ms200  aper.Enumerated = 15
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_ms300  aper.Enumerated = 16
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_ms400  aper.Enumerated = 17
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_ms500  aper.Enumerated = 18
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_ms600  aper.Enumerated = 19
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_ms800  aper.Enumerated = 20
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_ms1000 aper.Enumerated = 21
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_ms1200 aper.Enumerated = 22
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_ms1600 aper.Enumerated = 23
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_spare8 aper.Enumerated = 24
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_spare7 aper.Enumerated = 25
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_spare6 aper.Enumerated = 26
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_spare5 aper.Enumerated = 27
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_spare4 aper.Enumerated = 28
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_spare3 aper.Enumerated = 29
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_spare2 aper.Enumerated = 30
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds_Enum_spare1 aper.Enumerated = 31
)

type DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds struct {
	Value aper.Enumerated `lb:0,ub:31,madatory`
}

func (ie *DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Encode DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds", err)
	}
	return nil
}

func (ie *DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Decode DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
