package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CG_SDT_TA_ValidationConfig_r17_cg_SDT_RSRP_ChangeThreshold_r17_Enum_dB2    aper.Enumerated = 0
	CG_SDT_TA_ValidationConfig_r17_cg_SDT_RSRP_ChangeThreshold_r17_Enum_dB4    aper.Enumerated = 1
	CG_SDT_TA_ValidationConfig_r17_cg_SDT_RSRP_ChangeThreshold_r17_Enum_dB6    aper.Enumerated = 2
	CG_SDT_TA_ValidationConfig_r17_cg_SDT_RSRP_ChangeThreshold_r17_Enum_dB8    aper.Enumerated = 3
	CG_SDT_TA_ValidationConfig_r17_cg_SDT_RSRP_ChangeThreshold_r17_Enum_dB10   aper.Enumerated = 4
	CG_SDT_TA_ValidationConfig_r17_cg_SDT_RSRP_ChangeThreshold_r17_Enum_dB14   aper.Enumerated = 5
	CG_SDT_TA_ValidationConfig_r17_cg_SDT_RSRP_ChangeThreshold_r17_Enum_dB18   aper.Enumerated = 6
	CG_SDT_TA_ValidationConfig_r17_cg_SDT_RSRP_ChangeThreshold_r17_Enum_dB22   aper.Enumerated = 7
	CG_SDT_TA_ValidationConfig_r17_cg_SDT_RSRP_ChangeThreshold_r17_Enum_dB26   aper.Enumerated = 8
	CG_SDT_TA_ValidationConfig_r17_cg_SDT_RSRP_ChangeThreshold_r17_Enum_dB30   aper.Enumerated = 9
	CG_SDT_TA_ValidationConfig_r17_cg_SDT_RSRP_ChangeThreshold_r17_Enum_dB34   aper.Enumerated = 10
	CG_SDT_TA_ValidationConfig_r17_cg_SDT_RSRP_ChangeThreshold_r17_Enum_spare5 aper.Enumerated = 11
	CG_SDT_TA_ValidationConfig_r17_cg_SDT_RSRP_ChangeThreshold_r17_Enum_spare4 aper.Enumerated = 12
	CG_SDT_TA_ValidationConfig_r17_cg_SDT_RSRP_ChangeThreshold_r17_Enum_spare3 aper.Enumerated = 13
	CG_SDT_TA_ValidationConfig_r17_cg_SDT_RSRP_ChangeThreshold_r17_Enum_spare2 aper.Enumerated = 14
	CG_SDT_TA_ValidationConfig_r17_cg_SDT_RSRP_ChangeThreshold_r17_Enum_spare1 aper.Enumerated = 15
)

type CG_SDT_TA_ValidationConfig_r17_cg_SDT_RSRP_ChangeThreshold_r17 struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *CG_SDT_TA_ValidationConfig_r17_cg_SDT_RSRP_ChangeThreshold_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode CG_SDT_TA_ValidationConfig_r17_cg_SDT_RSRP_ChangeThreshold_r17", err)
	}
	return nil
}

func (ie *CG_SDT_TA_ValidationConfig_r17_cg_SDT_RSRP_ChangeThreshold_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode CG_SDT_TA_ValidationConfig_r17_cg_SDT_RSRP_ChangeThreshold_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
