# Solution

The solution is to create 1 endpoint that will return sum of `gas-used` * `gas-price` in ETH hourly.

With using database provided, only given single day (07/09/2020) the query is quite simple

```c
SELECT DATE_PART('epoch', DATE_TRUNC('hour', t.block_time))                   as time,
       ROUND(SUM(t.gas_used * t.gas_price / 1000000000000000000)::NUMERIC, 2) as gas
FROM transactions t
	LEFT OUTER JOIN contracts c ON t."from" = c.address
WHERE t."to" NOT LIKE '0x0000000000000000000000000000000000000000'
GROUP BY DATE_TRUNC('hour', t.block_time)
ORDER BY time
```

But if the data is multiple day, there's couple of endpoints that I think will be very useful

```
/v1/gas-price
{
    "date": "2020-10-25 16:12:55"
}
// will return a single time & value


/v1/gas-price/between
{
    "from": "2020-10-25 16:12:55"
    "to"  : "2020-10-25 23:44:10"
}
// will return a total of gas price between that time
```

## Endpoint

```
/v1/gas-price
```

Will return JSON format:

```
[
    {
        "time": 1599436800, // int
        "gas": 294.21       // float
    },
    {
        "time": 1599440400, // int
        "gas": 309.49       // float
    },
    ...
]
```

## Stack
* Golang
* Fiber
* Go-PG
