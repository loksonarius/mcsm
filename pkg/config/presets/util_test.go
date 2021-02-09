package presets

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func getDiff(err error, expected []string) (got, unexpected, missing []string) {
	got = make([]string, 0, 0)
	unexpected = make([]string, 0, 0)
	missing = make([]string, 0, 0)

	if err == nil {
		missing = expected
		return
	}

	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		missing = expected
		return
	}

	for _, err := range errs {
		got = append(got, err.Error())
	}

	for _, err := range errs {
		found := false
		for _, exp := range expected {
			if strings.Contains(err.Error(), exp) {
				found = true
				break
			}
		}

		if !found {
			unexpected = append(unexpected, err.Error())
		}
	}

	for _, exp := range expected {
		found := false
		for _, err := range errs {
			if strings.Contains(err.Error(), exp) {
				found = true
				break
			}
		}

		if !found {
			missing = append(missing, exp)
		}
	}

	return
}
