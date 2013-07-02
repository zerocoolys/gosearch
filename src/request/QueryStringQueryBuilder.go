package request

type Query struct {
	Query QueryStringQueryBuilder
}

type QueryStringQueryBuilder struct {
	QueryString QueryStringQuery
	// DefaultField           string
	// DefaultOperator        Operator
	// Analyzer               string
	// QuoteAnalyzer          string
	// QuoteFieldSuffix       string
	// AllowLeadingWildcard   bool
	// LowercaseExpandedTerms bool
	// Fields                 []string
}

type QueryStringQuery struct {
	Query string
}
type Operator struct {
	OR  string "OR"
	AND string "AND"
}
