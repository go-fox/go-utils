package entgo

import (
	"context"
	"github.com/go-fox/go-utils/entgo/test/ent"
	"github.com/go-fox/go-utils/pagination"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

func TestQuery(t *testing.T) {
	client, err := ent.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test_db?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	p, err := Pagination[*ent.UserQuery, *ent.User](context.Background(), client.User.Query(), &pagination.PagingRequest{
		Page:       1,
		Size:       1,
		Pagination: true,
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", p)
}

func createUser(client *ent.Client) error {
	_, err := client.User.
		Create().
		SetAge(18).
		SetName("张三").
		Save(context.Background())
	if err != nil {
		return err
	}
	_, err = client.User.
		Create().
		SetAge(22).
		SetName("李四").
		Save(context.Background())
	if err != nil {
		return err
	}
	_, err = client.User.
		Create().
		SetAge(30).
		SetName("王五").
		Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}
