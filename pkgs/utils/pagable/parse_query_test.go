package pagable_test

import (
	"github.com/tdatIT/who-sent-api/pkgs/utils/pagable"
	"reflect"
	"testing"
)

func TestORMConditions(t *testing.T) {
	q := &pagable.Query{
		ExpressionFilters: []pagable.Filter{
			{Field: "Name", Value: "John", Operation: pagable.Equal},
			{Field: "Age", Value: 30, Operation: pagable.GT},
		},
	}

	conds := map[string]int{
		"name": pagable.StringType,
		"age":  pagable.NumberType,
	}

	expected := []*pagable.QueryCondition{
		{Condition: "name = ?", Value: "John"},
		{Condition: "age > ?", Value: 30},
	}

	result := q.ORMConditions(conds)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestParseCondition(t *testing.T) {
	q := &pagable.Query{}
	filter := pagable.Filter{Field: "Name", Operation: pagable.Equal}
	expected := "Name = ?"

	result := q.ParseCondition(filter)
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestParseQueryParams(t *testing.T) {
	q := &pagable.Query{
		ExpressionFilters: []pagable.Filter{
			{Field: "Name", Value: "John", Operation: pagable.Equal},
		},
	}

	expected := map[string]string{
		"Name": "John",
	}

	result, err := q.ParseQueryParams()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
