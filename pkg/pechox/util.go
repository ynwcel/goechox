package pechox

func JsonOk(data any, msg ...string) map[string]any {
	json := map[string]any{
		"flag": true,
		"code": 200,
		"data": data,
		"msg":  "",
	}
	if len(msg) > 0 {
		json["msg"] = msg[0]
	}
	return json
}

func JsonFail(data any, msg ...string) map[string]any {
	json := map[string]any{
		"flag": false,
		"code": -200,
		"data": data,
		"msg":  "",
	}
	if len(msg) > 0 {
		json["msg"] = msg[0]
	}
	return json
}
