// run command: GOPATH=/Users/feynmanyuan/projects/blockchain/golang/riceis:$GOPATH qtdeploy test desktop .
package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/webview"
	"github.com/therecipe/qt/gui"
)

func NewMainFrame(args []string) {

	gui.NewQGuiApplication(len(args), args)
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)


	webview.QtWebView_Initialize()

	var view = qml.NewQQmlApplicationEngine(nil)
	view.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

	gui.QGuiApplication_Exec()

}
