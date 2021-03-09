package gotest

import (
	"testing"
)

var arr [1000][1000]int
func Loop1() {
	for i:=0;i<1000;i++{
		for j:=0;j<1000;j++ {
			arr[i][j] = i+j
		}
	}
}

func Loop2() {
	for i:=0;i<1000;i++{
		for j:=0;j<1000;j++ {
			arr[j][i] = i+j
		}
	}
}

func BenchmarkLoop2(b *testing.B) {
	b.ResetTimer()
	for i:=0;i<b.N;i++ {
		Loop2()
	}
}

func BenchmarkLoop1(b *testing.B) {
	b.ResetTimer()
	for i:=0;i<b.N;i++ {
		Loop1()
	}
}

