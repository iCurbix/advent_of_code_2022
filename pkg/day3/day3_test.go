package day3

import (
	"os"
	"path"
	"runtime"
	"testing"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func BenchmarkVariant1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Variant1()
	}
}
func BenchmarkVariant2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Variant2()
	}
}
