package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_ReportConfig_reportQuantity_cri_RI_i1_CQI struct {
	Pdsch_BundleSizeForCSI *CSI_ReportConfig_reportQuantity_cri_RI_i1_CQI_pdsch_BundleSizeForCSI `optional`
}

func (ie *CSI_ReportConfig_reportQuantity_cri_RI_i1_CQI) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Pdsch_BundleSizeForCSI != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Pdsch_BundleSizeForCSI != nil {
		if err = ie.Pdsch_BundleSizeForCSI.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_BundleSizeForCSI", err)
		}
	}
	return nil
}

func (ie *CSI_ReportConfig_reportQuantity_cri_RI_i1_CQI) Decode(r *uper.UperReader) error {
	var err error
	var Pdsch_BundleSizeForCSIPresent bool
	if Pdsch_BundleSizeForCSIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Pdsch_BundleSizeForCSIPresent {
		ie.Pdsch_BundleSizeForCSI = new(CSI_ReportConfig_reportQuantity_cri_RI_i1_CQI_pdsch_BundleSizeForCSI)
		if err = ie.Pdsch_BundleSizeForCSI.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_BundleSizeForCSI", err)
		}
	}
	return nil
}
