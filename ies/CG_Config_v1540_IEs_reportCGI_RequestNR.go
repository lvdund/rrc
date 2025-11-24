package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_Config_v1540_IEs_reportCGI_RequestNR struct {
	RequestedCellInfo *CG_Config_v1540_IEs_reportCGI_RequestNR_requestedCellInfo `optional`
}

func (ie *CG_Config_v1540_IEs_reportCGI_RequestNR) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.RequestedCellInfo != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.RequestedCellInfo != nil {
		if err = ie.RequestedCellInfo.Encode(w); err != nil {
			return utils.WrapError("Encode RequestedCellInfo", err)
		}
	}
	return nil
}

func (ie *CG_Config_v1540_IEs_reportCGI_RequestNR) Decode(r *uper.UperReader) error {
	var err error
	var RequestedCellInfoPresent bool
	if RequestedCellInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if RequestedCellInfoPresent {
		ie.RequestedCellInfo = new(CG_Config_v1540_IEs_reportCGI_RequestNR_requestedCellInfo)
		if err = ie.RequestedCellInfo.Decode(r); err != nil {
			return utils.WrapError("Decode RequestedCellInfo", err)
		}
	}
	return nil
}
