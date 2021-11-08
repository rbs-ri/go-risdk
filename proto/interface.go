// Общие объекты для плагина и клиентского кода
package proto

import (
	"context"
	"plugin"

	"google.golang.org/grpc"

	"github.com/hashicorp/go-plugin"
)

// RoboSdkApi Интерфейс предоставляемый подключаемым плагином
type RoboSdkApi interface {
	RI_SDK_InitSDK(logLevel int64) (errorText string, errorCode int64, err error)
	RI_SDK_CreateBasic() (descriptor int64, errorText string, errorCode int64, err error)
	RI_SDK_DestroyComponent(descriptor int64) (errorText string, errorCode int64, err error)
	RI_SDK_DestroySDK(isForce bool) (errorText string, errorCode int64, err error)
	RI_SDK_CreateGroupComponent(group string) (descriptor int64, errorText string, errorCode int64, err error)
	RI_SDK_CreateDeviceComponent(group, device string) (descriptor int64, errorText string, errorCode int64, err error)
	RI_SDK_CreateModelComponent(group, device, model string) (descriptor int64, errorText string, errorCode int64, err error)
	RI_SDK_LinkServodriveToController(servodriveDescriptor, pwmDescriptor, port int64) (errorText string, errorCode int64, err error)
	RI_SDK_LinkLedToController(ledDescriptor, pwmDescriptor, rport, gport, bport int64) (errorText string, errorCode int64, err error)
	RI_SDK_LinkPWMToController(pwmDescriptor, controllerDescriptor int64, addr uint64) (errorText string, errorCode int64, err error)

	RI_SDK_Executor_Extend(basic int64) (desrciptor int64, errorText string, errorCode int64, err error)
	RI_SDK_Executor_State(desrciptor int64) (state int64, errorText string, errorCode int64, err error)

	RI_SDK_Connector_Extend(basic int64) (desrciptor int64, errorText string, errorCode int64, err error)

	RI_SDK_Connector_I2C_Open(descriptor int64, addr uint8) (errorText string, errorCode int64, err error)
	RI_SDK_Connector_I2C_Extend(connectorDescriptor int64) (descriptor int64, errorText string, errorCode int64, err error)
	RI_SDK_Connector_I2C_ExtendToModel(baseDescriptor int64, modelName string) (descriptor int64, errorText string, errorCode int64, err error)
	RI_SDK_Connector_I2C_State(descriptor int64) (state int, errorText string, errorCode int64, err error)
	RI_SDK_Connector_I2C_CloseAll(descriptor int64) (errorText string, errorCode int64, err error)
	RI_SDK_Connector_I2C_Close(descriptor int64, addr uint8) (errorText string, errorCode int64, err error)
	RI_SDK_Connector_I2C_WriteBytes(descriptor int64, addr uint8, buf []byte, length int) (wroteBytesLen int64, errorText string, errorCode int64, err error)
	RI_SDK_Connector_I2C_ReadBytes(descriptor int64, addr uint8, buf []byte) (readedBytesLen int64, errorText string, errorCode int64, err error)
	RI_SDK_Connector_I2C_WriteByte(descriptor int64, addr uint8, value byte) (errorText string, errorCode int64, err error)
	RI_SDK_Connector_I2C_ReadByte(descriptor int64, addr uint8) (value byte, errorText string, errorCode int64, err error)

	RI_SDK_Sigmod_PWM_ExtendToModel(baseDescriptor int64, modelName string) (descriptor int64, errorText string, errorCode int64, err error)
	RI_SDK_Sigmod_PWM_GetResolution(descriptor int64) (resolution int, errorText string, errorCode int64, err error)
	RI_SDK_Sigmod_PWM_GetFreq(descriptor int64) (freq int64, errorText string, errorCode int64, err error)
	RI_SDK_Sigmod_PWM_SetFreq(descriptor int64, freq int64) (errorText string, errorCode int64, err error)
	RI_SDK_Sigmod_PWM_WriteRegBytes(descriptor int64, reg byte, buf []byte) (wroteBytesLen int64, errorText string, errorCode int64, err error)
	RI_SDK_Sigmod_PWM_ReadRegBytes(descriptor int64, reg byte, buf []byte) (readedBytesLen int64, errorText string, errorCode int64, err error)
	RI_SDK_Sigmod_PWM_WriteByte(descriptor int64, reg byte, value byte) (errorText string, errorCode int64, err error)
	RI_SDK_Sigmod_PWM_ReadByte(descriptor int64, reg byte) (value byte, errorText string, errorCode int64, err error)
	RI_SDK_Sigmod_PWM_Extend(con int64) (desrciptor int64, errorText string, errorCode int64, err error)
	RI_SDK_Sigmod_PWM_GetPortFreq(descriptor int64, port int64) (freq int64, errorText string, errorCode int64, err error)
	RI_SDK_Sigmod_PWM_SetPortFreq(descriptor int64, port int64, freq int64) (errorText string, errorCode int64, err error)
	RI_SDK_Sigmod_PWM_ResetAll(descriptor int64) (errorText string, errorCode int64, err error)
	RI_SDK_Sigmod_PWM_ResetPort(descriptor int64, port int64) (errorText string, errorCode int64, err error)
	RI_SDK_Sigmod_PWM_GetPortDutyCycle(descriptor int64, port int64) (on int64, off int64, errorText string, errorCode int64, err error)
	RI_SDK_Sigmod_PWM_SetPortDutyCycle(descriptor int64, port int64, on int64, off int64) (errorText string, errorCode int64, err error)
	RI_SDK_Sigmod_PWM_Close(descriptor int64) (errorText string, errorCode int64, err error)

	RI_SDK_Exec_ServoDrive_Extend(exec int64) (descriptor int64, errorText string, errorCode int64, err error)
	RI_SDK_Exec_ServoDrive_ExtendToModel(baseDescriptor int64, modelName string) (descriptor int64, errorText string, errorCode int64, err error)
	RI_SDK_Exec_ServoDrive_CustomDeviceInit(desrciptor, maxImpulse, maxSpeed, pulseRange int64) (errorText string, errorCode int64, err error)
	RI_SDK_Exec_ServoDrive_TurnByDutyCycle(desrciptor, steps int64) (errorText string, errorCode int64, err error)
	RI_SDK_Exec_ServoDrive_TurnByPulse(desrciptor, pulse int64) (errorText string, errorCode int64, err error)
	RI_SDK_Exec_ServoDrive_GetCurrentAngle(desrciptor int64) (angle int64, errorText string, errorCode int64, err error)
	RI_SDK_Exec_ServoDrive_GetState(desrciptor int64) (state int64, errorText string, errorCode int64, err error)
	RI_SDK_Exec_ServoDrive_MinStepRotate(descriptor, direction, speed int64, async bool) (errorText string, errorCode int64, err error)
	RI_SDK_Exec_ServoDrive_Turn(descriptor, angle, speed int64, async bool) (errorText string, errorCode int64, err error)
	RI_SDK_Exec_ServoDrive_Stop(descriptor int64) (errorText string, errorCode int64, err error)
	RI_SDK_Exec_ServoDrive_Rotate(descriptor, direction, speed int64, async bool) (errorText string, errorCode int64, err error)
	RI_SDK_Exec_ServoDrive_RotateWithRelativeSpeed(descriptor, angle, speed int64, async bool) (errorText string, errorCode int64, err error)
	RI_SDK_Exec_ServoDrive_TurnWithRelativeSpeed(descriptor, angle, speed int64, async bool) (errorText string, errorCode int64, err error)

	RI_SDK_Exec_RGB_LED_Extend(exec int64) (descriptor int64, errorText string, errorCode int64, err error)
	RI_SDK_Exec_RGB_LED_ExtendToModel(baseDescriptor int64, modelName string) (descriptor int64, errorText string, errorCode int64, err error)
	RI_SDK_Exec_RGB_LED_SinglePulse(descriptor, r, g, b, duration int64, async bool) (errorText string, errorCode int64, err error)
	RI_SDK_Exec_RGB_LED_Stop(descriptor int64) (errorText string, errorCode int64, err error)
	RI_SDK_Exec_RGB_LED_GetState(descriptor int64) (state int64, errorText string, errorCode int64, err error)
	RI_SDK_Exec_RGB_LED_GetColor(descriptor int64) (r, g, b int64, errorText string, errorCode int64, err error)
	RI_SDK_Exec_RGB_LED_FlashingWithFrequency(descriptor, r, g, b, frequency, qty int64, async bool) (errorText string, errorCode int64, err error)
	RI_SDK_Exec_RGB_LED_FlashingWithPause(descriptor, r, g, b, duration, pause, qty int64, async bool) (errorText string, errorCode int64, err error)
	RI_SDK_Exec_RGB_LED_Flicker(descriptor, r, g, b, duration, qty int64, async bool) (errorText string, errorCode int64, err error)
}

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"robosdk_grpc": &RoboSdkGRPCPlugin{},
}

// This is the implementation of plugin.GRPCPlugin so we can serve/consume this.
type RoboSdkGRPCPlugin struct {
	// GRPCPlugin must still implement the Plugin interface
	plugin.Plugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl RoboSdkApi
}

func (p *RoboSdkGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	RegisterRoboSdkServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *RoboSdkGRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: NewRoboSdkClient(c)}, nil
}
