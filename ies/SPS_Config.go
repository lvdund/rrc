package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SPS_Config struct {
	Periodicity                 SPS_Config_periodicity                  `madatory`
	NrofHARQ_Processes          int64                                   `lb:1,ub:8,madatory`
	N1PUCCH_AN                  *PUCCH_ResourceId                       `optional`
	Mcs_Table                   *SPS_Config_mcs_Table                   `optional`
	Sps_ConfigIndex_r16         *SPS_ConfigIndex_r16                    `optional,ext-1`
	Harq_ProcID_Offset_r16      *int64                                  `lb:0,ub:15,optional,ext-1`
	PeriodicityExt_r16          *int64                                  `lb:1,ub:5120,optional,ext-1`
	Harq_CodebookID_r16         *int64                                  `lb:1,ub:2,optional,ext-1`
	Pdsch_AggregationFactor_r16 *SPS_Config_pdsch_AggregationFactor_r16 `optional,ext-1`
	Sps_HARQ_Deferral_r17       *int64                                  `lb:1,ub:32,optional,ext-2`
	N1PUCCH_AN_PUCCHsSCell_r17  *PUCCH_ResourceId                       `optional,ext-2`
	PeriodicityExt_r17          *int64                                  `lb:1,ub:40960,optional,ext-2`
	NrofHARQ_Processes_v1710    *int64                                  `lb:9,ub:32,optional,ext-2`
	Harq_ProcID_Offset_v1700    *int64                                  `lb:16,ub:31,optional,ext-2`
}

