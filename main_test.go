package main

import (
	"testing"
)

/* old concurrent tests

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

/* old sequential tests

BenchmarkMain-4              224           4,834,421 ns/op
PASS
ok      client-concurrent       1.096s

BenchmarkMain-4              277           4,325,746 ns/op
PASS
ok      client-concurrent       1.207s



all done in 3 milliseconds
pkg: client-concurrent
BenchmarkMain4/testing_sequential_mode-4                     264     4,785,879 ns/op



all done in 5 milliseconds
BenchmarkMain4/testing_concurrent_mode_-4                    309     3,874,016 ns/op
ok      client-concurrent       2.477s


*/

func BenchmarkMain4(b *testing.B) {
	b.Run("testing sequential mode", func(b *testing.B) {
		for b.Loop() {
			seq()
		}
	})

	b.Run("testing concurrent mode ", func(b *testing.B) {
		for b.Loop() {
			cuncurr()
		}
	})
}
