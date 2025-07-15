package ticker

const (
	qPutAggregatedData = `	
	INSERT INTO %s (%s, %s, %s, %s, %s, %s)
		VALUES %s
	`
)
