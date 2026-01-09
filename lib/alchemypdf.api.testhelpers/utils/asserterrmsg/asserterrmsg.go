package asserterrmsg

import "fmt"

func BuildAssertErrorMessage(culprit string, expected, got any) string {
	return fmt.Sprintf("Assert fail - httpresponse.Default: expected `%s` = %s, got `%s` = %s ", culprit, expected, culprit, got)
}
