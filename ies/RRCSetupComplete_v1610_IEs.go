package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCSetupComplete_v1610_IEs struct {
	Iab_NodeIndication_r16       *RRCSetupComplete_v1610_IEs_iab_NodeIndication_r16   `optional`
	IdleMeasAvailable_r16        *RRCSetupComplete_v1610_IEs_idleMeasAvailable_r16    `optional`
	Ue_MeasurementsAvailable_r16 *UE_MeasurementsAvailable_r16                        `optional`
	MobilityHistoryAvail_r16     *RRCSetupComplete_v1610_IEs_mobilityHistoryAvail_r16 `optional`
	MobilityState_r16            *RRCSetupComplete_v1610_IEs_mobilityState_r16        `optional`
	NonCriticalExtension         *RRCSetupComplete_v1690_IEs                          `optional`
}

func (ie *RRCSetupComplete_v1610_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Iab_NodeIndication_r16 != nil, ie.IdleMeasAvailable_r16 != nil, ie.Ue_MeasurementsAvailable_r16 != nil, ie.MobilityHistoryAvail_r16 != nil, ie.MobilityState_r16 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Iab_NodeIndication_r16 != nil {
		if err = ie.Iab_NodeIndication_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Iab_NodeIndication_r16", err)
		}
	}
	if ie.IdleMeasAvailable_r16 != nil {
		if err = ie.IdleMeasAvailable_r16.Encode(w); err != nil {
			return utils.WrapError("Encode IdleMeasAvailable_r16", err)
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
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCSetupComplete_v1610_IEs) Decode(r *uper.UperReader) error {
	var err error
	var Iab_NodeIndication_r16Present bool
	if Iab_NodeIndication_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var IdleMeasAvailable_r16Present bool
	if IdleMeasAvailable_r16Present, err = r.ReadBool(); err != nil {
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
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Iab_NodeIndication_r16Present {
		ie.Iab_NodeIndication_r16 = new(RRCSetupComplete_v1610_IEs_iab_NodeIndication_r16)
		if err = ie.Iab_NodeIndication_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Iab_NodeIndication_r16", err)
		}
	}
	if IdleMeasAvailable_r16Present {
		ie.IdleMeasAvailable_r16 = new(RRCSetupComplete_v1610_IEs_idleMeasAvailable_r16)
		if err = ie.IdleMeasAvailable_r16.Decode(r); err != nil {
			return utils.WrapError("Decode IdleMeasAvailable_r16", err)
		}
	}
	if Ue_MeasurementsAvailable_r16Present {
		ie.Ue_MeasurementsAvailable_r16 = new(UE_MeasurementsAvailable_r16)
		if err = ie.Ue_MeasurementsAvailable_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ue_MeasurementsAvailable_r16", err)
		}
	}
	if MobilityHistoryAvail_r16Present {
		ie.MobilityHistoryAvail_r16 = new(RRCSetupComplete_v1610_IEs_mobilityHistoryAvail_r16)
		if err = ie.MobilityHistoryAvail_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MobilityHistoryAvail_r16", err)
		}
	}
	if MobilityState_r16Present {
		ie.MobilityState_r16 = new(RRCSetupComplete_v1610_IEs_mobilityState_r16)
		if err = ie.MobilityState_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MobilityState_r16", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(RRCSetupComplete_v1690_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
