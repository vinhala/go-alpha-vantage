package connection

// Base URL of the Alpha Vantage API
const API_BASE_URL = "https://alphavantage.co/query"

// Timeout for any Alpha Vantage API request
const API_TIMEOUT = 10

// Available query functions for the Alpha Vantage API
type QueryFunction string

const (
	// Time Series Intraday
	TIME_SERIES_INTRADAY QueryFunction = "TIME_SERIES_INTRADAY"
	// Time Series Daily
	TIME_SERIES_DAILY QueryFunction = "TIME_SERIES_DAILY"
	// Time Series Daily Adjusted
	TIME_SERIES_DAILY_ADJUSTED QueryFunction = "TIME_SERIES_DAILY_ADJUSTED"
	// Time Series Weekly
	TIME_SERIES_WEEKLY QueryFunction = "TIME_SERIES_WEEKLY"
	// Time Series Weekly Adjusted
	TIME_SERIES_WEEKLY_ADJUSTED QueryFunction = "TIME_SERIES_WEEKLY_ADJUSTED"
	// Time Series Monthly
	TIME_SERIES_MONTHLY QueryFunction = "TIME_SERIES_MONTHLY"
	// Time Series Monthly Adjusted
	TIME_SERIES_MONTHLY_ADJUSTED QueryFunction = "TIME_SERIES_MONTHLY_ADJUSTED"
	// Global Quote
	GLOBAL_QUOTE QueryFunction = "GLOBAL_QUOTE"
)
