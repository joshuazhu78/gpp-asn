package nrrrcdefinitions

var ReportIntervalMap = map[ReportInterval]int{
	ReportInterval_REPORT_INTERVAL_MS120:   120,
	ReportInterval_REPORT_INTERVAL_MS240:   240,
	ReportInterval_REPORT_INTERVAL_MS480:   480,
	ReportInterval_REPORT_INTERVAL_MS640:   640,
	ReportInterval_REPORT_INTERVAL_MS1024:  1024,
	ReportInterval_REPORT_INTERVAL_MS2048:  2048,
	ReportInterval_REPORT_INTERVAL_MS5120:  5120,
	ReportInterval_REPORT_INTERVAL_MS10240: 10240,
	ReportInterval_REPORT_INTERVAL_MS20480: 20480,
	ReportInterval_REPORT_INTERVAL_MS40960: 40960,
	ReportInterval_REPORT_INTERVAL_MIN1:    60000,
	ReportInterval_REPORT_INTERVAL_MIN6:    360000,
	ReportInterval_REPORT_INTERVAL_MIN12:   720000,
	ReportInterval_REPORT_INTERVAL_MIN30:   1800000,
}
