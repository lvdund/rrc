package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_ReportConfig_reportQuantity_cri_RI_i1_CQI_pdsch_BundleSizeForCSI_Enum_n2 aper.Enumerated = 0
	CSI_ReportConfig_reportQuantity_cri_RI_i1_CQI_pdsch_BundleSizeForCSI_Enum_n4 aper.Enumerated = 1
)

type CSI_ReportConfig_reportQuantity_cri_RI_i1_CQI_pdsch_BundleSizeForCSI struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *CSI_ReportConfig_reportQuantity_cri_RI_i1_CQI_pdsch_BundleSizeForCSI) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode CSI_ReportConfig_reportQuantity_cri_RI_i1_CQI_pdsch_BundleSizeForCSI", err)
	}
	return nil
}

func (ie *CSI_ReportConfig_reportQuantity_cri_RI_i1_CQI_pdsch_BundleSizeForCSI) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode CSI_ReportConfig_reportQuantity_cri_RI_i1_CQI_pdsch_BundleSizeForCSI", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
