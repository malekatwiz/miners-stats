package minerstats

import (
	"testing"
)

func Test_GetOperationWorkers_ReturnsAllOperationWorkers(t *testing.T) {
	workers := GetOperationWorkers("flux")
	if len(workers) == 0 {
		t.Errorf("expected workers")
	}
}

func Test_GetOperation_ReturnsOperationStats(t *testing.T) {
	expectedCurrency := "flux"
	operation := GetOperation(expectedCurrency)
	if operation.Currency != expectedCurrency {
		t.Errorf("expected currency '%s', received '%s'", expectedCurrency, operation.Currency)
	}
}
