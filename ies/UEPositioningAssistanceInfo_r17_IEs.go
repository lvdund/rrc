package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UEPositioningAssistanceInfo_r17_IEs struct {
	Ue_TxTEG_AssociationList_r17 *UE_TxTEG_AssociationList_r17          `optional`
	LateNonCriticalExtension     *[]byte                                `optional`
	NonCriticalExtension         *UEPositioningAssistanceInfo_v1720_IEs `optional`
}

func (ie *UEPositioningAssistanceInfo_r17_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ue_TxTEG_AssociationList_r17 != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ue_TxTEG_AssociationList_r17 != nil {
		if err = ie.Ue_TxTEG_AssociationList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ue_TxTEG_AssociationList_r17", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
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

func (ie *UEPositioningAssistanceInfo_r17_IEs) Decode(r *aper.AperReader) error {
	var err error
	var Ue_TxTEG_AssociationList_r17Present bool
	if Ue_TxTEG_AssociationList_r17Present, err = r.ReadBool(); err != nil {
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
	if Ue_TxTEG_AssociationList_r17Present {
		ie.Ue_TxTEG_AssociationList_r17 = new(UE_TxTEG_AssociationList_r17)
		if err = ie.Ue_TxTEG_AssociationList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ue_TxTEG_AssociationList_r17", err)
		}
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(UEPositioningAssistanceInfo_v1720_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
