package providers

type MessageSender interface {
	Send(msg []byte, topic string)
}
