package endpoints

import (
	"backend/postgres"
	"github.com/gofiber/fiber/v2"
)

type GasPriceHourly struct {
	Time int64   `json:"t"`
	Gas  float32 `json:"v"`
}

const GasPricesHourlyInEthQuery string = `
SELECT DATE_PART('epoch', DATE_TRUNC('hour', t.block_time))                   as time,
       ROUND(SUM(t.gas_used * t.gas_price / 1000000000000000000)::numeric, 2) as gas
FROM transactions t
	LEFT JOIN contracts cf ON t."from" = cf.address
	LEFT JOIN contracts ct ON t."to" = ct.address
WHERE cf.address IS NULL AND
      ct.address IS NULL AND
      t."to" NOT LIKE '0x0000000000000000000000000000000000000000'
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
