package proto

import (
	"context"
)

// ─── GRPCClient — клиентская сторона (вызывается из demo-программ) ───────────

// RI_SDK_LinkGPIOI2CAdapterToController связывает GPIO-расширитель по I2C с I2C-адаптером.
func (m *GRPCClient) RI_SDK_LinkGPIOI2CAdapterToController(gpioI2CDescriptor, i2cAdapterDescriptor int64, addr uint64) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_LinkGPIOI2CAdapterToController(context.Background(), &RI_SDK_LinkGPIOI2CAdapterToControllerParams{
		GpioI2CDescriptor:    gpioI2CDescriptor,
		I2CAdapterDescriptor: i2cAdapterDescriptor,
		Addr:                 addr,
	})
	if err != nil {
		return
	}
	errorText = resp.ErrorText
	errorCode = resp.ErrorCode
	return
}

// RI_SDK_LinkGPIODeviceToController связывает GPIO-устройство с GPIO-адаптером.
func (m *GRPCClient) RI_SDK_LinkGPIODeviceToController(deviceDescriptor, gpioAdapterDescriptor int64) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_LinkGPIODeviceToController(context.Background(), &RI_SDK_LinkGPIODeviceToControllerParams{
		DeviceDescriptor:      deviceDescriptor,
		GpioAdapterDescriptor: gpioAdapterDescriptor,
	})
	if err != nil {
		return
	}
	errorText = resp.ErrorText
	errorCode = resp.ErrorCode
	return
}

// RI_SDK_Connector_GPIO_I2C_Extend расширяет коннектор до I2C GPIO-адаптера.
func (m *GRPCClient) RI_SDK_Connector_GPIO_I2C_Extend(connectorDescriptor int64) (descriptor int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Connector_GPIO_I2C_Extend(context.Background(), &RI_SDK_Connector_GPIO_I2C_ExtendParams{
		ConnectorDescriptor: connectorDescriptor,
	})
	if err != nil {
		return
	}
	descriptor = resp.Descriptor_
	errorText = resp.ErrorText
	errorCode = resp.ErrorCode
	return
}

// RI_SDK_Connector_GPIO_I2C_ExtendToModel расширяет I2C GPIO-адаптер до конкретной модели.
func (m *GRPCClient) RI_SDK_Connector_GPIO_I2C_ExtendToModel(baseDescriptor int64, modelName string) (descriptor int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Connector_GPIO_I2C_ExtendToModel(context.Background(), &RI_SDK_Connector_GPIO_I2C_ExtendToModelParams{
		BaseDescriptor: baseDescriptor,
		ModelName:      modelName,
	})
	if err != nil {
		return
	}
	descriptor = resp.Descriptor_
	errorText = resp.ErrorText
	errorCode = resp.ErrorCode
	return
}

// RI_SDK_Connector_GPIO_I2C_SetAddr устанавливает I2C-адрес GPIO-расширителя.
func (m *GRPCClient) RI_SDK_Connector_GPIO_I2C_SetAddr(descriptor int64, addr uint64) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Connector_GPIO_I2C_SetAddr(context.Background(), &RI_SDK_Connector_GPIO_I2C_SetAddrParams{
		Descriptor_: descriptor,
		Addr:        addr,
	})
	if err != nil {
		return
	}
	errorText = resp.ErrorText
	errorCode = resp.ErrorCode
	return
}

// RI_SDK_Sensor_Extend расширяет базовый компонент до группы датчиков.
func (m *GRPCClient) RI_SDK_Sensor_Extend(basic int64) (descriptor int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sensor_Extend(context.Background(), &RI_SDK_Sensor_ExtendParams{
		Basic: basic,
	})
	if err != nil {
		return
	}
	descriptor = resp.Descriptor_
	errorText = resp.ErrorText
	errorCode = resp.ErrorCode
	return
}

// RI_SDK_Sensor_DistanceSensor_Extend расширяет датчик до DistanceSensor.
func (m *GRPCClient) RI_SDK_Sensor_DistanceSensor_Extend(exec int64) (descriptor int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sensor_DistanceSensor_Extend(context.Background(), &RI_SDK_Sensor_DistanceSensor_ExtendParams{
		Exec: exec,
	})
	if err != nil {
		return
	}
	descriptor = resp.Descriptor_
	errorText = resp.ErrorText
	errorCode = resp.ErrorCode
	return
}

