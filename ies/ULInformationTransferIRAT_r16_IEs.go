package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ULInformationTransferIRAT_r16_IEs struct {
	Ul_DCCH_MessageEUTRA_r16 *[]byte     `optional`
	LateNonCriticalExtension *[]byte     `optional`
	NonCriticalExtension     interface{} `optional`
}

func (ie *ULInformationTransferIRAT_r16_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Ul_DCCH_MessageEUTRA_r16 != nil, ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ul_DCCH_MessageEUTRA_r16 != nil {
		if err = w.WriteOctetString(*ie.Ul_DCCH_MessageEUTRA_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode Ul_DCCH_MessageEUTRA_r16", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *ULInformationTransferIRAT_r16_IEs) Decode(r *uper.UperReader) error {
	var err error
	var Ul_DCCH_MessageEUTRA_r16Present bool
	if Ul_DCCH_MessageEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Ul_DCCH_MessageEUTRA_r16Present {
		var tmp_os_Ul_DCCH_MessageEUTRA_r16 []byte
		if tmp_os_Ul_DCCH_MessageEUTRA_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode Ul_DCCH_MessageEUTRA_r16", err)
		}
		ie.Ul_DCCH_MessageEUTRA_r16 = &tmp_os_Ul_DCCH_MessageEUTRA_r16
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
