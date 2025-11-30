package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResults struct {
	MeasId                            MeasId                             `madatory`
	MeasResultServingMOList           MeasResultServMOList               `madatory`
	MeasResultNeighCells              *MeasResults_measResultNeighCells  `optional`
	MeasResultServFreqListEUTRA_SCG   *MeasResultServFreqListEUTRA_SCG   `optional,ext-1`
	MeasResultServFreqListNR_SCG      *MeasResultServFreqListNR_SCG      `optional,ext-1`
	MeasResultSFTD_EUTRA              *MeasResultSFTD_EUTRA              `optional,ext-1`
	MeasResultSFTD_NR                 *MeasResultCellSFTD_NR             `optional,ext-1`
	MeasResultCellListSFTD_NR         *MeasResultCellListSFTD_NR         `optional,ext-2`
	MeasResultForRSSI_r16             *MeasResultForRSSI_r16             `optional,ext-3`
	LocationInfo_r16                  *LocationInfo_r16                  `optional,ext-3`
	Ul_PDCP_DelayValueResultList_r16  *UL_PDCP_DelayValueResultList_r16  `optional,ext-3`
	MeasResultsSL_r16                 *MeasResultsSL_r16                 `optional,ext-3`
	MeasResultCLI_r16                 *MeasResultCLI_r16                 `optional,ext-3`
	MeasResultRxTxTimeDiff_r17        *MeasResultRxTxTimeDiff_r17        `optional,ext-4`
	Sl_MeasResultServingRelay_r17     *[]byte                            `optional,ext-4`
	Ul_PDCP_ExcessDelayResultList_r17 *UL_PDCP_ExcessDelayResultList_r17 `optional,ext-4`
	CoarseLocationInfo_r17            *[]byte                            `optional,ext-4`
}

