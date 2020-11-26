package endpoints

import (
	"backend/postgres"
	"github.com/gofiber/fiber/v2"
)

type GasPriceHourly struct {
	Time int64   `json:"time"`
	Gas  float32 `json:"gas"`
}

const GasPricesHourlyInEthQuery string = `
SELECT date_part('epoch', DATE_TRUNC('hour', t.block_time))                   as time,
       round(sum(t.gas_used * t.gas_price / 1000000000000000000)::numeric, 2) as gas
FROM transactions t
	LEFT OUTER JOIN contracts c ON t."from" = c.address
WHERE t."to" NOT LIKE '0x0000000000000000000000000000000000000000'
GROUP BY DATE_TRUNC('hour', t.block_time)
ORDER BY time`

func GasPrice(ctx *fiber.Ctx) error {
	db := postgres.InitDB()
	defer db.Close()

	var GasPrices []GasPriceHourly
	_, err := db.Query(&GasPrices, GasPricesHourlyInEthQuery)
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	return ctx.Status(200).JSON(GasPrices)

}
