// tickrate
// render screen

let rAF = null; // keep track of the animation frame so we can cancel it
const canvas = document.getElementById('game');
const context = canvas.getContext('2d');
const grid = 32;
const tetrominoSequence = [];

// keep track of what is in every cell of the game using a 2d array
// tetris playfield is 10x20, with a few rows offscreen
let playfield = [];
let isRun = true;
// // populate the empty state
// for (let row = -2; row < 20; row++) {
//   playfield[row] = [];
//
//   for (let col = 0; col < 10; col++) {
//     playfield[row][col] = 0;
//   }
// }

// game loop
function loop() {
  if (isRun) {
    rAF = requestAnimationFrame(loop);
    context.clearRect(0, 0, canvas.width, canvas.height);

    for (let row = 0; row < 20; row++) {
      playfield[row] = getBoard(row);
    }
    // console.log(playfield[5]);
    // draw the playfield
    for (let row = 0; row < 20; row++) {
      for (let col = 0; col < 15; col++) {
        if (playfield[row][col]) {
          const name = playfield[row][col];
          context.fillStyle = colors[name];

          // drawing 1 px smaller than the grid creates a grid effect
          context.fillRect(col * grid, row * grid, grid - 1, grid - 1);
        }
      }
    }
  }
}
// game state management
// control management
//
function stop() {
  isRun = false;
  console.log(playfield);
}
