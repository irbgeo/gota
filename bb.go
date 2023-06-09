package ta

import "github.com/shopspring/decimal"

// BB - Bollinger Bands.
// https://www.tradingview.com/pine-script-reference/v5/#fun_ta.bb
func BB(series Series, period, multi int) (middle Series, upper Series, lower Series) {
	middle = SMA(series, period)
	dev := StDev(series, period).MultiConst(decimal.NewFromInt(int64(multi)))

	upper = middle.Add(dev)
	lower = middle.Sub(dev)

	return
}
