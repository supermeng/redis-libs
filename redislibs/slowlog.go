package redislibs

import "strings"

func FetchSlowLogWithAddress(host, port string, flush bool) ([][]string, error) {
	t, err := BuildTalker(host, port)
	defer t.Close()
	if err != nil {
		return nil, err
	}
	return FetchSlowLogWithTalker(t, flush)
}

func FetchSlowLogWithTalker(t *Talker, flush bool) ([][]string, error) {
	resp, err := t.TalkForObject(Pack_command("SLOWLOG", "GET"))
	if err != nil {
		return nil, err
	}
	slowLogs := resp.([]interface{})

	logs := make([][]string, len(slowLogs))
	logLen := len(slowLogs)
	for i, iSlowLog := range slowLogs {
		slowLog := iSlowLog.([]interface{})
		if len(slowLog) != 4 {
			return nil, BADELEMENT
		}
		instructions := slowLog[3].([]interface{})
		instructionArray := make([]string, len(instructions))
		for i, instruction := range instructions {
			instructionArray[i] = instruction.(string)
		}
		instructionStr := strings.Join(instructionArray, " ")
		logs[logLen-i-1] = []string{slowLog[0].(string), slowLog[1].(string), slowLog[2].(string), instructionStr}
	}
	if flush {
		ResetSlowLogWithTalker(t)
	}

	return logs, nil

}

func ResetSlowLogWithTalker(t *Talker) error {
	_, err := t.TalkRaw(Pack_command("SLOWLOG", "RESET"))
	return err
}
