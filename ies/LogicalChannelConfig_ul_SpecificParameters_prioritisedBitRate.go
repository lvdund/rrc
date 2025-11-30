package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps0     aper.Enumerated = 0
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps8     aper.Enumerated = 1
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps16    aper.Enumerated = 2
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps32    aper.Enumerated = 3
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps64    aper.Enumerated = 4
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps128   aper.Enumerated = 5
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps256   aper.Enumerated = 6
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps512   aper.Enumerated = 7
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps1024  aper.Enumerated = 8
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps2048  aper.Enumerated = 9
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps4096  aper.Enumerated = 10
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps8192  aper.Enumerated = 11
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps16384 aper.Enumerated = 12
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps32768 aper.Enumerated = 13
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_kBps65536 aper.Enumerated = 14
	LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate_Enum_infinity  aper.Enumerated = 15
)

type LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate", err)
	}
	return nil
}

func (ie *LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
