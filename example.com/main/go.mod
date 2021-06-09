module example.com/main

go 1.16

replace example.com/study => ../study

replace example.com/house => ../house

require (
	example.com/house v0.0.0-00010101000000-000000000000
	example.com/study v0.0.0-00010101000000-000000000000
)
