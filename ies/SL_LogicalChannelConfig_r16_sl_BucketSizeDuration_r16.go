package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_LogicalChannelConfig_r16_sl_BucketSizeDuration_r16_Enum_ms5    aper.Enumerated = 0
	SL_LogicalChannelConfig_r16_sl_BucketSizeDuration_r16_Enum_ms10   aper.Enumerated = 1
	SL_LogicalChannelConfig_r16_sl_BucketSizeDuration_r16_Enum_ms20   aper.Enumerated = 2
	SL_LogicalChannelConfig_r16_sl_BucketSizeDuration_r16_Enum_ms50   aper.Enumerated = 3
	SL_LogicalChannelConfig_r16_sl_BucketSizeDuration_r16_Enum_ms100  aper.Enumerated = 4
	SL_LogicalChannelConfig_r16_sl_BucketSizeDuration_r16_Enum_ms150  aper.Enumerated = 5
	SL_LogicalChannelConfig_r16_sl_BucketSizeDuration_r16_Enum_ms300  aper.Enumerated = 6
	SL_LogicalChannelConfig_r16_sl_BucketSizeDuration_r16_Enum_ms500  aper.Enumerated = 7
	SL_LogicalChannelConfig_r16_sl_BucketSizeDuration_r16_Enum_ms1000 aper.Enumerated = 8
	SL_LogicalChannelConfig_r16_sl_BucketSizeDuration_r16_Enum_spare7 aper.Enumerated = 9
	SL_LogicalChannelConfig_r16_sl_BucketSizeDuration_r16_Enum_spare6 aper.Enumerated = 10
	SL_LogicalChannelConfig_r16_sl_BucketSizeDuration_r16_Enum_spare5 aper.Enumerated = 11
	SL_LogicalChannelConfig_r16_sl_BucketSizeDuration_r16_Enum_spare4 aper.Enumerated = 12
	SL_LogicalChannelConfig_r16_sl_BucketSizeDuration_r16_Enum_spare3 aper.Enumerated = 13
	SL_LogicalChannelConfig_r16_sl_BucketSizeDuration_r16_Enum_spare2 aper.Enumerated = 14
	SL_LogicalChannelConfig_r16_sl_BucketSizeDuration_r16_Enum_spare1 aper.Enumerated = 15
)

type SL_LogicalChannelConfig_r16_sl_BucketSizeDuration_r16 struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *SL_LogicalChannelConfig_r16_sl_BucketSizeDuration_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode SL_LogicalChannelConfig_r16_sl_BucketSizeDuration_r16", err)
	}
	return nil
}

func (ie *SL_LogicalChannelConfig_r16_sl_BucketSizeDuration_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode SL_LogicalChannelConfig_r16_sl_BucketSizeDuration_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
