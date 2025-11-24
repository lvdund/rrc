package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IAB_IP_AddressAndTraffic_r16 struct {
	All_Traffic_IAB_IP_Address_r16 []IAB_IP_Address_r16 `lb:1,ub:8,optional`
	F1_C_Traffic_IP_Address_r16    []IAB_IP_Address_r16 `lb:1,ub:8,optional`
	F1_U_Traffic_IP_Address_r16    []IAB_IP_Address_r16 `lb:1,ub:8,optional`
	Non_F1_Traffic_IP_Address_r16  []IAB_IP_Address_r16 `lb:1,ub:8,optional`
}

func (ie *IAB_IP_AddressAndTraffic_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.All_Traffic_IAB_IP_Address_r16) > 0, len(ie.F1_C_Traffic_IP_Address_r16) > 0, len(ie.F1_U_Traffic_IP_Address_r16) > 0, len(ie.Non_F1_Traffic_IP_Address_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.All_Traffic_IAB_IP_Address_r16) > 0 {
		tmp_All_Traffic_IAB_IP_Address_r16 := utils.NewSequence[*IAB_IP_Address_r16]([]*IAB_IP_Address_r16{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		for _, i := range ie.All_Traffic_IAB_IP_Address_r16 {
			tmp_All_Traffic_IAB_IP_Address_r16.Value = append(tmp_All_Traffic_IAB_IP_Address_r16.Value, &i)
		}
		if err = tmp_All_Traffic_IAB_IP_Address_r16.Encode(w); err != nil {
			return utils.WrapError("Encode All_Traffic_IAB_IP_Address_r16", err)
		}
	}
	if len(ie.F1_C_Traffic_IP_Address_r16) > 0 {
		tmp_F1_C_Traffic_IP_Address_r16 := utils.NewSequence[*IAB_IP_Address_r16]([]*IAB_IP_Address_r16{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		for _, i := range ie.F1_C_Traffic_IP_Address_r16 {
			tmp_F1_C_Traffic_IP_Address_r16.Value = append(tmp_F1_C_Traffic_IP_Address_r16.Value, &i)
		}
		if err = tmp_F1_C_Traffic_IP_Address_r16.Encode(w); err != nil {
			return utils.WrapError("Encode F1_C_Traffic_IP_Address_r16", err)
		}
	}
	if len(ie.F1_U_Traffic_IP_Address_r16) > 0 {
		tmp_F1_U_Traffic_IP_Address_r16 := utils.NewSequence[*IAB_IP_Address_r16]([]*IAB_IP_Address_r16{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		for _, i := range ie.F1_U_Traffic_IP_Address_r16 {
			tmp_F1_U_Traffic_IP_Address_r16.Value = append(tmp_F1_U_Traffic_IP_Address_r16.Value, &i)
		}
		if err = tmp_F1_U_Traffic_IP_Address_r16.Encode(w); err != nil {
			return utils.WrapError("Encode F1_U_Traffic_IP_Address_r16", err)
		}
	}
	if len(ie.Non_F1_Traffic_IP_Address_r16) > 0 {
		tmp_Non_F1_Traffic_IP_Address_r16 := utils.NewSequence[*IAB_IP_Address_r16]([]*IAB_IP_Address_r16{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		for _, i := range ie.Non_F1_Traffic_IP_Address_r16 {
			tmp_Non_F1_Traffic_IP_Address_r16.Value = append(tmp_Non_F1_Traffic_IP_Address_r16.Value, &i)
		}
		if err = tmp_Non_F1_Traffic_IP_Address_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Non_F1_Traffic_IP_Address_r16", err)
		}
	}
	return nil
}

func (ie *IAB_IP_AddressAndTraffic_r16) Decode(r *uper.UperReader) error {
	var err error
	var All_Traffic_IAB_IP_Address_r16Present bool
	if All_Traffic_IAB_IP_Address_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var F1_C_Traffic_IP_Address_r16Present bool
	if F1_C_Traffic_IP_Address_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var F1_U_Traffic_IP_Address_r16Present bool
	if F1_U_Traffic_IP_Address_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Non_F1_Traffic_IP_Address_r16Present bool
	if Non_F1_Traffic_IP_Address_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if All_Traffic_IAB_IP_Address_r16Present {
		tmp_All_Traffic_IAB_IP_Address_r16 := utils.NewSequence[*IAB_IP_Address_r16]([]*IAB_IP_Address_r16{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		fn_All_Traffic_IAB_IP_Address_r16 := func() *IAB_IP_Address_r16 {
			return new(IAB_IP_Address_r16)
		}
		if err = tmp_All_Traffic_IAB_IP_Address_r16.Decode(r, fn_All_Traffic_IAB_IP_Address_r16); err != nil {
			return utils.WrapError("Decode All_Traffic_IAB_IP_Address_r16", err)
		}
		ie.All_Traffic_IAB_IP_Address_r16 = []IAB_IP_Address_r16{}
		for _, i := range tmp_All_Traffic_IAB_IP_Address_r16.Value {
			ie.All_Traffic_IAB_IP_Address_r16 = append(ie.All_Traffic_IAB_IP_Address_r16, *i)
		}
	}
	if F1_C_Traffic_IP_Address_r16Present {
		tmp_F1_C_Traffic_IP_Address_r16 := utils.NewSequence[*IAB_IP_Address_r16]([]*IAB_IP_Address_r16{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		fn_F1_C_Traffic_IP_Address_r16 := func() *IAB_IP_Address_r16 {
			return new(IAB_IP_Address_r16)
		}
		if err = tmp_F1_C_Traffic_IP_Address_r16.Decode(r, fn_F1_C_Traffic_IP_Address_r16); err != nil {
			return utils.WrapError("Decode F1_C_Traffic_IP_Address_r16", err)
		}
		ie.F1_C_Traffic_IP_Address_r16 = []IAB_IP_Address_r16{}
		for _, i := range tmp_F1_C_Traffic_IP_Address_r16.Value {
			ie.F1_C_Traffic_IP_Address_r16 = append(ie.F1_C_Traffic_IP_Address_r16, *i)
		}
	}
	if F1_U_Traffic_IP_Address_r16Present {
		tmp_F1_U_Traffic_IP_Address_r16 := utils.NewSequence[*IAB_IP_Address_r16]([]*IAB_IP_Address_r16{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		fn_F1_U_Traffic_IP_Address_r16 := func() *IAB_IP_Address_r16 {
			return new(IAB_IP_Address_r16)
		}
		if err = tmp_F1_U_Traffic_IP_Address_r16.Decode(r, fn_F1_U_Traffic_IP_Address_r16); err != nil {
			return utils.WrapError("Decode F1_U_Traffic_IP_Address_r16", err)
		}
		ie.F1_U_Traffic_IP_Address_r16 = []IAB_IP_Address_r16{}
		for _, i := range tmp_F1_U_Traffic_IP_Address_r16.Value {
			ie.F1_U_Traffic_IP_Address_r16 = append(ie.F1_U_Traffic_IP_Address_r16, *i)
		}
	}
	if Non_F1_Traffic_IP_Address_r16Present {
		tmp_Non_F1_Traffic_IP_Address_r16 := utils.NewSequence[*IAB_IP_Address_r16]([]*IAB_IP_Address_r16{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		fn_Non_F1_Traffic_IP_Address_r16 := func() *IAB_IP_Address_r16 {
			return new(IAB_IP_Address_r16)
		}
		if err = tmp_Non_F1_Traffic_IP_Address_r16.Decode(r, fn_Non_F1_Traffic_IP_Address_r16); err != nil {
			return utils.WrapError("Decode Non_F1_Traffic_IP_Address_r16", err)
		}
		ie.Non_F1_Traffic_IP_Address_r16 = []IAB_IP_Address_r16{}
		for _, i := range tmp_Non_F1_Traffic_IP_Address_r16.Value {
			ie.Non_F1_Traffic_IP_Address_r16 = append(ie.Non_F1_Traffic_IP_Address_r16, *i)
		}
	}
	return nil
}
