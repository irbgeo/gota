package ta

import "github.com/shopspring/decimal"

func BB(series Series, period, multi int) (middle Series, upper Series, lower Series) {
	middle = SMA(series, period)
	dev := StDev(series, period).MulConst(decimal.NewFromInt(int64(multi)))

	upper = middle.Add(dev)
	lower = middle.Sub(dev)

	return
}
