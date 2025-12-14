package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"

	"github.com/Mr-Cheen1/go-reg-wails/backend/storage"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Создаем хранилище
	excelStorage := storage.NewExcelStorage()
	defer excelStorage.Close()

	// Создаем экземпляр приложения
	app := NewApp(excelStorage)

	// Создаем приложение Wails
	err := wails.Run(&options.App{
		Title:     "Редактор базы данных",
		Width:     800,
		Height:    650,
		MinWidth:  700,
		MinHeight: 500,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		OnDomReady:       app.OnDomReady,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}
