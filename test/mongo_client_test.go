package test

import (
	"context"
	"github.com/ramdanariadi/grocery-go-report/main/helpers"
	"github.com/ramdanariadi/grocery-go-report/main/utils"
	"testing"
)

func Test_Mongo_client(t *testing.T) {
	client := utils.GetClient()

	err := client.Ping(context.Background(), nil)
	helpers.LogIfError(err)

	client2 := utils.GetClient()
	err = client2.Ping(context.Background(), nil)
	helpers.LogIfError(err)
}
