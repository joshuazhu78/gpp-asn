package nrrrcdefinitions

import "github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"

func (c *CsiReportingBandreportFreqConfiguration) GetCsiReportingBandreportFreqConfigurationSubband() *asn1.BitString {
	switch c.CsiReportingBandreportFreqConfiguration.(type) {
	case *CsiReportingBandreportFreqConfiguration_Subbands3:
		return c.GetSubbands3()
	case *CsiReportingBandreportFreqConfiguration_Subbands4:
		return c.GetSubbands4()
	case *CsiReportingBandreportFreqConfiguration_Subbands5:
		return c.GetSubbands5()
	case *CsiReportingBandreportFreqConfiguration_Subbands6:
		return c.GetSubbands6()
	case *CsiReportingBandreportFreqConfiguration_Subbands7:
		return c.GetSubbands7()
	case *CsiReportingBandreportFreqConfiguration_Subbands8:
		return c.GetSubbands8()
	case *CsiReportingBandreportFreqConfiguration_Subbands9:
		return c.GetSubbands9()
	case *CsiReportingBandreportFreqConfiguration_Subbands10:
		return c.GetSubbands10()
	case *CsiReportingBandreportFreqConfiguration_Subbands11:
		return c.GetSubbands11()
	case *CsiReportingBandreportFreqConfiguration_Subbands12:
		return c.GetSubbands12()
	case *CsiReportingBandreportFreqConfiguration_Subbands13:
		return c.GetSubbands13()
	case *CsiReportingBandreportFreqConfiguration_Subbands14:
		return c.GetSubbands14()
	case *CsiReportingBandreportFreqConfiguration_Subbands15:
		return c.GetSubbands15()
	case *CsiReportingBandreportFreqConfiguration_Subbands16:
		return c.GetSubbands16()
	case *CsiReportingBandreportFreqConfiguration_Subbands17:
		return c.GetSubbands17()
	case *CsiReportingBandreportFreqConfiguration_Subbands18:
		return c.GetSubbands18()
	case *CsiReportingBandreportFreqConfiguration_Subbands19V1530:
		return c.GetSubbands19V1530()
	}
	return nil
}
