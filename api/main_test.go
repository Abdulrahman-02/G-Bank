package api

import (
	"os"
	"testing"
	"time"

	db "github.com/abdulrahman-02/G-Bank/db/sqlc"
	"github.com/abdulrahman-02/G-Bank/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:  util.RandomString(32),
		AccesTokenDuration: time.Minute,
	}
	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
