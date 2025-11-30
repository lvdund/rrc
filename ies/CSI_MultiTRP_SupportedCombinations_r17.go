package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CSI_MultiTRP_SupportedCombinations_r17 struct {
	MaxNumTx_Ports_r17                CSI_MultiTRP_SupportedCombinations_r17_maxNumTx_Ports_r17 `madatory`
	MaxTotalNumCMR_r17                int64                                                     `lb:2,ub:64,madatory`
	MaxTotalNumTx_PortsNZP_CSI_RS_r17 int64                                                     `lb:2,ub:256,madatory`
}

func (ie *CSI_MultiTRP_SupportedCombinations_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.MaxNumTx_Ports_r17.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumTx_Ports_r17", err)
	}
	if err = w.WriteInteger(ie.MaxTotalNumCMR_r17, &aper.Constraint{Lb: 2, Ub: 64}, false); err != nil {
		return utils.WrapError("WriteInteger MaxTotalNumCMR_r17", err)
	}
	if err = w.WriteInteger(ie.MaxTotalNumTx_PortsNZP_CSI_RS_r17, &aper.Constraint{Lb: 2, Ub: 256}, false); err != nil {
		return utils.WrapError("WriteInteger MaxTotalNumTx_PortsNZP_CSI_RS_r17", err)
	}
	return nil
}

func (ie *CSI_MultiTRP_SupportedCombinations_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.MaxNumTx_Ports_r17.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumTx_Ports_r17", err)
	}
	var tmp_int_MaxTotalNumCMR_r17 int64
	if tmp_int_MaxTotalNumCMR_r17, err = r.ReadInteger(&aper.Constraint{Lb: 2, Ub: 64}, false); err != nil {
		return utils.WrapError("ReadInteger MaxTotalNumCMR_r17", err)
	}
	ie.MaxTotalNumCMR_r17 = tmp_int_MaxTotalNumCMR_r17
	var tmp_int_MaxTotalNumTx_PortsNZP_CSI_RS_r17 int64
	if tmp_int_MaxTotalNumTx_PortsNZP_CSI_RS_r17, err = r.ReadInteger(&aper.Constraint{Lb: 2, Ub: 256}, false); err != nil {
		return utils.WrapError("ReadInteger MaxTotalNumTx_PortsNZP_CSI_RS_r17", err)
	}
	ie.MaxTotalNumTx_PortsNZP_CSI_RS_r17 = tmp_int_MaxTotalNumTx_PortsNZP_CSI_RS_r17
	return nil
}