// RI_SDK_Sensor_DistanceSensor_ExtendToModel расширяет DistanceSensor до конкретной модели.
func (m *GRPCClient) RI_SDK_Sensor_DistanceSensor_ExtendToModel(baseDescriptor int64, modelName string) (descriptor int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sensor_DistanceSensor_ExtendToModel(context.Background(), &RI_SDK_Sensor_DistanceSensor_ExtendToModelParams{
		BaseDescriptor: baseDescriptor,
		ModelName:      modelName,
	})
	if err != nil {
		return
	}
	descriptor = resp.Descriptor_
	errorText = resp.ErrorText
	errorCode = resp.ErrorCode
	return
}

// RI_SDK_Sensor_DistanceSensor_SetPins задаёт пины TRIG и ECHO.
func (m *GRPCClient) RI_SDK_Sensor_DistanceSensor_SetPins(descriptor, trigPin, echoPin int64) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sensor_DistanceSensor_SetPins(context.Background(), &RI_SDK_Sensor_DistanceSensor_SetPinsParams{
		Descriptor_: descriptor,
		TrigPin:     trigPin,
		EchoPin:     echoPin,
	})
	if err != nil {
		return
	}
	errorText = resp.ErrorText
	errorCode = resp.ErrorCode
	return
}

// RI_SDK_Sensor_DistanceSensor_Distance возвращает расстояние в сантиметрах.
func (m *GRPCClient) RI_SDK_Sensor_DistanceSensor_Distance(descriptor int64) (dist float64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sensor_DistanceSensor_Distance(context.Background(), &RI_SDK_Sensor_DistanceSensor_DistanceParams{
		Descriptor_: descriptor,
	})
	if err != nil {
		return
	}
	dist = resp.Dist
	errorText = resp.ErrorText
	errorCode = resp.ErrorCode
	return
}

// ─── GRPCServer — серверная сторона (реализуется плагином) ───────────────────

// RI_SDK_LinkGPIOI2CAdapterToController серверная сторона.
func (m *GRPCServer) RI_SDK_LinkGPIOI2CAdapterToController(ctx context.Context, req *RI_SDK_LinkGPIOI2CAdapterToControllerParams) (*RI_SDK_LinkGPIOI2CAdapterToControllerReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_LinkGPIOI2CAdapterToController(req.GpioI2CDescriptor, req.I2CAdapterDescriptor, req.Addr)
	return &RI_SDK_LinkGPIOI2CAdapterToControllerReturn{ErrorText: errorText, ErrorCode: errorCode}, err
}

// RI_SDK_LinkGPIODeviceToController серверная сторона.
func (m *GRPCServer) RI_SDK_LinkGPIODeviceToController(ctx context.Context, req *RI_SDK_LinkGPIODeviceToControllerParams) (*RI_SDK_LinkGPIODeviceToControllerReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_LinkGPIODeviceToController(req.DeviceDescriptor, req.GpioAdapterDescriptor)
	return &RI_SDK_LinkGPIODeviceToControllerReturn{ErrorText: errorText, ErrorCode: errorCode}, err
}

// RI_SDK_Connector_GPIO_I2C_Extend серверная сторона.
func (m *GRPCServer) RI_SDK_Connector_GPIO_I2C_Extend(ctx context.Context, req *RI_SDK_Connector_GPIO_I2C_ExtendParams) (*RI_SDK_Connector_GPIO_I2C_ExtendReturn, error) {
	descriptor, errorText, errorCode, err := m.Impl.RI_SDK_Connector_GPIO_I2C_Extend(req.ConnectorDescriptor)
	return &RI_SDK_Connector_GPIO_I2C_ExtendReturn{Descriptor_: descriptor, ErrorText: errorText, ErrorCode: errorCode}, err
}

