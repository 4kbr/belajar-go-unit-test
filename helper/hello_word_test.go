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

///testin.M , biasa digunakan untuk before -> (test) -> after 
func TestMain(m *testing.M) {
	fmt.Println("Before Unit test")
	m.Run()
	fmt.Println("After Unit test")
}
