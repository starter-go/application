package application

////////////////////////////////////////////////////////////////////////////////
// 定义各个生命周期管理函数

// OnCreateFunc ...
type OnCreateFunc func() error

// OnStartPreFunc ...
type OnStartPreFunc func() error

// OnStartFunc ...
type OnStartFunc func() error

// OnStartPostFunc ...
type OnStartPostFunc func() error

// OnLoopFunc ...
type OnLoopFunc func() error

// OnStopPreFunc ...
type OnStopPreFunc func() error

// OnStopFunc ...
type OnStopFunc func() error

// OnStopPostFunc ...
type OnStopPostFunc func() error

// OnDestroyFunc ...
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

// LifeManager ...
type LifeManager interface {
	Add(l *Life)

	GetMaster() *Life
}
