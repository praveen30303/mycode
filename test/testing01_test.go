package main

import (
	"testing"
	"regexp"
)

func TestHello(t *testing.T){
	testName := "RZFeeser"
    desiredResult := regexp.MustCompile("Hello there "+testName)

    result, err := Hello("RZFeeser")

    if !desiredResult.MatchString(result) || err != nil {
        t.Fatalf("Wanted %v, but got, %v, or got %v", desiredResult, result, err)
    }
}
