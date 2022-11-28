package stubs

import "uk.ac.bris.cs/gameoflife/util"

/*ReverseHandler
stubs.ReverseHandler : tell the remote server which method we're calling
ㄴ SecretStingOperation : registered type
ㄴ Reverse : name of the method
request, response are the argument
*/

var EvaluateAllHandler = "GameOfLifeOperation.EvaluateAll"
var KeyPressHandler = "GameOfLifeOperation.KeyPress"
var TickerHandler = "GameOfLifeOperation.Ticker"
var ShutDownHandler = "GameOfLifeOperation.ShutDown"

var EvaluateOneHandler = "GameOfLifeOperation.EvaluateOne"

var SendToServer = "Broker.SendToServer"
var TickerToServer = "Broker.TickerToServer"
var KeyPressToServer = "Broker.KeyPressToServer"
var ShutDown = "Broker.ShutDown"

/*stubs.go
client uses to call the remote methods on the server
need structure to send request to server and get response back
response struct: a 2d slice (final board state back to local controller)
request struct : a 2d slice (initial state of the board),
				   and other parameters(such as the number of turns, size of image -> so it can iterate the board correctly)
exported method name, exported type (going to be changed to something more appropriate like
									  game of life operations and process turns)
*/

type State struct {
	World [][]byte
	Turn  int
}

type None struct{}

type KeyPress struct {
	KeyPress rune
}

type AliveCellState struct {
	AliveCells []util.Cell
}

type EvaluationRequest struct {
	World          [][]byte
	ID             int
	NumberOfWorker int
}

type InitialRequest struct {
	World      [][]byte
	Turn       int
	InstanceIP string
	Ports      []string
}

/*

package stubs

var CreateChannel = "Broker.CreateChannel"
var Publish = "Broker.Publish"
var Subscribe = "Broker.Subscribe"

type Pair struct {
	X int
	Y int
}

type PublishRequest struct {
	Topic string
	Pair Pair
}


type ChannelRequest struct {
	Topic string
	Buffer int
}

type Subscription struct {
	Topic string
	FactoryAddress string
	Callback string
}

type JobReport struct {
	Result int
}

type StatusReport struct {
	Message string
}

*/
