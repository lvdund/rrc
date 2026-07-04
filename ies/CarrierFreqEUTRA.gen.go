// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CarrierFreqEUTRA (line 4133).

var carrierFreqEUTRAConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierFreq"},
		{Name: "eutra-multiBandInfoList", Optional: true},
		{Name: "eutra-FreqNeighCellList", Optional: true},
		{Name: "eutra-ExcludedCellList", Optional: true},
		{Name: "allowedMeasBandwidth"},
		{Name: "presenceAntennaPort1"},
		{Name: "cellReselectionPriority", Optional: true},
		{Name: "cellReselectionSubPriority", Optional: true},
		{Name: "threshX-High"},
		{Name: "threshX-Low"},
		{Name: "q-RxLevMin"},
		{Name: "q-QualMin"},
		{Name: "p-MaxEUTRA"},
		{Name: "threshX-Q", Optional: true},
	},
}

var carrierFreqEUTRAQRxLevMinConstraints = per.Constrained(-70, -22)

var carrierFreqEUTRAQQualMinConstraints = per.Constrained(-34, -3)

var carrierFreqEUTRAPMaxEUTRAConstraints = per.Constrained(-30, 33)

type CarrierFreqEUTRA struct {
	CarrierFreq                ARFCN_ValueEUTRA
	Eutra_MultiBandInfoList    *EUTRA_MultiBandInfoList
	Eutra_FreqNeighCellList    *EUTRA_FreqNeighCellList
	Eutra_ExcludedCellList     *EUTRA_FreqExcludedCellList
	AllowedMeasBandwidth       EUTRA_AllowedMeasBandwidth
	PresenceAntennaPort1       EUTRA_PresenceAntennaPort1
	CellReselectionPriority    *CellReselectionPriority
	CellReselectionSubPriority *CellReselectionSubPriority
	ThreshX_High               ReselectionThreshold
	ThreshX_Low                ReselectionThreshold
	Q_RxLevMin                 int64
	Q_QualMin                  int64
	P_MaxEUTRA                 int64
	ThreshX_Q                  *struct {
		ThreshX_HighQ ReselectionThresholdQ
		ThreshX_LowQ  ReselectionThresholdQ
	}
}

func (ie *CarrierFreqEUTRA) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(carrierFreqEUTRAConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Eutra_MultiBandInfoList != nil, ie.Eutra_FreqNeighCellList != nil, ie.Eutra_ExcludedCellList != nil, ie.CellReselectionPriority != nil, ie.CellReselectionSubPriority != nil, ie.ThreshX_Q != nil}); err != nil {
		return err
	}

	// 2. carrierFreq: ref
	{
		if err := ie.CarrierFreq.Encode(e); err != nil {
			return err
		}
	}

	// 3. eutra-multiBandInfoList: ref
	{
		if ie.Eutra_MultiBandInfoList != nil {
			if err := ie.Eutra_MultiBandInfoList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. eutra-FreqNeighCellList: ref
	{
		if ie.Eutra_FreqNeighCellList != nil {
			if err := ie.Eutra_FreqNeighCellList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. eutra-ExcludedCellList: ref
	{
		if ie.Eutra_ExcludedCellList != nil {
			if err := ie.Eutra_ExcludedCellList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. allowedMeasBandwidth: ref
	{
		if err := ie.AllowedMeasBandwidth.Encode(e); err != nil {
			return err
		}
	}

	// 7. presenceAntennaPort1: ref
	{
		if err := ie.PresenceAntennaPort1.Encode(e); err != nil {
			return err
		}
	}

	// 8. cellReselectionPriority: ref
	{
		if ie.CellReselectionPriority != nil {
			if err := ie.CellReselectionPriority.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. cellReselectionSubPriority: ref
	{
		if ie.CellReselectionSubPriority != nil {
			if err := ie.CellReselectionSubPriority.Encode(e); err != nil {
				return err
			}
		}
	}

	// 10. threshX-High: ref
	{
		if err := ie.ThreshX_High.Encode(e); err != nil {
			return err
		}
	}

	// 11. threshX-Low: ref
	{
		if err := ie.ThreshX_Low.Encode(e); err != nil {
			return err
		}
	}

	// 12. q-RxLevMin: integer
	{
		if err := e.EncodeInteger(ie.Q_RxLevMin, carrierFreqEUTRAQRxLevMinConstraints); err != nil {
			return err
		}
	}

	// 13. q-QualMin: integer
	{
		if err := e.EncodeInteger(ie.Q_QualMin, carrierFreqEUTRAQQualMinConstraints); err != nil {
			return err
		}
	}

	// 14. p-MaxEUTRA: integer
	{
		if err := e.EncodeInteger(ie.P_MaxEUTRA, carrierFreqEUTRAPMaxEUTRAConstraints); err != nil {
			return err
		}
	}

	// 15. threshX-Q: sequence
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

	return nil
}

func (ie *CarrierFreqEUTRA) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(carrierFreqEUTRAConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. carrierFreq: ref
	{
		if err := ie.CarrierFreq.Decode(d); err != nil {
			return err
		}
	}

	// 3. eutra-multiBandInfoList: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Eutra_MultiBandInfoList = new(EUTRA_MultiBandInfoList)
			if err := ie.Eutra_MultiBandInfoList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. eutra-FreqNeighCellList: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Eutra_FreqNeighCellList = new(EUTRA_FreqNeighCellList)
			if err := ie.Eutra_FreqNeighCellList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. eutra-ExcludedCellList: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Eutra_ExcludedCellList = new(EUTRA_FreqExcludedCellList)
			if err := ie.Eutra_ExcludedCellList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. allowedMeasBandwidth: ref
	{
		if err := ie.AllowedMeasBandwidth.Decode(d); err != nil {
			return err
		}
	}

	// 7. presenceAntennaPort1: ref
	{
		if err := ie.PresenceAntennaPort1.Decode(d); err != nil {
			return err
		}
	}

	// 8. cellReselectionPriority: ref
	{
		if seq.IsComponentPresent(6) {
			ie.CellReselectionPriority = new(CellReselectionPriority)
			if err := ie.CellReselectionPriority.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. cellReselectionSubPriority: ref
	{
		if seq.IsComponentPresent(7) {
			ie.CellReselectionSubPriority = new(CellReselectionSubPriority)
			if err := ie.CellReselectionSubPriority.Decode(d); err != nil {
				return err
			}
		}
	}

	// 10. threshX-High: ref
	{
		if err := ie.ThreshX_High.Decode(d); err != nil {
			return err
		}
	}

	// 11. threshX-Low: ref
	{
		if err := ie.ThreshX_Low.Decode(d); err != nil {
			return err
		}
	}

	// 12. q-RxLevMin: integer
	{
		v10, err := d.DecodeInteger(carrierFreqEUTRAQRxLevMinConstraints)
		if err != nil {
			return err
		}
		ie.Q_RxLevMin = v10
	}

	// 13. q-QualMin: integer
	{
		v11, err := d.DecodeInteger(carrierFreqEUTRAQQualMinConstraints)
		if err != nil {
			return err
		}
		ie.Q_QualMin = v11
	}

	// 14. p-MaxEUTRA: integer
	{
		v12, err := d.DecodeInteger(carrierFreqEUTRAPMaxEUTRAConstraints)
		if err != nil {
			return err
		}
		ie.P_MaxEUTRA = v12
	}

	// 15. threshX-Q: sequence
	{
		if seq.IsComponentPresent(13) {
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

	return nil
}
