package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ULInformationTransferMRDC_IEs struct {
	Ul_DCCH_MessageNR        *[]byte     `optional`
	Ul_DCCH_MessageEUTRA     *[]byte     `optional`
	LateNonCriticalExtension *[]byte     `optional`
	NonCriticalExtension     interface{} `optional`
}

func (ie *ULInformationTransferMRDC_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ul_DCCH_MessageNR != nil, ie.Ul_DCCH_MessageEUTRA != nil, ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ul_DCCH_MessageNR != nil {
		if err = w.WriteOctetString(*ie.Ul_DCCH_MessageNR, nil, false); err != nil {
			return utils.WrapError("Encode Ul_DCCH_MessageNR", err)
		}
	}
	if ie.Ul_DCCH_MessageEUTRA != nil {
		if err = w.WriteOctetString(*ie.Ul_DCCH_MessageEUTRA, nil, false); err != nil {
			return utils.WrapError("Encode Ul_DCCH_MessageEUTRA", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, nil, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *ULInformationTransferMRDC_IEs) Decode(r *aper.AperReader) error {
	var err error
	var Ul_DCCH_MessageNRPresent bool
	if Ul_DCCH_MessageNRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_DCCH_MessageEUTRAPresent bool
	if Ul_DCCH_MessageEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Ul_DCCH_MessageNRPresent {
		var tmp_os_Ul_DCCH_MessageNR []byte
		if tmp_os_Ul_DCCH_MessageNR, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode Ul_DCCH_MessageNR", err)
		}
		ie.Ul_DCCH_MessageNR = &tmp_os_Ul_DCCH_MessageNR
	}
	if Ul_DCCH_MessageEUTRAPresent {
		var tmp_os_Ul_DCCH_MessageEUTRA []byte
		if tmp_os_Ul_DCCH_MessageEUTRA, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode Ul_DCCH_MessageEUTRA", err)
		}
		ie.Ul_DCCH_MessageEUTRA = &tmp_os_Ul_DCCH_MessageEUTRA
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
