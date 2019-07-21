package remote_account

type Response struct {
	isFailed bool
}

func sendTransferToProtocolMsg(fromId, toId string, amount uint) {

}

func waitProtocolResp() Response {
	ok := Response{false}
	return ok
}

func sendProtocolResponse(resp Response) {

}