func (ie *MeasResults) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.MeasResultServFreqListEUTRA_SCG != nil || ie.MeasResultServFreqListNR_SCG != nil || ie.MeasResultSFTD_EUTRA != nil || ie.MeasResultSFTD_NR != nil || ie.MeasResultCellListSFTD_NR != nil || ie.MeasResultForRSSI_r16 != nil || ie.LocationInfo_r16 != nil || ie.Ul_PDCP_DelayValueResultList_r16 != nil || ie.MeasResultsSL_r16 != nil || ie.MeasResultCLI_r16 != nil || ie.MeasResultRxTxTimeDiff_r17 != nil || ie.Sl_MeasResultServingRelay_r17 != nil || ie.Ul_PDCP_ExcessDelayResultList_r17 != nil || ie.CoarseLocationInfo_r17 != nil
	preambleBits := []bool{hasExtensions, ie.MeasResultNeighCells != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.MeasId.Encode(w); err != nil {
		return utils.WrapError("Encode MeasId", err)
	}
	if err = ie.MeasResultServingMOList.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultServingMOList", err)
	}
	if ie.MeasResultNeighCells != nil {
		if err = ie.MeasResultNeighCells.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultNeighCells", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 4 bits for 4 extension groups
		extBitmap := []bool{ie.MeasResultServFreqListEUTRA_SCG != nil || ie.MeasResultServFreqListNR_SCG != nil || ie.MeasResultSFTD_EUTRA != nil || ie.MeasResultSFTD_NR != nil, ie.MeasResultCellListSFTD_NR != nil, ie.MeasResultForRSSI_r16 != nil || ie.LocationInfo_r16 != nil || ie.Ul_PDCP_DelayValueResultList_r16 != nil || ie.MeasResultsSL_r16 != nil || ie.MeasResultCLI_r16 != nil, ie.MeasResultRxTxTimeDiff_r17 != nil || ie.Sl_MeasResultServingRelay_r17 != nil || ie.Ul_PDCP_ExcessDelayResultList_r17 != nil || ie.CoarseLocationInfo_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MeasResults", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.MeasResultServFreqListEUTRA_SCG != nil, ie.MeasResultServFreqListNR_SCG != nil, ie.MeasResultSFTD_EUTRA != nil, ie.MeasResultSFTD_NR != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode MeasResultServFreqListEUTRA_SCG optional
			if ie.MeasResultServFreqListEUTRA_SCG != nil {
				if err = ie.MeasResultServFreqListEUTRA_SCG.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MeasResultServFreqListEUTRA_SCG", err)
				}
			}
			// encode MeasResultServFreqListNR_SCG optional
			if ie.MeasResultServFreqListNR_SCG != nil {
				if err = ie.MeasResultServFreqListNR_SCG.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MeasResultServFreqListNR_SCG", err)
				}
			}
			// encode MeasResultSFTD_EUTRA optional
			if ie.MeasResultSFTD_EUTRA != nil {
				if err = ie.MeasResultSFTD_EUTRA.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MeasResultSFTD_EUTRA", err)
				}
			}
			// encode MeasResultSFTD_NR optional
			if ie.MeasResultSFTD_NR != nil {
				if err = ie.MeasResultSFTD_NR.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MeasResultSFTD_NR", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.MeasResultCellListSFTD_NR != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode MeasResultCellListSFTD_NR optional
			if ie.MeasResultCellListSFTD_NR != nil {
				if err = ie.MeasResultCellListSFTD_NR.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MeasResultCellListSFTD_NR", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.MeasResultForRSSI_r16 != nil, ie.LocationInfo_r16 != nil, ie.Ul_PDCP_DelayValueResultList_r16 != nil, ie.MeasResultsSL_r16 != nil, ie.MeasResultCLI_r16 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode MeasResultForRSSI_r16 optional
			if ie.MeasResultForRSSI_r16 != nil {
				if err = ie.MeasResultForRSSI_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MeasResultForRSSI_r16", err)
				}
			}
			// encode LocationInfo_r16 optional
			if ie.LocationInfo_r16 != nil {
				if err = ie.LocationInfo_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode LocationInfo_r16", err)
				}
			}
			// encode Ul_PDCP_DelayValueResultList_r16 optional
			if ie.Ul_PDCP_DelayValueResultList_r16 != nil {
				if err = ie.Ul_PDCP_DelayValueResultList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ul_PDCP_DelayValueResultList_r16", err)
				}
			}
			// encode MeasResultsSL_r16 optional
			if ie.MeasResultsSL_r16 != nil {
				if err = ie.MeasResultsSL_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MeasResultsSL_r16", err)
				}
			}
			// encode MeasResultCLI_r16 optional
			if ie.MeasResultCLI_r16 != nil {
				if err = ie.MeasResultCLI_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MeasResultCLI_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 4
		if extBitmap[3] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 4
			optionals_ext_4 := []bool{ie.MeasResultRxTxTimeDiff_r17 != nil, ie.Sl_MeasResultServingRelay_r17 != nil, ie.Ul_PDCP_ExcessDelayResultList_r17 != nil, ie.CoarseLocationInfo_r17 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode MeasResultRxTxTimeDiff_r17 optional
			if ie.MeasResultRxTxTimeDiff_r17 != nil {
				if err = ie.MeasResultRxTxTimeDiff_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MeasResultRxTxTimeDiff_r17", err)
				}
			}
			// encode Sl_MeasResultServingRelay_r17 optional
			if ie.Sl_MeasResultServingRelay_r17 != nil {
				if err = extWriter.WriteOctetString(*ie.Sl_MeasResultServingRelay_r17, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Encode Sl_MeasResultServingRelay_r17", err)
				}
			}
			// encode Ul_PDCP_ExcessDelayResultList_r17 optional
			if ie.Ul_PDCP_ExcessDelayResultList_r17 != nil {
				if err = ie.Ul_PDCP_ExcessDelayResultList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ul_PDCP_ExcessDelayResultList_r17", err)
				}
			}
			// encode CoarseLocationInfo_r17 optional
			if ie.CoarseLocationInfo_r17 != nil {
				if err = extWriter.WriteOctetString(*ie.CoarseLocationInfo_r17, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Encode CoarseLocationInfo_r17", err)
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

func (ie *MeasResults) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultNeighCellsPresent bool
	if MeasResultNeighCellsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.MeasId.Decode(r); err != nil {
		return utils.WrapError("Decode MeasId", err)
	}
	if err = ie.MeasResultServingMOList.Decode(r); err != nil {
		return utils.WrapError("Decode MeasResultServingMOList", err)
	}
	if MeasResultNeighCellsPresent {
		ie.MeasResultNeighCells = new(MeasResults_measResultNeighCells)
		if err = ie.MeasResultNeighCells.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultNeighCells", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 4 bits for 4 extension groups
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

			MeasResultServFreqListEUTRA_SCGPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MeasResultServFreqListNR_SCGPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MeasResultSFTD_EUTRAPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MeasResultSFTD_NRPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode MeasResultServFreqListEUTRA_SCG optional
			if MeasResultServFreqListEUTRA_SCGPresent {
				ie.MeasResultServFreqListEUTRA_SCG = new(MeasResultServFreqListEUTRA_SCG)
				if err = ie.MeasResultServFreqListEUTRA_SCG.Decode(extReader); err != nil {
					return utils.WrapError("Decode MeasResultServFreqListEUTRA_SCG", err)
				}
			}
			// decode MeasResultServFreqListNR_SCG optional
			if MeasResultServFreqListNR_SCGPresent {
				ie.MeasResultServFreqListNR_SCG = new(MeasResultServFreqListNR_SCG)
				if err = ie.MeasResultServFreqListNR_SCG.Decode(extReader); err != nil {
					return utils.WrapError("Decode MeasResultServFreqListNR_SCG", err)
				}
			}
			// decode MeasResultSFTD_EUTRA optional
			if MeasResultSFTD_EUTRAPresent {
				ie.MeasResultSFTD_EUTRA = new(MeasResultSFTD_EUTRA)
				if err = ie.MeasResultSFTD_EUTRA.Decode(extReader); err != nil {
					return utils.WrapError("Decode MeasResultSFTD_EUTRA", err)
				}
			}
			// decode MeasResultSFTD_NR optional
			if MeasResultSFTD_NRPresent {
				ie.MeasResultSFTD_NR = new(MeasResultCellSFTD_NR)
				if err = ie.MeasResultSFTD_NR.Decode(extReader); err != nil {
					return utils.WrapError("Decode MeasResultSFTD_NR", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			MeasResultCellListSFTD_NRPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode MeasResultCellListSFTD_NR optional
			if MeasResultCellListSFTD_NRPresent {
				ie.MeasResultCellListSFTD_NR = new(MeasResultCellListSFTD_NR)
				if err = ie.MeasResultCellListSFTD_NR.Decode(extReader); err != nil {
					return utils.WrapError("Decode MeasResultCellListSFTD_NR", err)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			MeasResultForRSSI_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			LocationInfo_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ul_PDCP_DelayValueResultList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MeasResultsSL_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MeasResultCLI_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode MeasResultForRSSI_r16 optional
			if MeasResultForRSSI_r16Present {
				ie.MeasResultForRSSI_r16 = new(MeasResultForRSSI_r16)
				if err = ie.MeasResultForRSSI_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MeasResultForRSSI_r16", err)
				}
			}
			// decode LocationInfo_r16 optional
			if LocationInfo_r16Present {
				ie.LocationInfo_r16 = new(LocationInfo_r16)
				if err = ie.LocationInfo_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode LocationInfo_r16", err)
				}
			}
			// decode Ul_PDCP_DelayValueResultList_r16 optional
			if Ul_PDCP_DelayValueResultList_r16Present {
				ie.Ul_PDCP_DelayValueResultList_r16 = new(UL_PDCP_DelayValueResultList_r16)
				if err = ie.Ul_PDCP_DelayValueResultList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ul_PDCP_DelayValueResultList_r16", err)
				}
			}
			// decode MeasResultsSL_r16 optional
			if MeasResultsSL_r16Present {
				ie.MeasResultsSL_r16 = new(MeasResultsSL_r16)
				if err = ie.MeasResultsSL_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MeasResultsSL_r16", err)
				}
			}
			// decode MeasResultCLI_r16 optional
			if MeasResultCLI_r16Present {
				ie.MeasResultCLI_r16 = new(MeasResultCLI_r16)
				if err = ie.MeasResultCLI_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MeasResultCLI_r16", err)
				}
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			MeasResultRxTxTimeDiff_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sl_MeasResultServingRelay_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ul_PDCP_ExcessDelayResultList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CoarseLocationInfo_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode MeasResultRxTxTimeDiff_r17 optional
			if MeasResultRxTxTimeDiff_r17Present {
				ie.MeasResultRxTxTimeDiff_r17 = new(MeasResultRxTxTimeDiff_r17)
				if err = ie.MeasResultRxTxTimeDiff_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MeasResultRxTxTimeDiff_r17", err)
				}
			}
			// decode Sl_MeasResultServingRelay_r17 optional
			if Sl_MeasResultServingRelay_r17Present {
				var tmp_os_Sl_MeasResultServingRelay_r17 []byte
				if tmp_os_Sl_MeasResultServingRelay_r17, err = extReader.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Decode Sl_MeasResultServingRelay_r17", err)
				}
				ie.Sl_MeasResultServingRelay_r17 = &tmp_os_Sl_MeasResultServingRelay_r17
			}
			// decode Ul_PDCP_ExcessDelayResultList_r17 optional
			if Ul_PDCP_ExcessDelayResultList_r17Present {
				ie.Ul_PDCP_ExcessDelayResultList_r17 = new(UL_PDCP_ExcessDelayResultList_r17)
				if err = ie.Ul_PDCP_ExcessDelayResultList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ul_PDCP_ExcessDelayResultList_r17", err)
				}
			}
			// decode CoarseLocationInfo_r17 optional
			if CoarseLocationInfo_r17Present {
				var tmp_os_CoarseLocationInfo_r17 []byte
				if tmp_os_CoarseLocationInfo_r17, err = extReader.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Decode CoarseLocationInfo_r17", err)
				}
				ie.CoarseLocationInfo_r17 = &tmp_os_CoarseLocationInfo_r17
			}
		}
	}
	return nil
}
