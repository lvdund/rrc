package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReportSFTD_NR struct {
	ReportSFTD_Meas           bool                                `madatory`
	ReportRSRP                bool                                `madatory`
	ReportSFTD_NeighMeas      *ReportSFTD_NR_reportSFTD_NeighMeas `optional,ext-1`
	Drx_SFTD_NeighMeas        *ReportSFTD_NR_drx_SFTD_NeighMeas   `optional,ext-1`
	CellsForWhichToReportSFTD []PhysCellId                        `lb:1,ub:maxCellSFTD,optional,ext-1`
}

func (ie *ReportSFTD_NR) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.ReportSFTD_NeighMeas != nil || ie.Drx_SFTD_NeighMeas != nil || len(ie.CellsForWhichToReportSFTD) > 0
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteBoolean(ie.ReportSFTD_Meas); err != nil {
		return utils.WrapError("WriteBoolean ReportSFTD_Meas", err)
	}
	if err = w.WriteBoolean(ie.ReportRSRP); err != nil {
		return utils.WrapError("WriteBoolean ReportRSRP", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.ReportSFTD_NeighMeas != nil || ie.Drx_SFTD_NeighMeas != nil || len(ie.CellsForWhichToReportSFTD) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap ReportSFTD_NR", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.ReportSFTD_NeighMeas != nil, ie.Drx_SFTD_NeighMeas != nil, len(ie.CellsForWhichToReportSFTD) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ReportSFTD_NeighMeas optional
			if ie.ReportSFTD_NeighMeas != nil {
				if err = ie.ReportSFTD_NeighMeas.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ReportSFTD_NeighMeas", err)
				}
			}
			// encode Drx_SFTD_NeighMeas optional
			if ie.Drx_SFTD_NeighMeas != nil {
				if err = ie.Drx_SFTD_NeighMeas.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Drx_SFTD_NeighMeas", err)
				}
			}
			// encode CellsForWhichToReportSFTD optional
			if len(ie.CellsForWhichToReportSFTD) > 0 {
				tmp_CellsForWhichToReportSFTD := utils.NewSequence[*PhysCellId]([]*PhysCellId{}, uper.Constraint{Lb: 1, Ub: maxCellSFTD}, false)
				for _, i := range ie.CellsForWhichToReportSFTD {
					tmp_CellsForWhichToReportSFTD.Value = append(tmp_CellsForWhichToReportSFTD.Value, &i)
				}
				if err = tmp_CellsForWhichToReportSFTD.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CellsForWhichToReportSFTD", err)
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

func (ie *ReportSFTD_NR) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_bool_ReportSFTD_Meas bool
	if tmp_bool_ReportSFTD_Meas, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean ReportSFTD_Meas", err)
	}
	ie.ReportSFTD_Meas = tmp_bool_ReportSFTD_Meas
	var tmp_bool_ReportRSRP bool
	if tmp_bool_ReportRSRP, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean ReportRSRP", err)
	}
	ie.ReportRSRP = tmp_bool_ReportRSRP

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

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			ReportSFTD_NeighMeasPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Drx_SFTD_NeighMeasPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CellsForWhichToReportSFTDPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ReportSFTD_NeighMeas optional
			if ReportSFTD_NeighMeasPresent {
				ie.ReportSFTD_NeighMeas = new(ReportSFTD_NR_reportSFTD_NeighMeas)
				if err = ie.ReportSFTD_NeighMeas.Decode(extReader); err != nil {
					return utils.WrapError("Decode ReportSFTD_NeighMeas", err)
				}
			}
			// decode Drx_SFTD_NeighMeas optional
			if Drx_SFTD_NeighMeasPresent {
				ie.Drx_SFTD_NeighMeas = new(ReportSFTD_NR_drx_SFTD_NeighMeas)
				if err = ie.Drx_SFTD_NeighMeas.Decode(extReader); err != nil {
					return utils.WrapError("Decode Drx_SFTD_NeighMeas", err)
				}
			}
			// decode CellsForWhichToReportSFTD optional
			if CellsForWhichToReportSFTDPresent {
				tmp_CellsForWhichToReportSFTD := utils.NewSequence[*PhysCellId]([]*PhysCellId{}, uper.Constraint{Lb: 1, Ub: maxCellSFTD}, false)
				fn_CellsForWhichToReportSFTD := func() *PhysCellId {
					return new(PhysCellId)
				}
				if err = tmp_CellsForWhichToReportSFTD.Decode(extReader, fn_CellsForWhichToReportSFTD); err != nil {
					return utils.WrapError("Decode CellsForWhichToReportSFTD", err)
				}
				ie.CellsForWhichToReportSFTD = []PhysCellId{}
				for _, i := range tmp_CellsForWhichToReportSFTD.Value {
					ie.CellsForWhichToReportSFTD = append(ie.CellsForWhichToReportSFTD, *i)
				}
			}
		}
	}
	return nil
}
