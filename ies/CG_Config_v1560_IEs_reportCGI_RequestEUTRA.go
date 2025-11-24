package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_Config_v1560_IEs_reportCGI_RequestEUTRA struct {
	RequestedCellInfoEUTRA *CG_Config_v1560_IEs_reportCGI_RequestEUTRA_requestedCellInfoEUTRA `optional`
}

func (ie *CG_Config_v1560_IEs_reportCGI_RequestEUTRA) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.RequestedCellInfoEUTRA != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.RequestedCellInfoEUTRA != nil {
		if err = ie.RequestedCellInfoEUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode RequestedCellInfoEUTRA", err)
		}
	}
	return nil
}

func (ie *CG_Config_v1560_IEs_reportCGI_RequestEUTRA) Decode(r *uper.UperReader) error {
	var err error
	var RequestedCellInfoEUTRAPresent bool
	if RequestedCellInfoEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if RequestedCellInfoEUTRAPresent {
		ie.RequestedCellInfoEUTRA = new(CG_Config_v1560_IEs_reportCGI_RequestEUTRA_requestedCellInfoEUTRA)
		if err = ie.RequestedCellInfoEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode RequestedCellInfoEUTRA", err)
		}
	}
	return nil
}
