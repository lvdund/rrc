package ies

// MeasParameters-v1250 ::= SEQUENCE
type MeasparametersV1250 struct {
	Timert312R12                 *MeasparametersV1250Timert312R12
	AlternativetimetotriggerR12  *MeasparametersV1250AlternativetimetotriggerR12
	IncmoneutraR12               *MeasparametersV1250IncmoneutraR12
	IncmonutraR12                *MeasparametersV1250IncmonutraR12
	ExtendedmaxmeasidR12         *MeasparametersV1250ExtendedmaxmeasidR12
	ExtendedrsrqLowerrangeR12    *MeasparametersV1250ExtendedrsrqLowerrangeR12
	RsrqOnallsymbolsR12          *MeasparametersV1250RsrqOnallsymbolsR12
	CrsDiscoverysignalsmeasR12   *MeasparametersV1250CrsDiscoverysignalsmeasR12
	CsiRsDiscoverysignalsmeasR12 *MeasparametersV1250CsiRsDiscoverysignalsmeasR12
}
