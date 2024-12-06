package types

// Status 状态
// 1: 启用
// 2: 禁用
type Status int

const (
	Enabled  Status = 1
	Disabled Status = 2
)

func (s Status) String() string {
	switch s {
	case Enabled:
		return "enabled"
	case Disabled:
		return "disabled"
	}
	return ""
}

func (s Status) IsEnabled() bool {
	return s == Enabled
}

func (s Status) IsDisabled() bool {
	return s == Disabled
}

type Bool int

const (
	True  Bool = 1
	False Bool = 2
)

func (b Bool) String() string {
	switch b {
	case True:
		return "true"
	case False:
		return "false"
	}
	return ""
}

func (b Bool) Bool() bool {
	return b == 1
}
