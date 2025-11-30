package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_ms100  aper.Enumerated = 0
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_ms200  aper.Enumerated = 1
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_ms300  aper.Enumerated = 2
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_ms400  aper.Enumerated = 3
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_ms600  aper.Enumerated = 4
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_ms1000 aper.Enumerated = 5
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_ms2000 aper.Enumerated = 6
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_ms3000 aper.Enumerated = 7
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_ms4000 aper.Enumerated = 8
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_spare7 aper.Enumerated = 9
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_spare6 aper.Enumerated = 10
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_spare5 aper.Enumerated = 11
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_spare4 aper.Enumerated = 12
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_spare3 aper.Enumerated = 13
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_spare2 aper.Enumerated = 14
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_spare1 aper.Enumerated = 15
)

type SDT_ConfigCommonSIB_r17_t319a_r17 struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *SDT_ConfigCommonSIB_r17_t319a_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode SDT_ConfigCommonSIB_r17_t319a_r17", err)
	}
	return nil
}

func (ie *SDT_ConfigCommonSIB_r17_t319a_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode SDT_ConfigCommonSIB_r17_t319a_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
