package risdk

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"sync"
	"syscall"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/rbs-ri/go-risdk/proto"
)

// ClientRPC - объект для работы с API SDK по RPC
var clientRPC *ClientRPC
var mu sync.Mutex

type controlSignal int

const Stop controlSignal = 0

/*
	GenerateClientRPC - генератор уникального клиента.
	SDK находится в директории исполнения программы
*/
func GenerateClientRPC(path string) *ClientRPC {

	client, roboSdkApi := initClientWithPath(path)
	newClientRPC := &ClientRPC{
		Client:        client,
		RoboSdkApi:    roboSdkApi,
		controlSignal: make(chan controlSignal),
	}

	killAfterCloseSignal(newClientRPC)

	return newClientRPC
}

func GetClientRPC() *ClientRPC {
	mu.Lock()
	defer mu.Unlock()
	if clientRPC == nil {
		client, roboSdkApi := initClient()
		clientRPC = &ClientRPC{
			Client:        client,
			RoboSdkApi:    roboSdkApi,
			controlSignal: make(chan controlSignal),
		}

		killAfterCloseSignalGlobal()

	}

	return clientRPC
}

/*
	GetClientRPCWithPath - получаем объект для работы с API SDK по RPC.
	SDK находится в указанной директории
*/
func GetClientRPCWithPath(path string) *ClientRPC {
	mu.Lock()
	defer mu.Unlock()
	if clientRPC == nil {
		client, roboSdkApi := initClientWithPath(path)
		clientRPC = &ClientRPC{
			Client:        client,
			RoboSdkApi:    roboSdkApi,
			controlSignal: make(chan controlSignal),
		}

		killAfterCloseSignalGlobal()
	}

	return clientRPC
}

// clientRPC - Стуктура для для работы с API SDK по RPC
type ClientRPC struct {
	Client        *plugin.Client   //клиент
	RoboSdkApi    proto.RoboSdkApi //
	controlSignal chan controlSignal
}

/*
	StopClient - остановка работы клиента.
*/
func StopClient(client *ClientRPC) {
	client.controlSignal <- Stop

	return
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

	client, roboSdkApi = getClient(path, risdkName)
	return
}

// initClientWithPath - Создает клиента для получения доступа к API SDK по RPC, SDK находится в указанной директории
func initClientWithPath(path string) (client *plugin.Client, roboSdkApi proto.RoboSdkApi) {
	// We don't want to see the plugin logs.
	log.SetOutput(ioutil.Discard)

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

	client, roboSdkApi = getClient(path, risdkName)
	return
}

func getClient(path, risdkName string) (client *plugin.Client, roboSdkApi proto.RoboSdkApi) {
	// We're a host. Start by launching the plugin process.
	client = plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: proto.Handshake,
		Plugins:         proto.PluginMap,
		Cmd:             exec.Command(filepath.Join(path, risdkName)),
		Logger: hclog.New(&hclog.LoggerOptions{
			Output: ioutil.Discard,
			Level:  hclog.Trace,
			Name:   "plugin",
		}),
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

	return client, roboSdkApi
}

func killAfterCloseSignalGlobal() {
	// Закрытие соединение с RPC, когда выполнение программы было прекращено, до ее завершения.
	// Здесь отслеживается системное событие завершения программы.
	// Если не убивать клиент, то процесс остается и может блокировать подачу сигнала на устройство.
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		<-c
		select {
		case <-c:
			clientRPC.Client.Kill()
			os.Exit(0)
		case s := <-clientRPC.controlSignal:
			if s == Stop {
				clientRPC.Client.Kill()
				return
			}
		}
	}()
}

func killAfterCloseSignal(client *ClientRPC) {

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	go func(client *ClientRPC) {
		select {
		case <-c:
			client.Client.Kill()
			os.Exit(0)
		case <-client.controlSignal:
			return
		}

	}(client)
}
