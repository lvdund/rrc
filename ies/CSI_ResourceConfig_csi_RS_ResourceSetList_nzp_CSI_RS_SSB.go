package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_ResourceConfig_csi_RS_ResourceSetList_nzp_CSI_RS_SSB struct {
	Nzp_CSI_RS_ResourceSetList []NZP_CSI_RS_ResourceSetId `lb:1,ub:maxNrofNZP_CSI_RS_ResourceSetsPerConfig,optional`
	Csi_SSB_ResourceSetList    []CSI_SSB_ResourceSetId    `lb:1,ub:maxNrofCSI_SSB_ResourceSetsPerConfig,optional`
}

func (ie *CSI_ResourceConfig_csi_RS_ResourceSetList_nzp_CSI_RS_SSB) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Nzp_CSI_RS_ResourceSetList) > 0, len(ie.Csi_SSB_ResourceSetList) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Nzp_CSI_RS_ResourceSetList) > 0 {
		tmp_Nzp_CSI_RS_ResourceSetList := utils.NewSequence[*NZP_CSI_RS_ResourceSetId]([]*NZP_CSI_RS_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourceSetsPerConfig}, false)
		for _, i := range ie.Nzp_CSI_RS_ResourceSetList {
			tmp_Nzp_CSI_RS_ResourceSetList.Value = append(tmp_Nzp_CSI_RS_ResourceSetList.Value, &i)
		}
		if err = tmp_Nzp_CSI_RS_ResourceSetList.Encode(w); err != nil {
			return utils.WrapError("Encode Nzp_CSI_RS_ResourceSetList", err)
		}
	}
	if len(ie.Csi_SSB_ResourceSetList) > 0 {
		tmp_Csi_SSB_ResourceSetList := utils.NewSequence[*CSI_SSB_ResourceSetId]([]*CSI_SSB_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourceSetsPerConfig}, false)
		for _, i := range ie.Csi_SSB_ResourceSetList {
			tmp_Csi_SSB_ResourceSetList.Value = append(tmp_Csi_SSB_ResourceSetList.Value, &i)
		}
		if err = tmp_Csi_SSB_ResourceSetList.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_SSB_ResourceSetList", err)
		}
	}
	return nil
}

func (ie *CSI_ResourceConfig_csi_RS_ResourceSetList_nzp_CSI_RS_SSB) Decode(r *uper.UperReader) error {
	var err error
	var Nzp_CSI_RS_ResourceSetListPresent bool
	if Nzp_CSI_RS_ResourceSetListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_SSB_ResourceSetListPresent bool
	if Csi_SSB_ResourceSetListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Nzp_CSI_RS_ResourceSetListPresent {
		tmp_Nzp_CSI_RS_ResourceSetList := utils.NewSequence[*NZP_CSI_RS_ResourceSetId]([]*NZP_CSI_RS_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourceSetsPerConfig}, false)
		fn_Nzp_CSI_RS_ResourceSetList := func() *NZP_CSI_RS_ResourceSetId {
			return new(NZP_CSI_RS_ResourceSetId)
		}
		if err = tmp_Nzp_CSI_RS_ResourceSetList.Decode(r, fn_Nzp_CSI_RS_ResourceSetList); err != nil {
			return utils.WrapError("Decode Nzp_CSI_RS_ResourceSetList", err)
		}
		ie.Nzp_CSI_RS_ResourceSetList = []NZP_CSI_RS_ResourceSetId{}
		for _, i := range tmp_Nzp_CSI_RS_ResourceSetList.Value {
			ie.Nzp_CSI_RS_ResourceSetList = append(ie.Nzp_CSI_RS_ResourceSetList, *i)
		}
	}
	if Csi_SSB_ResourceSetListPresent {
		tmp_Csi_SSB_ResourceSetList := utils.NewSequence[*CSI_SSB_ResourceSetId]([]*CSI_SSB_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourceSetsPerConfig}, false)
		fn_Csi_SSB_ResourceSetList := func() *CSI_SSB_ResourceSetId {
			return new(CSI_SSB_ResourceSetId)
		}
		if err = tmp_Csi_SSB_ResourceSetList.Decode(r, fn_Csi_SSB_ResourceSetList); err != nil {
			return utils.WrapError("Decode Csi_SSB_ResourceSetList", err)
		}
		ie.Csi_SSB_ResourceSetList = []CSI_SSB_ResourceSetId{}
		for _, i := range tmp_Csi_SSB_ResourceSetList.Value {
			ie.Csi_SSB_ResourceSetList = append(ie.Csi_SSB_ResourceSetList, *i)
		}
	}
	return nil
}
