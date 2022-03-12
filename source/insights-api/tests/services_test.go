package minerstats

import (
	"testing"

	minerstats "github.com/malekatwiz/miners-stats-insights/services"
)

func Test_GetOperationWorkers_ReturnsAllOperationWorkers(t *testing.T) {
	workers := minerstats.GetOperationWorkers("flux")
	if len(workers) == 0 {
		t.Errorf("expected workers")
	}
}
