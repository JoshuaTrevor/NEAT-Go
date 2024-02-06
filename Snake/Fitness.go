package snake

import (
	ffneuralnet "github.com/JoshuaTrevor/Neat-Go/FFNeuralNet"
	snake "github.com/JoshuaTrevor/Neat-Go/Snake"
)

func evaluate(neuralNet *ffneuralnet.FFNeuralNet) float64 {
	fitnessSum := float64(0)
	fitnessCount := float64(0)

	for i := 0; i < 5; i++ {
		fitnessSum += evaluateSingleRun(neuralNet)
		fitnessCount += 1
	}
	return fitnessSum / fitnessCount
}

func evaluateSingleRun(neuralNet *ffneuralnet.FFNeuralNet) float64 {
	snakeGame := NewGame()
	moves := 0

	for !snakeGame.Snake.ShouldDie() {
		vision := snakeGame.getVision()
		nn_output := neuralNet.Feed(vision)

		// Convert to direction
		maxVal := float64(-1000)
		maxValIndex := -1
		for idx, val := range nn_output {
			if val > maxVal{
				maxVal = val
				maxValIndex = idx
			}
		}

		//Get prev manhattan first
		snakeGame.Move(maxValIndex)
		// Now get manhattan against and confirm destructive vs contstructive

		// Basically have everything we need to return fitness function here
		// TODO tomorrow: manhattan distance, fitness function, checkpoint saving
		// Then hopefully if all that works, can get it running in a browser using previous react app (after converting react app to accept json)
	}
}

// I think the game itself is essentially ready to go...
// So just porting the below should be good

//         float fitnessSum = 0;
//         int fitnessCount = 0;
//         for(int i = 0; i < thingsToDo; i++)
//         {
//             Snake snake = new Snake(false);
//             int moves = 0;
//             int productiveMoves = 0;
//             int destructiveMoves = 0;
//             while(!snake.dead)
//             {
//                 float[] stateDistances = snake.getStateDistances(); // This is getvision in go impl
//                 float[] output = nn.feed(stateDistances);
//
//
//                 //Convert to direction
//                 float maxVal = -1000;
//                 int maxValIndex = -1;
//                 for(int j = 0; j < output.length; j++)
//                 {
//                     if(output[j] > maxVal)
//                     {
//                         maxVal = output[j];
//                         maxValIndex = j;
//                     }
//                 }
//
//                 int prevManhattan = snake.manhattanDistanceToApple();
//                 snake.move(Snake.Direction.values()[maxValIndex]);
//                 if(snake.manhattanDistanceToApple() > prevManhattan)
//                     destructiveMoves++;
//                 else
//                     productiveMoves++;
//                 if(moves < 20)
//                     moves++;
//                 else if (snake.movesSinceApple > snake.rows * snake.cols * Math.min(0.4F + snake.applesEaten * 0.1F, 1F))
//                     snake.dead = true;
//             }
//             float suicidePenalty = moves < 5 ? -100F : 0;
//             fitnessSum += suicidePenalty + (snake.applesEaten*250) + (productiveMoves) - (destructiveMoves*1.3F);
//
//             fitnessCount++;
//         }
//
//         return (fitnessSum / fitnessCount);
//
