package validators

import "fmt"

func StringNotNull(name string) func(interface{}) error {
	return func(validateStr interface{}) error {
		if validateStr.(string) == "" {
			err := fmt.Sprintf("%s must be non-empty", name)
			return fmt.Errorf(err)
		}
		return nil
	}

}
