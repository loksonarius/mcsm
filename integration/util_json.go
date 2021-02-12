package integration

import (
	"fmt"
	"strings"

	"github.com/savaki/jq"
)

type jsonFieldCheck struct {
	path  string
	value string
}

func parseOp(p string) (jq.Op, error) {
	if p == "/" {
		return jq.Parse(".")
	}

	sub := strings.Split(p, "/")
	if len(sub) == 1 {
		return jq.Parse(p)
	}

	var dots []jq.Op
	for _, s := range sub {
		dots = append(dots, jq.Dot(s))
	}

	return jq.Chain(dots...), nil
}

func checkJsonFields(body string, tests []jsonFieldCheck) error {
	var errs ErrorList
	data := []byte(body)

	for _, test := range tests {
		op, err := parseOp(test.path)
		if err != nil {
			errs = append(errs, fmt.Errorf("failed to parse op: %s", err))
			continue
		}

		got, err := op.Apply(data)
		if err != nil {
			err = fmt.Errorf("failed to apply op %s: %s", test.path, err)
			errs = append(errs, err)
			continue
		}

		if string(got) != test.value {
			err = fmt.Errorf("bad value at path %s: got %s, expected: %s",
				test.path, string(got), test.value)
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
