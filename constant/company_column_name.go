package constant

import "errors"

type CompanyColumn string

const (
	CompanyColumnName CompanyColumn = "name"
)

var CompanyColumns = []CompanyColumn{
	CompanyColumnName,
}

func ParseCompanyColumnName(str string) (CompanyColumn, error) {
	for _, t := range CompanyColumns {
		if str == string(t) {
			return t, nil
		}
	}
	return "", errors.New(str)
}
