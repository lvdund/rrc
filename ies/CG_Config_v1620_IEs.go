package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CG_Config_v1620_IEs struct {
	UeAssistanceInformationSCG_r16 *[]byte              `optional`
	NonCriticalExtension           *CG_Config_v1630_IEs `optional`
}

func (ie *CG_Config_v1620_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.UeAssistanceInformationSCG_r16 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.UeAssistanceInformationSCG_r16 != nil {
		if err = w.WriteOctetString(*ie.UeAssistanceInformationSCG_r16, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode UeAssistanceInformationSCG_r16", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_Config_v1620_IEs) Decode(r *aper.AperReader) error {
	var err error
	var UeAssistanceInformationSCG_r16Present bool
	if UeAssistanceInformationSCG_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if UeAssistanceInformationSCG_r16Present {
		var tmp_os_UeAssistanceInformationSCG_r16 []byte
		if tmp_os_UeAssistanceInformationSCG_r16, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode UeAssistanceInformationSCG_r16", err)
		}
		ie.UeAssistanceInformationSCG_r16 = &tmp_os_UeAssistanceInformationSCG_r16
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(CG_Config_v1630_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
