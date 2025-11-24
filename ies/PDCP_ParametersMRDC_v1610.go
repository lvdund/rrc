package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCP_ParametersMRDC_v1610 struct {
	Scg_DRB_NR_IAB_r16 *PDCP_ParametersMRDC_v1610_scg_DRB_NR_IAB_r16 `optional`
}

func (ie *PDCP_ParametersMRDC_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Scg_DRB_NR_IAB_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Scg_DRB_NR_IAB_r16 != nil {
		if err = ie.Scg_DRB_NR_IAB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Scg_DRB_NR_IAB_r16", err)
		}
	}
	return nil
}

func (ie *PDCP_ParametersMRDC_v1610) Decode(r *uper.UperReader) error {
	var err error
	var Scg_DRB_NR_IAB_r16Present bool
	if Scg_DRB_NR_IAB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Scg_DRB_NR_IAB_r16Present {
		ie.Scg_DRB_NR_IAB_r16 = new(PDCP_ParametersMRDC_v1610_scg_DRB_NR_IAB_r16)
		if err = ie.Scg_DRB_NR_IAB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Scg_DRB_NR_IAB_r16", err)
		}
	}
	return nil
}
