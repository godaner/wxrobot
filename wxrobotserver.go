package wxrobot

var wxRobot *WXRobot
func init(){
	wxRobot = NewWeixin()
}
func SetServerHandler(messageHandler *Handler){
	wxRobot.handler = messageHandler
}
func StartServer() error {

	newLoginUri, err := wxRobot.GetNewLoginUrl()
	if err != nil {
		return err
	}

	err = wxRobot.NewLoginPage(newLoginUri)
	if err != nil {
		return err
	}

	err = wxRobot.Init()
	if err != nil {
		return err
	}

	err = wxRobot.GetContacts()
	if err != nil {
		return err
	}
	err = wxRobot.Listening()
	if err != nil {
		return err
	}
	return nil
}

