package proto

import "context"

// RI_SDK_Connector_Extend - расширяет базовый компонент до компонента группы коннекторов
func (m *GRPCClient) RI_SDK_Connector_Extend(basicDescriptor int64) (desrciptor int64, errorText string, errorCode int64, err error) {
	resp, err := m.client.RI_SDK_Connector_Extend(context.Background(), &RI_SDK_Connector_ExtendParams{
		Basic: basicDescriptor,
	})
	if err != nil {
		return
	}
	return resp.Desrciptor, resp.ErrorText, resp.ErrorCode, nil
}

// RI_SDK_Connector_Extend - расширяет компонент группы
func (m *GRPCServer) RI_SDK_Connector_Extend(
	ctx context.Context,
	req *RI_SDK_Connector_ExtendParams) (*RI_SDK_Connector_ExtendReturn, error) {
	descriptor, errorText, errorCode, err := m.Impl.RI_SDK_Executor_Extend(req.Basic)
	return &RI_SDK_Connector_ExtendReturn{
		Desrciptor: descriptor,
		ErrorText:  errorText,
		ErrorCode:  errorCode,
	}, err
}