// RI_SDK_Connector_GPIO_I2C_ExtendToModel серверная сторона.
func (m *GRPCServer) RI_SDK_Connector_GPIO_I2C_ExtendToModel(ctx context.Context, req *RI_SDK_Connector_GPIO_I2C_ExtendToModelParams) (*RI_SDK_Connector_GPIO_I2C_ExtendToModelReturn, error) {
	descriptor, errorText, errorCode, err := m.Impl.RI_SDK_Connector_GPIO_I2C_ExtendToModel(req.BaseDescriptor, req.ModelName)
	return &RI_SDK_Connector_GPIO_I2C_ExtendToModelReturn{Descriptor_: descriptor, ErrorText: errorText, ErrorCode: errorCode}, err
}

// RI_SDK_Connector_GPIO_I2C_SetAddr серверная сторона.
func (m *GRPCServer) RI_SDK_Connector_GPIO_I2C_SetAddr(ctx context.Context, req *RI_SDK_Connector_GPIO_I2C_SetAddrParams) (*RI_SDK_Connector_GPIO_I2C_SetAddrReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Connector_GPIO_I2C_SetAddr(req.Descriptor_, req.Addr)
	return &RI_SDK_Connector_GPIO_I2C_SetAddrReturn{ErrorText: errorText, ErrorCode: errorCode}, err
}

// RI_SDK_Sensor_Extend серверная сторона.
func (m *GRPCServer) RI_SDK_Sensor_Extend(ctx context.Context, req *RI_SDK_Sensor_ExtendParams) (*RI_SDK_Sensor_ExtendReturn, error) {
	descriptor, errorText, errorCode, err := m.Impl.RI_SDK_Sensor_Extend(req.Basic)
	return &RI_SDK_Sensor_ExtendReturn{Descriptor_: descriptor, ErrorText: errorText, ErrorCode: errorCode}, err
}

// RI_SDK_Sensor_DistanceSensor_Extend серверная сторона.
func (m *GRPCServer) RI_SDK_Sensor_DistanceSensor_Extend(ctx context.Context, req *RI_SDK_Sensor_DistanceSensor_ExtendParams) (*RI_SDK_Sensor_DistanceSensor_ExtendReturn, error) {
	descriptor, errorText, errorCode, err := m.Impl.RI_SDK_Sensor_DistanceSensor_Extend(req.Exec)
	return &RI_SDK_Sensor_DistanceSensor_ExtendReturn{Descriptor_: descriptor, ErrorText: errorText, ErrorCode: errorCode}, err
}

// RI_SDK_Sensor_DistanceSensor_ExtendToModel серверная сторона.
func (m *GRPCServer) RI_SDK_Sensor_DistanceSensor_ExtendToModel(ctx context.Context, req *RI_SDK_Sensor_DistanceSensor_ExtendToModelParams) (*RI_SDK_Sensor_DistanceSensor_ExtendToModelReturn, error) {
	descriptor, errorText, errorCode, err := m.Impl.RI_SDK_Sensor_DistanceSensor_ExtendToModel(req.BaseDescriptor, req.ModelName)
	return &RI_SDK_Sensor_DistanceSensor_ExtendToModelReturn{Descriptor_: descriptor, ErrorText: errorText, ErrorCode: errorCode}, err
}

// RI_SDK_Sensor_DistanceSensor_SetPins серверная сторона.
func (m *GRPCServer) RI_SDK_Sensor_DistanceSensor_SetPins(ctx context.Context, req *RI_SDK_Sensor_DistanceSensor_SetPinsParams) (*RI_SDK_Sensor_DistanceSensor_SetPinsReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Sensor_DistanceSensor_SetPins(req.Descriptor_, req.TrigPin, req.EchoPin)
	return &RI_SDK_Sensor_DistanceSensor_SetPinsReturn{ErrorText: errorText, ErrorCode: errorCode}, err
}

// RI_SDK_Sensor_DistanceSensor_Distance серверная сторона.
func (m *GRPCServer) RI_SDK_Sensor_DistanceSensor_Distance(ctx context.Context, req *RI_SDK_Sensor_DistanceSensor_DistanceParams) (*RI_SDK_Sensor_DistanceSensor_DistanceReturn, error) {
	dist, errorText, errorCode, err := m.Impl.RI_SDK_Sensor_DistanceSensor_Distance(req.Descriptor_)
	return &RI_SDK_Sensor_DistanceSensor_DistanceReturn{Dist: dist, ErrorText: errorText, ErrorCode: errorCode}, err
}
