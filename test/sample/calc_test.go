package sample

import (
	"log"
	"testing"
)

type args struct {
	a        int
	b        int
	operator string
}

func TestCalc(t *testing.T) {
	// テストケース
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// 正常系
		{
			name: "足し算",
			args: args{
				a:        10,
				b:        2,
				operator: "+",
			},
			want:    12,
			wantErr: false,
		},
		// 異常系
		{
			name: "不正な演算子を指定",
			args: args{
				a:        10,
				b:        2,
				operator: "?",
			},
			wantErr: true,
		},
	}

	log.Println(">> テスト関数の実行前")
	defer func() {
		log.Println(">> テスト関数の実行後")
	}()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println(">> テストケースの実行前")
			defer func() {
				log.Println(">> テストケースの実行後")
			}()

			got, err := Calc(tt.args.a, tt.args.b, tt.args.operator)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMain(m *testing.M) {
	setup()

	defer teardown()

	m.Run()
}

func setup() {
	log.Println(">> テスト全体の実行前")
}

func teardown() {
	log.Println(">> テスト全体の実行後")
}
