import QtQuick 2.9
import QtQuick.Window 2.9
import QtQuick.Controls 2.9
import QtWebEngine 1.6
import Goxui 1.0

Window {

    id: mainWindow
    visible: true
    width: 880
    height: 640



    WebEngineView {
        url: "https://bing.com"
        y: 50
        width: parent.width
        height: parent.height - y
    }

}
