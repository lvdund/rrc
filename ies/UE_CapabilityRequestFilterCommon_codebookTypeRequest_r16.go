package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_CapabilityRequestFilterCommon_codebookTypeRequest_r16 struct {
	Type1_SinglePanel_r16   *UE_CapabilityRequestFilterCommon_codebookTypeRequest_r16_type1_SinglePanel_r16   `optional`
	Type1_MultiPanel_r16    *UE_CapabilityRequestFilterCommon_codebookTypeRequest_r16_type1_MultiPanel_r16    `optional`
	Type2_r16               *UE_CapabilityRequestFilterCommon_codebookTypeRequest_r16_type2_r16               `optional`
	Type2_PortSelection_r16 *UE_CapabilityRequestFilterCommon_codebookTypeRequest_r16_type2_PortSelection_r16 `optional`
}

func (ie *UE_CapabilityRequestFilterCommon_codebookTypeRequest_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Type1_SinglePanel_r16 != nil, ie.Type1_MultiPanel_r16 != nil, ie.Type2_r16 != nil, ie.Type2_PortSelection_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Type1_SinglePanel_r16 != nil {
		if err = ie.Type1_SinglePanel_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Type1_SinglePanel_r16", err)
		}
	}
	if ie.Type1_MultiPanel_r16 != nil {
		if err = ie.Type1_MultiPanel_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Type1_MultiPanel_r16", err)
		}
	}
	if ie.Type2_r16 != nil {
		if err = ie.Type2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Type2_r16", err)
		}
	}
	if ie.Type2_PortSelection_r16 != nil {
		if err = ie.Type2_PortSelection_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Type2_PortSelection_r16", err)
		}
	}
	return nil
}

func (ie *UE_CapabilityRequestFilterCommon_codebookTypeRequest_r16) Decode(r *uper.UperReader) error {
	var err error
	var Type1_SinglePanel_r16Present bool
	if Type1_SinglePanel_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Type1_MultiPanel_r16Present bool
	if Type1_MultiPanel_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Type2_r16Present bool
	if Type2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Type2_PortSelection_r16Present bool
	if Type2_PortSelection_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Type1_SinglePanel_r16Present {
		ie.Type1_SinglePanel_r16 = new(UE_CapabilityRequestFilterCommon_codebookTypeRequest_r16_type1_SinglePanel_r16)
		if err = ie.Type1_SinglePanel_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Type1_SinglePanel_r16", err)
		}
	}
	if Type1_MultiPanel_r16Present {
		ie.Type1_MultiPanel_r16 = new(UE_CapabilityRequestFilterCommon_codebookTypeRequest_r16_type1_MultiPanel_r16)
		if err = ie.Type1_MultiPanel_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Type1_MultiPanel_r16", err)
		}
	}
	if Type2_r16Present {
		ie.Type2_r16 = new(UE_CapabilityRequestFilterCommon_codebookTypeRequest_r16_type2_r16)
		if err = ie.Type2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Type2_r16", err)
		}
	}
	if Type2_PortSelection_r16Present {
		ie.Type2_PortSelection_r16 = new(UE_CapabilityRequestFilterCommon_codebookTypeRequest_r16_type2_PortSelection_r16)
		if err = ie.Type2_PortSelection_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Type2_PortSelection_r16", err)
		}
	}
	return nil
}
