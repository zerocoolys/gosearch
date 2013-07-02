package query

type QueryStringQuery struct {
	Query                  string
	Default_field          string
	analyzer               string
	allow_leading_wildcard string
}

type TermsQuery struct {
}

type Query struct {
	QueryStringQuery QueryStringQuery
}
