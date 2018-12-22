## beego安装
     安装 GO 设置GOROOT GOPATH
     
     go get github.com/astaxie/beego
     
     go get github.com/beego/bee
     安装完后 在$gopath/bin路径下面会产生bee工具二进制文件
     将$gopath/bin添加到path环境变量中 才能使用bee命令
     
     
     
## bee 创建restful api

    bee api apiproject
    
    运行
    bee run -gendoc=true -downdoc=true
 