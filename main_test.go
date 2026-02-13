package main

import (
	"os"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	for b.Loop() {
		os.Args = []string{"city1", "city2"}
		main()
	}
}

/*

BenchmarkMain-4              603           1937730 ns/op
PASS
ok      client-concurrent       1.179s
-------------------------------------------------------------
BenchmarkMain-4              622           1926602 ns/op
PASS
ok      client-concurrent       1.208s
------------------------------------------------------------
BenchmarkMain-4              775           1577186 ns/op
PASS
ok      client-concurrent       1.232s
*/
