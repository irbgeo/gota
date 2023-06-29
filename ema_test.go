package ta_test

import (
	"testing"

	ta "github.com/irbgeo/gota"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

var (
	testPeriodEMA = 3

	testSrcSetEMA = []ta.Value{
		{
			Time:  1,
			Value: decimal.NewFromFloat(7),
		},
		{
			Time:  2,
			Value: decimal.NewFromFloat(8.1),
		},
		{
			Time:  3,
			Value: decimal.NewFromFloat(0.3993),
		},
		{
			Time:  4,
			Value: decimal.NewFromFloat(1.7712),
		},
		{
			Time:  5,
			Value: decimal.NewFromFloat(-4.0277),
		},
		{
			Time:  6,
			Value: decimal.NewFromFloat(900),
		},
		{
			Time:  7,
			Value: decimal.NewFromFloat(7),
		},
		{
			Time:  8,
			Value: decimal.NewFromFloat(2.71),
		},
		{
			Time:  9,
			Value: decimal.NewFromFloat(-2.71),
		},
		{
			Time:  10,
			Value: decimal.NewFromFloat(5.71),
		},
		{
			Time:  11,
			Value: decimal.NewFromFloat(77.11),
		},
	}
)

var expectedEMA = []ta.Value{
	{
		Time:  1,
		Value: decimal.NewFromFloat(0),
	},
	{
		Time:  2,
		Value: decimal.NewFromFloat(0),
	},
	{
		Time:  3,
		Value: decimal.NewFromFloat(5.1664333333333333),
	},
	{
		Time:  4,
		Value: decimal.NewFromFloat(4.0346888888888889),
	},
	{
		Time:  5,
		Value: decimal.NewFromFloat(1.3472259259259259),
	},
	{
		Time:  6,
		Value: decimal.NewFromFloat(300.8981506172839506),
	},
	{
		Time:  7,
		Value: decimal.NewFromFloat(202.9321004115226337),
	},
	{
		Time:  8,
		Value: decimal.NewFromFloat(136.1914002743484225),
	},
	{
		Time:  9,
		Value: decimal.NewFromFloat(89.8909335162322817),
	},
	{
		Time:  10,
		Value: decimal.NewFromFloat(61.8306223441548545),
	},
	{
		Time:  11,
		Value: decimal.NewFromFloat(66.9237482294365697),
	},
}

func TestEMA(t *testing.T) {
	actualEMA := ta.EMA(testSrcSetEMA, testPeriodEMA)

	require.Equal(t, len(expectedEMA), len(actualEMA), "compare len")
	for i := 0; i < len(actualEMA); i++ {
		require.Equalf(t, expectedEMA[i].Time, actualEMA[i].Time, "compare time: %d", i+1)
		require.Equalf(t, expectedEMA[i].Value.Round(4).String(), actualEMA[i].Value.Round(4).String(), "compare : %d", i+1)
	}
}
