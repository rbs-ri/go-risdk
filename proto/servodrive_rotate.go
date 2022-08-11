package proto

import "context"

// RI_SDK_Exec_RServoDrive_Extend - расширяет исполнитель до сервопривода вращения
func (m *GRPCClient) RI_SDK_Exec_RServoDrive_Extend(exec int64) (descriptor int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_RServoDrive_Extend(context.Background(), &RI_SDK_Exec_RServoDrive_ExtendParams{
		Exec: exec,
	})
	if err != nil {
		return
	}
	return resp.Descriptor_, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_RServoDrive_Extend - расширяет исполнитель до сервопривода вращения
func (m *GRPCServer) RI_SDK_Exec_RServoDrive_Extend(
	ctx context.Context,
	req *RI_SDK_Exec_RServoDrive_ExtendParams) (*RI_SDK_Exec_RServoDrive_ExtendReturn, error) {
	descriptor, errorText, errorCode, err := m.Impl.RI_SDK_Exec_RServoDrive_Extend(req.Exec)
	return &RI_SDK_Exec_RServoDrive_ExtendReturn{
		Descriptor_: descriptor,
		ErrorText:   errorText,
		ErrorCode:   errorCode,
	}, err
}

// RI_SDK_Exec_RServoDrive_ExtendToModel - расширяет сервопривод вращения до модели
func (m *GRPCClient) RI_SDK_Exec_RServoDrive_ExtendToModel(baseDescriptor int64, modelName string) (descriptor int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_RServoDrive_ExtendToModel(context.Background(), &RI_SDK_Exec_RServoDrive_ExtendToModelParams{
		BaseDescriptor: baseDescriptor,
		ModelName:      modelName,
	})
	if err != nil {
		return
	}
	return resp.Descriptor_, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_RServoDrive_ExtendToModel - расширяет сервопривод вращения до модели
func (m *GRPCServer) RI_SDK_Exec_RServoDrive_ExtendToModel(
	ctx context.Context,
	req *RI_SDK_Exec_RServoDrive_ExtendToModelParams) (*RI_SDK_Exec_RServoDrive_ExtendToModelReturn, error) {
	descriptor, errorText, errorCode, err := m.Impl.RI_SDK_Exec_RServoDrive_ExtendToModel(req.BaseDescriptor, req.ModelName)
	return &RI_SDK_Exec_RServoDrive_ExtendToModelReturn{
		Descriptor_: descriptor,
		ErrorText:   errorText,
		ErrorCode:   errorCode,
	}, err
}

// RI_SDK_Exec_RServoDrive_CustomDeviceInit - Инициализация кастомного сервопривода вращения
func (m *GRPCClient) RI_SDK_Exec_RServoDrive_CustomDeviceInit(desrciptor, minPulseClockwise, maxPulseClockwise, minPulseCounterClockwise, maxPulseCounterClockwise int64) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_RServoDrive_CustomDeviceInit(context.Background(), &RI_SDK_Exec_RServoDrive_CustomDeviceInitParams{
		Descriptor_:              desrciptor,
		MinPulseClockwise:        minPulseClockwise,
		MaxPulseClockwise:        maxPulseClockwise,
		MinPulseCounterClockwise: minPulseCounterClockwise,
		MaxPulseCounterClockwise: maxPulseCounterClockwise,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_RServoDrive_CustomDeviceInit - Инициализация кастомного сервопривода вращения
func (m *GRPCServer) RI_SDK_Exec_RServoDrive_CustomDeviceInit(
	ctx context.Context,
	req *RI_SDK_Exec_RServoDrive_CustomDeviceInitParams) (*RI_SDK_Exec_RServoDrive_CustomDeviceInitReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Exec_RServoDrive_CustomDeviceInit(req.Descriptor_, req.MinPulseClockwise, req.MaxPulseClockwise, req.MinPulseCounterClockwise, req.MaxPulseCounterClockwise)
	return &RI_SDK_Exec_RServoDrive_CustomDeviceInitReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Exec_RServoDrive_GetState - Получение состояния сервопривода вращения
func (m *GRPCClient) RI_SDK_Exec_RServoDrive_GetState(descriptor int64) (state int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_RServoDrive_GetState(context.Background(), &RI_SDK_Exec_RServoDrive_GetStateParams{
		Descriptor_: descriptor,
	})
	if err != nil {
		return
	}
	return resp.State, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_RServoDrive_GetState - Получение состояния сервопривода вращения
func (m *GRPCServer) RI_SDK_Exec_RServoDrive_GetState(
	ctx context.Context,
	req *RI_SDK_Exec_RServoDrive_GetStateParams) (*RI_SDK_Exec_RServoDrive_GetStateReturn, error) {
	state, errorText, errorCode, err := m.Impl.RI_SDK_Exec_RServoDrive_GetState(req.Descriptor_)
	return &RI_SDK_Exec_RServoDrive_GetStateReturn{
		State:     state,
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Exec_RServoDrive_Stop - Прекращает подачу сигнала к сервоприводу вращения
func (m *GRPCClient) RI_SDK_Exec_RServoDrive_Stop(descriptor int64) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_RServoDrive_Stop(context.Background(), &RI_SDK_Exec_RServoDrive_StopParams{
		Descriptor_: descriptor,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_RServoDrive_Stop - Прекращает подачу сигнала к сервоприводу вращения
func (m *GRPCServer) RI_SDK_Exec_RServoDrive_Stop(
	ctx context.Context,
	req *RI_SDK_Exec_RServoDrive_StopParams) (*RI_SDK_Exec_RServoDrive_StopReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Exec_RServoDrive_Stop(req.Descriptor_)
	return &RI_SDK_Exec_RServoDrive_StopReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Exec_RServoDrive_RotateByPulseOverTime - Выполняет вращение сервопривода, пока не будет вызван метод Stop
// или по истечению заданного времени, направление и скорость задаются импульсом.
func (m *GRPCClient) RI_SDK_Exec_RServoDrive_RotateByPulseOverTime(descriptor, pulse, timeout int64, async bool) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_RServoDrive_RotateByPulseOverTime(context.Background(), &RI_SDK_Exec_RServoDrive_RotateByPulseOverTimeParams{
		Descriptor_: descriptor,
		Pulse:       pulse,
		Timeout:     timeout,
		Async:       async,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_RServoDrive_RotateByPulseOverTime - Выполняет вращение сервопривода, пока не будет вызван метод Stop
// или по истечению заданного времени, направление и скорость задаются импульсом.
func (m *GRPCServer) RI_SDK_Exec_RServoDrive_RotateByPulseOverTime(
	ctx context.Context,
	req *RI_SDK_Exec_RServoDrive_RotateByPulseOverTimeParams) (*RI_SDK_Exec_RServoDrive_RotateByPulseOverTimeReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Exec_RServoDrive_RotateByPulseOverTime(req.Descriptor_, req.Pulse, req.Timeout, req.Async)
	return &RI_SDK_Exec_RServoDrive_RotateByPulseOverTimeReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Exec_RServoDrive_RotateByPulse - Выполняет вращение сервопривода, направление и скорость задаются импульсом.
func (m *GRPCClient) RI_SDK_Exec_RServoDrive_RotateByPulse(descriptor, pulse int64, async bool) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_RServoDrive_RotateByPulse(context.Background(), &RI_SDK_Exec_RServoDrive_RotateByPulseParams{
		Descriptor_: descriptor,
		Pulse:       pulse,
		Async:       async,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_RServoDrive_RotateByPulse - Выполняет вращение сервопривода, направление и скорость задаются импульсом.
func (m *GRPCServer) RI_SDK_Exec_RServoDrive_RotateByPulse(
	ctx context.Context,
	req *RI_SDK_Exec_RServoDrive_RotateByPulseParams) (*RI_SDK_Exec_RServoDrive_RotateByPulseReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Exec_RServoDrive_RotateByPulse(req.Descriptor_, req.Pulse, req.Async)
	return &RI_SDK_Exec_RServoDrive_RotateByPulseReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Exec_RServoDrive_RotateWithRelativeSpeedOverTime - Выполняет вращение сервопривода с заданной скоростью до
// тех пор, пока не будет вызван метод Stop или по истечению заданного времени.
func (m *GRPCClient) RI_SDK_Exec_RServoDrive_RotateWithRelativeSpeedOverTime(descriptor, direction, speed, timeout int64, async bool) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_RServoDrive_RotateWithRelativeSpeedOverTime(context.Background(), &RI_SDK_Exec_RServoDrive_RotateWithRelativeSpeedOverTimeParams{
		Descriptor_: descriptor,
		Direction:   direction,
		Speed:       speed,
		Timeout:     timeout,
		Async:       async,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_RServoDrive_RotateWithRelativeSpeedOverTime - Выполняет вращение сервопривода с заданной скоростью до
// тех пор, пока не будет вызван метод Stop или по истечению заданного времени.
func (m *GRPCServer) RI_SDK_Exec_RServoDrive_RotateWithRelativeSpeedOverTime(
	ctx context.Context,
	req *RI_SDK_Exec_RServoDrive_RotateWithRelativeSpeedOverTimeParams) (*RI_SDK_Exec_RServoDrive_RotateWithRelativeSpeedOverTimeReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Exec_RServoDrive_RotateWithRelativeSpeedOverTime(req.Descriptor_, req.Direction, req.Speed, req.Timeout, req.Async)
	return &RI_SDK_Exec_RServoDrive_RotateWithRelativeSpeedOverTimeReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}

// RI_SDK_Exec_RServoDrive_RotateWithRelativeSpeed - Выполняет вращение сервопривода с заданной скоростью до
// тех пор, пока не будет вызван метод Stop.
func (m *GRPCClient) RI_SDK_Exec_RServoDrive_RotateWithRelativeSpeed(descriptor, direction, speed int64, async bool) (errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Exec_RServoDrive_RotateWithRelativeSpeed(context.Background(), &RI_SDK_Exec_RServoDrive_RotateWithRelativeSpeedParams{
		Descriptor_: descriptor,
		Direction:   direction,
		Speed:       speed,
		Async:       async,
	})
	if err != nil {
		return
	}
	return resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Exec_RServoDrive_RotateWithRelativeSpeed - Выполняет вращение сервопривода с заданной скоростью до
// тех пор, пока не будет вызван метод Stop.
func (m *GRPCServer) RI_SDK_Exec_RServoDrive_RotateWithRelativeSpeed(
	ctx context.Context,
	req *RI_SDK_Exec_RServoDrive_RotateWithRelativeSpeedParams) (*RI_SDK_Exec_RServoDrive_RotateWithRelativeSpeedReturn, error) {
	errorText, errorCode, err := m.Impl.RI_SDK_Exec_RServoDrive_RotateWithRelativeSpeed(req.Descriptor_, req.Direction, req.Speed, req.Async)
	return &RI_SDK_Exec_RServoDrive_RotateWithRelativeSpeedReturn{
		ErrorText: errorText,
		ErrorCode: errorCode,
	}, err
}
