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

BenchmarkMain-4              603           1,937,730 ns/op
PASS
ok      client-concurrent       1.179s
-------------------------------------------------------------
BenchmarkMain-4              622           1,926,602 ns/op
PASS
ok      client-concurrent       1.208s
------------------------------------------------------------
BenchmarkMain-4              775           1,577,186 ns/op
PASS
ok      client-concurrent       1.232s
*/

/*

BenchmarkMain-4              224           4,834,421 ns/op
PASS
ok      client-concurrent       1.096s

BenchmarkMain-4              277           4,325,746 ns/op
PASS
ok      client-concurrent       1.207s
*/
