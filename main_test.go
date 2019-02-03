package main

import "testing"


// Obviously true
const ChanchiteIsMakingYouHappy = true

func TestMain(t *testing.T) {
    if  !ChanchiteIsMakingYouHappy {
        t.Fail()
    }
}
