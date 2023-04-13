package main

import (
	"fmt"

	"github.com/swxctx/gutil"
)

func main() {
	fmt.Println(gutil.NewRandom(gutil.RandLowerWordsAndNumber).RandomString(10))
	fmt.Println(gutil.NewRandom(gutil.RandUpperWordsAndNumber).RandomString(10))
	fmt.Println(gutil.NewRandom(gutil.RandLowerWords).RandomString(10))
	fmt.Println(gutil.NewRandom(gutil.RandUpperWords).RandomString(10))
	fmt.Println(gutil.NewRandom(gutil.RandLowerAndUpperWords).RandomString(10))
	fmt.Println(gutil.NewRandom(gutil.RandLowerAndUpperWordsAndNumber).RandomString(10))
	fmt.Println(gutil.NewRandom(gutil.RandNumbers).RandomString(10))
}
