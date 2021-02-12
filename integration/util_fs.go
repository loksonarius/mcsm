package integration

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type fileState uint

const (
	dne = fileState(iota)
	isfile
	isdir
)

type fileStateCheck struct {
	path  string
	state fileState
}

func checkFileStates(tests []fileStateCheck) error {
	var errs ErrorList

	for _, test := range tests {
		info, err := os.Stat(test.path)
		if err != nil {
			if os.IsNotExist(err) && test.state == dne {
			} else {
				err = fmt.Errorf("file %s missing: %s", test.path, err)
				errs = append(errs, err)
			}
			continue
		}

		if info.IsDir() != (test.state == isdir) {
			if info.IsDir() {
				err = fmt.Errorf("file %s is a dir", test.path)
			} else {
				err = fmt.Errorf("file %s isn't a dir", test.path)
			}
			errs = append(errs, err)
			continue
		}

		if info.IsDir() && test.state == isfile {
			err = fmt.Errorf("file %s isn't a file", test.path)
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

type lineInFileCheck struct {
	path string
	line string
}

func checkLinesInFiles(tests []lineInFileCheck) error {
	var errs ErrorList

	for _, test := range tests {
		f, err := os.Open(test.path)
		if err != nil {
			err = fmt.Errorf("error opening file %s: %s", test.path, err)
			errs = append(errs, err)
			continue
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		found := false
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, test.line) {
				found = true
				break
			}
		}

		if err = scanner.Err(); err != nil {
			err = fmt.Errorf("error reading file %s: %s", test.path, err)
			errs = append(errs, err)
			continue
		}

		if !found {
			err = fmt.Errorf("'%s' not in file %s", test.line, test.path)
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
