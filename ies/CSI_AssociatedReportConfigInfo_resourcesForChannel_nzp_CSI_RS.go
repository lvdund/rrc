package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CSI_AssociatedReportConfigInfo_resourcesForChannel_nzp_CSI_RS struct {
	ResourceSet int64         `lb:1,ub:maxNrofNZP_CSI_RS_ResourceSetsPerConfig,madatory`
	Qcl_info    []TCI_StateId `lb:1,ub:maxNrofAP_CSI_RS_ResourcesPerSet,optional`
}

func (ie *CSI_AssociatedReportConfigInfo_resourcesForChannel_nzp_CSI_RS) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Qcl_info) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.ResourceSet, &aper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourceSetsPerConfig}, false); err != nil {
		return utils.WrapError("WriteInteger ResourceSet", err)
	}
	if len(ie.Qcl_info) > 0 {
		tmp_Qcl_info := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, aper.Constraint{Lb: 1, Ub: maxNrofAP_CSI_RS_ResourcesPerSet}, false)
		for _, i := range ie.Qcl_info {
			tmp_Qcl_info.Value = append(tmp_Qcl_info.Value, &i)
		}
		if err = tmp_Qcl_info.Encode(w); err != nil {
			return utils.WrapError("Encode Qcl_info", err)
		}
	}
	return nil
}

func (ie *CSI_AssociatedReportConfigInfo_resourcesForChannel_nzp_CSI_RS) Decode(r *aper.AperReader) error {
	var err error
	var Qcl_infoPresent bool
	if Qcl_infoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_ResourceSet int64
	if tmp_int_ResourceSet, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourceSetsPerConfig}, false); err != nil {
		return utils.WrapError("ReadInteger ResourceSet", err)
	}
	ie.ResourceSet = tmp_int_ResourceSet
	if Qcl_infoPresent {
		tmp_Qcl_info := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, aper.Constraint{Lb: 1, Ub: maxNrofAP_CSI_RS_ResourcesPerSet}, false)
		fn_Qcl_info := func() *TCI_StateId {
			return new(TCI_StateId)
		}
		if err = tmp_Qcl_info.Decode(r, fn_Qcl_info); err != nil {
			return utils.WrapError("Decode Qcl_info", err)
		}
		ie.Qcl_info = []TCI_StateId{}
		for _, i := range tmp_Qcl_info.Value {
			ie.Qcl_info = append(ie.Qcl_info, *i)
		}
	}
	return nil
}
