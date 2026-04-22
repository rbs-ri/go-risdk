package proto

import "context"

// RI_SDK_InitSDK - Инициализация SDK
func (m *GRPCClient) RI_SDK_InitSDK(logLevel int64) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_InitSDK(context.Background(), &RI_SDK_InitSDKParams{
		LogLevel: logLevel,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Device_ModelList - Чтение списка доступных моделей
func (m *GRPCClient) RI_SDK_Device_ModelList(deviceType string) (modelList string, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Device_ModelList(context.Background(), &RI_SDK_Device_ModelListParams{
		DeviceType: deviceType,
	})
	if err != nil {
		return
	}
	return resp.ModelList, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_CreateBasic - Создание базового компонента
func (m *GRPCClient) RI_SDK_CreateBasic() (descriptor int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_CreateBasic(context.Background(), &RI_SDK_CreateBasicParams{})
	if err != nil {
		return
	}
	return resp.Descriptor_, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_CreateGroupComponent - функция создания компонента группы
func (m *GRPCClient) RI_SDK_CreateGroupComponent(group string) (descriptor int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_CreateGroupComponent(context.Background(), &RI_SDK_CreateGroupComponentParams{
		Group: group,
	})
	if err != nil {
		return
	}
	return resp.Descriptor_, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_CreateDeviceComponent - создание кмопнента устройства
func (m *GRPCClient) RI_SDK_CreateDeviceComponent(group, device string) (descriptor int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_CreateDeviceComponent(context.Background(), &RI_SDK_CreateDeviceComponentParams{
		Group:  group,
		Device: device,
	})
	if err != nil {
		return
	}
	return resp.Descriptor_, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_CreateModelComponent - создание кмопнента модели
func (m *GRPCClient) RI_SDK_CreateModelComponent(group, device, model string) (descriptor int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_CreateModelComponent(context.Background(), &RI_SDK_CreateModelComponentParams{
		Group:  group,
		Device: device,
		Model:  model,
	})
	if err != nil {
		return
	}
	return resp.Descriptor_, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_LinkPWMToController - связывает ШИМ с I2C адаптером.
func (m *GRPCClient) RI_SDK_LinkPWMToController(pwmDescriptor, controllerDescriptor int64, addr uint64) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_LinkPWMToController(context.Background(), &RI_SDK_LinkPWMToControllerParams{
		PwmDescriptor:        pwmDescriptor,
		ControllerDescriptor: controllerDescriptor,
		Addr:                 addr,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_LinkServodriveToController - связывает сервопривод с ШИМ.
func (m *GRPCClient) RI_SDK_LinkServodriveToController(servodriveDescriptor, pwmDescriptor, port int64) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_LinkServodriveToController(context.Background(), &RI_SDK_LinkServodriveToControllerParams{
		PwmDescriptor:        pwmDescriptor,
		ServodriveDescriptor: servodriveDescriptor,
		Port:                 port,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_LinkRServodriveToController - связывает сервопривод с ШИМ.
func (m *GRPCClient) RI_SDK_LinkRServodriveToController(servodriveDescriptor, pwmDescriptor, port int64) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_LinkRServodriveToController(context.Background(), &RI_SDK_LinkRServodriveToControllerParams{
		PwmDescriptor:        pwmDescriptor,
		ServodriveDescriptor: servodriveDescriptor,
		Port:                 port,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_LinkLedToController - связывает светодиод с ШИМ.
func (m *GRPCClient) RI_SDK_LinkLedToController(ledDescriptor, pwmDescriptor, rport, gport, bport int64) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_LinkLedToController(context.Background(), &RI_SDK_LinkLedToControllerParams{
		PwmDescriptor: pwmDescriptor,
		LedDescriptor: ledDescriptor,
		Rport:         rport,
		Gport:         gport,
		Bport:         bport,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_LinkVoltageSensorToController - связывает датчик тока с I2C адаптером.
func (m *GRPCClient) RI_SDK_LinkVoltageSensorToController(SensorDescriptor, I2CAdapterDescriptor int64, addr uint64) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_LinkVoltageSensorToController(context.Background(), &RI_SDK_LinkVoltageSensorToControllerParams{
		SensorDescriptor:     SensorDescriptor,
		I2CAdapterDescriptor: I2CAdapterDescriptor,
		Addr:                 addr,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_DestroyComponent - удаление компонента
func (m *GRPCClient) RI_SDK_DestroyComponent(descriptor int64) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_DestroyComponent(context.Background(), &RI_SDK_DestroyComponentParams{
		Descriptor_: descriptor,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_DestroySDK - удаление SDK
func (m *GRPCClient) RI_SDK_DestroySDK(isForce bool) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_DestroySDK(context.Background(), &RI_SDK_DestroySDKParams{
		IsForce: isForce,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_InitSDK - Инициализация SDK
func (m *GRPCServer) RI_SDK_InitSDK(
	ctx context.Context,
	req *RI_SDK_InitSDKParams) (*RI_SDK_InitSDKReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_InitSDK(req.LogLevel)
	return &RI_SDK_InitSDKReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Device_ModelList - Чтение списка доступных моделей
func (m *GRPCServer) RI_SDK_Device_ModelList(
	ctx context.Context,
	req *RI_SDK_Device_ModelListParams) (*RI_SDK_Device_ModelListReturn, error) {
	modelList, errorText, errorCode, err := m.Impl.RI_SDK_Device_ModelList(req.DeviceType)
	return &RI_SDK_Device_ModelListReturn{
		ModelList: modelList,
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_CreateBasic - Создание базового компонента
func (m *GRPCServer) RI_SDK_CreateBasic(
	ctx context.Context,
	req *RI_SDK_CreateBasicParams) (*RI_SDK_CreateBasicReturn, error) {
	descriptor, errorText, errorCode, err := m.Impl.RI_SDK_CreateBasic()
	return &RI_SDK_CreateBasicReturn{
		Descriptor_: descriptor,
		ErrorText:   errorText,
		ErrorCode:   errorCode,
	}, err
}

// RI_SDK_CreateGroupComponent - функция создания компонента группы
func (m *GRPCServer) RI_SDK_CreateGroupComponent(
	ctx context.Context,
	req *RI_SDK_CreateGroupComponentParams) (*RI_SDK_CreateGroupComponentReturn, error) {
	descriptor, errorText, errorCode, err := m.Impl.RI_SDK_CreateGroupComponent(req.Group)
	return &RI_SDK_CreateGroupComponentReturn{
		Descriptor_: descriptor,
		ErrorText:   errorText,
		ErrorCode:   errorCode,
	}, err
}

// RI_SDK_CreateDeviceComponent - создание кмопнента устройства
func (m *GRPCServer) RI_SDK_CreateDeviceComponent(
	ctx context.Context,
	req *RI_SDK_CreateDeviceComponentParams) (*RI_SDK_CreateDeviceComponentReturn, error) {
	descriptor, errorText, errorCode, err := m.Impl.RI_SDK_CreateDeviceComponent(req.Group, req.Device)
	return &RI_SDK_CreateDeviceComponentReturn{
		Descriptor_: descriptor,
		ErrorText:   errorText,
		ErrorCode:   errorCode,
	}, err
}

// RI_SDK_CreateModelComponent - создание кмопнента модели
func (m *GRPCServer) RI_SDK_CreateModelComponent(
	ctx context.Context,
	req *RI_SDK_CreateModelComponentParams) (*RI_SDK_CreateModelComponentReturn, error) {
	descriptor, errorText, errorCode, err := m.Impl.RI_SDK_CreateModelComponent(req.Group, req.Device, req.Model)
	return &RI_SDK_CreateModelComponentReturn{
		Descriptor_: descriptor,
		ErrorText:   errorText,
		ErrorCode:   errorCode,
	}, err
}

// RI_SDK_LinkPWMToController - связывает ШИМ с I2C адаптером.
func (m *GRPCServer) RI_SDK_LinkPWMToController(
	ctx context.Context,
	req *RI_SDK_LinkPWMToControllerParams) (*RI_SDK_LinkPWMToControllerReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_LinkPWMToController(req.PwmDescriptor, req.ControllerDescriptor, req.Addr)
	return &RI_SDK_LinkPWMToControllerReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_LinkServodriveToController - связывает сервопривод с ШИМ.
func (m *GRPCServer) RI_SDK_LinkServodriveToController(
	ctx context.Context,
	req *RI_SDK_LinkServodriveToControllerParams) (*RI_SDK_LinkServodriveToControllerReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_LinkServodriveToController(req.ServodriveDescriptor, req.PwmDescriptor, req.Port)
	return &RI_SDK_LinkServodriveToControllerReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_LinkRServodriveToController - связывает сервопривод с ШИМ.
func (m *GRPCServer) RI_SDK_LinkRServodriveToController(
	ctx context.Context,
	req *RI_SDK_LinkRServodriveToControllerParams) (*RI_SDK_LinkRServodriveToControllerReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_LinkRServodriveToController(req.ServodriveDescriptor, req.PwmDescriptor, req.Port)
	return &RI_SDK_LinkRServodriveToControllerReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_LinkLedToController - связывает светодиод с ШИМ.
func (m *GRPCServer) RI_SDK_LinkLedToController(
	ctx context.Context,
	req *RI_SDK_LinkLedToControllerParams) (*RI_SDK_LinkLedToControllerReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_LinkLedToController(req.LedDescriptor, req.PwmDescriptor, req.Rport, req.Gport, req.Bport)
	return &RI_SDK_LinkLedToControllerReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_LinkVoltageSensorToController - связывает датчик тока с I2C.
func (m *GRPCServer) RI_SDK_LinkVoltageSensorToController(
	ctx context.Context,
	req *RI_SDK_LinkVoltageSensorToControllerParams) (*RI_SDK_LinkVoltageSensorToControllerReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_LinkVoltageSensorToController(req.SensorDescriptor, req.I2CAdapterDescriptor, req.Addr)
	return &RI_SDK_LinkVoltageSensorToControllerReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_DestroyComponent - Удаление компонента
func (m *GRPCServer) RI_SDK_DestroyComponent(
	ctx context.Context,
	req *RI_SDK_DestroyComponentParams) (*RI_SDK_DestroyComponentReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_DestroyComponent(req.Descriptor_)
	return &RI_SDK_DestroyComponentReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_DestroySDK - Удаление SDK
func (m *GRPCServer) RI_SDK_DestroySDK(
	ctx context.Context,
	req *RI_SDK_DestroySDKParams) (*RI_SDK_DestroySDKReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_DestroySDK(req.IsForce)
	return &RI_SDK_DestroySDKReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}
