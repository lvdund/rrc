package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl0     aper.Enumerated = 0
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl1     aper.Enumerated = 1
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl2     aper.Enumerated = 2
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl4     aper.Enumerated = 3
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl6     aper.Enumerated = 4
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl8     aper.Enumerated = 5
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl16    aper.Enumerated = 6
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl24    aper.Enumerated = 7
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl33    aper.Enumerated = 8
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl40    aper.Enumerated = 9
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl64    aper.Enumerated = 10
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl80    aper.Enumerated = 11
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl96    aper.Enumerated = 12
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl112   aper.Enumerated = 13
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl128   aper.Enumerated = 14
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl160   aper.Enumerated = 15
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl320   aper.Enumerated = 16
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare15 aper.Enumerated = 17
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare14 aper.Enumerated = 18
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare13 aper.Enumerated = 19
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare12 aper.Enumerated = 20
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare11 aper.Enumerated = 21
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare10 aper.Enumerated = 22
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare9  aper.Enumerated = 23
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare8  aper.Enumerated = 24
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare7  aper.Enumerated = 25
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare6  aper.Enumerated = 26
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare5  aper.Enumerated = 27
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare4  aper.Enumerated = 28
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare3  aper.Enumerated = 29
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare2  aper.Enumerated = 30
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare1  aper.Enumerated = 31
)

type DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17 struct {
	Value aper.Enumerated `lb:0,ub:31,madatory`
}

func (ie *DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Encode DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17", err)
	}
	return nil
}

func (ie *DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Decode DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
