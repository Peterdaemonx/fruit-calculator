package app

import (
	"gitlab.cmpayments.local/creditcard/fruit-price-calculator/internal/entities"
	"reflect"
	"testing"
)

func Test_fruitService_CalculateAveragePrice(t *testing.T) {
	type args struct {
		fruits entities.Fruits
	}
	tests := []struct {
		name string
		args args
		want entities.Fruits
	}{
		{
			name: "should calculate fruit average",
			args: args{fruits: entities.Fruits{
				entities.Fruit{
					Type:  "banana",
					Price: 1.95,
				},
				entities.Fruit{
					Type:  "banana",
					Price: 7.42,
				},
				entities.Fruit{
					Type:  "banana",
					Price: 1.71,
				},
			}},
			want: entities.Fruits{
				entities.Fruit{
					Type:  "banana",
					Price: 3.693333333333333,
				}},
		},
		{
			name: "should return empty fruit average",
			args: args{fruits: entities.Fruits{}},
			want: entities.Fruits{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := fruitService{}
			if got := s.CalculateAveragePrice(tt.args.fruits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateAveragePrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
