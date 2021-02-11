// Package logging provides a consistent set of logging functions and
// configuration for package-wide use. It is intended to be shared across an
// entire package either privately or publicly. Private use would be done by
// having some package-wide var declaration like `log` that functions can call,
// for example. To allow package consumers to control things like the the
// destination files to write stdout and stderr messages to, or set the minimum
// log level per-package.
package logging
