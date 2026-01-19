package event

// DefaultEventHandler is an implementation of [EventHandler] interface with
// all handlers doing nothing. It is a good starting point to be embedded your
// own struct to be extended.
type DefaultEventHandler struct{}

func (e *DefaultEventHandler) Workspace(WorkspaceName)        {}
func (e *DefaultEventHandler) FocusedMonitor(FocusedMonitor)  {}
func (e *DefaultEventHandler) ActiveWindow(ActiveWindow)      {}
func (e *DefaultEventHandler) Fullscreen(Fullscreen)          {}
func (e *DefaultEventHandler) MonitorRemoved(MonitorName)     {}
func (e *DefaultEventHandler) MonitorAdded(MonitorName)       {}
func (e *DefaultEventHandler) CreateWorkspace(WorkspaceName)  {}
func (e *DefaultEventHandler) DestroyWorkspace(WorkspaceName) {}
func (e *DefaultEventHandler) MoveWorkspace(MoveWorkspace)    {}
func (e *DefaultEventHandler) ActiveLayout(ActiveLayout)      {}
func (e *DefaultEventHandler) OpenWindow(OpenWindow)          {}
func (e *DefaultEventHandler) CloseWindow(CloseWindow)        {}
func (e *DefaultEventHandler) MoveWindow(MoveWindow)          {}
func (e *DefaultEventHandler) OpenLayer(OpenLayer)            {}
func (e *DefaultEventHandler) CloseLayer(CloseLayer)          {}
func (e *DefaultEventHandler) SubMap(SubMap)                  {}
func (e *DefaultEventHandler) Screencast(Screencast)          {}
func (e *DefaultEventHandler) ToggleGroup(ToggleGroup)        {}
func (e *DefaultEventHandler) MoveOutofGroup(MoveOutofGroup)      {}
func (e *DefaultEventHandler) MoveIntogroup(MoveIntogroup)        {}
func (e *DefaultEventHandler) IgnoreGroupLock(IgnoreGroupLock)    {}
func (e *DefaultEventHandler) LockGroups(LockGroups)              {}
func (e *DefaultEventHandler) WorkspaceV2(WorkspaceV2)            {}
func (e *DefaultEventHandler) FocusedMonitorV2(FocusedMonitorV2)  {}
func (e *DefaultEventHandler) MonitorRemovedV2(MonitorRemovedV2)  {}
func (e *DefaultEventHandler) MonitorAddedV2(MonitorAddedV2)      {}
func (e *DefaultEventHandler) CreateWorkspaceV2(CreateWorkspaceV2)   {}
func (e *DefaultEventHandler) DestroyWorkspaceV2(DestroyWorkspaceV2) {}
func (e *DefaultEventHandler) MoveWorkspaceV2(MoveWorkspaceV2)    {}
func (e *DefaultEventHandler) RenameWorkspace(RenameWorkspace)    {}
func (e *DefaultEventHandler) ActiveSpecial(ActiveSpecial)        {}
func (e *DefaultEventHandler) ActiveSpecialV2(ActiveSpecialV2)    {}
func (e *DefaultEventHandler) MoveWindowV2(MoveWindowV2)          {}
func (e *DefaultEventHandler) ChangeFloatingMode(ChangeFloatingMode) {}
func (e *DefaultEventHandler) Urgent(Urgent)                      {}
func (e *DefaultEventHandler) WindowTitle(WindowTitle)            {}
func (e *DefaultEventHandler) WindowTitleV2(WindowTitleV2)        {}
func (e *DefaultEventHandler) ConfigReloaded()                    {}
func (e *DefaultEventHandler) Pin(Pin)                            {}
func (e *DefaultEventHandler) Minimize(Minimize)                  {}
func (e *DefaultEventHandler) Bell(Bell)                          {}
