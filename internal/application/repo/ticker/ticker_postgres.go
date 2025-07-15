package ticker

import (
	"app/market/internal/domain/entity"
	"app/market/internal/domain/repo"
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	TableName = "aggregated"

	PairNameColumn     = "pair_name"
	ExchangeColumn     = "exchange"
	TimestampColumn    = "timestamp"
	AveragePriceColumn = "average_price"
	MinPriceColumn     = "min_price"
	MaxPriceColumn     = "max_price"
)

type tickerPGRepository struct {
	pgDB *pgxpool.Pool
}

func NewTickerPGRepository(pgConn *pgxpool.Pool) repo.TickerPGRepository {
	return &tickerPGRepository{
		pgDB: pgConn,
	}
}

func (t *tickerPGRepository) PutAggregatedBatch(ctx context.Context, batch []entity.AggregatedPrice) error {
	if len(batch) == 0 {
		return nil
	}

	valueStrings := make([]string, 0, len(batch))
	valueArgs := make([]interface{}, 0, len(batch)*6)

	for i, item := range batch {
		// создаём плейсхолдеры вида: ($1,$2,$3,$4,$5,$6), ($7,$8,...)
		start := i*6 + 1
		valueStrings = append(valueStrings, fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d)",
			start, start+1, start+2, start+3, start+4, start+5))
		valueArgs = append(valueArgs,
			item.PairName,
			item.Exchange,
			item.Timestamp,
			item.AveragePrice,
			item.MinPrice,
			item.MaxPrice,
		)
	}

	query := fmt.Sprintf(qPutAggregatedData,
		TableName,
		PairNameColumn,
		ExchangeColumn,
		TimestampColumn,
		AveragePriceColumn,
		MinPriceColumn,
		MaxPriceColumn,
		strings.Join(valueStrings, ","),
	)

	_, err := t.pgDB.Exec(ctx, query, valueArgs...)
	return err
}
