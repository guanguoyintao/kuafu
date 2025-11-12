package eua

import erand "github.com/guanguoyintao/kuafu/math/rand"

func randomSelectElements(selectableSet []string) string {
	randomIndex := erand.RandomRange(0, len(selectableSet)-1)
	return selectableSet[randomIndex]
}

func userAgentHandler() string {
	userAgentValue := randomSelectElements(userAgentList)

	return userAgentValue
}

func HeaderGenerator(headers []string) map[string]string {
	headerMap := make(map[string]string)

	for _, header := range headers {
		switch header {
		case USER_AGENT:
			userAgentValue := userAgentHandler()
			headerMap[USER_AGENT] = userAgentValue
		default:
			break
		}
	}

	return headerMap
}
