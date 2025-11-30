package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17_nzp_CSI_RS2_r17 struct {
	ResourceSet2_r17 int64         `lb:1,ub:maxNrofNZP_CSI_RS_ResourceSetsPerConfig,madatory`
	Qcl_info2_r17    []TCI_StateId `lb:1,ub:maxNrofAP_CSI_RS_ResourcesPerSet,optional`
}

func (ie *CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17_nzp_CSI_RS2_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Qcl_info2_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.ResourceSet2_r17, &aper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourceSetsPerConfig}, false); err != nil {
		return utils.WrapError("WriteInteger ResourceSet2_r17", err)
	}
	if len(ie.Qcl_info2_r17) > 0 {
		tmp_Qcl_info2_r17 := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, aper.Constraint{Lb: 1, Ub: maxNrofAP_CSI_RS_ResourcesPerSet}, false)
		for _, i := range ie.Qcl_info2_r17 {
			tmp_Qcl_info2_r17.Value = append(tmp_Qcl_info2_r17.Value, &i)
		}
		if err = tmp_Qcl_info2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Qcl_info2_r17", err)
		}
	}
	return nil
}

func (ie *CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17_nzp_CSI_RS2_r17) Decode(r *aper.AperReader) error {
	var err error
	var Qcl_info2_r17Present bool
	if Qcl_info2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_ResourceSet2_r17 int64
	if tmp_int_ResourceSet2_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourceSetsPerConfig}, false); err != nil {
		return utils.WrapError("ReadInteger ResourceSet2_r17", err)
	}
	ie.ResourceSet2_r17 = tmp_int_ResourceSet2_r17
	if Qcl_info2_r17Present {
		tmp_Qcl_info2_r17 := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, aper.Constraint{Lb: 1, Ub: maxNrofAP_CSI_RS_ResourcesPerSet}, false)
		fn_Qcl_info2_r17 := func() *TCI_StateId {
			return new(TCI_StateId)
		}
		if err = tmp_Qcl_info2_r17.Decode(r, fn_Qcl_info2_r17); err != nil {
			return utils.WrapError("Decode Qcl_info2_r17", err)
		}
		ie.Qcl_info2_r17 = []TCI_StateId{}
		for _, i := range tmp_Qcl_info2_r17.Value {
			ie.Qcl_info2_r17 = append(ie.Qcl_info2_r17, *i)
		}
	}
	return nil
}
