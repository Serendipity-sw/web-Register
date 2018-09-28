# web-Register
前端寄存器 为了解决一些不会使用不会使用域名映射的场景

## 使用说明
    项目已经打包,三个带有系统名称的文件为对应操作系统的可执行文件. 程序会在当前目录下创建
    logs文件夹.其中存放了所有程序的日志. templates 为模板文件夹  内部请自行创建各个前端的文件夹名称 深度自便
    
    config.json 为配置文件
        serverListeningPort为程序需要监听的端口号
        logsDir 表明程序日志打印的文件夹, 不建议更改
        forwardingDomain 需要转发域名
        rootPrefix 程序运行的二级目录地址
        
###模板请求样例
    如目前文件夹已经存在样例
    http://localhost:8080/templates/webst/index.html
    html中的文件引用请使用相对路径
    
    需要转发的接口
    http://localhost:8080/api/gloomy/gloomysw.api
    参数请自行添加 url 传参请使用get提交   post 表单传参请使用post请求
    改地址进入后会转化为 forwardingDomain(配置的转发域名) + /gloomy/gloomysw.api
    所有需要转发的cookie会一并转发过去,转发接口返回的cookie会写到当前客户端域名下,以便后续请求可以正确携带cookie
    
###接口返回值说明(寄存器返回对象)
    {
        Code string 返回码(该返回码为寄存器返回)
        Data string 转发接口返回对象json字符串
    }
    
    Code 对照
    00000 成功
    00001 http转发构建失败
    00002 http转发请求失败
    00003 http转发请求返回值读取失败

        
###接口支持POST 及 get 转发请求 