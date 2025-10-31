package ies

// CSI-ReportConfig ::= SEQUENCE
// Extensible
type CsiReportconfig struct {
	Reportconfigid                             CsiReportconfigid
	Carrier                                    *Servcellindex
	Resourcesforchannelmeasurement             CsiResourceconfigid
	CsiImResourcesforinterference              *CsiResourceconfigid
	NzpCsiRsResourcesforinterference           *CsiResourceconfigid
	Reportconfigtype                           CsiReportconfigReportconfigtype
	Reportquantity                             *CsiReportconfigReportquantity
	Reportfreqconfiguration                    *CsiReportconfigReportfreqconfiguration
	Timerestrictionforchannelmeasurements      CsiReportconfigTimerestrictionforchannelmeasurements
	Timerestrictionforinterferencemeasurements CsiReportconfigTimerestrictionforinterferencemeasurements
	Codebookconfig                             *Codebookconfig
	Dummy                                      *CsiReportconfigDummy
	Groupbasedbeamreporting                    *CsiReportconfigGroupbasedbeamreporting
	CqiTable                                   *CsiReportconfigCqiTable
	Subbandsize                                CsiReportconfigSubbandsize
	NonPmiPortindication                       *[]Portindexfor8ranks                        `lb:1,ub:maxNrofNZPCsiRsResourcesperconfig`
	SemipersistentonpuschV1530                 *CsiReportconfigSemipersistentonpuschV1530   `ext`
	SemipersistentonpuschV1610                 *CsiReportconfigSemipersistentonpuschV1610   `ext`
	AperiodicV1610                             *CsiReportconfigAperiodicV1610               `ext`
	ReportquantityR16                          *CsiReportconfigReportquantityR16            `ext`
	CodebookconfigR16                          *CodebookconfigR16                           `ext`
	CqiBitspersubbandR17                       *CsiReportconfigCqiBitspersubbandR17         `ext`
	GroupbasedbeamreportingV1710               *CsiReportconfigGroupbasedbeamreportingV1710 `ext`
	CodebookconfigR17                          *CodebookconfigR17                           `ext`
	SharedcmrR17                               *CsiReportconfigSharedcmrR17                 `ext`
	CsiReportmodeR17                           *CsiReportconfigCsiReportmodeR17             `ext`
	NumberofsingletrpCsiMode1R17               *CsiReportconfigNumberofsingletrpCsiMode1R17 `ext`
	ReportquantityR17                          *CsiReportconfigReportquantityR17            `ext`
	SemipersistentonpuschV1720                 *CsiReportconfigSemipersistentonpuschV1720   `ext`
	AperiodicV1720                             *CsiReportconfigAperiodicV1720               `ext`
	CodebookconfigV1730                        *CodebookconfigV1730                         `ext`
}
