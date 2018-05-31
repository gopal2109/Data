package schema

import "net/url"

// Model interface should be satisfied by the db models
type Model interface {
	Save() (string, error)
	Find(interface{}, int, int) (interface{}, error)
	Delete(string, interface{}) []error
	Update(string, interface{}) []error
}

// Create is a generic Document upsertions function
func Create(m ...Model) ([]string, []error) {
	for _, x := range m {
		errs := ValidateFields("", x)
		if len(errs) > 0 {
			return []string{}, errs
		}
	}

	ids := make([]string, 0)
	allErrors := make([]error, 0)

	for _, x := range m {
		id, err := x.Save()
		if err != nil {
			allErrors = append(allErrors, err)
		}
		ids = append(ids, id)
	}

	return ids, nil
	// TODO: Update Meta data as a part of NWT-885
}

// Get fetches data from the struct which implements Model
func Get(m Model, filters url.Values, offset, limit int) (interface{}, error) {
	delete(filters, "offset")
	delete(filters, "limit")
	data, err := m.Find(filters, offset, limit)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetOne finds one Db Document
func GetOne(m Model, key string, value interface{}) (interface{}, error) {
	filters := map[string]interface{}{key: value}
	data, err := m.Find(filters, 0, 1)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Update used for put calls
func Update(m Model, selectorKey string, value interface{}) []error {
	errs := m.Update(selectorKey, value)
	return errs
	// TODO: Update Meta data as a part of NWT-885
}

// Delete ..
func Delete(m Model, key string, value interface{}) []error {
	return m.Delete(key, value)
}
