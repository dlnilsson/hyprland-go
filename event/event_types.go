package event

import (
	"context"
	"net"
)

// EventClient is the event struct from hyprland-go.
type EventClient struct {
	conn net.Conn
}

// Event Client interface, right now only used for testing.
type eventClient interface {
	Receive(_ context.Context) ([]ReceivedData, error)
}

type RawData string

type EventType string

type ReceivedData struct {
	Type EventType
	Data RawData
}

// EventHandler is the interface that defines all methods to handle each of
// events emitted by Hyprland.
// You can find move information about each event in the main Hyprland Wiki:
// https://wiki.hyprland.org/Plugins/Development/Event-list/.
type EventHandler interface {
	// Workspace emitted on workspace change. Is emitted ONLY when a user
	// requests a workspace change, and is not emitted on mouse movements.
	Workspace(w WorkspaceName)
	// FocusedMonitor emitted on the active monitor being changed.
	FocusedMonitor(m FocusedMonitor)
	// ActiveWindow emitted on the active window being changed.
	ActiveWindow(w ActiveWindow)
	// Fullscreen emitted when a fullscreen status of a window changes.
	Fullscreen(f Fullscreen)
	// MonitorRemoved emitted when a monitor is removed (disconnected).
	MonitorRemoved(m MonitorName)
	// MonitorAdded emitted when a monitor is added (connected).
	MonitorAdded(m MonitorName)
	// CreateWorkspace emitted when a workspace is created.
	CreateWorkspace(w WorkspaceName)
	// DestroyWorkspace emitted when a workspace is destroyed.
	DestroyWorkspace(w WorkspaceName)
	// MoveWorkspace emitted when a workspace is moved to a different
	// monitor.
	MoveWorkspace(w MoveWorkspace)
	// ActiveLayout emitted on a layout change of the active keyboard.
	ActiveLayout(l ActiveLayout)
	// OpenWindow emitted when a window is opened.
	OpenWindow(o OpenWindow)
	// CloseWindow emitted when a window is closed.
	CloseWindow(c CloseWindow)
	// MoveWindow emitted when a window is moved to a workspace.
	MoveWindow(m MoveWindow)
	// OpenLayer emitted when a layerSurface is mapped.
	OpenLayer(l OpenLayer)
	// CloseLayer emitted when a layerSurface is unmapped.
	CloseLayer(c CloseLayer)
	// SubMap emitted when a keybind submap changes. Empty means default.
	SubMap(s SubMap)
	// Screencast is fired when the screencopy state of a client changes.
	// Keep in mind there might be multiple separate clients.
	Screencast(s Screencast)
	// ToggleGroup emitted when a group is toggled.
	ToggleGroup(t ToggleGroup)
	// MoveIntogroup emitted when a window is moved into a group.
	MoveIntogroup(m MoveIntogroup)
	// MoveOutofGroup emitted when a window is moved out of a group.
	MoveOutofGroup(m MoveOutofGroup)
	// IgnoreGroupLock emitted when ignore group lock is toggled.
	IgnoreGroupLock(i IgnoreGroupLock)
	// LockGroups emitted when locked groups is toggled.
	LockGroups(l LockGroups)
	// WorkspaceV2 emitted on workspace change, includes workspace ID.
	WorkspaceV2(w WorkspaceV2)
	// FocusedMonitorV2 emitted on active monitor change, includes workspace ID.
	FocusedMonitorV2(m FocusedMonitorV2)
	// MonitorRemovedV2 emitted when a monitor is removed, includes ID and description.
	MonitorRemovedV2(m MonitorRemovedV2)
	// MonitorAddedV2 emitted when a monitor is added, includes ID and description.
	MonitorAddedV2(m MonitorAddedV2)
	// CreateWorkspaceV2 emitted when a workspace is created, includes workspace ID.
	CreateWorkspaceV2(w CreateWorkspaceV2)
	// DestroyWorkspaceV2 emitted when a workspace is destroyed, includes workspace ID.
	DestroyWorkspaceV2(w DestroyWorkspaceV2)
	// MoveWorkspaceV2 emitted when a workspace is moved, includes workspace ID.
	MoveWorkspaceV2(w MoveWorkspaceV2)
	// RenameWorkspace emitted when a workspace is renamed.
	RenameWorkspace(r RenameWorkspace)
	// ActiveSpecial emitted when special workspace opens/closes on a monitor.
	ActiveSpecial(a ActiveSpecial)
	// ActiveSpecialV2 emitted when special workspace opens/closes, includes ID.
	ActiveSpecialV2(a ActiveSpecialV2)
	// MoveWindowV2 emitted when window moves to workspace, includes workspace ID.
	MoveWindowV2(m MoveWindowV2)
	// ChangeFloatingMode emitted when window toggles between floating/tiled.
	ChangeFloatingMode(c ChangeFloatingMode)
	// Urgent emitted when a window requests urgent attention.
	Urgent(u Urgent)
	// WindowTitle emitted when a window title changes.
	WindowTitle(w WindowTitle)
	// WindowTitleV2 emitted when window title changes, includes title string.
	WindowTitleV2(w WindowTitleV2)
	// ConfigReloaded emitted when Hyprland config is reloaded.
	ConfigReloaded()
	// Pin emitted when a window is pinned or unpinned.
	Pin(p Pin)
	// Minimize emitted when a window is minimized or unminimized.
	Minimize(m Minimize)
	// Bell emitted when an app requests the system bell.
	Bell(b Bell)
}

