package application

////////////////////////////////////////////////////////////////////////////////
// 定义各个生命周期管理函数

type OnCreateFunc func() error

type OnStartPreFunc func() error

type OnStartFunc func() error

type OnStartPostFunc func() error

type OnLoopFunc func() error

type OnStopPreFunc func() error

type OnStopFunc func() error

type OnStopPostFunc func() error

type OnDestroyFunc func() error

////////////////////////////////////////////////////////////////////////////////

// Life ...
type Life struct {
	Order int // 表示初始化顺序， 值越大越靠后

	OnCreate    OnCreateFunc
	OnStartPre  OnStartPreFunc
	OnStart     OnStartFunc
	OnStartPost OnStartPostFunc
	OnLoop      OnLoopFunc
	OnStopPre   OnStopPreFunc
	OnStop      OnStopFunc
	OnStopPost  OnStopPostFunc
	OnDestroy   OnDestroyFunc
}

// Lifecycle ...
type Lifecycle interface {
	Life() *Life
}
