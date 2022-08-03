package proto

import "context"

// RI_SDK_Sensor_VoltageSensor_Extend -  расширяет датчик до датчика тока.
func (m *GRPCClient) RI_SDK_Sensor_VoltageSensor_Extend(exec int64) (descriptor int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sensor_VoltageSensor_Extend(context.Background(), &RI_SDK_Sensor_VoltageSensor_ExtendParams{
		Exec: exec,
	})
	if err != nil {
		return
	}
	return resp.Descriptor_, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Sensor_VoltageSensor_ExtendToModel - расширяет датчик тока до модели
func (m *GRPCClient) RI_SDK_Sensor_VoltageSensor_ExtendToModel(baseDescriptor int64, modelName string) (descriptor int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sensor_VoltageSensor_ExtendToModel(context.Background(), &RI_SDK_Sensor_VoltageSensor_ExtendToModelParams{
		BaseDescriptor: baseDescriptor,
		ModelName:      modelName,
	})
	if err != nil {
		return
	}
	return resp.Descriptor_, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Sensor_VoltageSensor_CustomDeviceInit - Инициализация кастомного датчика тока
func (m *GRPCClient) RI_SDK_Sensor_VoltageSensor_CustomDeviceInit(desrciptor int64, lsbBus, lsbShunt, shuntResist float64, regVoltageShunt, regVoltageBus byte) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sensor_VoltageSensor_CustomDeviceInit(context.Background(), &RI_SDK_Sensor_VoltageSensor_CustomDeviceInitParams{
		Desrciptor:      desrciptor,
		LsbBus:          lsbBus,
		LsbShunt:        lsbShunt,
		ShuntResist:     shuntResist,
		RegVoltageShunt: []byte{regVoltageShunt},
		RegVoltageBus:   []byte{regVoltageBus},
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Sensor_VoltageSensor_Voltage - Возвтращает значение напряжения в цепи.
func (m *GRPCClient) RI_SDK_Sensor_VoltageSensor_Voltage(descriptor int64) (voltage float32, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sensor_VoltageSensor_Voltage(context.Background(), &RI_SDK_Sensor_VoltageSensor_VoltageParams{
		Desrciptor: descriptor,
	})
	if err != nil {
		return
	}
	return resp.Result, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Sensor_VoltageSensor_VoltageShunt - Возвращает значение напряжения на шунте.
func (m *GRPCClient) RI_SDK_Sensor_VoltageSensor_VoltageShunt(descriptor int64) (voltageShunt float32, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sensor_VoltageSensor_VoltageShunt(context.Background(), &RI_SDK_Sensor_VoltageSensor_VoltageShuntParams{
		Desrciptor: descriptor,
	})
	if err != nil {
		return
	}
	return resp.Result, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Sensor_VoltageSensor_Current - Возвращает значение силы тока.
func (m *GRPCClient) RI_SDK_Sensor_VoltageSensor_Current(descriptor int64) (current float32, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sensor_VoltageSensor_Current(context.Background(), &RI_SDK_Sensor_VoltageSensor_CurrentParams{
		Desrciptor: descriptor,
	})
	if err != nil {
		return
	}
	return resp.Result, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Sensor_VoltageSensor_Power - Возвращает значение мощьности в цепи
func (m *GRPCClient) RI_SDK_Sensor_VoltageSensor_Power(descriptor int64) (power float32, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sensor_VoltageSensor_Power(context.Background(), &RI_SDK_Sensor_VoltageSensor_PowerParams{
		Desrciptor: descriptor,
	})
	if err != nil {
		return
	}
	return resp.Result, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Sensor_VoltageSensor_Sense - Возвращает все доступные показатели датчика тока
func (m *GRPCClient) RI_SDK_Sensor_VoltageSensor_Sense(descriptor int64) (voltage, voltageShunt, current, power float32, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Sensor_VoltageSensor_Sense(context.Background(), &RI_SDK_Sensor_VoltageSensor_SenseParams{
		Desrciptor: descriptor,
	})
	if err != nil {
		return
	}
	return resp.Voltage, resp.VoltageShunt, resp.Current, resp.Power, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Sensor_VoltageSensor_Extend - расширяет компонент группы
func (m *GRPCServer) RI_SDK_Sensor_VoltageSensor_Extend(
	ctx context.Context,
	req *RI_SDK_Sensor_VoltageSensor_ExtendParams) (*RI_SDK_Sensor_VoltageSensor_ExtendReturn, error) {
	descriptor, errorText, errorCode, err := m.Impl.RI_SDK_Sensor_VoltageSensor_Extend(req.Exec)
	return &RI_SDK_Sensor_VoltageSensor_ExtendReturn{
		Descriptor_: descriptor,
		ErrorText:   errorText,
		ErrorCode:   errorCode,
	}, err
}

// RI_SDK_Sensor_VoltageSensor_ExtendToModel - расширяет датчик тока до модели
func (m *GRPCServer) RI_SDK_Sensor_VoltageSensor_ExtendToModel(
	ctx context.Context,
	req *RI_SDK_Sensor_VoltageSensor_ExtendToModelParams) (*RI_SDK_Sensor_VoltageSensor_ExtendToModelReturn, error) {
	descriptor, errorText, errorCode, err := m.Impl.RI_SDK_Sensor_VoltageSensor_ExtendToModel(req.BaseDescriptor, req.ModelName)
	return &RI_SDK_Sensor_VoltageSensor_ExtendToModelReturn{
		Descriptor_: descriptor,
		ErrorText:   errorText,
		ErrorCode:   errorCode,
	}, err
}

// RI_SDK_Sensor_VoltageSensor_CustomDeviceInit - Инициализация кастомного датчика тока
func (m *GRPCServer) RI_SDK_Sensor_VoltageSensor_CustomDeviceInit(
	ctx context.Context,
	req *RI_SDK_Sensor_VoltageSensor_CustomDeviceInitParams) (*RI_SDK_Sensor_VoltageSensor_CustomDeviceInitReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Sensor_VoltageSensor_CustomDeviceInit(req.Desrciptor, req.LsbBus, req.LsbShunt, req.ShuntResist, req.RegVoltageShunt[0], req.RegVoltageBus[0])
	return &RI_SDK_Sensor_VoltageSensor_CustomDeviceInitReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Sensor_VoltageSensor_Voltage - Возвращает значение напряжения в цепи.
func (m *GRPCServer) RI_SDK_Sensor_VoltageSensor_Voltage(
	ctx context.Context,
	req *RI_SDK_Sensor_VoltageSensor_VoltageParams) (*RI_SDK_Sensor_VoltageSensor_VoltageReturn, error) {
	voltage, errorText, errorCode, err := m.Impl.RI_SDK_Sensor_VoltageSensor_Voltage(req.Desrciptor)
	return &RI_SDK_Sensor_VoltageSensor_VoltageReturn{
		Result:    voltage,
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Sensor_VoltageSensor_VoltageShunt - Возвращает значение напряжения на шунте.
func (m *GRPCServer) RI_SDK_Sensor_VoltageSensor_VoltageShunt(
	ctx context.Context,
	req *RI_SDK_Sensor_VoltageSensor_VoltageShuntParams) (*RI_SDK_Sensor_VoltageSensor_VoltageShuntReturn, error) {
	voltageShunt, errorText, errorCode, err := m.Impl.RI_SDK_Sensor_VoltageSensor_VoltageShunt(req.Desrciptor)
	return &RI_SDK_Sensor_VoltageSensor_VoltageShuntReturn{
		Result:    voltageShunt,
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Sensor_VoltageSensor_Current - Возвращает значение силы тока в цепи
func (m *GRPCServer) RI_SDK_Sensor_VoltageSensor_Current(
	ctx context.Context,
	req *RI_SDK_Sensor_VoltageSensor_CurrentParams) (*RI_SDK_Sensor_VoltageSensor_CurrentReturn, error) {
	current, errorText, errorCode, err := m.Impl.RI_SDK_Sensor_VoltageSensor_Current(req.Desrciptor)
	return &RI_SDK_Sensor_VoltageSensor_CurrentReturn{
		Result:    current,
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Sensor_VoltageSensor_Power - Возвращает значение мощности в цепи
func (m *GRPCServer) RI_SDK_Sensor_VoltageSensor_Power(
	ctx context.Context,
	req *RI_SDK_Sensor_VoltageSensor_PowerParams) (*RI_SDK_Sensor_VoltageSensor_PowerReturn, error) {
	power, errorText, errorCode, err := m.Impl.RI_SDK_Sensor_VoltageSensor_Power(req.Desrciptor)
	return &RI_SDK_Sensor_VoltageSensor_PowerReturn{
		Result:    power,
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Sensor_VoltageSensor_Sense - Возвращает значение всех показателей датчика тока
func (m *GRPCServer) RI_SDK_Sensor_VoltageSensor_Sense(
	ctx context.Context,
	req *RI_SDK_Sensor_VoltageSensor_SenseParams) (*RI_SDK_Sensor_VoltageSensor_SenseReturn, error) {
	voltage, voltageShunt, current, power, errorText, errorCode, err := m.Impl.RI_SDK_Sensor_VoltageSensor_Sense(req.Desrciptor)
	return &RI_SDK_Sensor_VoltageSensor_SenseReturn{
		Voltage:      voltage,
		VoltageShunt: voltageShunt,
		Current:      current,
		Power:        power,
		ErrorText:    errorText,
		ErrorCode:    errorCode,
	}, err
}
