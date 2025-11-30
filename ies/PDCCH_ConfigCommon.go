package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PDCCH_ConfigCommon struct {
	ControlResourceSetZero                  *ControlResourceSetZero                          `optional`
	CommonControlResourceSet                *ControlResourceSet                              `optional`
	SearchSpaceZero                         *SearchSpaceZero                                 `optional`
	CommonSearchSpaceList                   []SearchSpace                                    `lb:1,ub:4,optional`
	SearchSpaceSIB1                         *SearchSpaceId                                   `optional`
	SearchSpaceOtherSystemInformation       *SearchSpaceId                                   `optional`
	PagingSearchSpace                       *SearchSpaceId                                   `optional`
	Ra_SearchSpace                          *SearchSpaceId                                   `optional`
	FirstPDCCH_MonitoringOccasionOfPO       []int64                                          `lb:1,ub:maxPO_perPF,e_lb:0,e_ub:139,optional,ext-1`
	CommonSearchSpaceListExt_r16            []SearchSpaceExt_r16                             `lb:1,ub:4,optional,ext-2`
	Sdt_SearchSpace_r17                     *PDCCH_ConfigCommon_sdt_SearchSpace_r17          `optional,ext-3`
	SearchSpaceMCCH_r17                     *SearchSpaceId                                   `optional,ext-3`
	SearchSpaceMTCH_r17                     *SearchSpaceId                                   `optional,ext-3`
	CommonSearchSpaceListExt2_r17           []SearchSpaceExt_v1700                           `lb:1,ub:4,optional,ext-3`
	FirstPDCCH_MonitoringOccasionOfPO_v1710 []int64                                          `lb:1,ub:maxPO_perPF,e_lb:0,e_ub:35839,optional,ext-3`
	Pei_ConfigBWP_r17                       []int64                                          `lb:1,ub:maxPEI_perPF_r17,e_lb:0,e_ub:139,optional,ext-3`
	FollowUnifiedTCI_State_v1720            *PDCCH_ConfigCommon_followUnifiedTCI_State_v1720 `optional,ext-4`
}

