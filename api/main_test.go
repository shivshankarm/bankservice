package api

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/shivshankarm/bankservice/db/sqlc"
	"github.com/shivshankarm/bankservice/util"
	"github.com/stretchr/testify/require"

	_ "github.com/lib/pq"

	_ "github.com/shivshankarm/bankservice/util"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
