package handler

import (
	mocks "fullstack_api_test/mocks/service"
	"fullstack_api_test/pkg/config"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := config.LoadWithPath("./../../../configs/config-test.yml")
	if err != nil {
		log.Fatal("Load config error: ", err)
	}
	code := m.Run()
	os.Exit(code)
}

func TestRegisterHandlers(t *testing.T) {
	h := NewHandler(
		&mocks.EmployeeService{},
	)
	RegisterHandlers(echo.New(), h)
}
