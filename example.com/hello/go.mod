module example.com/hello

go 1.16

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	github.com/bigwhite/functrace v0.0.0-20210622013229-318a19dbb29a
	rsc.io/quote v1.5.2
)

replace example.com/greetings => ../greetings
