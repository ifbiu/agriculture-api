package yearbooks

func SwitchChangeCity(city string) int {
	var res = 0
	switch city {
	case "huhehaote":
		res = 150100
	case "baotou":
		res = 150200
	case "wulanchabu":
		res = 150900
	case "tongliao":
		res = 150500
	case "chifeng":
		res = 150400
	case "eerduosi":
		res = 150600
	case "bayannaoer":
		res = 150800
	case "xilinguole":
		res = 152500
	case "hulunberer":
		res = 150700
	case "xinganmeng":
		res = 152200
	case "alashan":
		res = 152900
	case "wuhai":
		res = 150300
	}
	return res
}
