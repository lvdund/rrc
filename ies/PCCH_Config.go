package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PCCH_Config struct {
	DefaultPagingCycle                          PagingCycle                        `madatory`
	NAndPagingFrameOffset                       PCCH_Config_nAndPagingFrameOffset  `lb:0,ub:1,madatory`
	Ns                                          PCCH_Config_ns                     `madatory`
	FirstPDCCH_MonitoringOccasionOfPO           []int64                            `lb:1,ub:maxPO_perPF,e_lb:0,e_ub:139,optional`
	NrofPDCCH_MonitoringOccasionPerSSB_InPO_r16 *int64                             `lb:2,ub:4,optional,ext-1`
	RanPagingInIdlePO_r17                       *PCCH_Config_ranPagingInIdlePO_r17 `optional,ext-2`
	FirstPDCCH_MonitoringOccasionOfPO_v1710     []int64                            `lb:1,ub:maxPO_perPF,e_lb:0,e_ub:35839,optional,ext-2`
}

func (ie *PCCH_Config) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.NrofPDCCH_MonitoringOccasionPerSSB_InPO_r16 != nil || ie.RanPagingInIdlePO_r17 != nil || len(ie.FirstPDCCH_MonitoringOccasionOfPO_v1710) > 0
	preambleBits := []bool{hasExtensions, len(ie.FirstPDCCH_MonitoringOccasionOfPO) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.DefaultPagingCycle.Encode(w); err != nil {
		return utils.WrapError("Encode DefaultPagingCycle", err)
	}
	if err = ie.NAndPagingFrameOffset.Encode(w); err != nil {
		return utils.WrapError("Encode NAndPagingFrameOffset", err)
	}
	if err = ie.Ns.Encode(w); err != nil {
		return utils.WrapError("Encode Ns", err)
	}
	if len(ie.FirstPDCCH_MonitoringOccasionOfPO) > 0 {
		tmp_FirstPDCCH_MonitoringOccasionOfPO := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxPO_perPF}, false)
		for _, i := range ie.FirstPDCCH_MonitoringOccasionOfPO {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: 139}, false)
			tmp_FirstPDCCH_MonitoringOccasionOfPO.Value = append(tmp_FirstPDCCH_MonitoringOccasionOfPO.Value, &tmp_ie)
		}
		if err = tmp_FirstPDCCH_MonitoringOccasionOfPO.Encode(w); err != nil {
			return utils.WrapError("Encode FirstPDCCH_MonitoringOccasionOfPO", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.NrofPDCCH_MonitoringOccasionPerSSB_InPO_r16 != nil, ie.RanPagingInIdlePO_r17 != nil || len(ie.FirstPDCCH_MonitoringOccasionOfPO_v1710) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PCCH_Config", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.NrofPDCCH_MonitoringOccasionPerSSB_InPO_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode NrofPDCCH_MonitoringOccasionPerSSB_InPO_r16 optional
			if ie.NrofPDCCH_MonitoringOccasionPerSSB_InPO_r16 != nil {
				if err = extWriter.WriteInteger(*ie.NrofPDCCH_MonitoringOccasionPerSSB_InPO_r16, &aper.Constraint{Lb: 2, Ub: 4}, false); err != nil {
					return utils.WrapError("Encode NrofPDCCH_MonitoringOccasionPerSSB_InPO_r16", err)
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
			optionals_ext_2 := []bool{ie.RanPagingInIdlePO_r17 != nil, len(ie.FirstPDCCH_MonitoringOccasionOfPO_v1710) > 0}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode RanPagingInIdlePO_r17 optional
			if ie.RanPagingInIdlePO_r17 != nil {
				if err = ie.RanPagingInIdlePO_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RanPagingInIdlePO_r17", err)
				}
			}
			// encode FirstPDCCH_MonitoringOccasionOfPO_v1710 optional
			if len(ie.FirstPDCCH_MonitoringOccasionOfPO_v1710) > 0 {
				tmp_FirstPDCCH_MonitoringOccasionOfPO_v1710 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxPO_perPF}, false)
				for _, i := range ie.FirstPDCCH_MonitoringOccasionOfPO_v1710 {
					tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: 35839}, false)
					tmp_FirstPDCCH_MonitoringOccasionOfPO_v1710.Value = append(tmp_FirstPDCCH_MonitoringOccasionOfPO_v1710.Value, &tmp_ie)
				}
				if err = tmp_FirstPDCCH_MonitoringOccasionOfPO_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FirstPDCCH_MonitoringOccasionOfPO_v1710", err)
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

