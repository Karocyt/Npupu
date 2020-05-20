module github.com/Karocyt/Npupu

go 1.14

replace github.com/Karocyt/Npupu/internal/solver => ./internal/solver

replace github.com/Karocyt/Npupu/internal/parser => ./internal/parser

replace github.com/Karocyt/Npupu/internal/heuristics => ./internal/heuristics

require (
	github.com/Karocyt/Npupu/internal/heuristics v0.0.0-00010101000000-000000000000
	github.com/Karocyt/Npupu/internal/parser v0.0.0-00010101000000-000000000000
	github.com/Karocyt/Npupu/internal/solver v0.0.0-00010101000000-000000000000
)
