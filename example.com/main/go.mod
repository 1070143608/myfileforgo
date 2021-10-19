module example.com/main

go 1.16

replace example.com/study => ../study

require (
	example.com/study v0.0.0-00010101000000-000000000000
	github.com/bigwhite/functrace v0.0.0-20210622013229-318a19dbb29a
)