package main

import (
    "log"
    "runtime"

    "github.com/go-gl/gl/v4.1-core/gl" // OR: github.com/go-gl/gl/v2.1/gl
    "github.com/go-gl/glfw/v3.2/glfw"
	"github.com/xyproto/convexhull"
)

const (
    width  = 840
    height = 630
	title  = "Convex Hull in 2D"
	HW     = width / 2
	HH     = height / 2
)

var (
	running, drawHull bool
	points, hull convexhull.Points
	px, py float64
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

// initGlfw initializes glfw and returns a Window to use.
func initGlfw() *glfw.Window {
    if err := glfw.Init(); err != nil {
            panic(err)
    }

    glfw.WindowHint(glfw.Resizable, glfw.False)
    glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
    glfw.WindowHint(glfw.ContextVersionMinor, 1)
    glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
    glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

    window, err := glfw.CreateWindow(width, height, title, nil, nil)
    if err != nil {
            panic(err)
    }
    window.MakeContextCurrent()

    return window
}

// initOpenGL initializes OpenGL and returns an intiialized program.
func initOpenGL() uint32 {
    if err := gl.Init(); err != nil {
            panic(err)
    }
    version := gl.GoStr(gl.GetString(gl.VERSION))
    log.Println("OpenGL version", version)

    prog := gl.CreateProgram()
    gl.LinkProgram(prog)
    return prog
}

func draw(window *glfw.Window, program uint32) {
    gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
    gl.UseProgram(program)

    glfw.PollEvents()
    window.SwapBuffers()
}

func main() {
    window := initGlfw()
    defer glfw.Terminate()

	program := initOpenGL()

//	initGL()

	for !window.ShouldClose() {
		draw(window, program)
	}

	//window.SetKeyCallback(onKey)

//	for !window.ShouldClose() {
		// Do OpenGL stuff.
//		drawScene()

//		window.SwapBuffers()
//		glfw.PollEvents()
//	}
}

//	glfw.SetSwapInterval(1)
//	glfw.SetWindowSizeCallback(onResize)
//	glfw.SetKeyCallback(onKey)
//	glfw.SetMouseButtonCallback(onMouse)
//	glfw.SetMousePosCallback(onCursor)
//
//	initGL()
//
//	for running && !window.ShouldClose() {
//		drawScene()
//		window.SwapBuffers()
//		glfw.PollEvents()
//	}
//}

//func onResize(w, h int32) {
//	if h == 0 {
//		h = 1
//	}
//
//	gl.Viewport(0, 0, w, h)
//	gl.MatrixMode(gl.PROJECTION)
//	gl.LoadIdentity()
//	gl.Ortho(0 ,float64(w), float64(h), 0, -1, 1)
//	//glu.Ortho2D(0, float64(w), float64(h), 0)
//	gl.MatrixMode(gl.MODELVIEW)
//	gl.LoadIdentity()
//}
//
//func onKey(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
//	switch key {
//	case glfw.KeyEscape:
//		running = false
//
//	case 'H':
//		drawHull = !drawHull
//
//	case 'C':
//		points, hull = nil, nil
//		points = make(convexhull.Points, 0)
//		hull = make(convexhull.Points, 0)
//	}
//}
//
//func onCursor(x, y int) {
//	px, py = float64(x), float64(y)
//}
//
//func onMouse(button, state int) {
//	var ok bool
//	if state == 1 {
//		points = append(points, convexhull.MakePoint(px, py))
//		hull, ok = points.Compute()
//		if !ok {
//			fmt.Println("does not compute")
//		}
//	}
//}
//
//func initGL() {
//	gl.ClearColor(1, 1, 1, 0)
//	gl.ClearDepth(1)
//	gl.Enable(gl.DEPTH_TEST)
//	gl.DepthFunc(gl.LEQUAL)
//	gl.Hint(gl.PERSPECTIVE_CORRECTION_HINT, gl.NICEST)
//
//	gl.LineWidth(3)
//	gl.Enable(gl.LINE_SMOOTH)
//
//	gl.PointSize(5)
//	gl.Enable(gl.POINT_SMOOTH)
//
//	gl.Hint(gl.POINT_SMOOTH, gl.NICEST)
//	gl.Enable(gl.BLEND)
//	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
//
//	points = make(convexhull.Points, 0)
//}

//func drawCartesian() {
//	//Horizontal Line
//	gl.Begin(gl.LINES)
//	gl.Color3f(0, 0, 0)
//	gl.Vertex2f(0, float32(HH))
//	gl.Vertex2f(Width, float32(HH))
//	gl.End()
//
//	//Vertical line
//	gl.Begin(gl.LINES)
//	gl.Color3f(0, 0, 0)
//	gl.Vertex2f(float32(HW), 0)
//	gl.Vertex2f(float32(HW), Height)
//	gl.End()
//
//	//Origin
//	gl.Begin(gl.POINTS)
//	gl.Color3f(0, 1, 1)
//	gl.Vertex2f(float32(HW), float32(HH))
//	gl.End()
//}
//
//func drawScene() {
//	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
//	gl.LoadIdentity()
//
//	DrawPoints(points)
//	DrawLowestPoint(points)
//
//	if drawHull {
//		DrawLines(hull)
//	}
//
//	//Print cartesian
//	drawCartesian()
//}

//func DrawLowestPoint(points convexhull.Points) {
//	if len(points) <= 0 {
//		return
//	}
//
//	gl.Begin(gl.POINTS)
//	gl.Color3f(0, 0, 0)
//	gl.Vertex2f(float32(points[0].X), float32(points[0].Y))
//	gl.End()
//}
//
//func DrawPoints(points convexhull.Points) {
//	gl.Begin(gl.POINTS)
//	for _, p := range points {
//		gl.Color3f(1, 0, 0)
//		gl.Vertex2f(float32(p.X), float32(p.Y))
//	}
//	gl.End()
//}
//
//func DrawLines(points convexhull.Points) {
//	gl.Begin(gl.LINE_LOOP)
//	for _, p := range points {
//		gl.Color3f(0, 0, 1)
//		gl.Vertex2f(float32(p.X), float32(p.Y))
//	}
//	gl.End()
//}


