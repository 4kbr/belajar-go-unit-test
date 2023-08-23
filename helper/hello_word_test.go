package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
● Selain aturan nama file, di Go-Lang juga sudah diatur untuk nama function unit test
● Nama function untuk unit test harus diawali dengan nama Test
● Misal jika kita ingin mengetest function HelloWorld, maka kita akan membuat function unit test
dengan nama TestHelloWorld
● Tak ada aturan untuk nama belakang function unit test harus sama dengan nama function yang
akan di test, yang penting adalah harus diawalin dengan kata Test
● Selanjutnya harus memiliki parameter (t *testing.T) dan tidak mengembalikan return value
*/

func TestHelloWord(t *testing.T) {
	result := HelloWord("Aku")
	if result != "Hello Aku" {
		//error
		// panic("Result is not Hello Aku")
		// t.Fail()
		t.Error("Result is not Hello Aku")
	}
	fmt.Printf("TestHelloWord Done") //tetap dirunning
}

func TestHelloWordKamu(t *testing.T) {
	result := HelloWord("Kamu")
	if result != "Hello Kamu" {
		//error
		// panic("Result is not Hello Kamu")
		// t.FailNow()
		t.Fatal("Result is not Hello Kamu")
	}
	fmt.Printf("TestHelloWordKamu Done") //tidak dirunning
}

//running dengan go test
/*
PASS
ok      belajar-go-unit-test/helper     1.120s
*/

// go test -v
/*
=== RUN   TestHelloWord
--- PASS: TestHelloWord (0.00s)
=== RUN   TestHelloWordKamu
--- PASS: TestHelloWordKamu (0.00s)
PASS
ok      belajar-go-unit-test/helper     0.461s
*/

// go test -v -run=TestHelloWordKamu
/*
=== RUN   TestHelloWordKamu
--- PASS: TestHelloWordKamu (0.00s)
PASS
ok      belajar-go-unit-test/helper     0.141s
*/

//salah satu library yang paling populer di Go adalah Testify
// github.com/stretchr/testify

func TestHelloWordAssert(t *testing.T) {
	result := HelloWord("Kita")
	assert.Equal(t, "Hello Kita", result, "Result must be 'Hello Kita'")
	fmt.Printf("TestHelloWordAssert Done")
}

// bedanya assert vs require
// kalau assert jika gagal, akan memanggil Fail()
// kalau require akan memanggil FailNow()
func TestHelloWordRequire(t *testing.T) {
	result := HelloWord("Kita")
	require.Equal(t, "Hello Kita", result, "Result must be 'Hello Kita'") //langsung berhenti jika gagal
	fmt.Printf("TestHelloWordAssert Done")
}

// /penggunaan Skip
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWord("Amboi")
	require.Equal(t, "Hello Amboi", result, "Result must be 'Hello Amboi'")
}

// /testin.M , biasa digunakan untuk before -> (test) -> after
// func TestMain(m *testing.M) {
// 	fmt.Println("Before Unit test")
// 	m.Run()
// 	fmt.Println("After Unit test")
// }

// Membuat subtest didalam test
// untuk menjalankan subtest
// go test -v -run=TestSubTest/Aku //untuk menjalankan salah satu subtest
// go test -v -run=/Aku //untuk semua subtest yang ada kata aku
func TestSubTest(t *testing.T) {
	t.Run("Aku", func(t *testing.T) {
		result := HelloWord("Aku")
		require.Equal(t, "Hello Aku", result, "Result must be 'Hello Aku'")
	})
	t.Run("Kita", func(t *testing.T) {
		result := HelloWord("Kita")
		require.Equal(t, "Hello Kita", result, "Result must be 'Hello Kita'")
	})
}

// TableTest
func TestTableHelloWord(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name: "Aku", request: "Aku", expected: "Hello Aku",
		},
		{
			name: "Kamu", request: "Kamu", expected: "Hello Kamu",
		},
		{
			name: "Kita", request: "Kita", expected: "Hello Kita",
		},
		{
			"Mereka", "Mereka", "Hello Mereka",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWord(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

//Mock, sebagai pengganti third party saat melakukan uji coba

// /Benchmark
// perlu menamai functionnya dengan awal Benchmark
// go test -v -bench=
func BenchmarkHelloWordAndi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWord("Andi")
	}
}

// go test -v -bench=
// go test -v -bench=BenchmarkHelloWordAndi
/*
goos: darwin
goarch: amd64
pkg: belajar-go-unit-test/helper
cpu: Intel(R) Core(TM) i5-5350U CPU @ 1.80GHz
BenchmarkHelloWordAndi
BenchmarkHelloWordAndi-4        45876061                26.88 ns/op
*/

func BenchmarkHelloWordBudi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWord("Budi")
	}
}

// go test -v -run=BenchmarkHelloWordAndi -bench=BenchmarkHelloWordAndi
/*
goos: darwin
goarch: amd64
pkg: belajar-go-unit-test/helper
cpu: Intel(R) Core(TM) i5-5350U CPU @ 1.80GHz
BenchmarkHelloWordAndi
BenchmarkHelloWordAndi-4        51590070                23.26 ns/op
PASS
ok      belajar-go-unit-test/helper     2.654s
*/

// Sub Benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("Andi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWord("Andi")
		}
	})
	b.Run("Budi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWord("Budi")
		}
	})
}

/*
go test -v -run=TIdakAda -bench=BenchmarkSub
goos: darwin
goarch: amd64
pkg: belajar-go-unit-test/helper
cpu: Intel(R) Core(TM) i5-5350U CPU @ 1.80GHz
BenchmarkSub
BenchmarkSub/Andi
BenchmarkSub/Andi-4             49374896                26.01 ns/op
BenchmarkSub/Budi
BenchmarkSub/Budi-4             42356473                25.22 ns/op
PASS
ok      belajar-go-unit-test/helper     5.317s
*/

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{name: "Asep", request: "Asep"},
		{name: "Dio", request: "Dio"},
		{name: "Jojo", request: "Jojo"},
		{name: "Johan", request: "Johan"},
	}

	for _, bench := range benchmarks {
		b.Run(bench.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWord(bench.request)
			}
		})
	}
}
/*
go test -v -run=TIdakAda -bench=BenchmarkTable              
goos: darwin
goarch: amd64
pkg: belajar-go-unit-test/helper
cpu: Intel(R) Core(TM) i5-5350U CPU @ 1.80GHz
BenchmarkTable
BenchmarkTable/Asep
BenchmarkTable/Asep-4           46823917                24.14 ns/op
BenchmarkTable/Dio
BenchmarkTable/Dio-4            41451638                42.13 ns/op
BenchmarkTable/Jojo
BenchmarkTable/Jojo-4           37881945                28.02 ns/op
BenchmarkTable/Johan
BenchmarkTable/Johan-4          42517900                26.11 ns/op
PASS
ok      belajar-go-unit-test/helper     8.352s
*/
