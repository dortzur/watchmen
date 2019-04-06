package watcherParams

import "time"

func getCommonParams() map[string]string {
	return map[string]string{
		"allowremarks":  "1",
		"msgfound":      "0",
		"thetask":       "0",
		"teamleader":    "0",
		"speccomp":      "",
		"remark":        "",
		"tasks":         "",
		"taskdescr":     "",
		"prevtask":      "0",
		"prevtaskdescr": "",
		"withtasks":     "0",
	}
}

func getDynamicParams(user string, company string, ixee string) map[string]string {
	layout := "2006-01-02 15:04:05"
	now := time.Now().Format(layout)
	return map[string]string{
		"comp": company,
		"name": user,
		"ts":   now,
		"ix":   ixee,
	}
}
func getParams(user string, company string, ixee string) map[string]string {
	params := getCommonParams()
	for k, v := range getDynamicParams(user, company, ixee) {
		params[k] = v
	}
	return params
}
func GetCheckinParams(user string, company string, ixee string) map[string]string {
	params := map[string]string{
		"B1":    "כניסה",
		"tflag": "",
	}

	for k, v := range getParams(user, company, ixee) {
		params[k] = v
	}
	return params
}

func GetCheckoutParams(user string, company string, ixee string) map[string]string {
	params := map[string]string{
		"B1":    "יציאה",
		"tflag": "1",
	}

	for k, v := range getParams(user, company, ixee) {
		params[k] = v
	}
	return params
}
