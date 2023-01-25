package test

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

type TestParameters struct {
	BaseUrl string
	ApiKey  string
	Logging bool
	AuthTLS bool
}

var (
	TestArgs        TestParameters
	TestInitPending bool        = true
	pkgLog          *log.Logger = log.Default()
)

func init() {
	if TestInitPending {
		TestInitPending = false
		os.Mkdir("test_output", os.ModeDir)
		jsonFile, err := os.Open("test-config.json")
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)

		json.Unmarshal(byteValue, &TestArgs)
		if TestArgs.Logging {
			pkgLog.Println(TestArgs)
		} else {
			log.SetOutput(io.Discard)
			pkgLog.Println("Discarding...")
			pkgLog.SetOutput(io.Discard)
			pkgLog.Println("Discarded.")
		}
	}
}

// This method is for init and cleanup
func TestMain(m *testing.M) {
	// Init test cases
	//InitTest()

	// Run test cases
	c := m.Run()

	// Any possible cleanup goes here.

	// Exit test framework
	os.Exit(c)
}
