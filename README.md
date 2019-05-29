# http-file-upload-download
代号 `mushroom` ，主要目的为了“搞事情”，顺便学习下相关知识点（文件上传下载）

## 项目结构
```
├─asset 静态文件(主要为vue、element-ui)
│
├─cmd main.go
│
├─internal 打包成 `.go` 的相关文件以及部分内部函数和类型
│ ├─asset 打包成 `.go` 的静态文件
│ ├─template 打包成 `.go` 的模版文件
│
├─mushroom 保存上传文件的文件夹
│
├─script 打包和交叉编译的脚本
│
├─template 页面的模版文件
```
