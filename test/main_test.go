package testsample

import (
	"testing"
)

func Add(a, b int) int {
	return a + b
}

// ファイル名の末尾が_test.goにする。
// パッケージ名はそのフォルダにあるファイルと同じもの。もしくは_testを末尾につけたもの（公開要素しかテストできない）
// テスト関数名はTestから始める
func TestAdd(t *testing.T) {
	got := Add(1, 2)
	if got != 3 {
		t.Errorf("expect 3, but %d", got)
	}
}
