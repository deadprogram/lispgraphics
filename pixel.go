package main

import (
    "github.com/deosjr/whistle/lisp"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

// wrapper around faiface/pixel libraries for use in lisp

func LoadPixel(env *lisp.Env) {
    // window
    env.AddBuiltin("window", newWindow)
    env.AddBuiltin("closed?", isClosed)
    env.AddBuiltin("clear", clear)
    env.AddBuiltin("update", update)

    // imdraw
    env.AddBuiltin("imdraw", newIMDraw)
    env.AddBuiltin("im-set-color!", setColor)
    env.AddBuiltin("im-push", push)
    env.AddBuiltin("im-draw", drawIMDraw)

    // vector
    env.AddBuiltin("vec2d", newVector)

    // colors
    env.Add("black", lisp.NewPrimitive(pixel.RGB(0, 0, 0)))
    env.Add("red",   lisp.NewPrimitive(pixel.RGB(1, 0, 0)))
    env.Add("green", lisp.NewPrimitive(pixel.RGB(0, 1, 0)))
    env.Add("blue",  lisp.NewPrimitive(pixel.RGB(0, 0, 1)))

    // shapes
    env.AddBuiltin("line", line)
    env.AddBuiltin("polygon", polygon)

    // canvas
    env.AddBuiltin("canvas", newCanvas)
    env.AddBuiltin("canvas-draw", drawCanvas)
}

func newWindow(args []lisp.SExpression) (lisp.SExpression, error) {
	cfg := pixelgl.WindowConfig{
		Title:  "Logo",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
    return lisp.NewPrimitive(win), err
}

func isClosed(args []lisp.SExpression) (lisp.SExpression, error) {
    win := args[0].AsPrimitive().(*pixelgl.Window)
    return lisp.NewPrimitive(win.Closed()), nil
}

func clear(args []lisp.SExpression) (lisp.SExpression, error) {
    win := args[0].AsPrimitive().(*pixelgl.Window)
    color := args[1].AsPrimitive().(pixel.RGBA)
    win.Clear(color)
    return lisp.NewPrimitive(true), nil
}

func update(args []lisp.SExpression) (lisp.SExpression, error) {
    win := args[0].AsPrimitive().(*pixelgl.Window)
    win.Update()
    return lisp.NewPrimitive(true), nil
}

func newIMDraw(args []lisp.SExpression) (lisp.SExpression, error) {
    imd := imdraw.New(nil)
    return lisp.NewPrimitive(imd), nil
}

func setColor(args []lisp.SExpression) (lisp.SExpression, error) {
    imd := args[0].AsPrimitive().(*imdraw.IMDraw)
    color := args[1].AsPrimitive().(pixel.RGBA)
    imd.Color = color
    return lisp.NewPrimitive(true), nil
}

func push(args []lisp.SExpression) (lisp.SExpression, error) {
    imd := args[0].AsPrimitive().(*imdraw.IMDraw)
    vecs := []pixel.Vec{}
    for _, arg := range args[1:] {
        if !arg.IsPrimitive() {
            break
        }
        v, ok := arg.AsPrimitive().(pixel.Vec)
        if !ok {
            break
        }
        vecs = append(vecs, v)
    }
    if len(vecs) == 0 {
        return lisp.NewPrimitive(false), nil
    }
    imd.Push(vecs...)
    return lisp.NewPrimitive(true), nil
}

func drawIMDraw(args []lisp.SExpression) (lisp.SExpression, error) {
    imd := args[0].AsPrimitive().(*imdraw.IMDraw)
    target := args[1].AsPrimitive().(pixel.Target)
    imd.Draw(target)
    return lisp.NewPrimitive(true), nil
}

func newVector(args []lisp.SExpression) (lisp.SExpression, error) {
    x, y := args[0].AsNumber(), args[1].AsNumber()
    return lisp.NewPrimitive(pixel.V(x,y)), nil
}

func line(args []lisp.SExpression) (lisp.SExpression, error) {
    imd := args[0].AsPrimitive().(*imdraw.IMDraw)
    thickness := args[1].AsNumber()
    imd.Line(thickness)
    return lisp.NewPrimitive(true), nil
}

func polygon(args []lisp.SExpression) (lisp.SExpression, error) {
    imd := args[0].AsPrimitive().(*imdraw.IMDraw)
    thickness := args[1].AsNumber()
    imd.Polygon(thickness)
    return lisp.NewPrimitive(true), nil
}

func newCanvas(args []lisp.SExpression) (lisp.SExpression, error) {
    windowBounds := pixel.R(0, 0, 1024, 768)
    canvas := pixelgl.NewCanvas(windowBounds)
    return lisp.NewPrimitive(canvas), nil
}

func drawCanvas(args []lisp.SExpression) (lisp.SExpression, error) {
    canvas := args[0].AsPrimitive().(*pixelgl.Canvas)
    target := args[1].AsPrimitive().(pixel.Target)
    canvas.Draw(target, pixel.IM)
    return lisp.NewPrimitive(true), nil
}
