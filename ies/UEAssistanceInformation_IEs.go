package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UEAssistanceInformation_IEs struct {
	DelayBudgetReport        *DelayBudgetReport                 `optional`
	LateNonCriticalExtension *[]byte                            `optional`
	NonCriticalExtension     *UEAssistanceInformation_v1540_IEs `optional`
}

func (ie *UEAssistanceInformation_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.DelayBudgetReport != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.DelayBudgetReport != nil {
		if err = ie.DelayBudgetReport.Encode(w); err != nil {
			return utils.WrapError("Encode DelayBudgetReport", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, nil, false); err != nil {
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

func (ie *UEAssistanceInformation_IEs) Decode(r *aper.AperReader) error {
	var err error
	var DelayBudgetReportPresent bool
	if DelayBudgetReportPresent, err = r.ReadBool(); err != nil {
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
	if DelayBudgetReportPresent {
		ie.DelayBudgetReport = new(DelayBudgetReport)
		if err = ie.DelayBudgetReport.Decode(r); err != nil {
			return utils.WrapError("Decode DelayBudgetReport", err)
		}
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(UEAssistanceInformation_v1540_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
