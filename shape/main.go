package shape

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"math"
)

type IMDraw imdraw.IMDraw

func DrawRoundedRectangle(minVec pixel.Vec, maxVec pixel.Vec, borderRadius float64, imd *imdraw.IMDraw) {
	innerCornerVecs := [4]pixel.Vec{
		pixel.V(minVec.X+borderRadius, minVec.Y+borderRadius),
		pixel.V(maxVec.X-borderRadius, minVec.Y+borderRadius),
		pixel.V(maxVec.X-borderRadius, maxVec.Y-borderRadius),
		pixel.V(minVec.X+borderRadius, maxVec.Y-borderRadius),
	}

	vecs := [12]pixel.Vec{
		pixel.V(minVec.X+borderRadius, minVec.Y+borderRadius),
		pixel.V(minVec.X+borderRadius, minVec.Y),
		pixel.V(maxVec.X-borderRadius, minVec.Y),
		pixel.V(maxVec.X-borderRadius, minVec.Y+borderRadius),
		pixel.V(maxVec.X, minVec.Y+borderRadius),
		pixel.V(maxVec.X, maxVec.Y-borderRadius),
		pixel.V(maxVec.X-borderRadius, maxVec.Y-borderRadius),
		pixel.V(maxVec.X-borderRadius, maxVec.Y),
		pixel.V(minVec.X+borderRadius, maxVec.Y),
		pixel.V(minVec.X+borderRadius, maxVec.Y-borderRadius),
		pixel.V(minVec.X, maxVec.Y-borderRadius),
		pixel.V(minVec.X, minVec.Y+borderRadius),
	}

	imd.Push(vecs[:]...)
	imd.Polygon(0)

	imd.Push(innerCornerVecs[0])
	imd.CircleArc(borderRadius, -math.Pi/2, -math.Pi, 0)

	imd.Push(innerCornerVecs[1])
	imd.CircleArc(borderRadius, -math.Pi/2, 0, 0)

	imd.Push(innerCornerVecs[2])
	imd.CircleArc(borderRadius, 0, math.Pi/2, 0)

	imd.Push(innerCornerVecs[3])
	imd.CircleArc(borderRadius, math.Pi/2, math.Pi, 0)
}
