package init

func init() {
	go initToken()
	go initReplyTable()
}
