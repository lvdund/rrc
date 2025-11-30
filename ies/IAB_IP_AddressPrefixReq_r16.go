package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type IAB_IP_AddressPrefixReq_r16 struct {
	All_Traffic_PrefixReq_r16    *IAB_IP_AddressPrefixReq_r16_all_Traffic_PrefixReq_r16    `optional`
	F1_C_Traffic_PrefixReq_r16   *IAB_IP_AddressPrefixReq_r16_f1_C_Traffic_PrefixReq_r16   `optional`
	F1_U_Traffic_PrefixReq_r16   *IAB_IP_AddressPrefixReq_r16_f1_U_Traffic_PrefixReq_r16   `optional`
	Non_F1_Traffic_PrefixReq_r16 *IAB_IP_AddressPrefixReq_r16_non_F1_Traffic_PrefixReq_r16 `optional`
}

func (ie *IAB_IP_AddressPrefixReq_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.All_Traffic_PrefixReq_r16 != nil, ie.F1_C_Traffic_PrefixReq_r16 != nil, ie.F1_U_Traffic_PrefixReq_r16 != nil, ie.Non_F1_Traffic_PrefixReq_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.All_Traffic_PrefixReq_r16 != nil {
		if err = ie.All_Traffic_PrefixReq_r16.Encode(w); err != nil {
			return utils.WrapError("Encode All_Traffic_PrefixReq_r16", err)
		}
	}
	if ie.F1_C_Traffic_PrefixReq_r16 != nil {
		if err = ie.F1_C_Traffic_PrefixReq_r16.Encode(w); err != nil {
			return utils.WrapError("Encode F1_C_Traffic_PrefixReq_r16", err)
		}
	}
	if ie.F1_U_Traffic_PrefixReq_r16 != nil {
		if err = ie.F1_U_Traffic_PrefixReq_r16.Encode(w); err != nil {
			return utils.WrapError("Encode F1_U_Traffic_PrefixReq_r16", err)
		}
	}
	if ie.Non_F1_Traffic_PrefixReq_r16 != nil {
		if err = ie.Non_F1_Traffic_PrefixReq_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Non_F1_Traffic_PrefixReq_r16", err)
		}
	}
	return nil
}

func (ie *IAB_IP_AddressPrefixReq_r16) Decode(r *aper.AperReader) error {
	var err error
	var All_Traffic_PrefixReq_r16Present bool
	if All_Traffic_PrefixReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var F1_C_Traffic_PrefixReq_r16Present bool
	if F1_C_Traffic_PrefixReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var F1_U_Traffic_PrefixReq_r16Present bool
	if F1_U_Traffic_PrefixReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Non_F1_Traffic_PrefixReq_r16Present bool
	if Non_F1_Traffic_PrefixReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if All_Traffic_PrefixReq_r16Present {
		ie.All_Traffic_PrefixReq_r16 = new(IAB_IP_AddressPrefixReq_r16_all_Traffic_PrefixReq_r16)
		if err = ie.All_Traffic_PrefixReq_r16.Decode(r); err != nil {
			return utils.WrapError("Decode All_Traffic_PrefixReq_r16", err)
		}
	}
	if F1_C_Traffic_PrefixReq_r16Present {
		ie.F1_C_Traffic_PrefixReq_r16 = new(IAB_IP_AddressPrefixReq_r16_f1_C_Traffic_PrefixReq_r16)
		if err = ie.F1_C_Traffic_PrefixReq_r16.Decode(r); err != nil {
			return utils.WrapError("Decode F1_C_Traffic_PrefixReq_r16", err)
		}
	}
	if F1_U_Traffic_PrefixReq_r16Present {
		ie.F1_U_Traffic_PrefixReq_r16 = new(IAB_IP_AddressPrefixReq_r16_f1_U_Traffic_PrefixReq_r16)
		if err = ie.F1_U_Traffic_PrefixReq_r16.Decode(r); err != nil {
			return utils.WrapError("Decode F1_U_Traffic_PrefixReq_r16", err)
		}
	}
	if Non_F1_Traffic_PrefixReq_r16Present {
		ie.Non_F1_Traffic_PrefixReq_r16 = new(IAB_IP_AddressPrefixReq_r16_non_F1_Traffic_PrefixReq_r16)
		if err = ie.Non_F1_Traffic_PrefixReq_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Non_F1_Traffic_PrefixReq_r16", err)
		}
	}
	return nil
}
