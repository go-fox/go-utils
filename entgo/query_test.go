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
	whereSelector, _, err := BuildQuerySelect(&pagination.Where{
		LogicalOperator: pagination.LogicalOperatorAnd,
		Conditions: []pagination.Condition{
			{
				Field:           "age",
				Operator:        pagination.QueryOperatorGreater,
				Value:           10,
				LogicalOperator: pagination.LogicalOperatorOr,
				Conditions: []pagination.Condition{
					{
						Field:    "name",
						Operator: pagination.QueryOperatorEqual,
						Value:    "张三",
					},
					{
						Field:    "name",
						Operator: pagination.QueryOperatorEqual,
						Value:    "李四",
					},
				},
			},
		},
	})
	print(whereSelector)
	if err != nil {
		log.Fatalf("failed building query: %v", err)
	}
	all, err := client.Debug().User.Query().Modify(whereSelector).All(context.Background())
	if err != nil {
		log.Fatalf("failed querying users: %v", err)
	}
	for _, u := range all {
		log.Printf("user: %s", u.Name)
	}
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
