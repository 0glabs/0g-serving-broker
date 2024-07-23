package server

import (
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/ethereum/go-ethereum/common"

	"github.com/0glabs/0g-serving-agent/common/config"
	"github.com/0glabs/0g-serving-agent/common/contract"
	database "github.com/0glabs/0g-serving-agent/user/internal/db"
	"github.com/0glabs/0g-serving-agent/user/internal/handler"
)

func Main() {
	config := config.GetConfig()

	db, err := gorm.Open(mysql.Open(config.Database.User), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	if err := database.Migrate(db); err != nil {
		panic(err)
	}

	c, err := contract.NewServingContract(common.HexToAddress(config.ContractAddress), config, os.Getenv("NETWORK"))
	if err != nil {
		panic(err)
	}
	defer c.Close()

	r := gin.New()
	h := handler.New(db, c, config.ServingUrl, config.SigningKey, config.Address)
	h.Register(r)

	// Listen and Serve, config port with PORT=X
	if err := r.Run(); err != nil {
		panic(err)
	}
}
