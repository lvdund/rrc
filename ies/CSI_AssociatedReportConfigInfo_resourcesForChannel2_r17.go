package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17_Choice_nothing uint64 = iota
	CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17_Choice_Nzp_CSI_RS2_r17
	CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17_Choice_Csi_SSB_ResourceSet2_r17
)

type CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17 struct {
	Choice                   uint64
	Nzp_CSI_RS2_r17          *CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17_nzp_CSI_RS2_r17
	Csi_SSB_ResourceSet2_r17 int64 `lb:1,ub:maxNrofCSI_SSB_ResourceSetsPerConfigExt,madatory`
}

func (ie *CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17_Choice_Nzp_CSI_RS2_r17:
		if err = ie.Nzp_CSI_RS2_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode Nzp_CSI_RS2_r17", err)
		}
	case CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17_Choice_Csi_SSB_ResourceSet2_r17:
		if err = w.WriteInteger(int64(ie.Csi_SSB_ResourceSet2_r17), &uper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourceSetsPerConfigExt}, false); err != nil {
			err = utils.WrapError("Encode Csi_SSB_ResourceSet2_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17_Choice_Nzp_CSI_RS2_r17:
		ie.Nzp_CSI_RS2_r17 = new(CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17_nzp_CSI_RS2_r17)
		if err = ie.Nzp_CSI_RS2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Nzp_CSI_RS2_r17", err)
		}
	case CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17_Choice_Csi_SSB_ResourceSet2_r17:
		var tmp_int_Csi_SSB_ResourceSet2_r17 int64
		if tmp_int_Csi_SSB_ResourceSet2_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourceSetsPerConfigExt}, false); err != nil {
			return utils.WrapError("Decode Csi_SSB_ResourceSet2_r17", err)
		}
		ie.Csi_SSB_ResourceSet2_r17 = tmp_int_Csi_SSB_ResourceSet2_r17
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
