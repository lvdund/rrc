package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BeamManagementSSB_CSI_RS_maxNumberSSB_CSI_RS_ResourceOneTx_Enum_n0  aper.Enumerated = 0
	BeamManagementSSB_CSI_RS_maxNumberSSB_CSI_RS_ResourceOneTx_Enum_n8  aper.Enumerated = 1
	BeamManagementSSB_CSI_RS_maxNumberSSB_CSI_RS_ResourceOneTx_Enum_n16 aper.Enumerated = 2
	BeamManagementSSB_CSI_RS_maxNumberSSB_CSI_RS_ResourceOneTx_Enum_n32 aper.Enumerated = 3
	BeamManagementSSB_CSI_RS_maxNumberSSB_CSI_RS_ResourceOneTx_Enum_n64 aper.Enumerated = 4
)

type BeamManagementSSB_CSI_RS_maxNumberSSB_CSI_RS_ResourceOneTx struct {
	Value aper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *BeamManagementSSB_CSI_RS_maxNumberSSB_CSI_RS_ResourceOneTx) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode BeamManagementSSB_CSI_RS_maxNumberSSB_CSI_RS_ResourceOneTx", err)
	}
	return nil
}

func (ie *BeamManagementSSB_CSI_RS_maxNumberSSB_CSI_RS_ResourceOneTx) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode BeamManagementSSB_CSI_RS_maxNumberSSB_CSI_RS_ResourceOneTx", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
