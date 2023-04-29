package processor

import (
	"fmt"
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

type Functions struct {
	DebugLog func(string)
}

func (p Processor) Execute(status Status, entityIndex int, index int) float64 {
	nodes := p.Data.EngineData.Nodes
	p.prepareMemory(entityIndex, status)
	return p.execute(
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

func (p Processor) execute(stacks []Stack) float64 {
	nodes := p.Data.EngineData.Nodes
	for len(stacks) > 0 {
		stack := stacks[len(stacks)-1]
		indent := strings.Repeat("  ", len(stacks)-1)
		for i, arg := range stack.Args {
			if i < stack.Current {
				continue
			}
			log.Infof("%sProcessing stack %s%v %v#%d", indent, stack.Func, stack.Args, stack.Data, stack.Current)

			if nodes[arg].Func == "" { // value
				stack.Data[i] = nodes[arg].Value
				stack.Current++
				log.Infof("%s  Value %f", indent, stack.Data[i])
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
			log.Infof("%sStack complete %s%v %v#%d", indent, stack.Func, stack.Args, stack.Data, stack.Current)
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
				memory, ok := p.Memories[id]
				if !ok {
					panic(fmt.Sprintf("Memory %d not found", id))
				}
				permissions, ok := p.MemoryPermissions[id]
				if !ok {
					panic(fmt.Sprintf("Memory %d permissions not found", id))
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
				memory, ok := p.Memories[id]
				if !ok {
					panic(fmt.Sprintf("Memory %d not found", id))
				}
				permissions, ok := p.MemoryPermissions[id]
				if !ok {
					panic(fmt.Sprintf("Memory %d permissions not found", id))
				}
				if !permissions.Read {
					panic(fmt.Sprintf("Memory %d is not readable", id))
				}

				if index >= len(memory) {
					panic(fmt.Sprintf("Memory %d index %d out of range %d", id, index, len(memory)))
				}
				result = memory[index]
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