const (
	EventWorkspace          EventType = "workspace"
	EventFocusedMonitor     EventType = "focusedmon"
	EventActiveWindow       EventType = "activewindow"
	EventActiveWindowV2     EventType = "activewindowv2"
	EventFullscreen         EventType = "fullscreen"
	EventMonitorRemoved     EventType = "monitorremoved"
	EventMonitorAdded       EventType = "monitoradded"
	EventCreateWorkspace    EventType = "createworkspace"
	EventDestroyWorkspace   EventType = "destroyworkspace"
	EventMoveWorkspace      EventType = "moveworkspace"
	EventActiveLayout       EventType = "activelayout"
	EventOpenWindow         EventType = "openwindow"
	EventCloseWindow        EventType = "closewindow"
	EventMoveWindow         EventType = "movewindow"
	EventOpenLayer          EventType = "openlayer"
	EventCloseLayer         EventType = "closelayer"
	EventSubMap             EventType = "submap"
	EventScreencast         EventType = "screencast"
	EventToggleGroup        EventType = "togglegroup"
	EventMoveIntogroup      EventType = "moveintogroup"
	EventMoveOutofGroup     EventType = "moveoutofgroup"
	EventIgnoreGroupLock    EventType = "ignoregrouplock"
	EventLockGroups         EventType = "lockgroups"
	EventWorkspaceV2        EventType = "workspacev2"
	EventFocusedMonitorV2   EventType = "focusedmonv2"
	EventMonitorRemovedV2   EventType = "monitorremovedv2"
	EventMonitorAddedV2     EventType = "monitoraddedv2"
	EventCreateWorkspaceV2  EventType = "createworkspacev2"
	EventDestroyWorkspaceV2 EventType = "destroyworkspacev2"
	EventMoveWorkspaceV2    EventType = "moveworkspacev2"
	EventRenameWorkspace    EventType = "renameworkspace"
	EventActiveSpecial      EventType = "activespecial"
	EventActiveSpecialV2    EventType = "activespecialv2"
	EventMoveWindowV2       EventType = "movewindowv2"
	EventChangeFloatingMode EventType = "changefloatingmode"
	EventUrgent             EventType = "urgent"
	EventWindowTitle        EventType = "windowtitle"
	EventWindowTitleV2      EventType = "windowtitlev2"
	EventConfigReloaded     EventType = "configreloaded"
	EventPin                EventType = "pin"
	EventMinimize           EventType = "minimize"
	EventBell               EventType = "bell"
)

