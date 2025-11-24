package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability_v1690 struct {
	Ul_RRC_Segmentation_r16 *UE_NR_Capability_v1690_ul_RRC_Segmentation_r16 `optional`
	NonCriticalExtension    *UE_NR_Capability_v1700                         `optional`
}

func (ie *UE_NR_Capability_v1690) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Ul_RRC_Segmentation_r16 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ul_RRC_Segmentation_r16 != nil {
		if err = ie.Ul_RRC_Segmentation_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_RRC_Segmentation_r16", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability_v1690) Decode(r *uper.UperReader) error {
	var err error
	var Ul_RRC_Segmentation_r16Present bool
	if Ul_RRC_Segmentation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Ul_RRC_Segmentation_r16Present {
		ie.Ul_RRC_Segmentation_r16 = new(UE_NR_Capability_v1690_ul_RRC_Segmentation_r16)
		if err = ie.Ul_RRC_Segmentation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_RRC_Segmentation_r16", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(UE_NR_Capability_v1700)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
