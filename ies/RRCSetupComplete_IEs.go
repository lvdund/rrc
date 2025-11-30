package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RRCSetupComplete_IEs struct {
	SelectedPLMN_Identity    int64                                    `lb:1,ub:maxPLMN,madatory`
	RegisteredAMF            *RegisteredAMF                           `optional`
	Guami_Type               *RRCSetupComplete_IEs_guami_Type         `optional`
	S_NSSAI_List             []S_NSSAI                                `lb:1,ub:maxNrofS_NSSAI,optional`
	DedicatedNAS_Message     DedicatedNAS_Message                     `madatory`
	Ng_5G_S_TMSI_Value       *RRCSetupComplete_IEs_ng_5G_S_TMSI_Value `lb:9,ub:9,optional`
	LateNonCriticalExtension []byte                                   `optional`
	NonCriticalExtension     *RRCSetupComplete_v1610_IEs              `optional`
}

func (ie *RRCSetupComplete_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.RegisteredAMF != nil, ie.Guami_Type != nil, len(ie.S_NSSAI_List) > 0, ie.Ng_5G_S_TMSI_Value != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.SelectedPLMN_Identity, &aper.Constraint{Lb: 1, Ub: maxPLMN}, false); err != nil {
		return utils.WrapError("WriteInteger SelectedPLMN_Identity", err)
	}
	if ie.RegisteredAMF != nil {
		if err = ie.RegisteredAMF.Encode(w); err != nil {
			return utils.WrapError("Encode RegisteredAMF", err)
		}
	}
	if ie.Guami_Type != nil {
		if err = ie.Guami_Type.Encode(w); err != nil {
			return utils.WrapError("Encode Guami_Type", err)
		}
	}
	if len(ie.S_NSSAI_List) > 0 {
		tmp_S_NSSAI_List := utils.NewSequence[*S_NSSAI]([]*S_NSSAI{}, aper.Constraint{Lb: 1, Ub: maxNrofS_NSSAI}, false)
		for _, i := range ie.S_NSSAI_List {
			tmp_S_NSSAI_List.Value = append(tmp_S_NSSAI_List.Value, &i)
		}
		if err = tmp_S_NSSAI_List.Encode(w); err != nil {
			return utils.WrapError("Encode S_NSSAI_List", err)
		}
	}
	if err = ie.DedicatedNAS_Message.Encode(w); err != nil {
		return utils.WrapError("Encode DedicatedNAS_Message", err)
	}
	if ie.Ng_5G_S_TMSI_Value != nil {
		if err = ie.Ng_5G_S_TMSI_Value.Encode(w); err != nil {
			return utils.WrapError("Encode Ng_5G_S_TMSI_Value", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(ie.LateNonCriticalExtension, nil, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCSetupComplete_IEs) Decode(r *aper.AperReader) error {
	var err error
	var RegisteredAMFPresent bool
	if RegisteredAMFPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Guami_TypePresent bool
	if Guami_TypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var S_NSSAI_ListPresent bool
	if S_NSSAI_ListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ng_5G_S_TMSI_ValuePresent bool
	if Ng_5G_S_TMSI_ValuePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_SelectedPLMN_Identity int64
	if tmp_int_SelectedPLMN_Identity, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxPLMN}, false); err != nil {
		return utils.WrapError("ReadInteger SelectedPLMN_Identity", err)
	}
	ie.SelectedPLMN_Identity = tmp_int_SelectedPLMN_Identity
	if RegisteredAMFPresent {
		ie.RegisteredAMF = new(RegisteredAMF)
		if err = ie.RegisteredAMF.Decode(r); err != nil {
			return utils.WrapError("Decode RegisteredAMF", err)
		}
	}
	if Guami_TypePresent {
		ie.Guami_Type = new(RRCSetupComplete_IEs_guami_Type)
		if err = ie.Guami_Type.Decode(r); err != nil {
			return utils.WrapError("Decode Guami_Type", err)
		}
	}
	if S_NSSAI_ListPresent {
		tmp_S_NSSAI_List := utils.NewSequence[*S_NSSAI]([]*S_NSSAI{}, aper.Constraint{Lb: 1, Ub: maxNrofS_NSSAI}, false)
		fn_S_NSSAI_List := func() *S_NSSAI {
			return new(S_NSSAI)
		}
		if err = tmp_S_NSSAI_List.Decode(r, fn_S_NSSAI_List); err != nil {
			return utils.WrapError("Decode S_NSSAI_List", err)
		}
		ie.S_NSSAI_List = []S_NSSAI{}
		for _, i := range tmp_S_NSSAI_List.Value {
			ie.S_NSSAI_List = append(ie.S_NSSAI_List, *i)
		}
	}
	if err = ie.DedicatedNAS_Message.Decode(r); err != nil {
		return utils.WrapError("Decode DedicatedNAS_Message", err)
	}
	if Ng_5G_S_TMSI_ValuePresent {
		ie.Ng_5G_S_TMSI_Value = new(RRCSetupComplete_IEs_ng_5G_S_TMSI_Value)
		if err = ie.Ng_5G_S_TMSI_Value.Decode(r); err != nil {
			return utils.WrapError("Decode Ng_5G_S_TMSI_Value", err)
		}
	}
	if LateNonCriticalExtensionPresent {
		fmt.Printf("Begin decode -> LateNonCriticalExtensionPresent\n")
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = tmp_os_LateNonCriticalExtension
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(RRCSetupComplete_v1610_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
