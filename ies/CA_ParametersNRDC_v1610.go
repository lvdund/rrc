package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNRDC_v1610 struct {
	IntraFR_NR_DC_PwrSharingMode1_r16   *CA_ParametersNRDC_v1610_intraFR_NR_DC_PwrSharingMode1_r16   `optional`
	IntraFR_NR_DC_PwrSharingMode2_r16   *CA_ParametersNRDC_v1610_intraFR_NR_DC_PwrSharingMode2_r16   `optional`
	IntraFR_NR_DC_DynamicPwrSharing_r16 *CA_ParametersNRDC_v1610_intraFR_NR_DC_DynamicPwrSharing_r16 `optional`
	AsyncNRDC_r16                       *CA_ParametersNRDC_v1610_asyncNRDC_r16                       `optional`
}

func (ie *CA_ParametersNRDC_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.IntraFR_NR_DC_PwrSharingMode1_r16 != nil, ie.IntraFR_NR_DC_PwrSharingMode2_r16 != nil, ie.IntraFR_NR_DC_DynamicPwrSharing_r16 != nil, ie.AsyncNRDC_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.IntraFR_NR_DC_PwrSharingMode1_r16 != nil {
		if err = ie.IntraFR_NR_DC_PwrSharingMode1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode IntraFR_NR_DC_PwrSharingMode1_r16", err)
		}
	}
	if ie.IntraFR_NR_DC_PwrSharingMode2_r16 != nil {
		if err = ie.IntraFR_NR_DC_PwrSharingMode2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode IntraFR_NR_DC_PwrSharingMode2_r16", err)
		}
	}
	if ie.IntraFR_NR_DC_DynamicPwrSharing_r16 != nil {
		if err = ie.IntraFR_NR_DC_DynamicPwrSharing_r16.Encode(w); err != nil {
			return utils.WrapError("Encode IntraFR_NR_DC_DynamicPwrSharing_r16", err)
		}
	}
	if ie.AsyncNRDC_r16 != nil {
		if err = ie.AsyncNRDC_r16.Encode(w); err != nil {
			return utils.WrapError("Encode AsyncNRDC_r16", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNRDC_v1610) Decode(r *uper.UperReader) error {
	var err error
	var IntraFR_NR_DC_PwrSharingMode1_r16Present bool
	if IntraFR_NR_DC_PwrSharingMode1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var IntraFR_NR_DC_PwrSharingMode2_r16Present bool
	if IntraFR_NR_DC_PwrSharingMode2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var IntraFR_NR_DC_DynamicPwrSharing_r16Present bool
	if IntraFR_NR_DC_DynamicPwrSharing_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AsyncNRDC_r16Present bool
	if AsyncNRDC_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if IntraFR_NR_DC_PwrSharingMode1_r16Present {
		ie.IntraFR_NR_DC_PwrSharingMode1_r16 = new(CA_ParametersNRDC_v1610_intraFR_NR_DC_PwrSharingMode1_r16)
		if err = ie.IntraFR_NR_DC_PwrSharingMode1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode IntraFR_NR_DC_PwrSharingMode1_r16", err)
		}
	}
	if IntraFR_NR_DC_PwrSharingMode2_r16Present {
		ie.IntraFR_NR_DC_PwrSharingMode2_r16 = new(CA_ParametersNRDC_v1610_intraFR_NR_DC_PwrSharingMode2_r16)
		if err = ie.IntraFR_NR_DC_PwrSharingMode2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode IntraFR_NR_DC_PwrSharingMode2_r16", err)
		}
	}
	if IntraFR_NR_DC_DynamicPwrSharing_r16Present {
		ie.IntraFR_NR_DC_DynamicPwrSharing_r16 = new(CA_ParametersNRDC_v1610_intraFR_NR_DC_DynamicPwrSharing_r16)
		if err = ie.IntraFR_NR_DC_DynamicPwrSharing_r16.Decode(r); err != nil {
			return utils.WrapError("Decode IntraFR_NR_DC_DynamicPwrSharing_r16", err)
		}
	}
	if AsyncNRDC_r16Present {
		ie.AsyncNRDC_r16 = new(CA_ParametersNRDC_v1610_asyncNRDC_r16)
		if err = ie.AsyncNRDC_r16.Decode(r); err != nil {
			return utils.WrapError("Decode AsyncNRDC_r16", err)
		}
	}
	return nil
}
