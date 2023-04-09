/*
Copyright © 2023 tienvu461@gmail.com
*/
package tetris

import (
	"math"
	"math/rand"
)

type vector struct {
	y, x int
}
type block struct {
	shape     []vector
	canRotate bool
	color     int
}

func (b *block) rotateBack() {
	var rotateAngle (float64) = 90
	rotateRadian := rotateAngle * math.Pi / 180
	b.rotateWithRad(rotateRadian)
}

func (b *block) rotate() {
	var rotateAngle (float64) = 270
	rotateRadian := rotateAngle * math.Pi / 180
	b.rotateWithRad(rotateRadian)
}

func (b *block) rotateWithRad(rad float64) {
	if !b.canRotate {
		return
	}
	cos := int(math.Round(math.Cos(rad)))
	sin := int(math.Round(math.Sin(rad)))

	for i, e := range b.shape {
		ny := e.y*cos - e.x*sin
		nx := e.y*sin - e.x*cos

		b.shape[i] = vector{ny, nx}
	}
}

func (b *block) ShapeMinMax() (int, int, int, int) {
	var xmax int = b.shape[0].x
	var xmin int = b.shape[0].x
	var ymax int = b.shape[0].y
	var ymin int = b.shape[0].y
	for _, value := range b.shape {
		if xmax < value.x {
			xmax = value.x
		}
		if xmin > value.x {
			xmin = value.x
		}
		if ymax < value.y {
			ymax = value.y
		}
		if ymin > value.y {
			ymin = value.y
		}
	}
	return ymin, ymax, xmin, xmax
}

// https://qph.cf2.quoracdn.net/main-qimg-356e2b21c801381db2890dab49a9ea88
var blocks = []block{
	{
		shape:     []vector{{0, 0}},
		color:     0,
		canRotate: false,
	},
	// 1. L block - Orange Ricky
	//    X
	// XXXX
	{
		shape:     []vector{{0, 0}, {0, -1}, {0, 1}, {-1, 1}},
		color:     1,
		canRotate: true,
	},
	// 2. Oposite L block - Blue Ricky
	// X
	// XXXX
	{
		shape:     []vector{{0, 0}, {0, -1}, {0, 1}, {-1, -1}},
		color:     2,
		canRotate: true,
	},
	// 3. Z Block - Cleverland Z
	{
		shape:     []vector{{-1, -1}, {-1, 0}, {0, 0}, {0, 1}},
		color:     3,
		canRotate: true,
	},
	// 4. S block - Rhode Island Z
	{
		shape:     []vector{{0, 0}, {0, -1}, {-1, 0}, {-1, 1}},
		color:     4,
		canRotate: true,
	},
	// 5. I Block - Hero
	{
		shape:     []vector{{0, -1}, {0, 0}, {0, 1}, {0, 2}},
		color:     5,
		canRotate: true,
	},
	// 6. Upsidedown T Block - Teewee
	{
		shape:     []vector{{0, 0}, {-1, 0}, {0, -1}, {0, 1}},
		color:     6,
		canRotate: true,
	},
	// 7. Square Block - Smashboy
	{
		shape:     []vector{{0, 0}, {1, 0}, {0, 1}, {1, 1}},
		color:     7,
		canRotate: false,
	},
}

func randBlock() block {
	idx := rand.Intn(len(blocks)-1) + 1
	blk := blocks[idx]
	return block{
		shape:     append([]vector(nil), blk.shape...),
		canRotate: blk.canRotate,
		color:     blk.color,
	}
}
