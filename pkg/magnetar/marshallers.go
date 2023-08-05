package magnetar

type Marshaller interface {
	MarshalRegistrationReq(req RegistrationReq) ([]byte, error)
	UnmarshalRegistrationReq(reqMarshalled []byte) (*RegistrationReq, error)
	MarshalRegistrationResp(resp RegistrationResp) ([]byte, error)
	UnmarshalRegistrationResp(resp []byte) (*RegistrationResp, error)
	MarshalLabel(label Label) ([]byte, error)
	UnmarshalLabel(labelMarshalled []byte) (Label, error)
	MarshalNode(node Node) ([]byte, error)
	UnmarshalNode(nodeMarshalled []byte) (*Node, error)
}
