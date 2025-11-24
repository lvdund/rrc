package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UECapabilityEnquirySidelink_r16_IEs struct {
	FrequencyBandListFilterSidelink_r16  *FreqBandList `optional`
	Ue_CapabilityInformationSidelink_r16 *[]byte       `optional`
	LateNonCriticalExtension             *[]byte       `optional`
	NonCriticalExtension                 interface{}   `optional`
}

func (ie *UECapabilityEnquirySidelink_r16_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.FrequencyBandListFilterSidelink_r16 != nil, ie.Ue_CapabilityInformationSidelink_r16 != nil, ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.FrequencyBandListFilterSidelink_r16 != nil {
		if err = ie.FrequencyBandListFilterSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode FrequencyBandListFilterSidelink_r16", err)
		}
	}
	if ie.Ue_CapabilityInformationSidelink_r16 != nil {
		if err = w.WriteOctetString(*ie.Ue_CapabilityInformationSidelink_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode Ue_CapabilityInformationSidelink_r16", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UECapabilityEnquirySidelink_r16_IEs) Decode(r *uper.UperReader) error {
	var err error
	var FrequencyBandListFilterSidelink_r16Present bool
	if FrequencyBandListFilterSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ue_CapabilityInformationSidelink_r16Present bool
	if Ue_CapabilityInformationSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if FrequencyBandListFilterSidelink_r16Present {
		ie.FrequencyBandListFilterSidelink_r16 = new(FreqBandList)
		if err = ie.FrequencyBandListFilterSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode FrequencyBandListFilterSidelink_r16", err)
		}
	}
	if Ue_CapabilityInformationSidelink_r16Present {
		var tmp_os_Ue_CapabilityInformationSidelink_r16 []byte
		if tmp_os_Ue_CapabilityInformationSidelink_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode Ue_CapabilityInformationSidelink_r16", err)
		}
		ie.Ue_CapabilityInformationSidelink_r16 = &tmp_os_Ue_CapabilityInformationSidelink_r16
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	return nil
}
