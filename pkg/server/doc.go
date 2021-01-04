// Package server contains the general server interface containing
// operational commmands for use by clients along with implementations of the
// server interface for various server kinds, as well as the server definition
// file spec that should be used by clients hoping to request a server interface
// implementation. Using a server interface implementation, a client should be
// capable of performing operational tasks (such as starting, installing,
// configuring, etc) on a given server so long as they have the server
// definition file. The server package operates expecting local filesystem
// access to the server definition's root directory for operations such as
// installing and configuring.
package server
