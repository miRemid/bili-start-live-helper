package main

const (
	QRCODE_URL       = "https://passport.bilibili.com/qrcode/getLoginUrl"
	CHECK_QRCODE_URL = "https://passport.bilibili.com/qrcode/getLoginInfo"

	USER_INFO_URL = "https://api.bilibili.com/x/web-interface/nav"
	GO_URL        = "https://live.bilibili.com/"
	ORIGIN_URL    = "https://passport.bilibili.com"
	ROOMID_URL    = "https://api.live.bilibili.com/room/v1/Room/getRoomInfoOld?mid=%d"

	INFO_URL         = "https://api.bilibili.com/x/space/acc/info?mid=%d&jsonp=jsonp"
	SEND_DANMAKU_URL = "https://api.live.bilibili.com/msg/send"
	LIVE_LIST_URL    = "https://api.live.bilibili.com/room/v1/Area/getList"
	UPDATE_TITLE_URL = "https://api.live.bilibili.com/room/v1/Room/update"
	START_LIVE_URL   = "https://api.live.bilibili.com/room/v1/Room/startLive"
	STOP_LIVE_URL    = "https://api.live.bilibili.com/room/v1/Room/stopLive"

	COOKIES_SAVE_PATH = "static/user/cookies.json"
)
