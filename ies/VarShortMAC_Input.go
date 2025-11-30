package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type VarShortMAC_Input struct {
	SourcePhysCellId   PhysCellId   `madatory`
	TargetCellIdentity CellIdentity `madatory`
	Source_c_RNTI      RNTI_Value   `madatory`
}

func (ie *VarShortMAC_Input) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.SourcePhysCellId.Encode(w); err != nil {
		return utils.WrapError("Encode SourcePhysCellId", err)
	}
	if err = ie.TargetCellIdentity.Encode(w); err != nil {
		return utils.WrapError("Encode TargetCellIdentity", err)
	}
	if err = ie.Source_c_RNTI.Encode(w); err != nil {
		return utils.WrapError("Encode Source_c_RNTI", err)
	}
	return nil
}

func (ie *VarShortMAC_Input) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.SourcePhysCellId.Decode(r); err != nil {
		return utils.WrapError("Decode SourcePhysCellId", err)
	}
	if err = ie.TargetCellIdentity.Decode(r); err != nil {
		return utils.WrapError("Decode TargetCellIdentity", err)
	}
	if err = ie.Source_c_RNTI.Decode(r); err != nil {
		return utils.WrapError("Decode Source_c_RNTI", err)
	}
	return nil
}
