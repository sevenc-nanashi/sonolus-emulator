package processor

import (
	"fmt"
	"math"
	"math/rand"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Stack struct {
	Index   int
	Data    []float64
	Current int

	Args []int
	Func string
}

type DrawCall struct {
	Id    int
	X1    float64
	Y1    float64
	X2    float64
	Y2    float64
	X3    float64
	Y3    float64
	X4    float64
	Y4    float64
	Z     float64
	Alpha float64
}

type Functions struct {
	DebugLog func(string)
}

func (p *Processor) Execute(status Status, entityIndex int, index int) float64 {
	nodes := p.data.EngineData.Nodes
	p.prepareMemory(entityIndex, status)
	return p.execute(
		entityIndex,
		[]Stack{
			{
				Index:   0,
				Data:    make([]float64, len(nodes[index].Args)),
				Current: 0,

				Args: nodes[index].Args,
				Func: nodes[index].Func,
			},
		})
}

func (p *Processor) execute(entityIndex int, stacks []Stack) float64 {
	nodes := p.data.EngineData.Nodes
	for len(stacks) > 0 {
		stack := stacks[len(stacks)-1]
		indent := strings.Repeat("  ", len(stacks)-1) + "  "
		shouldSkip := false
		for i := stack.Current; i < len(stack.Data); i++ {
			if shouldSkip {
				break
			}
			arg := stack.Args[i]
			switch stack.Func {
			case "If":
				if stack.Current == 1 { // Processed condition
					if stack.Data[0] == 0 { // False
						stack.Current = 2
						continue
					}
					shouldSkip = true
				}
				break
			case "Or":
				for i := 0; i < stack.Current; i++ {
					if stack.Data[i] != 0 {
						shouldSkip = true
						break
					}
				}
				if shouldSkip {
					break
				}
				break
			case "And":
				for i := 0; i < stack.Current; i++ {
					if stack.Data[i] == 0 {
						stack.Current = len(stack.Args)
						shouldSkip = true
						break
					}
				}
				if shouldSkip {
					break
				}
				break
			}
			// log.Infof("%sProcessing stack %s%v %v#%d", indent, stack.Func, stack.Args, stack.Data, stack.Current)

			if nodes[arg].Func == "" { // value
				stack.Data[i] = nodes[arg].Value
				stack.Current++
				// log.Infof("%s  Value %f", indent, stack.Data[i])
			} else { // function
				stacks = append(stacks, Stack{
					Index:   i,
					Data:    make([]float64, len(nodes[arg].Args)),
					Current: 0,

					Args: nodes[arg].Args,
					Func: nodes[arg].Func,
				})
				break
			}
		}
		if stack.Current == len(stack.Data) {
			var result float64 = 0
			// log.Infof("%sStack complete %s%v#%d", indent, stack.Func, stack.Data, stack.Current)
			switch stack.Func {
			case "Add":
				for _, value := range stack.Data {
					result += value
				}
				break
			case "Subtract":
				result = stack.Data[0] * 2
				for _, value := range stack.Data {
					result -= value
				}
				break
			case "Execute":
				result = stack.Data[len(stack.Data)-1]
				break
			case "Set":
				result = 0
				id := int(stack.Data[0])
				index := int(stack.Data[1])
				permissions, ok := p.memoryPermissions[id]
				if !ok {
					panic(fmt.Sprintf("Memory %d permissions not found", id))
				}
				memory, ok := p.memories[id]
				if !ok {
					if id == 22 { // EntityData
						id = 11
						index = 32*entityIndex + index
					} else if id == 24 { // EntitySharedMemory
						id = 12
						index = 32*entityIndex + index
					} else {
						panic(fmt.Sprintf("Memory %d not found", id))
					}
				}
				if !permissions.Write {
					panic(fmt.Sprintf("Memory %d is not writable", id))
				}

				if index >= len(memory) {
					panic(fmt.Sprintf("Memory %d index %d out of range %d", id, index, len(memory)))
				}
				memory[index] = stack.Data[2]
				break
			case "Get":
				id := int(stack.Data[0])
				index := int(stack.Data[1])
				permissions, ok := p.memoryPermissions[id]
				if !ok {
					panic(fmt.Sprintf("Memory %d permissions not found", id))
				}
				memory, ok := p.memories[id]
				if !ok {
					if id == 22 { // EntityData
						id = 11
						index = 32*entityIndex + index
					} else if id == 24 { // EntitySharedMemory
						id = 12
						index = 32*entityIndex + index
					} else {
						panic(fmt.Sprintf("Memory %d not found", id))
					}
				}
				if !permissions.Read {
					panic(fmt.Sprintf("Memory %d is not readable", id))
				}

				if index >= len(memory) {
					panic(fmt.Sprintf("Memory %d index %d out of range %d", id, index, len(memory)))
				}
				result = memory[index]
				break
			case "If":
				if stack.Data[0] == 0 {
					result = stack.Data[2]
				} else {
					result = stack.Data[1]
				}
				break
			case "And":
				result = stack.Data[len(stack.Data)-1]
				break
			case "Spawn":
				script := int(stack.Data[0])
				data := stack.Data[1:]

				p.spawnQueue = append(p.spawnQueue, SpawnQueueItem{
					Scripts: p.data.EngineData.Scripts[script],
					Data:    data,
				})
				break
			case "Sin":
				result = math.Sin(stack.Data[0])
				break
			case "Cos":
				result = math.Cos(stack.Data[0])
				break
			case "Tan":
				result = math.Tan(stack.Data[0])
				break
			case "Draw":
				result = 0
				p.DrawCalls = append(p.DrawCalls, DrawCall{
					Id:    int(stack.Data[0]),
					X1:    stack.Data[1],
					Y1:    stack.Data[2],
					X2:    stack.Data[3],
					Y2:    stack.Data[4],
					X3:    stack.Data[5],
					Y3:    stack.Data[6],
					X4:    stack.Data[7],
					Y4:    stack.Data[8],
					Z:     stack.Data[9],
					Alpha: stack.Data[10],
				})
				break
			case "Random":
				a := stack.Data[0]
				b := stack.Data[1]
				// x := stack.Data[2]
				r := rand.Float64()
				result = a + (b-a)*r
			case "DebugLog":
				p.DebugLog = stack.Data[0]
				break
			default:
				log.Warnf("%sUnknown function %s", indent, stack.Func)
			}
			stacks = stacks[:len(stacks)-1]
			if len(stacks) == 0 {
				return result
			}
			stacks[len(stacks)-1].Current++
			stacks[len(stacks)-1].Data[stack.Index] = result
			continue
		}
	}
	panic("Unreachable")
}
