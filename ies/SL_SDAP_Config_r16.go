package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_SDAP_Config_r16 struct {
	Sl_SDAP_Header_r16     SL_SDAP_Config_r16_sl_SDAP_Header_r16      `madatory`
	Sl_DefaultRB_r16       bool                                       `madatory`
	Sl_MappedQoS_Flows_r16 *SL_SDAP_Config_r16_sl_MappedQoS_Flows_r16 `lb:1,ub:maxNrofSL_QFIs_r16,optional`
	Sl_CastType_r16        *SL_SDAP_Config_r16_sl_CastType_r16        `optional`
}

func (ie *SL_SDAP_Config_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_MappedQoS_Flows_r16 != nil, ie.Sl_CastType_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Sl_SDAP_Header_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_SDAP_Header_r16", err)
	}
	if err = w.WriteBoolean(ie.Sl_DefaultRB_r16); err != nil {
		return utils.WrapError("WriteBoolean Sl_DefaultRB_r16", err)
	}
	if ie.Sl_MappedQoS_Flows_r16 != nil {
		if err = ie.Sl_MappedQoS_Flows_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_MappedQoS_Flows_r16", err)
		}
	}
	if ie.Sl_CastType_r16 != nil {
		if err = ie.Sl_CastType_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_CastType_r16", err)
		}
	}
	return nil
}

func (ie *SL_SDAP_Config_r16) Decode(r *uper.UperReader) error {
	var err error
	var Sl_MappedQoS_Flows_r16Present bool
	if Sl_MappedQoS_Flows_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_CastType_r16Present bool
	if Sl_CastType_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Sl_SDAP_Header_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_SDAP_Header_r16", err)
	}
	var tmp_bool_Sl_DefaultRB_r16 bool
	if tmp_bool_Sl_DefaultRB_r16, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Sl_DefaultRB_r16", err)
	}
	ie.Sl_DefaultRB_r16 = tmp_bool_Sl_DefaultRB_r16
	if Sl_MappedQoS_Flows_r16Present {
		ie.Sl_MappedQoS_Flows_r16 = new(SL_SDAP_Config_r16_sl_MappedQoS_Flows_r16)
		if err = ie.Sl_MappedQoS_Flows_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_MappedQoS_Flows_r16", err)
		}
	}
	if Sl_CastType_r16Present {
		ie.Sl_CastType_r16 = new(SL_SDAP_Config_r16_sl_CastType_r16)
		if err = ie.Sl_CastType_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_CastType_r16", err)
		}
	}
	return nil
}
