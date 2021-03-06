package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func start() {
	var t int
	mustReadLineOfInts(&t)
	for i := 0; i < t; i++ {
		fmt.Fprintf(stdout, "Case #%d: ", i+1)
		test()
	}
}

type Sign struct {
	d          int
	a, b       int
	m, n       int
	mLen, nLen int
}

func test() {
	var numSigns int
	mustReadLineOfInts(&numSigns)
	signs := make([]Sign, numSigns)
	for i := 0; i < numSigns; i++ {
		var d, a, b int
		mustReadLineOfInts(&d, &a, &b)
		m := d + a
		n := d - b
		signs[i] = Sign{d, a, b, m, n, 0, 0}
		// debug(fmt.Sprintf("m=%v, n=%v", m, n))
	}
	maxValidSetLen := 0
	maxSetCount := 0
	for i := 0; i < len(signs); i++ {
		sign := &signs[i]
		m, n := sign.m, sign.n
		mLen := 0
		nLen := 0
		mBreak := false
		nBreak := false
		for j := i; j < len(signs); j++ {
			if signs[j].m == m && !mBreak {
				mLen++
			}
			if signs[j].n == n && !nBreak {
				nLen++
			}
			if signs[j].m != m {
				mBreak = true
			}
			if signs[j].n != n {
				nBreak = true
			}
			if mBreak && nBreak {
				break
			}
		}
		sign.mLen = mLen
		sign.nLen = nLen
	}
	for i, startSign := range signs {
		var setSize, m, n int
		var mSet, nSet bool
		continueWith := func(checkM bool) int {
			sLen := 0
			for i+sLen < len(signs) {
				s := signs[i+sLen]
				if checkM {
					if mSet && s.m != m {
						break
					} else if !mSet {
						m = s.m
						mSet = true
					}
					sLen += s.mLen
				} else {
					if nSet && s.n != n {
						break
					} else if !nSet {
						n = s.n
						nSet = true
					}
					sLen += s.nLen
				}
				checkM = !checkM
			}
			return sLen
		}
		m = startSign.m
		mSet = true
		nSet = false
		a := continueWith(true)
		n = startSign.n
		nSet = true
		mSet = false
		b := continueWith(false)
		setSize = max(a, b)
		if maxValidSetLen < setSize {
			maxValidSetLen = setSize
			maxSetCount = 1
		} else if maxValidSetLen == setSize {
			maxSetCount++
		}
	}
	fmt.Fprintf(stdout, "%d %d\n", maxValidSetLen, maxSetCount)
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

/*********Start boilerplate***********/

var stdin *bufio.Reader
var stdout *bufio.Writer

func main() {
	readFrom := os.Stdin
	if len(os.Args) == 2 {
		var err error
		readFrom, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
	}
	if len(os.Args) > 2 {
		panic("Too much arguments.")
	}
	stdin = bufio.NewReader(readFrom)
	stdout = bufio.NewWriter(os.Stdout)
	defer stdout.Flush()
	start()
}

func mustReadLine() string {
	str, err := stdin.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	str = strings.TrimRight(str, "\n")
	return str
}

func mustAtoi(in string) int {
	i, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return i
}

func debug(str string) {
	stdout.Flush()
	os.Stderr.WriteString("[debug] " + str + "\n")
}

func mustReadLineOfInts(a ...*int) {
	line := mustReadLine()
	strs := strings.Split(line, " ")
	if len(strs) != len(a) {
		panic("Expected " + strconv.Itoa(len(a)) + " numbers, got " + strconv.Itoa(len(strs)) + ".")
	}
	for i := 0; i < len(a); i++ {
		(*a[i]) = mustAtoi(strs[i])
	}
}

func mustReadLineOfIntsIntoArray() []int {
	line := mustReadLine()
	strs := strings.Split(line, " ")
	res := make([]int, len(strs))
	for i := 0; i < len(strs); i++ {
		res[i] = mustAtoi(strs[i])
	}
	return res
}

func assert(t bool) {
	if !t {
		panic("Assertion failed.")
	}
}
