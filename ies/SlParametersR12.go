package ies

// SL-Parameters-r12 ::= SEQUENCE
type SlParametersR12 struct {
	CommsimultaneoustxR12          *SlParametersR12CommsimultaneoustxR12
	CommsupportedbandsR12          *FreqbandindicatorlisteutraR12
	DiscsupportedbandsR12          *SupportedbandinfolistR12
	DiscscheduledresourceallocR12  *SlParametersR12DiscscheduledresourceallocR12
	DiscUeSelectedresourceallocR12 *SlParametersR12DiscUeSelectedresourceallocR12
	DiscSlssR12                    *SlParametersR12DiscSlssR12
	DiscsupportedprocR12           *SlParametersR12DiscsupportedprocR12
}
