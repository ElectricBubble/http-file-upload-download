package main

import (
	"bytes"
	"flag"
	"http-file-upload-download/internal"
	"http-file-upload-download/internal/asset"
	"http-file-upload-download/internal/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

var mushroom = "mushroom"
var defaultDir = ""
var mrStr = ""

func main() {
	wd := flag.String("wd", *internal.WorkingDirectory, "指定当前工作路径")
	flag.Parse()
	if *wd != "" {
		internal.WorkingDirectory = wd
	}
	// log.Println(*internal.WorkingDirectory)

	// 🍄默认路径
	defaultDir = filepath.Join(*internal.WorkingDirectory, mushroom)
	// log.Println("defaultDir =", defaultDir)

	// 🌌静态文件(Vue.js、Element UI)
	http.Handle("/asset/", http.FileServer(asset.FS(false)))

	// 📄上传页面
	http.HandleFunc("/up", uploadPageHandler)
	// 🕹处理文件上传的操作，只处理POST
	http.HandleFunc("/upload/", uploadHandler)
	// 🕹处理文件下载的操作
	http.HandleFunc("/download/", downloadHandler)
	// 🕹飞鸽传书
	http.HandleFunc("/text", textHandler)

	// 📄主页 http://127.0.0.1:8080/mushroom
	http.HandleFunc("/"+mushroom+"/", mushroomHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err)
	}
}

func mushroomHandler(writer http.ResponseWriter, request *http.Request) {
	// fmt.Println("当前请求URI:", request.RequestURI)
	nowDir := request.RequestURI[len(mushroom)+2:]
	relativeDir := nowDir
	// fmt.Println("当前的相对路径:", relativeDir)
	if nowDir == "" {
		nowDir = defaultDir
	} else {
		nowDir = filepath.Join(defaultDir, nowDir)
	}
	// fmt.Println("当前文件夹路径:", nowDir)

	var fsItems []internal.FSItem
	if dir, err := os.Open(nowDir); err == nil {
		if fis, err := dir.Readdir(0); err == nil {
			for _, fi := range fis {
				// fmt.Println("文件的相对路径", path.Join(relativeDir, fi.Name()))
				fsItems = append(fsItems, internal.FSItem{
					Name:  fi.Name(),
					IsDir: fi.IsDir(),
					// 使用URL中的 `/` , 避免js的转义报错, 且windows也支持该分隔符
					Filename: path.Join(relativeDir, fi.Name()),
					Size:     strconv.FormatFloat(float64(fi.Size())/1024.0/1024.0, 'f', 2, 64) + " MB",
				})
			}
		}
	}
	// log.Println("当前文件夹的全部文件", fsItems)

	buffer := new(bytes.Buffer)
	template.PageMushroom(fsItems, buffer)
	if _, err := writer.Write(buffer.Bytes()); err != nil {
		log.Println(err)
	}
}

func downloadHandler(writer http.ResponseWriter, request *http.Request) {
	// fmt.Println("下载请求开始⬇️")
	// fmt.Println("当前请求的URI =", request.RequestURI)
	relativeDir := request.RequestURI[len("/download/"):]
	relativeDir, _ = url.QueryUnescape(relativeDir)
	// 因为使用URL的 `/` 所以使用path包
	name := path.Base(relativeDir)
	// fmt.Println("文件名 =", name)
	// fmt.Println("该文件的相对路径 =", relativeDir)
	fullpath := filepath.Join(defaultDir, relativeDir)
	// fmt.Println("该文件的绝对路径 =", fullpath)
	if info, err := os.Stat(fullpath); err == nil {
		writer.Header().Set("Content-type", "application/octet-stream")
		writer.Header().Set("Content-disposition", "attachment; filename="+name)
		writer.Header().Set("Content-Length", strconv.Itoa(int(info.Size())))
		http.ServeFile(writer, request, fullpath)
	}
}

func uploadPageHandler(writer http.ResponseWriter, request *http.Request) {
	buffer := new(bytes.Buffer)
	template.PageUpload(buffer)
	if _, err := writer.Write(buffer.Bytes()); err != nil {
		log.Println(err)
	}
}

func uploadHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		_ = request.ParseMultipartForm(10 << 20)
		for _, v := range request.MultipartForm.File {
			for i := 0; i < len(v); i++ {
				f, _ := v[i].Open()
				all, _ := ioutil.ReadAll(f)
				newPath := filepath.Join(defaultDir, v[i].Filename)
				// 忽略创建文件错误的可能性
				newFile, _ := os.Create(newPath)
				if _, err := newFile.Write(all); err != nil || newFile.Close() != nil {
					writer.WriteHeader(http.StatusInternalServerError)
					_, _ = writer.Write([]byte("文件上传写入失败。" + err.Error()))
					if newFile != nil {
						_ = newFile.Close()
					}
					return
				}
				_ = f.Close()
				_ = newFile.Close()
			}
		}
		_, _ = writer.Write([]byte("上传成功。"))
	}
}

func textHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		all, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Println("text -> ", err)
		}
		mrStr = string(all)
		sr := func(oStr, nStr string) {
			mrStr = strings.Replace(mrStr, oStr, nStr, -1)
		}
		sr(`\`, `\\`)
		sr(`'`, `\'`)
		sr("\n", `\n`)
	}

	buffer := new(bytes.Buffer)
	template.PageText(mrStr, buffer)
	if _, err := writer.Write(buffer.Bytes()); err != nil {
		log.Println(err)
	}
}