func (ie *SPS_Config) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.Sps_ConfigIndex_r16 != nil || ie.Harq_ProcID_Offset_r16 != nil || ie.PeriodicityExt_r16 != nil || ie.Harq_CodebookID_r16 != nil || ie.Pdsch_AggregationFactor_r16 != nil || ie.Sps_HARQ_Deferral_r17 != nil || ie.N1PUCCH_AN_PUCCHsSCell_r17 != nil || ie.PeriodicityExt_r17 != nil || ie.NrofHARQ_Processes_v1710 != nil || ie.Harq_ProcID_Offset_v1700 != nil
	preambleBits := []bool{hasExtensions, ie.N1PUCCH_AN != nil, ie.Mcs_Table != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Periodicity.Encode(w); err != nil {
		return utils.WrapError("Encode Periodicity", err)
	}
	if err = w.WriteInteger(ie.NrofHARQ_Processes, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger NrofHARQ_Processes", err)
	}
	if ie.N1PUCCH_AN != nil {
		if err = ie.N1PUCCH_AN.Encode(w); err != nil {
			return utils.WrapError("Encode N1PUCCH_AN", err)
		}
	}
	if ie.Mcs_Table != nil {
		if err = ie.Mcs_Table.Encode(w); err != nil {
			return utils.WrapError("Encode Mcs_Table", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.Sps_ConfigIndex_r16 != nil || ie.Harq_ProcID_Offset_r16 != nil || ie.PeriodicityExt_r16 != nil || ie.Harq_CodebookID_r16 != nil || ie.Pdsch_AggregationFactor_r16 != nil, ie.Sps_HARQ_Deferral_r17 != nil || ie.N1PUCCH_AN_PUCCHsSCell_r17 != nil || ie.PeriodicityExt_r17 != nil || ie.NrofHARQ_Processes_v1710 != nil || ie.Harq_ProcID_Offset_v1700 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SPS_Config", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Sps_ConfigIndex_r16 != nil, ie.Harq_ProcID_Offset_r16 != nil, ie.PeriodicityExt_r16 != nil, ie.Harq_CodebookID_r16 != nil, ie.Pdsch_AggregationFactor_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Sps_ConfigIndex_r16 optional
			if ie.Sps_ConfigIndex_r16 != nil {
				if err = ie.Sps_ConfigIndex_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sps_ConfigIndex_r16", err)
				}
			}
			// encode Harq_ProcID_Offset_r16 optional
			if ie.Harq_ProcID_Offset_r16 != nil {
				if err = extWriter.WriteInteger(*ie.Harq_ProcID_Offset_r16, &aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
					return utils.WrapError("Encode Harq_ProcID_Offset_r16", err)
				}
			}
			// encode PeriodicityExt_r16 optional
			if ie.PeriodicityExt_r16 != nil {
				if err = extWriter.WriteInteger(*ie.PeriodicityExt_r16, &aper.Constraint{Lb: 1, Ub: 5120}, false); err != nil {
					return utils.WrapError("Encode PeriodicityExt_r16", err)
				}
			}
			// encode Harq_CodebookID_r16 optional
			if ie.Harq_CodebookID_r16 != nil {
				if err = extWriter.WriteInteger(*ie.Harq_CodebookID_r16, &aper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
					return utils.WrapError("Encode Harq_CodebookID_r16", err)
				}
			}
			// encode Pdsch_AggregationFactor_r16 optional
			if ie.Pdsch_AggregationFactor_r16 != nil {
				if err = ie.Pdsch_AggregationFactor_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_AggregationFactor_r16", err)
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
			optionals_ext_2 := []bool{ie.Sps_HARQ_Deferral_r17 != nil, ie.N1PUCCH_AN_PUCCHsSCell_r17 != nil, ie.PeriodicityExt_r17 != nil, ie.NrofHARQ_Processes_v1710 != nil, ie.Harq_ProcID_Offset_v1700 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Sps_HARQ_Deferral_r17 optional
			if ie.Sps_HARQ_Deferral_r17 != nil {
				if err = extWriter.WriteInteger(*ie.Sps_HARQ_Deferral_r17, &aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
					return utils.WrapError("Encode Sps_HARQ_Deferral_r17", err)
				}
			}
			// encode N1PUCCH_AN_PUCCHsSCell_r17 optional
			if ie.N1PUCCH_AN_PUCCHsSCell_r17 != nil {
				if err = ie.N1PUCCH_AN_PUCCHsSCell_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode N1PUCCH_AN_PUCCHsSCell_r17", err)
				}
			}
			// encode PeriodicityExt_r17 optional
			if ie.PeriodicityExt_r17 != nil {
				if err = extWriter.WriteInteger(*ie.PeriodicityExt_r17, &aper.Constraint{Lb: 1, Ub: 40960}, false); err != nil {
					return utils.WrapError("Encode PeriodicityExt_r17", err)
				}
			}
			// encode NrofHARQ_Processes_v1710 optional
			if ie.NrofHARQ_Processes_v1710 != nil {
				if err = extWriter.WriteInteger(*ie.NrofHARQ_Processes_v1710, &aper.Constraint{Lb: 9, Ub: 32}, false); err != nil {
					return utils.WrapError("Encode NrofHARQ_Processes_v1710", err)
				}
			}
			// encode Harq_ProcID_Offset_v1700 optional
			if ie.Harq_ProcID_Offset_v1700 != nil {
				if err = extWriter.WriteInteger(*ie.Harq_ProcID_Offset_v1700, &aper.Constraint{Lb: 16, Ub: 31}, false); err != nil {
					return utils.WrapError("Encode Harq_ProcID_Offset_v1700", err)
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

func (ie *SPS_Config) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var N1PUCCH_ANPresent bool
	if N1PUCCH_ANPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Mcs_TablePresent bool
	if Mcs_TablePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Periodicity.Decode(r); err != nil {
		return utils.WrapError("Decode Periodicity", err)
	}
	var tmp_int_NrofHARQ_Processes int64
	if tmp_int_NrofHARQ_Processes, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger NrofHARQ_Processes", err)
	}
	ie.NrofHARQ_Processes = tmp_int_NrofHARQ_Processes
	if N1PUCCH_ANPresent {
		ie.N1PUCCH_AN = new(PUCCH_ResourceId)
		if err = ie.N1PUCCH_AN.Decode(r); err != nil {
			return utils.WrapError("Decode N1PUCCH_AN", err)
		}
	}
	if Mcs_TablePresent {
		ie.Mcs_Table = new(SPS_Config_mcs_Table)
		if err = ie.Mcs_Table.Decode(r); err != nil {
			return utils.WrapError("Decode Mcs_Table", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
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

			Sps_ConfigIndex_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Harq_ProcID_Offset_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PeriodicityExt_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Harq_CodebookID_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdsch_AggregationFactor_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Sps_ConfigIndex_r16 optional
			if Sps_ConfigIndex_r16Present {
				ie.Sps_ConfigIndex_r16 = new(SPS_ConfigIndex_r16)
				if err = ie.Sps_ConfigIndex_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sps_ConfigIndex_r16", err)
				}
			}
			// decode Harq_ProcID_Offset_r16 optional
			if Harq_ProcID_Offset_r16Present {
				var tmp_int_Harq_ProcID_Offset_r16 int64
				if tmp_int_Harq_ProcID_Offset_r16, err = extReader.ReadInteger(&aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
					return utils.WrapError("Decode Harq_ProcID_Offset_r16", err)
				}
				ie.Harq_ProcID_Offset_r16 = &tmp_int_Harq_ProcID_Offset_r16
			}
			// decode PeriodicityExt_r16 optional
			if PeriodicityExt_r16Present {
				var tmp_int_PeriodicityExt_r16 int64
				if tmp_int_PeriodicityExt_r16, err = extReader.ReadInteger(&aper.Constraint{Lb: 1, Ub: 5120}, false); err != nil {
					return utils.WrapError("Decode PeriodicityExt_r16", err)
				}
				ie.PeriodicityExt_r16 = &tmp_int_PeriodicityExt_r16
			}
			// decode Harq_CodebookID_r16 optional
			if Harq_CodebookID_r16Present {
				var tmp_int_Harq_CodebookID_r16 int64
				if tmp_int_Harq_CodebookID_r16, err = extReader.ReadInteger(&aper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
					return utils.WrapError("Decode Harq_CodebookID_r16", err)
				}
				ie.Harq_CodebookID_r16 = &tmp_int_Harq_CodebookID_r16
			}
			// decode Pdsch_AggregationFactor_r16 optional
			if Pdsch_AggregationFactor_r16Present {
				ie.Pdsch_AggregationFactor_r16 = new(SPS_Config_pdsch_AggregationFactor_r16)
				if err = ie.Pdsch_AggregationFactor_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdsch_AggregationFactor_r16", err)
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

			Sps_HARQ_Deferral_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			N1PUCCH_AN_PUCCHsSCell_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PeriodicityExt_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			NrofHARQ_Processes_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Harq_ProcID_Offset_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Sps_HARQ_Deferral_r17 optional
			if Sps_HARQ_Deferral_r17Present {
				var tmp_int_Sps_HARQ_Deferral_r17 int64
				if tmp_int_Sps_HARQ_Deferral_r17, err = extReader.ReadInteger(&aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
					return utils.WrapError("Decode Sps_HARQ_Deferral_r17", err)
				}
				ie.Sps_HARQ_Deferral_r17 = &tmp_int_Sps_HARQ_Deferral_r17
			}
			// decode N1PUCCH_AN_PUCCHsSCell_r17 optional
			if N1PUCCH_AN_PUCCHsSCell_r17Present {
				ie.N1PUCCH_AN_PUCCHsSCell_r17 = new(PUCCH_ResourceId)
				if err = ie.N1PUCCH_AN_PUCCHsSCell_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode N1PUCCH_AN_PUCCHsSCell_r17", err)
				}
			}
			// decode PeriodicityExt_r17 optional
			if PeriodicityExt_r17Present {
				var tmp_int_PeriodicityExt_r17 int64
				if tmp_int_PeriodicityExt_r17, err = extReader.ReadInteger(&aper.Constraint{Lb: 1, Ub: 40960}, false); err != nil {
					return utils.WrapError("Decode PeriodicityExt_r17", err)
				}
				ie.PeriodicityExt_r17 = &tmp_int_PeriodicityExt_r17
			}
			// decode NrofHARQ_Processes_v1710 optional
			if NrofHARQ_Processes_v1710Present {
				var tmp_int_NrofHARQ_Processes_v1710 int64
				if tmp_int_NrofHARQ_Processes_v1710, err = extReader.ReadInteger(&aper.Constraint{Lb: 9, Ub: 32}, false); err != nil {
					return utils.WrapError("Decode NrofHARQ_Processes_v1710", err)
				}
				ie.NrofHARQ_Processes_v1710 = &tmp_int_NrofHARQ_Processes_v1710
			}
			// decode Harq_ProcID_Offset_v1700 optional
			if Harq_ProcID_Offset_v1700Present {
				var tmp_int_Harq_ProcID_Offset_v1700 int64
				if tmp_int_Harq_ProcID_Offset_v1700, err = extReader.ReadInteger(&aper.Constraint{Lb: 16, Ub: 31}, false); err != nil {
					return utils.WrapError("Decode Harq_ProcID_Offset_v1700", err)
				}
				ie.Harq_ProcID_Offset_v1700 = &tmp_int_Harq_ProcID_Offset_v1700
			}
		}
	}
	return nil
}
