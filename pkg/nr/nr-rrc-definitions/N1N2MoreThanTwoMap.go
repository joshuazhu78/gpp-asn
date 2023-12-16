package nrrrcdefinitions

import "github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"

func (c *N1N2MoreThanTwo) GetN1N2() (int, int, *asn1.BitString) {
	switch c.N1N2MoreThanTwo.(type) {
	case *N1N2MoreThanTwo_TwoOneTypeISinglePanelRestriction:
		return 2, 1, c.GetTwoOneTypeISinglePanelRestriction()
	case *N1N2MoreThanTwo_TwoTwoTypeISinglePanelRestriction:
		return 2, 2, c.GetTwoTwoTypeISinglePanelRestriction()
	case *N1N2MoreThanTwo_FourOneTypeISinglePanelRestriction:
		return 4, 1, c.GetFourOneTypeISinglePanelRestriction()
	case *N1N2MoreThanTwo_ThreeTwoTypeISinglePanelRestriction:
		return 3, 2, c.GetThreeTwoTypeISinglePanelRestriction()
	case *N1N2MoreThanTwo_SixOneTypeISinglePanelRestriction:
		return 6, 1, c.GetSixOneTypeISinglePanelRestriction()
	case *N1N2MoreThanTwo_FourTwoTypeISinglePanelRestriction:
		return 4, 2, c.GetFourTwoTypeISinglePanelRestriction()
	case *N1N2MoreThanTwo_EightOneTypeISinglePanelRestriction:
		return 8, 1, c.GetEightOneTypeISinglePanelRestriction()
	case *N1N2MoreThanTwo_FourThreeTypeISinglePanelRestriction:
		return 4, 3, c.GetFourThreeTypeISinglePanelRestriction()
	case *N1N2MoreThanTwo_SixTwoTypeISinglePanelRestriction:
		return 6, 2, c.GetSixTwoTypeISinglePanelRestriction()
	case *N1N2MoreThanTwo_TwelveOneTypeISinglePanelRestriction:
		return 12, 1, c.GetTwelveOneTypeISinglePanelRestriction()
	case *N1N2MoreThanTwo_FourFourTypeISinglePanelRestriction:
		return 4, 4, c.GetFourFourTypeISinglePanelRestriction()
	case *N1N2MoreThanTwo_EightTwoTypeISinglePanelRestriction:
		return 8, 2, c.GetEightTwoTypeISinglePanelRestriction()
	case *N1N2MoreThanTwo_SixteenOneTypeISinglePanelRestriction:
		return 16, 1, c.GetSixteenOneTypeISinglePanelRestriction()
	}
	return 0, 0, nil
}
