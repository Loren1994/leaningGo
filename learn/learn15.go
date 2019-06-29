package learn

func Learn15() {
	//只能发送的通道
	sender := make(chan<- interface{})
	//只能接收的通道
	//一个不能填充数据（发送）只能读取的通道是毫无意义的
	receiver := make(<-chan interface{})
	<-receiver
	sender <- "message"
}
