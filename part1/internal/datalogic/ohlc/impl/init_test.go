package impl

import (
	"testing"

	redisRepo "part1/internal/repository/redis"
)

func TestNew(t *testing.T) {
	redisRepoMock := new(redisRepo.MockRepository)

	datalogic := New(redisRepoMock)
	if datalogic == nil {
		t.Errorf("fail init datalogic")
	}
}
