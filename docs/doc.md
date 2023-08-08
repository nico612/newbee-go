
### endless 平滑重启、热更新库
`go get -u github.com/fvbock/endless`

endless 是一个 Go 语言库，它用于在 HTTP 服务器中实现平滑重启和热更新的功能，无需停止正在运行的服务器即可更新代码。使用 endless 可以在不影响现有连接的情况下更新服务器，提供更好的用户体验和更高的可用性。

主要特点和用途：

平滑重启：endless 允许您在不中断现有连接的情况下重新启动 HTTP 服务器。当您更新服务器的代码或配置时，可以使用平滑重启，以便新代码和配置能够立即生效，而不会影响当前正在处理的请求。

热更新：通过 endless 实现热更新功能，您可以在运行中的服务器上动态地加载和更新代码，而无需停止服务器。这在开发过程中特别有用，可以避免频繁重启服务器，加快开发速度。

无需第三方负载均衡器：endless 通过在单个进程中启动多个 goroutine 来处理并发请求，而无需使用负载均衡器。这使得服务器的管理和维护更加简单。

使用 endless 可以让您的服务器始终在线，并在更新代码时不中断现有连接。您只需使用 endless.ListenAndServe 替换标准库的 http.ListenAndServe，即可实现平滑重启和热更新。
```go
package main

import (
    "fmt"
    "net/http"
    "time"
    "github.com/fvbock/endless"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, Endless!")
}

func main() {
    server := &http.Server{
        Addr:    ":8080",
        Handler: http.HandlerFunc(handler),
    }

	// 使用 endless.ListenAndServe 替换标准库的 http.ListenAndServe
	// 可以实现平滑重启和热更新的功能
	err := endless.ListenAndServe(":8080", server)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
```
在上述示例中，我们使用 endless.ListenAndServe 替代了标准库的 http.ListenAndServe 来启动 HTTP 服务器。这样，在更新服务器代码时，可以实现平滑重启和热更新的功能。

总结：endless 是一个 Go 语言库，用于实现 HTTP 服务器的平滑重启和热更新。通过使用 endless，您可以在不中断现有连接的情况下更新服务器代码，提供更好的用户体验和更高的可用性。

### Viper 配置文件读取
`go get -u github.com/spf13/viper`

```go
package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/nico612/newbee-goo/global"
	"github.com/spf13/viper"
)

func Viper(path string) *viper.Viper {

	v := viper.New()
	v.SetConfigFile(path) //设置文件路径
	v.SetConfigType("yaml")  //设置文件类型

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal err config file : %s \n", err))
	}

	v.WatchConfig() //监听配置文件变化
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed: ", e.Name)

		//	当配置发生变化时解析
		if err := v.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
		}
	})
    
	// 解析到Config结构体
	if err := v.Unmarshal(&global.Config); err != nil {
		fmt.Println(err)
	}

	return v
}

```









