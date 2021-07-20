package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

var tests = []struct {
	n    int
	want string
}{
	{1, "1\n"},
	{2, "1, 2\n"},
	{3, "1, 2, Fizz\n"},
	{4, "1, 2, Fizz, 4\n"},
	{5, "1, 2, Fizz, 4, Buzz\n"},
	{6, "1, 2, Fizz, 4, Buzz, Fizz\n"},
	{7, "1, 2, Fizz, 4, Buzz, Fizz, 7\n"},
	{8, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8\n"},
	{9, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz\n"},
	{10, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz\n"},
	{11, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11\n"},
	{12, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz\n"},
	{13, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13\n"},
	{14, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14\n"},
	{15, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz\n"},
	{16, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz, 16\n"},
	{17, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz, 16, 17\n"},
	{18, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz, 16, 17, Fizz\n"},
	{19, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz, 16, 17, Fizz, 19\n"},
	{20, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz, 16, 17, Fizz, 19, Buzz\n"},
}

func TestFizzBuzz(t *testing.T) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("N=%d", test.n), func(t *testing.T) {
			testStdOut, writer, err := os.Pipe()
			if err != nil {
				t.Fatalf("os.Pipe() err = %v; want %v", err, nil)
			}
			backupOsOut := os.Stdout // keep backup of the real stdout
			defer func() {
				os.Stdout = backupOsOut
			}()
			os.Stdout = writer

			FizzBuzz(test.n)
			writer.Close()

			var buf bytes.Buffer
			// should check copy err here
			io.Copy(&buf, testStdOut)
			got := buf.String()
			if got != test.want {
				t.Errorf("FizzBuzz(%d) = %q; want %q", test.n, got, test.want)
			}
		})
	}
}

func TestFizzBuzzWithoutJoin(t *testing.T) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("N=%d", test.n), func(t *testing.T) {
			testStdOut, writer, err := os.Pipe()
			if err != nil {
				t.Fatalf("os.Pipe() err = %v; want %v", err, nil)
			}
			backupOsOut := os.Stdout // keep backup of the real stdout
			defer func() {
				os.Stdout = backupOsOut
			}()
			os.Stdout = writer

			FizzBuzzWithoutJoin(test.n)
			writer.Close()

			var buf bytes.Buffer
			// should check copy err here
			io.Copy(&buf, testStdOut)
			got := buf.String()
			if got != test.want {
				t.Errorf("FizzBuzz(%d) = %q; want %q", test.n, got, test.want)
			}
		})
	}
}