func (ie *PDCCH_ConfigCommon) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := len(ie.FirstPDCCH_MonitoringOccasionOfPO) > 0 || len(ie.CommonSearchSpaceListExt_r16) > 0 || ie.Sdt_SearchSpace_r17 != nil || ie.SearchSpaceMCCH_r17 != nil || ie.SearchSpaceMTCH_r17 != nil || len(ie.CommonSearchSpaceListExt2_r17) > 0 || len(ie.FirstPDCCH_MonitoringOccasionOfPO_v1710) > 0 || len(ie.Pei_ConfigBWP_r17) > 0 || ie.FollowUnifiedTCI_State_v1720 != nil
	preambleBits := []bool{hasExtensions, ie.ControlResourceSetZero != nil, ie.CommonControlResourceSet != nil, ie.SearchSpaceZero != nil, len(ie.CommonSearchSpaceList) > 0, ie.SearchSpaceSIB1 != nil, ie.SearchSpaceOtherSystemInformation != nil, ie.PagingSearchSpace != nil, ie.Ra_SearchSpace != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ControlResourceSetZero != nil {
		if err = ie.ControlResourceSetZero.Encode(w); err != nil {
			return utils.WrapError("Encode ControlResourceSetZero", err)
		}
	}
	if ie.CommonControlResourceSet != nil {
		if err = ie.CommonControlResourceSet.Encode(w); err != nil {
			return utils.WrapError("Encode CommonControlResourceSet", err)
		}
	}
	if ie.SearchSpaceZero != nil {
		if err = ie.SearchSpaceZero.Encode(w); err != nil {
			return utils.WrapError("Encode SearchSpaceZero", err)
		}
	}
	if len(ie.CommonSearchSpaceList) > 0 {
		tmp_CommonSearchSpaceList := utils.NewSequence[*SearchSpace]([]*SearchSpace{}, aper.Constraint{Lb: 1, Ub: 4}, false)
		for _, i := range ie.CommonSearchSpaceList {
			tmp_CommonSearchSpaceList.Value = append(tmp_CommonSearchSpaceList.Value, &i)
		}
		if err = tmp_CommonSearchSpaceList.Encode(w); err != nil {
			return utils.WrapError("Encode CommonSearchSpaceList", err)
		}
	}
	if ie.SearchSpaceSIB1 != nil {
		if err = ie.SearchSpaceSIB1.Encode(w); err != nil {
			return utils.WrapError("Encode SearchSpaceSIB1", err)
		}
	}
	if ie.SearchSpaceOtherSystemInformation != nil {
		if err = ie.SearchSpaceOtherSystemInformation.Encode(w); err != nil {
			return utils.WrapError("Encode SearchSpaceOtherSystemInformation", err)
		}
	}
	if ie.PagingSearchSpace != nil {
		if err = ie.PagingSearchSpace.Encode(w); err != nil {
			return utils.WrapError("Encode PagingSearchSpace", err)
		}
	}
	if ie.Ra_SearchSpace != nil {
		if err = ie.Ra_SearchSpace.Encode(w); err != nil {
			return utils.WrapError("Encode Ra_SearchSpace", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 4 bits for 4 extension groups
		extBitmap := []bool{len(ie.FirstPDCCH_MonitoringOccasionOfPO) > 0, len(ie.CommonSearchSpaceListExt_r16) > 0, ie.Sdt_SearchSpace_r17 != nil || ie.SearchSpaceMCCH_r17 != nil || ie.SearchSpaceMTCH_r17 != nil || len(ie.CommonSearchSpaceListExt2_r17) > 0 || len(ie.FirstPDCCH_MonitoringOccasionOfPO_v1710) > 0 || len(ie.Pei_ConfigBWP_r17) > 0, ie.FollowUnifiedTCI_State_v1720 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PDCCH_ConfigCommon", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{len(ie.FirstPDCCH_MonitoringOccasionOfPO) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode FirstPDCCH_MonitoringOccasionOfPO optional
			if len(ie.FirstPDCCH_MonitoringOccasionOfPO) > 0 {
				tmp_FirstPDCCH_MonitoringOccasionOfPO := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxPO_perPF}, false)
				for _, i := range ie.FirstPDCCH_MonitoringOccasionOfPO {
					tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: 139}, false)
					tmp_FirstPDCCH_MonitoringOccasionOfPO.Value = append(tmp_FirstPDCCH_MonitoringOccasionOfPO.Value, &tmp_ie)
				}
				if err = tmp_FirstPDCCH_MonitoringOccasionOfPO.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FirstPDCCH_MonitoringOccasionOfPO", err)
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
			optionals_ext_2 := []bool{len(ie.CommonSearchSpaceListExt_r16) > 0}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode CommonSearchSpaceListExt_r16 optional
			if len(ie.CommonSearchSpaceListExt_r16) > 0 {
				tmp_CommonSearchSpaceListExt_r16 := utils.NewSequence[*SearchSpaceExt_r16]([]*SearchSpaceExt_r16{}, aper.Constraint{Lb: 1, Ub: 4}, false)
				for _, i := range ie.CommonSearchSpaceListExt_r16 {
					tmp_CommonSearchSpaceListExt_r16.Value = append(tmp_CommonSearchSpaceListExt_r16.Value, &i)
				}
				if err = tmp_CommonSearchSpaceListExt_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CommonSearchSpaceListExt_r16", err)
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
			optionals_ext_3 := []bool{ie.Sdt_SearchSpace_r17 != nil, ie.SearchSpaceMCCH_r17 != nil, ie.SearchSpaceMTCH_r17 != nil, len(ie.CommonSearchSpaceListExt2_r17) > 0, len(ie.FirstPDCCH_MonitoringOccasionOfPO_v1710) > 0, len(ie.Pei_ConfigBWP_r17) > 0}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Sdt_SearchSpace_r17 optional
			if ie.Sdt_SearchSpace_r17 != nil {
				if err = ie.Sdt_SearchSpace_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sdt_SearchSpace_r17", err)
				}
			}
			// encode SearchSpaceMCCH_r17 optional
			if ie.SearchSpaceMCCH_r17 != nil {
				if err = ie.SearchSpaceMCCH_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SearchSpaceMCCH_r17", err)
				}
			}
			// encode SearchSpaceMTCH_r17 optional
			if ie.SearchSpaceMTCH_r17 != nil {
				if err = ie.SearchSpaceMTCH_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SearchSpaceMTCH_r17", err)
				}
			}
			// encode CommonSearchSpaceListExt2_r17 optional
			if len(ie.CommonSearchSpaceListExt2_r17) > 0 {
				tmp_CommonSearchSpaceListExt2_r17 := utils.NewSequence[*SearchSpaceExt_v1700]([]*SearchSpaceExt_v1700{}, aper.Constraint{Lb: 1, Ub: 4}, false)
				for _, i := range ie.CommonSearchSpaceListExt2_r17 {
					tmp_CommonSearchSpaceListExt2_r17.Value = append(tmp_CommonSearchSpaceListExt2_r17.Value, &i)
				}
				if err = tmp_CommonSearchSpaceListExt2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CommonSearchSpaceListExt2_r17", err)
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
			// encode Pei_ConfigBWP_r17 optional
			if len(ie.Pei_ConfigBWP_r17) > 0 {
				tmp_Pei_ConfigBWP_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxPEI_perPF_r17}, false)
				for _, i := range ie.Pei_ConfigBWP_r17 {
					tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: 139}, false)
					tmp_Pei_ConfigBWP_r17.Value = append(tmp_Pei_ConfigBWP_r17.Value, &tmp_ie)
				}
				if err = tmp_Pei_ConfigBWP_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pei_ConfigBWP_r17", err)
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
			optionals_ext_4 := []bool{ie.FollowUnifiedTCI_State_v1720 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode FollowUnifiedTCI_State_v1720 optional
			if ie.FollowUnifiedTCI_State_v1720 != nil {
				if err = ie.FollowUnifiedTCI_State_v1720.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FollowUnifiedTCI_State_v1720", err)
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

func (ie *PDCCH_ConfigCommon) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var ControlResourceSetZeroPresent bool
	if ControlResourceSetZeroPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var CommonControlResourceSetPresent bool
	if CommonControlResourceSetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SearchSpaceZeroPresent bool
	if SearchSpaceZeroPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var CommonSearchSpaceListPresent bool
	if CommonSearchSpaceListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SearchSpaceSIB1Present bool
	if SearchSpaceSIB1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SearchSpaceOtherSystemInformationPresent bool
	if SearchSpaceOtherSystemInformationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var PagingSearchSpacePresent bool
	if PagingSearchSpacePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ra_SearchSpacePresent bool
	if Ra_SearchSpacePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ControlResourceSetZeroPresent {
		ie.ControlResourceSetZero = new(ControlResourceSetZero)
		if err = ie.ControlResourceSetZero.Decode(r); err != nil {
			return utils.WrapError("Decode ControlResourceSetZero", err)
		}
	}
	if CommonControlResourceSetPresent {
		ie.CommonControlResourceSet = new(ControlResourceSet)
		if err = ie.CommonControlResourceSet.Decode(r); err != nil {
			return utils.WrapError("Decode CommonControlResourceSet", err)
		}
	}
	if SearchSpaceZeroPresent {
		ie.SearchSpaceZero = new(SearchSpaceZero)
		if err = ie.SearchSpaceZero.Decode(r); err != nil {
			return utils.WrapError("Decode SearchSpaceZero", err)
		}
	}
	if CommonSearchSpaceListPresent {
		tmp_CommonSearchSpaceList := utils.NewSequence[*SearchSpace]([]*SearchSpace{}, aper.Constraint{Lb: 1, Ub: 4}, false)
		fn_CommonSearchSpaceList := func() *SearchSpace {
			return new(SearchSpace)
		}
		if err = tmp_CommonSearchSpaceList.Decode(r, fn_CommonSearchSpaceList); err != nil {
			return utils.WrapError("Decode CommonSearchSpaceList", err)
		}
		ie.CommonSearchSpaceList = []SearchSpace{}
		for _, i := range tmp_CommonSearchSpaceList.Value {
			ie.CommonSearchSpaceList = append(ie.CommonSearchSpaceList, *i)
		}
	}
	if SearchSpaceSIB1Present {
		ie.SearchSpaceSIB1 = new(SearchSpaceId)
		if err = ie.SearchSpaceSIB1.Decode(r); err != nil {
			return utils.WrapError("Decode SearchSpaceSIB1", err)
		}
	}
	if SearchSpaceOtherSystemInformationPresent {
		ie.SearchSpaceOtherSystemInformation = new(SearchSpaceId)
		if err = ie.SearchSpaceOtherSystemInformation.Decode(r); err != nil {
			return utils.WrapError("Decode SearchSpaceOtherSystemInformation", err)
		}
	}
	if PagingSearchSpacePresent {
		ie.PagingSearchSpace = new(SearchSpaceId)
		if err = ie.PagingSearchSpace.Decode(r); err != nil {
			return utils.WrapError("Decode PagingSearchSpace", err)
		}
	}
	if Ra_SearchSpacePresent {
		ie.Ra_SearchSpace = new(SearchSpaceId)
		if err = ie.Ra_SearchSpace.Decode(r); err != nil {
			return utils.WrapError("Decode Ra_SearchSpace", err)
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

			FirstPDCCH_MonitoringOccasionOfPOPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode FirstPDCCH_MonitoringOccasionOfPO optional
			if FirstPDCCH_MonitoringOccasionOfPOPresent {
				tmp_FirstPDCCH_MonitoringOccasionOfPO := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxPO_perPF}, false)
				fn_FirstPDCCH_MonitoringOccasionOfPO := func() *utils.INTEGER {
					ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 139}, false)
					return &ie
				}
				if err = tmp_FirstPDCCH_MonitoringOccasionOfPO.Decode(extReader, fn_FirstPDCCH_MonitoringOccasionOfPO); err != nil {
					return utils.WrapError("Decode FirstPDCCH_MonitoringOccasionOfPO", err)
				}
				ie.FirstPDCCH_MonitoringOccasionOfPO = []int64{}
				for _, i := range tmp_FirstPDCCH_MonitoringOccasionOfPO.Value {
					ie.FirstPDCCH_MonitoringOccasionOfPO = append(ie.FirstPDCCH_MonitoringOccasionOfPO, int64(i.Value))
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

			CommonSearchSpaceListExt_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode CommonSearchSpaceListExt_r16 optional
			if CommonSearchSpaceListExt_r16Present {
				tmp_CommonSearchSpaceListExt_r16 := utils.NewSequence[*SearchSpaceExt_r16]([]*SearchSpaceExt_r16{}, aper.Constraint{Lb: 1, Ub: 4}, false)
				fn_CommonSearchSpaceListExt_r16 := func() *SearchSpaceExt_r16 {
					return new(SearchSpaceExt_r16)
				}
				if err = tmp_CommonSearchSpaceListExt_r16.Decode(extReader, fn_CommonSearchSpaceListExt_r16); err != nil {
					return utils.WrapError("Decode CommonSearchSpaceListExt_r16", err)
				}
				ie.CommonSearchSpaceListExt_r16 = []SearchSpaceExt_r16{}
				for _, i := range tmp_CommonSearchSpaceListExt_r16.Value {
					ie.CommonSearchSpaceListExt_r16 = append(ie.CommonSearchSpaceListExt_r16, *i)
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

			Sdt_SearchSpace_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SearchSpaceMCCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SearchSpaceMTCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CommonSearchSpaceListExt2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			FirstPDCCH_MonitoringOccasionOfPO_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pei_ConfigBWP_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Sdt_SearchSpace_r17 optional
			if Sdt_SearchSpace_r17Present {
				ie.Sdt_SearchSpace_r17 = new(PDCCH_ConfigCommon_sdt_SearchSpace_r17)
				if err = ie.Sdt_SearchSpace_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sdt_SearchSpace_r17", err)
				}
			}
			// decode SearchSpaceMCCH_r17 optional
			if SearchSpaceMCCH_r17Present {
				ie.SearchSpaceMCCH_r17 = new(SearchSpaceId)
				if err = ie.SearchSpaceMCCH_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SearchSpaceMCCH_r17", err)
				}
			}
			// decode SearchSpaceMTCH_r17 optional
			if SearchSpaceMTCH_r17Present {
				ie.SearchSpaceMTCH_r17 = new(SearchSpaceId)
				if err = ie.SearchSpaceMTCH_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SearchSpaceMTCH_r17", err)
				}
			}
			// decode CommonSearchSpaceListExt2_r17 optional
			if CommonSearchSpaceListExt2_r17Present {
				tmp_CommonSearchSpaceListExt2_r17 := utils.NewSequence[*SearchSpaceExt_v1700]([]*SearchSpaceExt_v1700{}, aper.Constraint{Lb: 1, Ub: 4}, false)
				fn_CommonSearchSpaceListExt2_r17 := func() *SearchSpaceExt_v1700 {
					return new(SearchSpaceExt_v1700)
				}
				if err = tmp_CommonSearchSpaceListExt2_r17.Decode(extReader, fn_CommonSearchSpaceListExt2_r17); err != nil {
					return utils.WrapError("Decode CommonSearchSpaceListExt2_r17", err)
				}
				ie.CommonSearchSpaceListExt2_r17 = []SearchSpaceExt_v1700{}
				for _, i := range tmp_CommonSearchSpaceListExt2_r17.Value {
					ie.CommonSearchSpaceListExt2_r17 = append(ie.CommonSearchSpaceListExt2_r17, *i)
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
			// decode Pei_ConfigBWP_r17 optional
			if Pei_ConfigBWP_r17Present {
				tmp_Pei_ConfigBWP_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxPEI_perPF_r17}, false)
				fn_Pei_ConfigBWP_r17 := func() *utils.INTEGER {
					ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 139}, false)
					return &ie
				}
				if err = tmp_Pei_ConfigBWP_r17.Decode(extReader, fn_Pei_ConfigBWP_r17); err != nil {
					return utils.WrapError("Decode Pei_ConfigBWP_r17", err)
				}
				ie.Pei_ConfigBWP_r17 = []int64{}
				for _, i := range tmp_Pei_ConfigBWP_r17.Value {
					ie.Pei_ConfigBWP_r17 = append(ie.Pei_ConfigBWP_r17, int64(i.Value))
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

			FollowUnifiedTCI_State_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode FollowUnifiedTCI_State_v1720 optional
			if FollowUnifiedTCI_State_v1720Present {
				ie.FollowUnifiedTCI_State_v1720 = new(PDCCH_ConfigCommon_followUnifiedTCI_State_v1720)
				if err = ie.FollowUnifiedTCI_State_v1720.Decode(extReader); err != nil {
					return utils.WrapError("Decode FollowUnifiedTCI_State_v1720", err)
				}
			}
		}
	}
	return nil
}
