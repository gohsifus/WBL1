package main

import (
	"bytes"
	"strings"
	"testing"
)

func concat1(inp []string) string {
	ret := ""
	for _, v := range inp {
		ret += v
	}

	return ret
}

func concat2(inp []string) string {
	ret := strings.Builder{}
	for _, v := range inp {
		ret.WriteString(v)
	}

	return ret.String()
}

func concat3(inp []string) string {
	ret := bytes.Buffer{}
	for _, v := range inp {
		ret.WriteString(v)
	}

	return ret.String()
}

func BenchmarkConcat1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concat1(inpData)
	}
}

func BenchmarkConcat2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concat2(inpData)
	}
}

func BenchmarkConcat3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concat3(inpData)
	}
}

var inpData = []string{
	"qqqqqqqqqqqwwwwwwwwwwwwwwrrrrrrrrrrrrkmvodfvodjodfmv",
	"qqqqqqqqqqqwwwwwwwwwwwwwwrrrrrrrrrrrrkmvodfvodjodfmv",
	"qqqqqqqqqqqwwwwwwwwwwwwwwrrrrrrrrrrrrkmvodfvodjodfmv",
	"qqqqqqqqовтлысымтщавмотваwwwwrrrrrrrrrrrrkmvodfvodjodfmv",
	"qqqqqqqqqqqwwwwwwwwwwwwwwrrrrrrrrrrrrkmvodfvodjodfmv",
	"qqqqqqqqqqqлатвоамтвалwwwrrrrrrrrrrrrkmvodfvodjodfmv",
	"qqqqqqqqqqqwwwwwwwwwwwwwwrrrrrrrrrrrrkmvodfvodjodfmv",
	"qqqqqqqqqqqwwwwwwwwwwwwwwrrrrrrrrrrrrkmvodfvodjodfmv",
	"qqqqqqqqqqqwwwwwwwwwwwwwwrrrrrrrrrrrrkmvodfvodjodfmv",
	"qqqqqqqqqПривет влриа влоаптвао rrrrrrrrrrrkmvodfvodjodfmv",
	"qqqqqqqqqqqwwwwwwwwwwwwwwrrrrrrrrrrrrkmvodfvodjodfmv",
	"qqqqqqqqqqqwwwwwwwwwwwwwwrrrrrrrrrrrrkmvodfvodjodfmv",
	"qqqqqqqqqqqwwwwwwwwwwwwwwrrrrrrrrrrrrkmvodfvodjodfmv",
	"qqqqqqqqqqqwwwaaaaaaaaaaaaaaaaaaaaaaarrrrrrrkmvodfvodjodfmv",
}