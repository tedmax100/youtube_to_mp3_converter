package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// 可以在這裡設置全局測試環境
	code := m.Run()
	os.Exit(code)
}

// 注意：main() 函數本身很難直接測試，因為它使用 os.Exit()
// 但我們已經通過單元測試覆蓋了所有主要組件：
// - config 包
// - validator 包
// - downloader 包
// 並且通過集成測試驗證了端到端流程
