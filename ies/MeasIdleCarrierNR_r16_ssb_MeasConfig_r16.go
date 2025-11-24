package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasIdleCarrierNR_r16_ssb_MeasConfig_r16 struct {
	NrofSS_BlocksToAverage_r16          *int64               `lb:2,ub:maxNrofSS_BlocksToAverage,optional`
	AbsThreshSS_BlocksConsolidation_r16 *ThresholdNR         `optional`
	Smtc_r16                            *SSB_MTC             `optional`
	Ssb_ToMeasure_r16                   *SSB_ToMeasure       `optional`
	DeriveSSB_IndexFromCell_r16         bool                 `madatory`
	Ss_RSSI_Measurement_r16             *SS_RSSI_Measurement `optional`
}

func (ie *MeasIdleCarrierNR_r16_ssb_MeasConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.NrofSS_BlocksToAverage_r16 != nil, ie.AbsThreshSS_BlocksConsolidation_r16 != nil, ie.Smtc_r16 != nil, ie.Ssb_ToMeasure_r16 != nil, ie.Ss_RSSI_Measurement_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.NrofSS_BlocksToAverage_r16 != nil {
		if err = w.WriteInteger(*ie.NrofSS_BlocksToAverage_r16, &uper.Constraint{Lb: 2, Ub: maxNrofSS_BlocksToAverage}, false); err != nil {
			return utils.WrapError("Encode NrofSS_BlocksToAverage_r16", err)
		}
	}
	if ie.AbsThreshSS_BlocksConsolidation_r16 != nil {
		if err = ie.AbsThreshSS_BlocksConsolidation_r16.Encode(w); err != nil {
			return utils.WrapError("Encode AbsThreshSS_BlocksConsolidation_r16", err)
		}
	}
	if ie.Smtc_r16 != nil {
		if err = ie.Smtc_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Smtc_r16", err)
		}
	}
	if ie.Ssb_ToMeasure_r16 != nil {
		if err = ie.Ssb_ToMeasure_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_ToMeasure_r16", err)
		}
	}
	if err = w.WriteBoolean(ie.DeriveSSB_IndexFromCell_r16); err != nil {
		return utils.WrapError("WriteBoolean DeriveSSB_IndexFromCell_r16", err)
	}
	if ie.Ss_RSSI_Measurement_r16 != nil {
		if err = ie.Ss_RSSI_Measurement_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ss_RSSI_Measurement_r16", err)
		}
	}
	return nil
}

func (ie *MeasIdleCarrierNR_r16_ssb_MeasConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var NrofSS_BlocksToAverage_r16Present bool
	if NrofSS_BlocksToAverage_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AbsThreshSS_BlocksConsolidation_r16Present bool
	if AbsThreshSS_BlocksConsolidation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Smtc_r16Present bool
	if Smtc_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ssb_ToMeasure_r16Present bool
	if Ssb_ToMeasure_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ss_RSSI_Measurement_r16Present bool
	if Ss_RSSI_Measurement_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if NrofSS_BlocksToAverage_r16Present {
		var tmp_int_NrofSS_BlocksToAverage_r16 int64
		if tmp_int_NrofSS_BlocksToAverage_r16, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: maxNrofSS_BlocksToAverage}, false); err != nil {
			return utils.WrapError("Decode NrofSS_BlocksToAverage_r16", err)
		}
		ie.NrofSS_BlocksToAverage_r16 = &tmp_int_NrofSS_BlocksToAverage_r16
	}
	if AbsThreshSS_BlocksConsolidation_r16Present {
		ie.AbsThreshSS_BlocksConsolidation_r16 = new(ThresholdNR)
		if err = ie.AbsThreshSS_BlocksConsolidation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode AbsThreshSS_BlocksConsolidation_r16", err)
		}
	}
	if Smtc_r16Present {
		ie.Smtc_r16 = new(SSB_MTC)
		if err = ie.Smtc_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Smtc_r16", err)
		}
	}
	if Ssb_ToMeasure_r16Present {
		ie.Ssb_ToMeasure_r16 = new(SSB_ToMeasure)
		if err = ie.Ssb_ToMeasure_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_ToMeasure_r16", err)
		}
	}
	var tmp_bool_DeriveSSB_IndexFromCell_r16 bool
	if tmp_bool_DeriveSSB_IndexFromCell_r16, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean DeriveSSB_IndexFromCell_r16", err)
	}
	ie.DeriveSSB_IndexFromCell_r16 = tmp_bool_DeriveSSB_IndexFromCell_r16
	if Ss_RSSI_Measurement_r16Present {
		ie.Ss_RSSI_Measurement_r16 = new(SS_RSSI_Measurement)
		if err = ie.Ss_RSSI_Measurement_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ss_RSSI_Measurement_r16", err)
		}
	}
	return nil
}
