package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCResumeComplete_v1610_IEs struct {
	IdleMeasAvailable_r16        *RRCResumeComplete_v1610_IEs_idleMeasAvailable_r16    `optional`
	MeasResultIdleEUTRA_r16      *MeasResultIdleEUTRA_r16                              `optional`
	MeasResultIdleNR_r16         *MeasResultIdleNR_r16                                 `optional`
	Scg_Response_r16             *RRCResumeComplete_v1610_IEs_scg_Response_r16         `optional`
	Ue_MeasurementsAvailable_r16 *UE_MeasurementsAvailable_r16                         `optional`
	MobilityHistoryAvail_r16     *RRCResumeComplete_v1610_IEs_mobilityHistoryAvail_r16 `optional`
	MobilityState_r16            *RRCResumeComplete_v1610_IEs_mobilityState_r16        `optional`
	NeedForGapsInfoNR_r16        *NeedForGapsInfoNR_r16                                `optional`
	NonCriticalExtension         *RRCResumeComplete_v1640_IEs                          `optional`
}

func (ie *RRCResumeComplete_v1610_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.IdleMeasAvailable_r16 != nil, ie.MeasResultIdleEUTRA_r16 != nil, ie.MeasResultIdleNR_r16 != nil, ie.Scg_Response_r16 != nil, ie.Ue_MeasurementsAvailable_r16 != nil, ie.MobilityHistoryAvail_r16 != nil, ie.MobilityState_r16 != nil, ie.NeedForGapsInfoNR_r16 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.IdleMeasAvailable_r16 != nil {
		if err = ie.IdleMeasAvailable_r16.Encode(w); err != nil {
			return utils.WrapError("Encode IdleMeasAvailable_r16", err)
		}
	}
	if ie.MeasResultIdleEUTRA_r16 != nil {
		if err = ie.MeasResultIdleEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultIdleEUTRA_r16", err)
		}
	}
	if ie.MeasResultIdleNR_r16 != nil {
		if err = ie.MeasResultIdleNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultIdleNR_r16", err)
		}
	}
	if ie.Scg_Response_r16 != nil {
		if err = ie.Scg_Response_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Scg_Response_r16", err)
		}
	}
	if ie.Ue_MeasurementsAvailable_r16 != nil {
		if err = ie.Ue_MeasurementsAvailable_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ue_MeasurementsAvailable_r16", err)
		}
	}
	if ie.MobilityHistoryAvail_r16 != nil {
		if err = ie.MobilityHistoryAvail_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MobilityHistoryAvail_r16", err)
		}
	}
	if ie.MobilityState_r16 != nil {
		if err = ie.MobilityState_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MobilityState_r16", err)
		}
	}
	if ie.NeedForGapsInfoNR_r16 != nil {
		if err = ie.NeedForGapsInfoNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode NeedForGapsInfoNR_r16", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCResumeComplete_v1610_IEs) Decode(r *uper.UperReader) error {
	var err error
	var IdleMeasAvailable_r16Present bool
	if IdleMeasAvailable_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultIdleEUTRA_r16Present bool
	if MeasResultIdleEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultIdleNR_r16Present bool
	if MeasResultIdleNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scg_Response_r16Present bool
	if Scg_Response_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ue_MeasurementsAvailable_r16Present bool
	if Ue_MeasurementsAvailable_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MobilityHistoryAvail_r16Present bool
	if MobilityHistoryAvail_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MobilityState_r16Present bool
	if MobilityState_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NeedForGapsInfoNR_r16Present bool
	if NeedForGapsInfoNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if IdleMeasAvailable_r16Present {
		ie.IdleMeasAvailable_r16 = new(RRCResumeComplete_v1610_IEs_idleMeasAvailable_r16)
		if err = ie.IdleMeasAvailable_r16.Decode(r); err != nil {
			return utils.WrapError("Decode IdleMeasAvailable_r16", err)
		}
	}
	if MeasResultIdleEUTRA_r16Present {
		ie.MeasResultIdleEUTRA_r16 = new(MeasResultIdleEUTRA_r16)
		if err = ie.MeasResultIdleEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultIdleEUTRA_r16", err)
		}
	}
	if MeasResultIdleNR_r16Present {
		ie.MeasResultIdleNR_r16 = new(MeasResultIdleNR_r16)
		if err = ie.MeasResultIdleNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultIdleNR_r16", err)
		}
	}
	if Scg_Response_r16Present {
		ie.Scg_Response_r16 = new(RRCResumeComplete_v1610_IEs_scg_Response_r16)
		if err = ie.Scg_Response_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Scg_Response_r16", err)
		}
	}
	if Ue_MeasurementsAvailable_r16Present {
		ie.Ue_MeasurementsAvailable_r16 = new(UE_MeasurementsAvailable_r16)
		if err = ie.Ue_MeasurementsAvailable_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ue_MeasurementsAvailable_r16", err)
		}
	}
	if MobilityHistoryAvail_r16Present {
		ie.MobilityHistoryAvail_r16 = new(RRCResumeComplete_v1610_IEs_mobilityHistoryAvail_r16)
		if err = ie.MobilityHistoryAvail_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MobilityHistoryAvail_r16", err)
		}
	}
	if MobilityState_r16Present {
		ie.MobilityState_r16 = new(RRCResumeComplete_v1610_IEs_mobilityState_r16)
		if err = ie.MobilityState_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MobilityState_r16", err)
		}
	}
	if NeedForGapsInfoNR_r16Present {
		ie.NeedForGapsInfoNR_r16 = new(NeedForGapsInfoNR_r16)
		if err = ie.NeedForGapsInfoNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode NeedForGapsInfoNR_r16", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(RRCResumeComplete_v1640_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
