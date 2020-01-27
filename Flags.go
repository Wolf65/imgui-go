package imgui

// Flags for Begin()
const (
	WindowFlagsNone                   			= 0			// Default
	WindowFlagsNoTitleBar             			= 1 << 0   	// Disable title-bar
	WindowFlagsNoResize               			= 1 << 1   	// Disable user resizing with the lower-right grip
	WindowFlagsNoMove                 			= 1 << 2   	// Disable user moving the window
	WindowFlagsNoScrollbar            			= 1 << 3   	// Disable scrollbars (window can still scroll with mouse or programmatically)
	WindowFlagsNoScrollWithMouse      			= 1 << 4   	// Disable user vertically scrolling with mouse wheel. On child window, mouse wheel will be forwarded to the parent unless NoScrollbar is also set.
	WindowFlagsNoCollapse             			= 1 << 5   	// Disable user collapsing window by double-clicking on it
	WindowFlagsAlwaysAutoResize       			= 1 << 6   	// Resize every window to its content every frame
	WindowFlagsNoBackground           			= 1 << 7   	// Disable drawing background color (WindowBg, etc.) and outside border. Similar as using SetNextWindowBgAlpha(0.0f).
	WindowFlagsNoSavedSettings        			= 1 << 8   	// Never load/save settings in .ini file
	WindowFlagsNoMouseInputs          			= 1 << 9   	// Disable catching mouse, hovering test with pass through.
	WindowFlagsMenuBar                			= 1 << 10  	// Has a menu-bar
	WindowFlagsHorizontalScrollbar    			= 1 << 11  	// Allow horizontal scrollbar to appear (off by default). You may use SetNextWindowContentSize(ImVec2(width,0.0f)); prior to calling Begin() to specify width. Read code in demo in the "Horizontal Scrolling" section.
	WindowFlagsNoFocusOnAppearing     			= 1 << 12  	// Disable taking focus when transitioning from hidden to visible state
	WindowFlagsNoBringToFrontOnFocus  			= 1 << 13  	// Disable bringing window to front when taking focus (e.g. clicking on it or programmatically giving it focus)
	WindowFlagsAlwaysVerticalScrollbar			= 1 << 14  	// Always show vertical scrollbar (even if ContentSize.y < Size.y)
	WindowFlagsAlwaysHorizontalScrollbar		= 1 << 15	// Always show horizontal scrollbar (even if ContentSize.x < Size.x)
	WindowFlagsAlwaysUseWindowPadding 			= 1 << 16  	// Ensure child windows without border uses style.WindowPadding (ignored by default for non-bordered child windows, because more convenient)
	WindowFlagsNoNavInputs            			= 1 << 18  	// No gamepad/keyboard navigation within the window
	WindowFlagsNoNavFocus             			= 1 << 19  	// No focusing toward this window with gamepad/keyboard navigation (e.g. skipped by CTRL+TAB)
	WindowFlagsUnsavedDocument        			= 1 << 20  	// Append '*' to title without affecting the ID, as a convenience to avoid using the ### operator. When used in a tab/docking context, tab is selected on closure and closure is deferred by one frame to allow code to cancel the closure (with a confirmation popup, etc.) without flicker.
	WindowFlagsNoNav                  			= WindowFlagsNoNavInputs | WindowFlagsNoNavFocus
	WindowFlagsNoDecoration           			= WindowFlagsNoTitleBar | WindowFlagsNoResize | WindowFlagsNoScrollbar | WindowFlagsNoCollapse
	WindowFlagsNoInputs               			= WindowFlagsNoMouseInputs | WindowFlagsNoNavInputs | WindowFlagsNoNavFocus
)
// Flags for InputText()
const (
	InputTextFlagsNone                			= 0			// Default
	InputTextFlagsCharsDecimal        			= 1 << 0   	// Allow 0123456789.+-*/
	InputTextFlagsCharsHexadecimal    			= 1 << 1   	// Allow 0123456789ABCDEFabcdef
	InputTextFlagsCharsUppercase      			= 1 << 2   	// Turn a..z into A..Z
	InputTextFlagsCharsNoBlank        			= 1 << 3   	// Filter out spaces, tabs
	InputTextFlagsAutoSelectAll       			= 1 << 4   	// Select entire text when first taking mouse focus
	InputTextFlagsEnterReturnsTrue    			= 1 << 5   	// Return 'true' when Enter is pressed (as opposed to every time the value was modified). Consider looking at the IsItemDeactivatedAfterEdit() function.
	InputTextFlagsCallbackCompletion  			= 1 << 6   	// Callback on pressing TAB (for completion handling)
	InputTextFlagsCallbackHistory     			= 1 << 7   	// Callback on pressing Up/Down arrows (for history handling)
	InputTextFlagsCallbackAlways      			= 1 << 8   	// Callback on each iteration. User code may query cursor position, modify text buffer.
	InputTextFlagsCallbackCharFilter  			= 1 << 9   	// Callback on character inputs to replace or discard them. Modify 'EventChar' to replace or discard, or return 1 in callback to discard.
	InputTextFlagsAllowTabInput       			= 1 << 10  	// Pressing TAB input a '\t' character into the text field
	InputTextFlagsCtrlEnterForNewLine 			= 1 << 11  	// In multi-line mode, unfocus with Enter, add new line with Ctrl+Enter (default is opposite: unfocus with Ctrl+Enter, add line with Enter).
	InputTextFlagsNoHorizontalScroll  			= 1 << 12  	// Disable following the cursor horizontally
	InputTextFlagsAlwaysInsertMode    			= 1 << 13  	// Insert mode
	InputTextFlagsReadOnly            			= 1 << 14  	// Read-only mode
	InputTextFlagsPassword            			= 1 << 15  	// Password mode, display all characters as '*'
	InputTextFlagsNoUndoRedo          			= 1 << 16  	// Disable undo/redo. Note that input text owns the text data while active, if you want to provide your own undo/redo stack you need e.g. to call ClearActiveID().
	InputTextFlagsCharsScientific     			= 1 << 17  	// Allow 0123456789.+-*/eE (Scientific notation input)
	InputTextFlagsCallbackResize      			= 1 << 18  	// Callback on buffer capacity changes request (beyond 'bufsize' parameter value), allowing the string to grow. Notify when the string wants to be resized (for string types which hold a cache of their Size). You will be provided a new BufSize in the callback and NEED to honor it. (see misc/cpp/stdlib.h for an example of using this)
)
// Flags for TreeNodeEx(), CollapsingHeader*()
const (
	TreeNodeFlagsNone                 			= 0			// Default
	TreeNodeFlagsSelected             			= 1 << 0   	// Draw as selected
	TreeNodeFlagsFramed               			= 1 << 1   	// Full colored frame (e.g. for CollapsingHeader)
	TreeNodeFlagsAllowItemOverlap     			= 1 << 2   	// Hit testing to allow subsequent widgets to overlap this one
	TreeNodeFlagsNoTreePushOnOpen     			= 1 << 3   	// Don't do a TreePush() when open (e.g. for CollapsingHeader) = no extra indent nor pushing on ID stack
	TreeNodeFlagsNoAutoOpenOnLog      			= 1 << 4   	// Don't automatically and temporarily open node when Logging is active (by default logging will automatically open tree nodes)
	TreeNodeFlagsDefaultOpen          			= 1 << 5   	// Default node to be open
	TreeNodeFlagsOpenOnDoubleClick    			= 1 << 6   	// Need double-click to open node
	TreeNodeFlagsOpenOnArrow          			= 1 << 7   	// Only open when clicking on the arrow part. If TreeNodeFlagsOpenOnDoubleClick is also set, single-click arrow or double-click all box to open.
	TreeNodeFlagsLeaf                 			= 1 << 8   	// No collapsing, no arrow (use as a convenience for leaf nodes).
	TreeNodeFlagsBullet               			= 1 << 9   	// Display a bullet instead of arrow
	TreeNodeFlagsFramePadding         			= 1 << 10  	// Use FramePadding (even for an unframed text node) to vertically align text baseline to regular widget height. Equivalent to calling AlignTextToFramePadding().
	TreeNodeFlagsSpanAvailWidth       			= 1 << 11  	// Extend hit box to the right-most edge, even if not framed. This is not the default in order to allow adding other items on the same line. In the future we may refactor the hit system to be front-to-back, allowing natural overlaps and then this can become the default.
	TreeNodeFlagsSpanFullWidth        			= 1 << 12  	// Extend hit box to the left-most and right-most edges (bypass the indented area).
	TreeNodeFlagsNavLeftJumpsBackHere 			= 1 << 13  	// (WIP) Nav: left direction may move to this TreeNode() from any of its child (items submitted between TreeNode and TreePop)
	//TreeNodeFlagsNoScrollOnOpen     			= 1 << 14  	// FIXME: TODO: Disable automatic scroll on TreePop() if node got just open and contents is not visible
	TreeNodeFlagsCollapsingHeader     			= TreeNodeFlagsFramed | TreeNodeFlagsNoTreePushOnOpen | TreeNodeFlagsNoAutoOpenOnLog
)
// Flags for Selectable()
const (
	SelectableFlagsNone               			= 0			// Default
	SelectableFlagsDontClosePopups    			= 1 << 0   	// Clicking this don't close parent popup window
	SelectableFlagsSpanAllColumns     			= 1 << 1   	// Selectable frame can span all columns (text will still fit in current column)
	SelectableFlagsAllowDoubleClick   			= 1 << 2   	// Generate press events on double clicks too
	SelectableFlagsDisabled           			= 1 << 3   	// Cannot be selected, display grayed out text
	SelectableFlagsAllowItemOverlap   			= 1 << 4    // (WIP) Hit testing to allow subsequent widgets to overlap this one
)
// Flags for BeginCombo()
const (
	ComboFlagsNone                    			= 0			// Default
	ComboFlagsPopupAlignLeft          			= 1 << 0   	// Align the popup toward the left by default
	ComboFlagsHeightSmall             			= 1 << 1   	// Max ~4 items visible. Tip: If you want your combo popup to be a specific size you can use SetNextWindowSizeConstraints() prior to calling BeginCombo()
	ComboFlagsHeightRegular           			= 1 << 2   	// Max ~8 items visible (default)
	ComboFlagsHeightLarge             			= 1 << 3   	// Max ~20 items visible
	ComboFlagsHeightLargest           			= 1 << 4   	// As many fitting items as possible
	ComboFlagsNoArrowButton           			= 1 << 5   	// Display on the preview box without the square arrow button
	ComboFlagsNoPreview               			= 1 << 6   	// Display only a square arrow button
	ComboFlagsHeightMask              			= ComboFlagsHeightSmall | ComboFlagsHeightRegular | ComboFlagsHeightLarge | ComboFlagsHeightLargest
)
// Flags for BeginTabBar()
const (
	TabBarFlagsNone                           	= 0			// Default
	TabBarFlagsReorderable                    	= 1 << 0   	// Allow manually dragging tabs to re-order them + New tabs are appended at the end of list
	TabBarFlagsAutoSelectNewTabs              	= 1 << 1   	// Automatically select new tabs when they appear
	TabBarFlagsTabListPopupButton             	= 1 << 2   	// Disable buttons to open the tab list popup
	TabBarFlagsNoCloseWithMiddleMouseButton   	= 1 << 3   	// Disable behavior of closing tabs (that are submitted with popen != NULL) with middle mouse button. You can still repro this behavior on user's side with if (IsItemHovered() && IsMouseClicked(2)) *popen = false.
	TabBarFlagsNoTabListScrollingButtons      	= 1 << 4   	// Disable scrolling buttons (apply when fitting policy is TabBarFlagsFittingPolicyScroll)
	TabBarFlagsNoTooltip                      	= 1 << 5   	// Disable tooltips when hovering a tab
	TabBarFlagsFittingPolicyResizeDown        	= 1 << 6   	// Resize tabs when they don't fit
	TabBarFlagsFittingPolicyScroll            	= 1 << 7   	// Add scroll buttons when tabs don't fit
	TabBarFlagsFittingPolicyMask              	= TabBarFlagsFittingPolicyResizeDown | TabBarFlagsFittingPolicyScroll
	TabBarFlagsFittingPolicyDefault           	= TabBarFlagsFittingPolicyResizeDown
)
// Flags for BeginTabItem()
const (
	TabItemFlagsNone                          	= 0			// Default
	TabItemFlagsUnsavedDocument               	= 1 << 0   	// Append '*' to title without affecting the ID, as a convenience to avoid using the ### operator. Also: tab is selected on closure and closure is deferred by one frame to allow code to undo it without flicker.
	TabItemFlagsSetSelected                   	= 1 << 1  	// Trigger flag to programmatically make the tab selected when calling BeginTabItem()
	TabItemFlagsNoCloseWithMiddleMouseButton  	= 1 << 2  	// Disable behavior of closing tabs (that are submitted with popen != NULL) with middle mouse button. You can still repro this behavior on user's side with if (IsItemHovered() && IsMouseClicked(2)) *popen = false.
	TabItemFlagsNoPushId                      	= 1 << 3    // Don't call PushID(tab->ID)/PopID() on BeginTabItem()/EndTabItem()
)
// Flags for IsWindowFocused()
const (
	FocusedFlagsNone                          	= 0			// Default
	FocusedFlagsChildWindows                  	= 1 << 0   	// IsWindowFocused(): Return true if any children of the window is focused
	FocusedFlagsRootWindow                    	= 1 << 1   	// IsWindowFocused(): Test from root window (top most parent of the current hierarchy)
	FocusedFlagsAnyWindow                     	= 1 << 2   	// IsWindowFocused(): Return true if any window is focused. Important: If you are trying to tell how to dispatch your low-level inputs, do NOT use this. Use ::GetIO().WantCaptureMouse instead.
	FocusedFlagsRootAndChildWindows           	= FocusedFlagsRootWindow | FocusedFlagsChildWindows
)
// Flags for IsItemHovered(), IsWindowHovered()
// Note: if you are trying to check whether your mouse should be dispatched to  or to your app, you should use the 'io.WantCaptureMouse' boolean for that. Please read the FAQ!
// Note: windows with the WindowFlagsNoInputs flag are ignored by IsWindowHovered() calls.
const (
	HoveredFlagsNone                          	= 0        	// Return true if directly over the item/window, not obstructed by another window, not obstructed by an active popup or modal blocking inputs under them.
	HoveredFlagsChildWindows                  	= 1 << 0   	// IsWindowHovered() only: Return true if any children of the window is hovered
	HoveredFlagsRootWindow                    	= 1 << 1   	// IsWindowHovered() only: Test from root window (top most parent of the current hierarchy)
	HoveredFlagsAnyWindow                     	= 1 << 2   	// IsWindowHovered() only: Return true if any window is hovered
	HoveredFlagsAllowWhenBlockedByPopup       	= 1 << 3   	// Return true even if a popup window is normally blocking access to this item/window
	//HoveredFlagsAllowWhenBlockedByModal     	= 1 << 4   	// Return true even if a modal popup window is normally blocking access to this item/window. FIXME-TODO: Unavailable yet.
	HoveredFlagsAllowWhenBlockedByActiveItem  	= 1 << 5   	// Return true even if an active item is blocking access to this item/window. Useful for Drag and Drop patterns.
	HoveredFlagsAllowWhenOverlapped           	= 1 << 6   	// Return true even if the position is obstructed or overlapped by another window
	HoveredFlagsAllowWhenDisabled             	= 1 << 7   	// Return true even if the item is disabled
	HoveredFlagsRectOnly                      	= HoveredFlagsAllowWhenBlockedByPopup | HoveredFlagsAllowWhenBlockedByActiveItem | HoveredFlagsAllowWhenOverlapped
	HoveredFlagsRootAndChildWindows           	= HoveredFlagsRootWindow | HoveredFlagsChildWindows
)
// Flags for BeginDragDropSource(), AcceptDragDropPayload()
const (
	DragDropFlagsNone                         	= 0			// Default
	// BeginDragDropSource() flags
	DragDropFlagsSourceNoPreviewTooltip       	= 1 << 0   	// By default, a successful call to BeginDragDropSource opens a tooltip so you can display a preview or description of the source contents. This flag disable this behavior.
	DragDropFlagsSourceNoDisableHover         	= 1 << 1   	// By default, when dragging we clear data so that IsItemHovered() will return false, to avoid subsequent user code submitting tooltips. This flag disable this behavior so you can still call IsItemHovered() on the source item.
	DragDropFlagsSourceNoHoldToOpenOthers     	= 1 << 2   	// Disable the behavior that allows to open tree nodes and collapsing header by holding over them while dragging a source item.
	DragDropFlagsSourceAllowNullID            	= 1 << 3   	// Allow items such as Text(), Image() that have no unique identifier to be used as drag source, by manufacturing a temporary identifier based on their window-relative position. This is extremely unusual within the dear  ecosystem and so we made it explicit.
	DragDropFlagsSourceExtern                 	= 1 << 4   	// External source (from outside of dear ), won't attempt to read current item/window info. Will always return true. Only one Extern source can be active simultaneously.
	DragDropFlagsSourceAutoExpirePayload      	= 1 << 5   	// Automatically expire the payload if the source cease to be submitted (otherwise payloads are persisting while being dragged)
	// AcceptDragDropPayload() flags
	DragDropFlagsAcceptBeforeDelivery         	= 1 << 10  	// AcceptDragDropPayload() will returns true even before the mouse button is released. You can then call IsDelivery() to test if the payload needs to be delivered.
	DragDropFlagsAcceptNoDrawDefaultRect      	= 1 << 11  	// Do not draw the default highlight rectangle when hovering over target.
	DragDropFlagsAcceptNoPreviewTooltip       	= 1 << 12  	// Request hiding the BeginDragDropSource tooltip from the BeginDragDropTarget site.
	DragDropFlagsAcceptPeekOnly               	= DragDropFlagsAcceptBeforeDelivery | DragDropFlagsAcceptNoDrawDefaultRect  // For peeking ahead and inspecting the payload before delivery.
)
// Configuration flags stored in io.ConfigFlags. Set by user/application.
const (
	ConfigFlagsNone                   			= 0			// Default
	ConfigFlagsNavEnableKeyboard      			= 1 << 0   	// Master keyboard navigation enable flag. NewFrame() will automatically fill io.NavInputs[] based on io.KeysDown[].
	ConfigFlagsNavEnableGamepad       			= 1 << 1   	// Master gamepad navigation enable flag. This is mostly to instruct your  back-end to fill io.NavInputs[]. Back-end also needs to set BackendFlagsHasGamepad.
	ConfigFlagsNavEnableSetMousePos   			= 1 << 2   	// Instruct navigation to move the mouse cursor. May be useful on TV/console systems where moving a virtual mouse is awkward. Will update io.MousePos and set io.WantSetMousePos=true. If enabled you MUST honor io.WantSetMousePos requests in your binding, otherwise  will react as if the mouse is jumping around back and forth.
	ConfigFlagsNavNoCaptureKeyboard   			= 1 << 3   	// Instruct navigation to not set the io.WantCaptureKeyboard flag when io.NavActive is set.
	ConfigFlagsNoMouse                			= 1 << 4   	// Instruct  to clear mouse position/buttons in NewFrame(). This allows ignoring the mouse information set by the back-end.
	ConfigFlagsNoMouseCursorChange    			= 1 << 5   	// Instruct back-end to not alter mouse cursor shape and visibility. Use if the back-end cursor changes are interfering with yours and you don't want to use SetMouseCursor() to change mouse cursor. You may want to honor requests from  by reading GetMouseCursor() yourself instead.

	// User storage (to allow your back-end/engine to communicate to code that may be shared between multiple projects. Those flags are not used by core Dear )
	ConfigFlagsIsSRGB                 			= 1 << 20  	// Application is SRGB-aware.
	ConfigFlagsIsTouchScreen          			= 1 << 21   // Application is using a touch screen instead of a mouse.
)
// Back-end capabilities flags stored in io.BackendFlags. Set by back-end.
const (
	BackendFlagsNone                  			= 0
	BackendFlagsHasGamepad            			= 1 << 0   	// Back-end Platform supports gamepad and currently has one connected.
	BackendFlagsHasMouseCursors       			= 1 << 1   	// Back-end Platform supports honoring GetMouseCursor() value to change the OS cursor shape.
	BackendFlagsHasSetMousePos        			= 1 << 2  	// Back-end Platform supports io.WantSetMousePos requests to reposition the OS mouse position (only used if ConfigFlagsNavEnableSetMousePos is set).
	BackendFlagsRendererHasVtxOffset  			= 1 << 3   	// Back-end Renderer supports ImDrawCmd::VtxOffset. This enables output of large meshes (64K+ vertices) while still using 16-bit indices.
)
// Flags for ColorEdit3() / ColorEdit4() / ColorPicker3() / ColorPicker4() / ColorButton()
const (
	ColorEditFlagsNone            				= 0			//				// Default
	ColorEditFlagsNoAlpha         				= 1 << 1   	//              // ColorEdit, ColorPicker, ColorButton: ignore Alpha component (will only read 3 components from the input pointer).
	ColorEditFlagsNoPicker        				= 1 << 2   	//              // ColorEdit: disable picker when clicking on colored square.
	ColorEditFlagsNoOptions       				= 1 << 3   	//              // ColorEdit: disable toggling options menu when right-clicking on inputs/small preview.
	ColorEditFlagsNoSmallPreview  				= 1 << 4   	//              // ColorEdit, ColorPicker: disable colored square preview next to the inputs. (e.g. to show only the inputs)
	ColorEditFlagsNoInputs        				= 1 << 5   	//              // ColorEdit, ColorPicker: disable inputs sliders/text widgets (e.g. to show only the small preview colored square).
	ColorEditFlagsNoTooltip       				= 1 << 6   	//              // ColorEdit, ColorPicker, ColorButton: disable tooltip when hovering the preview.
	ColorEditFlagsNoLabel         				= 1 << 7   	//              // ColorEdit, ColorPicker: disable display of inline text label (the label is still forwarded to the tooltip and picker).
	ColorEditFlagsNoSidePreview   				= 1 << 8   	//              // ColorPicker: disable bigger color preview on right side of the picker, use small colored square preview instead.
	ColorEditFlagsNoDragDrop      				= 1 << 9   	//              // ColorEdit: disable drag and drop target. ColorButton: disable drag and drop source.

	// User Options (right-click on widget to change some of them).
	ColorEditFlagsAlphaBar        				= 1 << 16  	//              // ColorEdit, ColorPicker: show vertical alpha bar/gradient in picker.
	ColorEditFlagsAlphaPreview    				= 1 << 17  	//              // ColorEdit, ColorPicker, ColorButton: display preview as a transparent color over a checkerboard, instead of opaque.
	ColorEditFlagsAlphaPreviewHalf				= 1 << 18  	//              // ColorEdit, ColorPicker, ColorButton: display half opaque / half checkerboard, instead of opaque.
	ColorEditFlagsHDR             				= 1 << 19  	//              // (WIP) ColorEdit: Currently only disable 0.0f..1.0f limits in RGBA edition (note: you probably want to use ColorEditFlagsFloat flag as well).
	ColorEditFlagsDisplayRGB      				= 1 << 20  	// [Display]    // ColorEdit: override display type among RGB/HSV/Hex. ColorPicker: select any combination using one or more of RGB/HSV/Hex.
	ColorEditFlagsDisplayHSV      				= 1 << 21  	// [Display]    // "
	ColorEditFlagsDisplayHex      				= 1 << 22  	// [Display]    // "
	ColorEditFlagsUint8           				= 1 << 23  	// [DataType]   // ColorEdit, ColorPicker, ColorButton: display values formatted as 0..255.
	ColorEditFlagsFloat           				= 1 << 24  	// [DataType]   // ColorEdit, ColorPicker, ColorButton: display values formatted as 0.0f..1.0f floats instead of 0..255 integers. No round-trip of value via integers.
	ColorEditFlagsPickerHueBar    				= 1 << 25  	// [Picker]     // ColorPicker: bar for Hue, rectangle for Sat/Value.
	ColorEditFlagsPickerHueWheel  				= 1 << 26  	// [Picker]     // ColorPicker: wheel for Hue, triangle for Sat/Value.
	ColorEditFlagsInputRGB        				= 1 << 27  	// [Input]      // ColorEdit, ColorPicker: input and output data in RGB format.
	ColorEditFlagsInputHSV        				= 1 << 28  	// [Input]      // ColorEdit, ColorPicker: input and output data in HSV format.

	// Defaults Options. You can set application defaults using SetColorEditOptions(). The intent is that you probably don't want to
	// override them in most of your calls. Let the user choose via the option menu and/or call SetColorEditOptions() once during startup.
	ColorEditFlagsOptionsDefault 				= ColorEditFlagsUint8|ColorEditFlagsDisplayRGB|ColorEditFlagsInputRGB|ColorEditFlagsPickerHueBar
)