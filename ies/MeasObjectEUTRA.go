package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasObjectEUTRA struct {
	CarrierFreq                     ARFCN_ValueEUTRA           `madatory`
	AllowedMeasBandwidth            EUTRA_AllowedMeasBandwidth `madatory`
	CellsToRemoveListEUTRAN         *EUTRA_CellIndexList       `optional`
	CellsToAddModListEUTRAN         []EUTRA_Cell               `lb:1,ub:maxCellMeasEUTRA,optional`
	ExcludedCellsToRemoveListEUTRAN *EUTRA_CellIndexList       `optional`
	ExcludedCellsToAddModListEUTRAN []EUTRA_ExcludedCell       `lb:1,ub:maxCellMeasEUTRA,optional`
	Eutra_PresenceAntennaPort1      EUTRA_PresenceAntennaPort1 `madatory`
	Eutra_Q_OffsetRange             *EUTRA_Q_OffsetRange       `optional`
	WidebandRSRQ_Meas               bool                       `madatory`
	AssociatedMeasGap_r17           *MeasGapId_r17             `optional,ext-1`
}

func (ie *MeasObjectEUTRA) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.AssociatedMeasGap_r17 != nil
	preambleBits := []bool{hasExtensions, ie.CellsToRemoveListEUTRAN != nil, len(ie.CellsToAddModListEUTRAN) > 0, ie.ExcludedCellsToRemoveListEUTRAN != nil, len(ie.ExcludedCellsToAddModListEUTRAN) > 0, ie.Eutra_Q_OffsetRange != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.CarrierFreq.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierFreq", err)
	}
	if err = ie.AllowedMeasBandwidth.Encode(w); err != nil {
		return utils.WrapError("Encode AllowedMeasBandwidth", err)
	}
	if ie.CellsToRemoveListEUTRAN != nil {
		if err = ie.CellsToRemoveListEUTRAN.Encode(w); err != nil {
			return utils.WrapError("Encode CellsToRemoveListEUTRAN", err)
		}
	}
	if len(ie.CellsToAddModListEUTRAN) > 0 {
		tmp_CellsToAddModListEUTRAN := utils.NewSequence[*EUTRA_Cell]([]*EUTRA_Cell{}, aper.Constraint{Lb: 1, Ub: maxCellMeasEUTRA}, false)
		for _, i := range ie.CellsToAddModListEUTRAN {
			tmp_CellsToAddModListEUTRAN.Value = append(tmp_CellsToAddModListEUTRAN.Value, &i)
		}
		if err = tmp_CellsToAddModListEUTRAN.Encode(w); err != nil {
			return utils.WrapError("Encode CellsToAddModListEUTRAN", err)
		}
	}
	if ie.ExcludedCellsToRemoveListEUTRAN != nil {
		if err = ie.ExcludedCellsToRemoveListEUTRAN.Encode(w); err != nil {
			return utils.WrapError("Encode ExcludedCellsToRemoveListEUTRAN", err)
		}
	}
	if len(ie.ExcludedCellsToAddModListEUTRAN) > 0 {
		tmp_ExcludedCellsToAddModListEUTRAN := utils.NewSequence[*EUTRA_ExcludedCell]([]*EUTRA_ExcludedCell{}, aper.Constraint{Lb: 1, Ub: maxCellMeasEUTRA}, false)
		for _, i := range ie.ExcludedCellsToAddModListEUTRAN {
			tmp_ExcludedCellsToAddModListEUTRAN.Value = append(tmp_ExcludedCellsToAddModListEUTRAN.Value, &i)
		}
		if err = tmp_ExcludedCellsToAddModListEUTRAN.Encode(w); err != nil {
			return utils.WrapError("Encode ExcludedCellsToAddModListEUTRAN", err)
		}
	}
	if err = ie.Eutra_PresenceAntennaPort1.Encode(w); err != nil {
		return utils.WrapError("Encode Eutra_PresenceAntennaPort1", err)
	}
	if ie.Eutra_Q_OffsetRange != nil {
		if err = ie.Eutra_Q_OffsetRange.Encode(w); err != nil {
			return utils.WrapError("Encode Eutra_Q_OffsetRange", err)
		}
	}
	if err = w.WriteBoolean(ie.WidebandRSRQ_Meas); err != nil {
		return utils.WrapError("WriteBoolean WidebandRSRQ_Meas", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.AssociatedMeasGap_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MeasObjectEUTRA", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.AssociatedMeasGap_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode AssociatedMeasGap_r17 optional
			if ie.AssociatedMeasGap_r17 != nil {
				if err = ie.AssociatedMeasGap_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode AssociatedMeasGap_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *MeasObjectEUTRA) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var CellsToRemoveListEUTRANPresent bool
	if CellsToRemoveListEUTRANPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var CellsToAddModListEUTRANPresent bool
	if CellsToAddModListEUTRANPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ExcludedCellsToRemoveListEUTRANPresent bool
	if ExcludedCellsToRemoveListEUTRANPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ExcludedCellsToAddModListEUTRANPresent bool
	if ExcludedCellsToAddModListEUTRANPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Eutra_Q_OffsetRangePresent bool
	if Eutra_Q_OffsetRangePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.CarrierFreq.Decode(r); err != nil {
		return utils.WrapError("Decode CarrierFreq", err)
	}
	if err = ie.AllowedMeasBandwidth.Decode(r); err != nil {
		return utils.WrapError("Decode AllowedMeasBandwidth", err)
	}
	if CellsToRemoveListEUTRANPresent {
		ie.CellsToRemoveListEUTRAN = new(EUTRA_CellIndexList)
		if err = ie.CellsToRemoveListEUTRAN.Decode(r); err != nil {
			return utils.WrapError("Decode CellsToRemoveListEUTRAN", err)
		}
	}
	if CellsToAddModListEUTRANPresent {
		tmp_CellsToAddModListEUTRAN := utils.NewSequence[*EUTRA_Cell]([]*EUTRA_Cell{}, aper.Constraint{Lb: 1, Ub: maxCellMeasEUTRA}, false)
		fn_CellsToAddModListEUTRAN := func() *EUTRA_Cell {
			return new(EUTRA_Cell)
		}
		if err = tmp_CellsToAddModListEUTRAN.Decode(r, fn_CellsToAddModListEUTRAN); err != nil {
			return utils.WrapError("Decode CellsToAddModListEUTRAN", err)
		}
		ie.CellsToAddModListEUTRAN = []EUTRA_Cell{}
		for _, i := range tmp_CellsToAddModListEUTRAN.Value {
			ie.CellsToAddModListEUTRAN = append(ie.CellsToAddModListEUTRAN, *i)
		}
	}
	if ExcludedCellsToRemoveListEUTRANPresent {
		ie.ExcludedCellsToRemoveListEUTRAN = new(EUTRA_CellIndexList)
		if err = ie.ExcludedCellsToRemoveListEUTRAN.Decode(r); err != nil {
			return utils.WrapError("Decode ExcludedCellsToRemoveListEUTRAN", err)
		}
	}
	if ExcludedCellsToAddModListEUTRANPresent {
		tmp_ExcludedCellsToAddModListEUTRAN := utils.NewSequence[*EUTRA_ExcludedCell]([]*EUTRA_ExcludedCell{}, aper.Constraint{Lb: 1, Ub: maxCellMeasEUTRA}, false)
		fn_ExcludedCellsToAddModListEUTRAN := func() *EUTRA_ExcludedCell {
			return new(EUTRA_ExcludedCell)
		}
		if err = tmp_ExcludedCellsToAddModListEUTRAN.Decode(r, fn_ExcludedCellsToAddModListEUTRAN); err != nil {
			return utils.WrapError("Decode ExcludedCellsToAddModListEUTRAN", err)
		}
		ie.ExcludedCellsToAddModListEUTRAN = []EUTRA_ExcludedCell{}
		for _, i := range tmp_ExcludedCellsToAddModListEUTRAN.Value {
			ie.ExcludedCellsToAddModListEUTRAN = append(ie.ExcludedCellsToAddModListEUTRAN, *i)
		}
	}
	if err = ie.Eutra_PresenceAntennaPort1.Decode(r); err != nil {
		return utils.WrapError("Decode Eutra_PresenceAntennaPort1", err)
	}
	if Eutra_Q_OffsetRangePresent {
		ie.Eutra_Q_OffsetRange = new(EUTRA_Q_OffsetRange)
		if err = ie.Eutra_Q_OffsetRange.Decode(r); err != nil {
			return utils.WrapError("Decode Eutra_Q_OffsetRange", err)
		}
	}
	var tmp_bool_WidebandRSRQ_Meas bool
	if tmp_bool_WidebandRSRQ_Meas, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean WidebandRSRQ_Meas", err)
	}
	ie.WidebandRSRQ_Meas = tmp_bool_WidebandRSRQ_Meas

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			AssociatedMeasGap_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode AssociatedMeasGap_r17 optional
			if AssociatedMeasGap_r17Present {
				ie.AssociatedMeasGap_r17 = new(MeasGapId_r17)
				if err = ie.AssociatedMeasGap_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode AssociatedMeasGap_r17", err)
				}
			}
		}
	}
	return nil
}
