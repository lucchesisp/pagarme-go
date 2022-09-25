package types

type Split struct {
	enable bool        `json:"enable"`
	rules  []SplitRule `json:"rules"`
}
