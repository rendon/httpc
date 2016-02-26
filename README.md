httpc
=====
[![GoDoc](https://godoc.org/github.com/rendon/httpc?status.svg)](https://godoc.org/github.com/rendon/httpc)

httpc package provides an HTTP client interface for creating API clients. It comes with a default implementation that covers most cases.

Using an interface for your API client allows you to stub requests in tests, just create a type that implements the interface with your custom behavior and the rest of your application will work seamlessly.

This package aims to be minimal, just a bit more than the standard HTTP package so you can still be in control of your software. 
