package risdk

import (
	"fmt"
	"github.com/hashicorp/go-plugin"
	"github.com/rbs-ri/go-risdk/proto"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"sync"
	"syscall"
)

// ClientRPC - объект для работы с API SDK по RPC
var clientRPC *ClientRPC
var mu sync.Mutex

// GetClientRPC - получаем объект для работы с API SDK по RPC
func GetClientRPC() *ClientRPC {
	mu.Lock()
	defer mu.Unlock()
	if clientRPC == nil {
		client, roboSdkApi := initClient()
		clientRPC = &ClientRPC{
			Client:     client,
			RoboSdkApi: roboSdkApi,
		}
	}

	// Закрытие соединение с RPC, когда выполнение программы было прекращено, до ее завершения.
	// Здесь отслеживается системное событие завершения программы.
	// Если не убивать килент, то процесс остается и может блокировать подачу сигнала на устройство.
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		<-c
		clientRPC.Client.Kill()
		os.Exit(0)
	}()

	return clientRPC
}

// clientRPC - Стуктура для для работы с API SDK по RPC
type ClientRPC struct {
	Client     *plugin.Client   //клиент
	RoboSdkApi proto.RoboSdkApi //
}

// initClient - Создает клиента для получения доступа к API SDK по RPC
func initClient() (client *plugin.Client, roboSdkApi proto.RoboSdkApi) {
	// We don't want to see the plugin logs.
	log.SetOutput(ioutil.Discard)

	// Determine executable path
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Current path: ", path)

	// Determine os AND select risdk filename
	var risdkName string
	switch runtime.GOOS {
	case "windows":
		risdkName = "risdk.exe"
	case "linux":
		risdkName = "risdk.bin"
	default:
		fmt.Println("Incompatible OS (Support windows, linux only)")
		os.Exit(1)
	}

	// We're a host. Start by launching the plugin process.
	client = plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: proto.Handshake,
		Plugins:         proto.PluginMap,
		Cmd:             exec.Command(filepath.Join(path, risdkName)),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC, plugin.ProtocolGRPC},
	})

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("robosdk_grpc")
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	roboSdkApi = raw.(proto.RoboSdkApi)
	return
}
