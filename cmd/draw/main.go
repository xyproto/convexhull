package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/xyproto/convexhull"
)

const (
	title  = "Convex Hull in 2D"
	width  = 840
	height = 630
	HW     = width / 2
	HH     = height / 2
)

var (
	running, drawHull bool
	points, hull      convexhull.Points
	px, py            float64
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	runtime.LockOSThread()
}

// initGlfw initializes glfw and returns a Window to use.
func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

func main() {
	panic("work in progress")

	window := initGlfw()
	defer glfw.Terminate()

	initGL()

	window.SetKeyCallback(onKey)
	window.SetMouseButtonCallback(onMouse)
	window.SetCursorPosCallback(onCursor)
	window.SetSizeCallback(onResize)

	glfw.SwapInterval(1)

	for !window.ShouldClose() {
		drawScene(window)
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func onResize(win *glfw.Window, w, h int) {
	if h == 0 {
		h = 1
	}

	gl.Viewport(0, 0, int32(w), int32(h))
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, float64(w), float64(h), 0, -1, 1)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}

func onKey(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	switch key {
	case glfw.KeyEscape:
		running = false

	case glfw.KeyH:
		drawHull = !drawHull

	case glfw.KeyC:
		points, hull = nil, nil
		points = make(convexhull.Points, 0)
		hull = make(convexhull.Points, 0)
	}
}

func onCursor(w *glfw.Window, xpos, ypos float64) {
	px, py = xpos, ypos
}

func onMouse(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) { // button, state int) {
	var err error
	if button == glfw.MouseButtonLeft {
		points = append(points, convexhull.New(px, py))
		hull, err = points.Compute()
		if err != nil {
			panic(err)
		}
	}
}

func initGL() {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	gl.ClearColor(1, 1, 1, 0)
	gl.ClearDepth(1)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LEQUAL)
	gl.Hint(gl.PERSPECTIVE_CORRECTION_HINT, gl.NICEST)

	gl.LineWidth(3)
	gl.Enable(gl.LINE_SMOOTH)

	gl.PointSize(5)
	gl.Enable(gl.POINT_SMOOTH)

	gl.Hint(gl.POINT_SMOOTH, gl.NICEST)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	points = make(convexhull.Points, 0)
}

func drawCartesian() {
	//Horizontal Line
	gl.Begin(gl.LINES)
	gl.Color3f(0, 0, 0)
	gl.Vertex2f(0, float32(HH))
	gl.Vertex2f(width, float32(HH))
	gl.End()

	//Vertical line
	gl.Begin(gl.LINES)
	gl.Color3f(0, 0, 0)
	gl.Vertex2f(float32(HW), 0)
	gl.Vertex2f(float32(HW), height)
	gl.End()

	//Origin
	gl.Begin(gl.POINTS)
	gl.Color3f(0, 1, 1)
	gl.Vertex2f(float32(HW), float32(HH))
	gl.End()
}

func drawScene(window *glfw.Window) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.LoadIdentity()

	DrawPoints(points)
	DrawLowestPoint(points)

	if drawHull {
		DrawLines(hull)
	}

	//Print cartesian
	drawCartesian()
}

func DrawLowestPoint(points convexhull.Points) {
	if len(points) <= 0 {
		return
	}

	gl.Begin(gl.POINTS)
	gl.Color3f(0, 0, 0)
	gl.Vertex2f(float32(points[0].X), float32(points[0].Y))
	gl.End()
}

func DrawPoints(points convexhull.Points) {
	gl.Begin(gl.POINTS)
	for _, p := range points {
		gl.Color3f(1, 0, 0)
		gl.Vertex2f(float32(p.X), float32(p.Y))
	}
	gl.End()
}

func DrawLines(points convexhull.Points) {
	gl.Begin(gl.LINE_LOOP)
	for _, p := range points {
		gl.Color3f(0, 0, 1)
		gl.Vertex2f(float32(p.X), float32(p.Y))
	}
	gl.End()
}
