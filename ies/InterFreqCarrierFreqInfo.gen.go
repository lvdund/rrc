// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: InterFreqCarrierFreqInfo (line 3974).

var interFreqCarrierFreqInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dl-CarrierFreq"},
		{Name: "frequencyBandList", Optional: true},
		{Name: "frequencyBandListSUL", Optional: true},
		{Name: "nrofSS-BlocksToAverage", Optional: true},
		{Name: "absThreshSS-BlocksConsolidation", Optional: true},
		{Name: "smtc", Optional: true},
		{Name: "ssbSubcarrierSpacing"},
		{Name: "ssb-ToMeasure", Optional: true},
		{Name: "deriveSSB-IndexFromCell"},
		{Name: "ss-RSSI-Measurement", Optional: true},
		{Name: "q-RxLevMin"},
		{Name: "q-RxLevMinSUL", Optional: true},
		{Name: "q-QualMin", Optional: true},
		{Name: "p-Max", Optional: true},
		{Name: "t-ReselectionNR"},
		{Name: "t-ReselectionNR-SF", Optional: true},
		{Name: "threshX-HighP"},
		{Name: "threshX-LowP"},
		{Name: "threshX-Q", Optional: true},
		{Name: "cellReselectionPriority", Optional: true},
		{Name: "cellReselectionSubPriority", Optional: true},
		{Name: "q-OffsetFreq", Optional: true},
		{Name: "interFreqNeighCellList", Optional: true},
		{Name: "interFreqExcludedCellList", Optional: true},
	},
}

var interFreqCarrierFreqInfoNrofSSBlocksToAverageConstraints = per.Constrained(2, common.MaxNrofSS_BlocksToAverage)

type InterFreqCarrierFreqInfo struct {
	Dl_CarrierFreq                  ARFCN_ValueNR
	FrequencyBandList               *MultiFrequencyBandListNR_SIB
	FrequencyBandListSUL            *MultiFrequencyBandListNR_SIB
	NrofSS_BlocksToAverage          *int64
	AbsThreshSS_BlocksConsolidation *ThresholdNR
	Smtc                            *SSB_MTC
	SsbSubcarrierSpacing            SubcarrierSpacing
	Ssb_ToMeasure                   *SSB_ToMeasure
	DeriveSSB_IndexFromCell         bool
	Ss_RSSI_Measurement             *SS_RSSI_Measurement
	Q_RxLevMin                      Q_RxLevMin
	Q_RxLevMinSUL                   *Q_RxLevMin
	Q_QualMin                       *Q_QualMin
	P_Max                           *P_Max
	T_ReselectionNR                 T_Reselection
	T_ReselectionNR_SF              *SpeedStateScaleFactors
	ThreshX_HighP                   ReselectionThreshold
	ThreshX_LowP                    ReselectionThreshold
	ThreshX_Q                       *struct {
		ThreshX_HighQ ReselectionThresholdQ
		ThreshX_LowQ  ReselectionThresholdQ
	}
	CellReselectionPriority    *CellReselectionPriority
	CellReselectionSubPriority *CellReselectionSubPriority
	Q_OffsetFreq               *Q_OffsetRange
	InterFreqNeighCellList     *InterFreqNeighCellList
	InterFreqExcludedCellList  *InterFreqExcludedCellList
}

func (ie *InterFreqCarrierFreqInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(interFreqCarrierFreqInfoConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FrequencyBandList != nil, ie.FrequencyBandListSUL != nil, ie.NrofSS_BlocksToAverage != nil, ie.AbsThreshSS_BlocksConsolidation != nil, ie.Smtc != nil, ie.Ssb_ToMeasure != nil, ie.Ss_RSSI_Measurement != nil, ie.Q_RxLevMinSUL != nil, ie.Q_QualMin != nil, ie.P_Max != nil, ie.T_ReselectionNR_SF != nil, ie.ThreshX_Q != nil, ie.CellReselectionPriority != nil, ie.CellReselectionSubPriority != nil, ie.Q_OffsetFreq != nil && ie.Q_OffsetFreq.Value != Q_OffsetRange_DB0, ie.InterFreqNeighCellList != nil, ie.InterFreqExcludedCellList != nil}); err != nil {
		return err
	}

	// 3. dl-CarrierFreq: ref
	{
		if err := ie.Dl_CarrierFreq.Encode(e); err != nil {
			return err
		}
	}

	// 4. frequencyBandList: ref
	{
		if ie.FrequencyBandList != nil {
			if err := ie.FrequencyBandList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. frequencyBandListSUL: ref
	{
		if ie.FrequencyBandListSUL != nil {
			if err := ie.FrequencyBandListSUL.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. nrofSS-BlocksToAverage: integer
	{
		if ie.NrofSS_BlocksToAverage != nil {
			if err := e.EncodeInteger(*ie.NrofSS_BlocksToAverage, interFreqCarrierFreqInfoNrofSSBlocksToAverageConstraints); err != nil {
				return err
			}
		}
	}

	// 7. absThreshSS-BlocksConsolidation: ref
	{
		if ie.AbsThreshSS_BlocksConsolidation != nil {
			if err := ie.AbsThreshSS_BlocksConsolidation.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. smtc: ref
	{
		if ie.Smtc != nil {
			if err := ie.Smtc.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. ssbSubcarrierSpacing: ref
	{
		if err := ie.SsbSubcarrierSpacing.Encode(e); err != nil {
			return err
		}
	}

	// 10. ssb-ToMeasure: ref
	{
		if ie.Ssb_ToMeasure != nil {
			if err := ie.Ssb_ToMeasure.Encode(e); err != nil {
				return err
			}
		}
	}

	// 11. deriveSSB-IndexFromCell: boolean
	{
		if err := e.EncodeBoolean(ie.DeriveSSB_IndexFromCell); err != nil {
			return err
		}
	}

	// 12. ss-RSSI-Measurement: ref
	{
		if ie.Ss_RSSI_Measurement != nil {
			if err := ie.Ss_RSSI_Measurement.Encode(e); err != nil {
				return err
			}
		}
	}

	// 13. q-RxLevMin: ref
	{
		if err := ie.Q_RxLevMin.Encode(e); err != nil {
			return err
		}
	}

	// 14. q-RxLevMinSUL: ref
	{
		if ie.Q_RxLevMinSUL != nil {
			if err := ie.Q_RxLevMinSUL.Encode(e); err != nil {
				return err
			}
		}
	}

	// 15. q-QualMin: ref
	{
		if ie.Q_QualMin != nil {
			if err := ie.Q_QualMin.Encode(e); err != nil {
				return err
			}
		}
	}

	// 16. p-Max: ref
	{
		if ie.P_Max != nil {
			if err := ie.P_Max.Encode(e); err != nil {
				return err
			}
		}
	}

	// 17. t-ReselectionNR: ref
	{
		if err := ie.T_ReselectionNR.Encode(e); err != nil {
			return err
		}
	}

	// 18. t-ReselectionNR-SF: ref
	{
		if ie.T_ReselectionNR_SF != nil {
			if err := ie.T_ReselectionNR_SF.Encode(e); err != nil {
				return err
			}
		}
	}

	// 19. threshX-HighP: ref
	{
		if err := ie.ThreshX_HighP.Encode(e); err != nil {
			return err
		}
	}

	// 20. threshX-LowP: ref
	{
		if err := ie.ThreshX_LowP.Encode(e); err != nil {
			return err
		}
	}

	// 21. threshX-Q: sequence
	{
		if ie.ThreshX_Q != nil {
			c := ie.ThreshX_Q
			if err := c.ThreshX_HighQ.Encode(e); err != nil {
				return err
			}
			if err := c.ThreshX_LowQ.Encode(e); err != nil {
				return err
			}
		}
	}

	// 22. cellReselectionPriority: ref
	{
		if ie.CellReselectionPriority != nil {
			if err := ie.CellReselectionPriority.Encode(e); err != nil {
				return err
			}
		}
	}

	// 23. cellReselectionSubPriority: ref
	{
		if ie.CellReselectionSubPriority != nil {
			if err := ie.CellReselectionSubPriority.Encode(e); err != nil {
				return err
			}
		}
	}

	// 24. q-OffsetFreq: ref
	{
		if ie.Q_OffsetFreq != nil && ie.Q_OffsetFreq.Value != Q_OffsetRange_DB0 {
			if err := ie.Q_OffsetFreq.Encode(e); err != nil {
				return err
			}
		}
	}

	// 25. interFreqNeighCellList: ref
	{
		if ie.InterFreqNeighCellList != nil {
			if err := ie.InterFreqNeighCellList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 26. interFreqExcludedCellList: ref
	{
		if ie.InterFreqExcludedCellList != nil {
			if err := ie.InterFreqExcludedCellList.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *InterFreqCarrierFreqInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(interFreqCarrierFreqInfoConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. dl-CarrierFreq: ref
	{
		if err := ie.Dl_CarrierFreq.Decode(d); err != nil {
			return err
		}
	}

	// 4. frequencyBandList: ref
	{
		if seq.IsComponentPresent(1) {
			ie.FrequencyBandList = new(MultiFrequencyBandListNR_SIB)
			if err := ie.FrequencyBandList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. frequencyBandListSUL: ref
	{
		if seq.IsComponentPresent(2) {
			ie.FrequencyBandListSUL = new(MultiFrequencyBandListNR_SIB)
			if err := ie.FrequencyBandListSUL.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. nrofSS-BlocksToAverage: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(interFreqCarrierFreqInfoNrofSSBlocksToAverageConstraints)
			if err != nil {
				return err
			}
			ie.NrofSS_BlocksToAverage = &v
		}
	}

	// 7. absThreshSS-BlocksConsolidation: ref
	{
		if seq.IsComponentPresent(4) {
			ie.AbsThreshSS_BlocksConsolidation = new(ThresholdNR)
			if err := ie.AbsThreshSS_BlocksConsolidation.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. smtc: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Smtc = new(SSB_MTC)
			if err := ie.Smtc.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. ssbSubcarrierSpacing: ref
	{
		if err := ie.SsbSubcarrierSpacing.Decode(d); err != nil {
			return err
		}
	}

	// 10. ssb-ToMeasure: ref
	{
		if seq.IsComponentPresent(7) {
			ie.Ssb_ToMeasure = new(SSB_ToMeasure)
			if err := ie.Ssb_ToMeasure.Decode(d); err != nil {
				return err
			}
		}
	}

	// 11. deriveSSB-IndexFromCell: boolean
	{
		v8, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.DeriveSSB_IndexFromCell = v8
	}

	// 12. ss-RSSI-Measurement: ref
	{
		if seq.IsComponentPresent(9) {
			ie.Ss_RSSI_Measurement = new(SS_RSSI_Measurement)
			if err := ie.Ss_RSSI_Measurement.Decode(d); err != nil {
				return err
			}
		}
	}

	// 13. q-RxLevMin: ref
	{
		if err := ie.Q_RxLevMin.Decode(d); err != nil {
			return err
		}
	}

	// 14. q-RxLevMinSUL: ref
	{
		if seq.IsComponentPresent(11) {
			ie.Q_RxLevMinSUL = new(Q_RxLevMin)
			if err := ie.Q_RxLevMinSUL.Decode(d); err != nil {
				return err
			}
		}
	}

	// 15. q-QualMin: ref
	{
		if seq.IsComponentPresent(12) {
			ie.Q_QualMin = new(Q_QualMin)
			if err := ie.Q_QualMin.Decode(d); err != nil {
				return err
			}
		}
	}

	// 16. p-Max: ref
	{
		if seq.IsComponentPresent(13) {
			ie.P_Max = new(P_Max)
			if err := ie.P_Max.Decode(d); err != nil {
				return err
			}
		}
	}

	// 17. t-ReselectionNR: ref
	{
		if err := ie.T_ReselectionNR.Decode(d); err != nil {
			return err
		}
	}

	// 18. t-ReselectionNR-SF: ref
	{
		if seq.IsComponentPresent(15) {
			ie.T_ReselectionNR_SF = new(SpeedStateScaleFactors)
			if err := ie.T_ReselectionNR_SF.Decode(d); err != nil {
				return err
			}
		}
	}

	// 19. threshX-HighP: ref
	{
		if err := ie.ThreshX_HighP.Decode(d); err != nil {
			return err
		}
	}

	// 20. threshX-LowP: ref
	{
		if err := ie.ThreshX_LowP.Decode(d); err != nil {
			return err
		}
	}

	// 21. threshX-Q: sequence
	{
		if seq.IsComponentPresent(18) {
			ie.ThreshX_Q = &struct {
				ThreshX_HighQ ReselectionThresholdQ
				ThreshX_LowQ  ReselectionThresholdQ
			}{}
			c := ie.ThreshX_Q
			{
				if err := c.ThreshX_HighQ.Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.ThreshX_LowQ.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 22. cellReselectionPriority: ref
	{
		if seq.IsComponentPresent(19) {
			ie.CellReselectionPriority = new(CellReselectionPriority)
			if err := ie.CellReselectionPriority.Decode(d); err != nil {
				return err
			}
		}
	}

	// 23. cellReselectionSubPriority: ref
	{
		if seq.IsComponentPresent(20) {
			ie.CellReselectionSubPriority = new(CellReselectionSubPriority)
			if err := ie.CellReselectionSubPriority.Decode(d); err != nil {
				return err
			}
		}
	}

	// 24. q-OffsetFreq: ref
	{
		if seq.IsComponentPresent(21) {
			ie.Q_OffsetFreq = new(Q_OffsetRange)
			if err := ie.Q_OffsetFreq.Decode(d); err != nil {
				return err
			}
		} else {
			ie.Q_OffsetFreq = &Q_OffsetRange{Value: Q_OffsetRange_DB0}
		}
	}

	// 25. interFreqNeighCellList: ref
	{
		if seq.IsComponentPresent(22) {
			ie.InterFreqNeighCellList = new(InterFreqNeighCellList)
			if err := ie.InterFreqNeighCellList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 26. interFreqExcludedCellList: ref
	{
		if seq.IsComponentPresent(23) {
			ie.InterFreqExcludedCellList = new(InterFreqExcludedCellList)
			if err := ie.InterFreqExcludedCellList.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
