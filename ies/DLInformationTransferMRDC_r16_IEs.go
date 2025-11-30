package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DLInformationTransferMRDC_r16_IEs struct {
	Dl_DCCH_MessageNR_r16    *[]byte     `optional`
	Dl_DCCH_MessageEUTRA_r16 *[]byte     `optional`
	LateNonCriticalExtension *[]byte     `optional`
	NonCriticalExtension     interface{} `optional`
}

func (ie *DLInformationTransferMRDC_r16_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Dl_DCCH_MessageNR_r16 != nil, ie.Dl_DCCH_MessageEUTRA_r16 != nil, ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Dl_DCCH_MessageNR_r16 != nil {
		if err = w.WriteOctetString(*ie.Dl_DCCH_MessageNR_r16, nil, false); err != nil {
			return utils.WrapError("Encode Dl_DCCH_MessageNR_r16", err)
		}
	}
	if ie.Dl_DCCH_MessageEUTRA_r16 != nil {
		if err = w.WriteOctetString(*ie.Dl_DCCH_MessageEUTRA_r16, nil, false); err != nil {
			return utils.WrapError("Encode Dl_DCCH_MessageEUTRA_r16", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, nil, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *DLInformationTransferMRDC_r16_IEs) Decode(r *aper.AperReader) error {
	var err error
	var Dl_DCCH_MessageNR_r16Present bool
	if Dl_DCCH_MessageNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dl_DCCH_MessageEUTRA_r16Present bool
	if Dl_DCCH_MessageEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Dl_DCCH_MessageNR_r16Present {
		var tmp_os_Dl_DCCH_MessageNR_r16 []byte
		if tmp_os_Dl_DCCH_MessageNR_r16, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode Dl_DCCH_MessageNR_r16", err)
		}
		ie.Dl_DCCH_MessageNR_r16 = &tmp_os_Dl_DCCH_MessageNR_r16
	}
	if Dl_DCCH_MessageEUTRA_r16Present {
		var tmp_os_Dl_DCCH_MessageEUTRA_r16 []byte
		if tmp_os_Dl_DCCH_MessageEUTRA_r16, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode Dl_DCCH_MessageEUTRA_r16", err)
		}
		ie.Dl_DCCH_MessageEUTRA_r16 = &tmp_os_Dl_DCCH_MessageEUTRA_r16
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	return nil
}
