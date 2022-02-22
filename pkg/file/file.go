package file

import (
	"goblog/pkg/logger"
	"io"
	"net/http"
	"os"
	"path"
)

/**
 * 保存文件的方法
 */
func SaveFile(parameterName string, setPath string, r *http.Request) string {

	/*通过*http.Request 获取上传的文件*/
	file, handler, err := r.FormFile(parameterName)
	if err != nil {
		logger.LogError(err)
		return ""
	}
	defer file.Close()

	savePath := path.Join("./public/upload/", setPath)                            //定义文件存储目录
	saveFileNamePath := path.Join(savePath, handler.Filename)                     //定义文件存储
	os.MkdirAll(savePath, os.ModePerm)                                            //创建目录为了防止没有目录报错
	f, _ := os.OpenFile(saveFileNamePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777) //打开文件如果没有就创建

	if err != nil {
		logger.LogError(err)
		return ""
	}
	defer f.Close()

	io.Copy(f, file)
	return saveFileNamePath
}
