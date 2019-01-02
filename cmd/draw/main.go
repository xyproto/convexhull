package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl" // OR: github.com/go-gl/gl/v2.1/gl
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/xyproto/convexhull"
)

const (
	width  = 500 // 840
	height = 500 // 630
	title  = "Convex Hull in 2D"
	HW     = width / 2
	HH     = height / 2

	vertexShaderSource = `
    #version 410
    in vec3 vp;
    void main() {
        gl_Position = vec4(vp, 1.0);
    }
` + "\x00"

	fragmentShaderSource = `
    #version 410
    out vec4 frag_colour;
    void main() {
        frag_colour = vec4(1, 1, 1, 1);
    }
` + "\x00"
)

var (
	running, drawHull bool
	points, hull      convexhull.Points
	px, py            float64

	triangle = []float32{
		0, 0.5, 0, // top
		-0.5, -0.5, 0, // left
		0.5, -0.5, 0, // right
	}
)

// makeVao initializes and returns a vertex array from the points provided.
func makeVao(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

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

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	return prog
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func draw(vao uint32, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	glfw.PollEvents()
	window.SwapBuffers()
}

func main() {
	window := initGlfw()
	defer glfw.Terminate()

	program := initOpenGL()

	vao := makeVao(triangle)

	for !window.ShouldClose() {
		draw(vao, window, program)
		window.SwapBuffers()
	}
}

// ---------------------------------------------

// window.SetKeyCallback(onKey)
//	for !window.ShouldClose() {
//		drawScene()
//		window.SwapBuffers()
//      ...
//		glfw.PollEvents()
//	}

//	glfw.SetSwapInterval(1)
//	glfw.SetWindowSizeCallback(onResize)
//	glfw.SetKeyCallback(onKey)
//	glfw.SetMouseButtonCallback(onMouse)
//	glfw.SetMousePosCallback(onCursor)
//

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
