package pf2e

type InvalidDbCategoryError struct{}

func (ic InvalidDbCategoryError) Error() string {
	return "Invalid db category"
}
