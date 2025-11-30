package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte32    aper.Enumerated = 0
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte100   aper.Enumerated = 1
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte200   aper.Enumerated = 2
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte400   aper.Enumerated = 3
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte600   aper.Enumerated = 4
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte800   aper.Enumerated = 5
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte1000  aper.Enumerated = 6
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte2000  aper.Enumerated = 7
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte4000  aper.Enumerated = 8
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte8000  aper.Enumerated = 9
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte9000  aper.Enumerated = 10
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte10000 aper.Enumerated = 11
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte12000 aper.Enumerated = 12
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte24000 aper.Enumerated = 13
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte48000 aper.Enumerated = 14
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte96000 aper.Enumerated = 15
)

type SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17 struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17", err)
	}
	return nil
}

func (ie *SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