func (ie *PCCH_Config) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var FirstPDCCH_MonitoringOccasionOfPOPresent bool
	if FirstPDCCH_MonitoringOccasionOfPOPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.DefaultPagingCycle.Decode(r); err != nil {
		return utils.WrapError("Decode DefaultPagingCycle", err)
	}
	if err = ie.NAndPagingFrameOffset.Decode(r); err != nil {
		return utils.WrapError("Decode NAndPagingFrameOffset", err)
	}
	if err = ie.Ns.Decode(r); err != nil {
		return utils.WrapError("Decode Ns", err)
	}
	if FirstPDCCH_MonitoringOccasionOfPOPresent {
		tmp_FirstPDCCH_MonitoringOccasionOfPO := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxPO_perPF}, false)
		fn_FirstPDCCH_MonitoringOccasionOfPO := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 139}, false)
			return &ie
		}
		if err = tmp_FirstPDCCH_MonitoringOccasionOfPO.Decode(r, fn_FirstPDCCH_MonitoringOccasionOfPO); err != nil {
			return utils.WrapError("Decode FirstPDCCH_MonitoringOccasionOfPO", err)
		}
		ie.FirstPDCCH_MonitoringOccasionOfPO = []int64{}
		for _, i := range tmp_FirstPDCCH_MonitoringOccasionOfPO.Value {
			ie.FirstPDCCH_MonitoringOccasionOfPO = append(ie.FirstPDCCH_MonitoringOccasionOfPO, int64(i.Value))
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

			NrofPDCCH_MonitoringOccasionPerSSB_InPO_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode NrofPDCCH_MonitoringOccasionPerSSB_InPO_r16 optional
			if NrofPDCCH_MonitoringOccasionPerSSB_InPO_r16Present {
				var tmp_int_NrofPDCCH_MonitoringOccasionPerSSB_InPO_r16 int64
				if tmp_int_NrofPDCCH_MonitoringOccasionPerSSB_InPO_r16, err = extReader.ReadInteger(&aper.Constraint{Lb: 2, Ub: 4}, false); err != nil {
					return utils.WrapError("Decode NrofPDCCH_MonitoringOccasionPerSSB_InPO_r16", err)
				}
				ie.NrofPDCCH_MonitoringOccasionPerSSB_InPO_r16 = &tmp_int_NrofPDCCH_MonitoringOccasionPerSSB_InPO_r16
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			RanPagingInIdlePO_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			FirstPDCCH_MonitoringOccasionOfPO_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode RanPagingInIdlePO_r17 optional
			if RanPagingInIdlePO_r17Present {
				ie.RanPagingInIdlePO_r17 = new(PCCH_Config_ranPagingInIdlePO_r17)
				if err = ie.RanPagingInIdlePO_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode RanPagingInIdlePO_r17", err)
				}
			}
			// decode FirstPDCCH_MonitoringOccasionOfPO_v1710 optional
			if FirstPDCCH_MonitoringOccasionOfPO_v1710Present {
				tmp_FirstPDCCH_MonitoringOccasionOfPO_v1710 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxPO_perPF}, false)
				fn_FirstPDCCH_MonitoringOccasionOfPO_v1710 := func() *utils.INTEGER {
					ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 35839}, false)
					return &ie
				}
				if err = tmp_FirstPDCCH_MonitoringOccasionOfPO_v1710.Decode(extReader, fn_FirstPDCCH_MonitoringOccasionOfPO_v1710); err != nil {
					return utils.WrapError("Decode FirstPDCCH_MonitoringOccasionOfPO_v1710", err)
				}
				ie.FirstPDCCH_MonitoringOccasionOfPO_v1710 = []int64{}
				for _, i := range tmp_FirstPDCCH_MonitoringOccasionOfPO_v1710.Value {
					ie.FirstPDCCH_MonitoringOccasionOfPO_v1710 = append(ie.FirstPDCCH_MonitoringOccasionOfPO_v1710, int64(i.Value))
				}
			}
		}
	}
	return nil
}
