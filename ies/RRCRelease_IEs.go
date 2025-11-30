package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RRCRelease_IEs struct {
	RedirectedCarrierInfo     *RedirectedCarrierInfo              `optional`
	CellReselectionPriorities *CellReselectionPriorities          `optional`
	SuspendConfig             *SuspendConfig                      `optional`
	DeprioritisationReq       *RRCRelease_IEs_deprioritisationReq `optional`
	LateNonCriticalExtension  *[]byte                             `optional`
	NonCriticalExtension      *RRCRelease_v1540_IEs               `optional`
}

func (ie *RRCRelease_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.RedirectedCarrierInfo != nil, ie.CellReselectionPriorities != nil, ie.SuspendConfig != nil, ie.DeprioritisationReq != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.RedirectedCarrierInfo != nil {
		if err = ie.RedirectedCarrierInfo.Encode(w); err != nil {
			return utils.WrapError("Encode RedirectedCarrierInfo", err)
		}
	}
	if ie.CellReselectionPriorities != nil {
		if err = ie.CellReselectionPriorities.Encode(w); err != nil {
			return utils.WrapError("Encode CellReselectionPriorities", err)
		}
	}
	if ie.SuspendConfig != nil {
		if err = ie.SuspendConfig.Encode(w); err != nil {
			return utils.WrapError("Encode SuspendConfig", err)
		}
	}
	if ie.DeprioritisationReq != nil {
		if err = ie.DeprioritisationReq.Encode(w); err != nil {
			return utils.WrapError("Encode DeprioritisationReq", err)
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

func (ie *RRCRelease_IEs) Decode(r *aper.AperReader) error {
	var err error
	var RedirectedCarrierInfoPresent bool
	if RedirectedCarrierInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var CellReselectionPrioritiesPresent bool
	if CellReselectionPrioritiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SuspendConfigPresent bool
	if SuspendConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DeprioritisationReqPresent bool
	if DeprioritisationReqPresent, err = r.ReadBool(); err != nil {
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
	if RedirectedCarrierInfoPresent {
		ie.RedirectedCarrierInfo = new(RedirectedCarrierInfo)
		if err = ie.RedirectedCarrierInfo.Decode(r); err != nil {
			return utils.WrapError("Decode RedirectedCarrierInfo", err)
		}
	}
	if CellReselectionPrioritiesPresent {
		ie.CellReselectionPriorities = new(CellReselectionPriorities)
		if err = ie.CellReselectionPriorities.Decode(r); err != nil {
			return utils.WrapError("Decode CellReselectionPriorities", err)
		}
	}
	if SuspendConfigPresent {
		ie.SuspendConfig = new(SuspendConfig)
		if err = ie.SuspendConfig.Decode(r); err != nil {
			return utils.WrapError("Decode SuspendConfig", err)
		}
	}
	if DeprioritisationReqPresent {
		ie.DeprioritisationReq = new(RRCRelease_IEs_deprioritisationReq)
		if err = ie.DeprioritisationReq.Decode(r); err != nil {
			return utils.WrapError("Decode DeprioritisationReq", err)
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
		ie.NonCriticalExtension = new(RRCRelease_v1540_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
