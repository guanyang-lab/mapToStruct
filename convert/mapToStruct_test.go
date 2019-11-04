package convert

import (
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"testing"
	"time"
)

type User struct {
	gorm.Model
	Name   string
	Age    int8
	Date   time.Time
	Money  decimal.Decimal
	DeptID uint
	Dept   Dept
	A      map[string]interface{}
}
type Dept struct {
	gorm.Model
	DeptName string
}

func TestMapToStruct(t *testing.T) {
	type args struct {
		data map[string]interface{}
		obj  interface{}
	}
	d := Dept{}
	d.ID = 1
	d.DeptName = "djkd"
	data := map[string]interface{}{
		"ID":     10,
		"Name":   "张三",
		"Age":    26,
		"Money":  30,
		"Date":   time.Now(),
		"DeptID": 1,
		"A": map[string]interface{}{
			"AA": 12,
			"dd": "41",
		},
		"Dept": d,
	}
	data1 := map[string]interface{}{
		"ID":     10,
		"Name":   "张三",
		"Age":    26,
		"Money":  30,
		"Date":   time.Now(),
		"DeptID": 1,
		"A":      11,
		"Dept":   d,
	}
	result := &User{}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"1", args{data, result}, false},
		{"2", args{data1, result}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MapToStruct(tt.args.data, tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("MapToStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDataToAnyData(t *testing.T) {
	m := map[string]interface{}{
		"ID":       10,
		"DeptName": "是什么",
	}
	m1 := map[string]interface{}{
		"ID":       "fdas",
		"DeptName": "是什么",
	}

	type args struct {
		data    interface{}
		newData interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"1", args{m, &Dept{}}, false},
		{"2", args{m1, &Dept{}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DataToAnyData(tt.args.data, tt.args.newData); (err != nil) != tt.wantErr {
				t.Errorf("DataToAnyData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
