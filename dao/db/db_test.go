package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitDB(t *testing.T) {
	// 连接数据库
	err := InitDB()
	assert.NoError(t, err, "数据库连接失败！")
}
