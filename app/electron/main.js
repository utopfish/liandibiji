const {app, BrowserWindow, shell} = require('electron')

function createWindow () {
  // 创建浏览器窗口
  let mainWindow = new BrowserWindow({
    width: 1024,
    height: 768,
    webPreferences: {
      nodeIntegration: true,
    },
  })

  // 加载index.html文件
  mainWindow.loadFile('../public/index.html')

  // 组织当前页面链接跳转
  mainWindow.webContents.on('will-navigate', (e, url) => {
    e.preventDefault()
    console.log(e, url)
    shell.openExternal(url)
  })
}

app.on('ready', createWindow)