// AllEvents is the combination of all event types, useful if you want to
// subscribe to all supported events at the same time.
// Keep in mind that generally explicit declaring which events you want to
// subscribe is better, since new events will be added in future.
var AllEvents = []EventType{
	EventWorkspace,
	EventWorkspaceV2,
	EventFocusedMonitor,
	EventFocusedMonitorV2,
	EventActiveWindow,
	EventActiveWindowV2,
	EventFullscreen,
	EventMonitorRemoved,
	EventMonitorRemovedV2,
	EventMonitorAdded,
	EventMonitorAddedV2,
	EventCreateWorkspace,
	EventCreateWorkspaceV2,
	EventDestroyWorkspace,
	EventDestroyWorkspaceV2,
	EventMoveWorkspace,
	EventMoveWorkspaceV2,
	EventRenameWorkspace,
	EventActiveSpecial,
	EventActiveSpecialV2,
	EventActiveLayout,
	EventOpenWindow,
	EventCloseWindow,
	EventMoveWindow,
	EventMoveWindowV2,
	EventOpenLayer,
	EventCloseLayer,
	EventSubMap,
	EventChangeFloatingMode,
	EventUrgent,
	EventScreencast,
	EventWindowTitle,
	EventWindowTitleV2,
	EventToggleGroup,
	EventMoveIntogroup,
	EventMoveOutofGroup,
	EventIgnoreGroupLock,
	EventLockGroups,
	EventConfigReloaded,
	EventPin,
	EventMinimize,
	EventBell,
}

type MoveWorkspace struct {
	WorkspaceName
	MonitorName
}

type Fullscreen bool

type MonitorName string

type FocusedMonitor struct {
	MonitorName
	WorkspaceName
}

type WorkspaceName string

type SubMap string

type CloseLayer string

type OpenLayer string

type MoveWindow struct {
	Address string
	WorkspaceName
}

type CloseWindow struct {
	Address string
}

type OpenWindow struct {
	Address, Class, Title string
	WorkspaceName
}

type ActiveLayout struct {
	Type, Name string
}

type ActiveWindow struct {
	Name, Title string
}

type ActiveWorkspace WorkspaceName

type Screencast struct {
	// True if a screen or window is being shared.
	Sharing bool

	// "0" if monitor is shared, "1" if window is shared.
	Owner string
}

type ToggleGroup struct {
	Toggle  bool
	Address string
}

type MoveIntogroup struct {
	Address string
}

type MoveOutofGroup struct {
	Address string
}

type WorkspaceV2 struct {
	ID   string
	Name WorkspaceName
}

type FocusedMonitorV2 struct {
	MonitorName
	WorkspaceID string
}

type MonitorRemovedV2 struct {
	ID          string
	Name        MonitorName
	Description string
}

type MonitorAddedV2 struct {
	ID          string
	Name        MonitorName
	Description string
}

type CreateWorkspaceV2 struct {
	ID   string
	Name WorkspaceName
}

type DestroyWorkspaceV2 struct {
	ID   string
	Name WorkspaceName
}

type MoveWorkspaceV2 struct {
	ID   string
	Name WorkspaceName
	MonitorName
}

type RenameWorkspace struct {
	ID      string
	NewName WorkspaceName
}

type ActiveSpecial struct {
	Name WorkspaceName
	MonitorName
}

type ActiveSpecialV2 struct {
	ID   string
	Name WorkspaceName
	MonitorName
}

type MoveWindowV2 struct {
	Address     string
	WorkspaceID string
	WorkspaceName
}

type ChangeFloatingMode struct {
	Address  string
	Floating bool
}

type Urgent struct {
	Address string
}

type WindowTitle struct {
	Address string
}

type WindowTitleV2 struct {
	Address string
	Title   string
}

type Pin struct {
	Address string
	Pinned  bool
}

type Minimize struct {
	Address   string
	Minimized bool
}

type Bell struct {
	Address string
}

type IgnoreGroupLock bool

type LockGroups bool
