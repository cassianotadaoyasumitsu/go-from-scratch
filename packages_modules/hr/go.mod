module example.com/hr

go 1.17

replace example.com/employee => ../employee

require example.com/employee v0.0.0-00010101000000-000000000000 // indirect
