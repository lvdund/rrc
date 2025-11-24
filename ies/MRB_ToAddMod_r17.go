package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MRB_ToAddMod_r17 struct {
	Mbs_SessionId_r17   *TMGI_r17                             `optional`
	Mrb_Identity_r17    MRB_Identity_r17                      `madatory`
	Mrb_IdentityNew_r17 *MRB_Identity_r17                     `optional`
	ReestablishPDCP_r17 *MRB_ToAddMod_r17_reestablishPDCP_r17 `optional`
	RecoverPDCP_r17     *MRB_ToAddMod_r17_recoverPDCP_r17     `optional`
	Pdcp_Config_r17     *PDCP_Config                          `optional`
}

func (ie *MRB_ToAddMod_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Mbs_SessionId_r17 != nil, ie.Mrb_IdentityNew_r17 != nil, ie.ReestablishPDCP_r17 != nil, ie.RecoverPDCP_r17 != nil, ie.Pdcp_Config_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Mbs_SessionId_r17 != nil {
		if err = ie.Mbs_SessionId_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Mbs_SessionId_r17", err)
		}
	}
	if err = ie.Mrb_Identity_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Mrb_Identity_r17", err)
	}
	if ie.Mrb_IdentityNew_r17 != nil {
		if err = ie.Mrb_IdentityNew_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Mrb_IdentityNew_r17", err)
		}
	}
	if ie.ReestablishPDCP_r17 != nil {
		if err = ie.ReestablishPDCP_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ReestablishPDCP_r17", err)
		}
	}
	if ie.RecoverPDCP_r17 != nil {
		if err = ie.RecoverPDCP_r17.Encode(w); err != nil {
			return utils.WrapError("Encode RecoverPDCP_r17", err)
		}
	}
	if ie.Pdcp_Config_r17 != nil {
		if err = ie.Pdcp_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcp_Config_r17", err)
		}
	}
	return nil
}

func (ie *MRB_ToAddMod_r17) Decode(r *uper.UperReader) error {
	var err error
	var Mbs_SessionId_r17Present bool
	if Mbs_SessionId_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mrb_IdentityNew_r17Present bool
	if Mrb_IdentityNew_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ReestablishPDCP_r17Present bool
	if ReestablishPDCP_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var RecoverPDCP_r17Present bool
	if RecoverPDCP_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcp_Config_r17Present bool
	if Pdcp_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Mbs_SessionId_r17Present {
		ie.Mbs_SessionId_r17 = new(TMGI_r17)
		if err = ie.Mbs_SessionId_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Mbs_SessionId_r17", err)
		}
	}
	if err = ie.Mrb_Identity_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Mrb_Identity_r17", err)
	}
	if Mrb_IdentityNew_r17Present {
		ie.Mrb_IdentityNew_r17 = new(MRB_Identity_r17)
		if err = ie.Mrb_IdentityNew_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Mrb_IdentityNew_r17", err)
		}
	}
	if ReestablishPDCP_r17Present {
		ie.ReestablishPDCP_r17 = new(MRB_ToAddMod_r17_reestablishPDCP_r17)
		if err = ie.ReestablishPDCP_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ReestablishPDCP_r17", err)
		}
	}
	if RecoverPDCP_r17Present {
		ie.RecoverPDCP_r17 = new(MRB_ToAddMod_r17_recoverPDCP_r17)
		if err = ie.RecoverPDCP_r17.Decode(r); err != nil {
			return utils.WrapError("Decode RecoverPDCP_r17", err)
		}
	}
	if Pdcp_Config_r17Present {
		ie.Pdcp_Config_r17 = new(PDCP_Config)
		if err = ie.Pdcp_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcp_Config_r17", err)
		}
	}
	return nil
}
