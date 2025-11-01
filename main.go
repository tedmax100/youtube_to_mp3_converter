package main

import (
	"fmt"
	"os"

	"youtube_to_mp3/pkg/config"
	"youtube_to_mp3/pkg/downloader"
	"youtube_to_mp3/pkg/validator"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("使用方法: go run main.go <YouTube URL>")
		fmt.Println("範例: go run main.go https://www.youtube.com/watch?v=dQw4w9WgXcQ")
		os.Exit(1)
	}

	youtubeURL := os.Args[1]

	fmt.Println("開始處理 YouTube 視頻...")
	fmt.Printf("URL: %s\n\n", youtubeURL)

	// 檢查依賴
	systemValidator := validator.NewSystemValidator(nil)
	if err := systemValidator.ValidateDependencies(); err != nil {
		fmt.Printf("錯誤: %v\n", err)
		os.Exit(1)
	}

	// 創建配置
	cfg := config.NewConfig()

	// 創建下載器
	dl := downloader.NewYtDlpDownloader(cfg, nil)

	// 下載並轉換為 MP3
	fmt.Println("正在下載並轉換...")
	fmt.Println("(大文件轉換可能需要幾分鐘，請耐心等待...)")

	if err := dl.Download(youtubeURL); err != nil {
		fmt.Printf("\n錯誤: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n轉換完成！正在查找輸出文件...")

	// 顯示輸出文件
	files, err := dl.GetOutputFiles()
	if err != nil {
		fmt.Printf("警告: 無法查找輸出文件: %v\n", err)
	} else if len(files) > 0 {
		fmt.Printf("\n成功！MP3 文件已保存到: %s\n", files[len(files)-1])
	} else {
		fmt.Println("警告: 未找到輸出文件，但轉換過程已完成")
		fmt.Printf("請檢查 %s 目錄\n", cfg.OutputDir)
	}

	fmt.Println("\n✓ 全部完成！")
}
