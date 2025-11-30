package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps0     aper.Enumerated = 0
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps8     aper.Enumerated = 1
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps16    aper.Enumerated = 2
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps32    aper.Enumerated = 3
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps64    aper.Enumerated = 4
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps128   aper.Enumerated = 5
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps256   aper.Enumerated = 6
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps512   aper.Enumerated = 7
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps1024  aper.Enumerated = 8
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps2048  aper.Enumerated = 9
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps4096  aper.Enumerated = 10
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps8192  aper.Enumerated = 11
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps16384 aper.Enumerated = 12
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps32768 aper.Enumerated = 13
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_kBps65536 aper.Enumerated = 14
	SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16_Enum_infinity  aper.Enumerated = 15
)

type SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16 struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16", err)
	}
	return nil
}

func (ie *SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
