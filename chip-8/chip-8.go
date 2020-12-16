package main

import (
    "fmt"
    "os"
    "encoding/binary"
    "encoding/json"
    "math/rand"
    "time"
    "net/http"
    
)

const FILE_PATH = "" //Path of the program to execute 

var memory []byte
var registers struct {
	general [16]byte
	i, pc uint16
	sp, dt, st uint8
}
var stack [16]uint16
var screen[32][64] uint8
var keys [16]bool

func reverseByte(b byte) byte {
	b = (b & 0xF0) >> 4 | (b & 0x0F) << 4
	b = (b & 0xCC) >> 2 | (b & 0x33) << 2
	b = (b & 0xAA) >> 1 | (b & 0x55) << 1
	return b
 }

func handleKeys	(w http.ResponseWriter, r* http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	switch r.URL.Query()["4"][0] {
	case "d":
		keys[4] = true
	case "u":
		keys[4] = false
	}

	switch r.URL.Query()["1"][0] {
	case "d":
		keys[1] = true
	case "u":
		keys[1] = false
	}	

	switch r.URL.Query()["5"][0] {
	case "d":
		keys[5] = true
	case "u":
		keys[5] = false
	}

	switch r.URL.Query()["6"][0] {
	case "d":
		keys[6] = true
	case "u":
		keys[6] = false
	}

	switch r.URL.Query()["7"][0] {
	case "d":
		keys[7] = true	
	case "u":
		keys[7] = false
	}
}

func mainPage(w http.ResponseWriter, r* http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	enc.Encode(screen)
}

func initWebServer() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/input", handleKeys)
	http.ListenAndServe("127.0.0.1:8080",nil)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	go initWebServer()
	
	memory = make([]byte, 4096)	
	
	fd, err := os.Open(FILE_PATH)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR:::os.Open:::%s", err)
	}
	defer fd.Close()

	_, err = fd.Read(memory[512:])
	if err != nil {		
		fmt.Fprintf(os.Stderr, "ERROR:::fd.Read:::%s", err)	
	}
	
	registers.pc = 512
	
	for {
		//start := time.Now().UnixNano()
		
		opcode := binary.BigEndian.Uint16(memory[registers.pc:registers.pc+2])
		registers.pc += 2
			
		switch opcode {
		case 0x00E0: //CLS
			for i := 0; i < 32; i++ {
				for j := 0; j < 64; j++ {
					screen[i][j] = 0
				}
			}
			//fmt.Print("\u001b[2J")
			continue
		case 0x00EE: //RET
			registers.pc = stack[registers.sp-1]
			registers.sp--
			continue
		}
		
		data := opcode & 0x0FFF
		opcode = opcode & 0xF000

		switch opcode {
		case 0xA000: //LD I, addr
			registers.i = data
		case 0x6000: //LD Vx, byte
			registers.general[(data & 0x0F00) >> 8] = byte(data & 0x00FF)
		case 0xd000: //DRW Vx, Vy, nibble
			x := registers.general[(data & 0x0F00) >> 8]
			y := registers.general[(data & 0x00F0) >> 4]
			n := data & 0x000F

			
			collisionFlag := false
			for i := uint16(0); i < n; i++ {
				b := memory[registers.i + i]
				b = reverseByte(b)
				
				shiftCounter := 0
				
				for j := uint8(1);; j *= 2 {
					
					if screen[y%32][x%64] == 0 {
						screen[y%32][x%64] = (uint8(b) & j) >> shiftCounter
					} else if (uint8(b) & j) >> shiftCounter == 1 {
						screen[y%32][x%64] = 0
						collisionFlag = true
					}
					shiftCounter++
					x++
					
					if j == 128 {
						break
					}
				} 
				y++
				
				x = registers.general[(data & 0x0F00) >> 8]
			}
			if collisionFlag {
				registers.general[15] = 1
			} else {
				registers.general[15] = 0
			}
		case 0x7000: //ADD Vx, byte
			registers.general[(data & 0x0F00)>>8] += byte(data & 0x00FF)
		case 0x4000: //SNE Vx, byte
			if registers.general[(data & 0x0F00)>>8] != byte(data & 0x00FF) {
				registers.pc += 2
			}
		case 0x1000: //JP addr
			registers.pc = data
		case 0x3000: //SE Vx, byte
			if registers.general[(data & 0x0F00)>>8] == byte(data & 0x00FF) {
				registers.pc += 2
			}
		case 0x9000: //SNE Vx, Vy
			x := (data & 0x0F00) >> 8
			y := (data & 0x00F0) >> 4
			if registers.general[x] != registers.general[y] {
				registers.pc += 2
			}
		case 0x2000: // CALL addr
			registers.sp++
			stack[registers.sp-1] = registers.pc
			registers.pc = data & 0x0FFF
		case 0xF000:
			which := data & 0x00FF
			data = (data & 0x0F00) >> 8
			switch which {
			case 0x15: //LD DT, Vx
				registers.dt = registers.general[data]
			case 0x07: //LD Vx, DT
				registers.general[data] = registers.dt
			case 0x1E:
				registers.i += uint16(registers.general[data])
			case 0x0A: //LD Vx, K
				keyFlag := false
				for !keyFlag {
					for i := 0; i < 16; i++ {
						if keys[i] {
							registers.general[data] = byte(i)
							keyFlag = true
							break
						}
					}
					time.Sleep(66)
				}
			default:
				fmt.Println("MANCA 0xF", which)
			}
		case 0x8000:
			which := data & 0x000F
			x := (data & 0x0F00) >> 8
			y := (data & 0x00F0) >> 4
			switch which {
			case 0x6: //SHR Vx {, Vy}
				if registers.general[x] & 0b00000001 == 1 {
					registers.general[15] = 1
				} else {
					registers.general[15] = 0
				}
				registers.general[x] /= 2
			case 0xe: //SHL Vx {, Vy}
				if registers.general[x] & 0b10000000 == 1 {
					registers.general[15] = 1
				} else {
					registers.general[15] = 0
				}
				registers.general[x] *= 2
			case 0x2: //AND Vx, Vy
				registers.general[x] &= registers.general[y]
			case 0x4: //ADD Vx, Vy
				if int(registers.general[x]) + int(registers.general[y]) > 255 {
					registers.general[15] = 1
				} else {
					registers.general[15] = 0
				}
				registers.general[x] += registers.general[y]
			case 0x0: //LD Vx, Vy
				registers.general[x] = registers.general[y]
			case 0x5: //SUB Vx, Vy
				if registers.general[x] > registers.general[y] {
					registers.general[15] = 1
				} else {
					registers.general[15] = 0
				}
				registers.general[x] -= registers.general[y]
			default:
				fmt.Println("MANCA 0x8", which)
			}
		case 0xE000:
			which := data & 0x00FF
			data = (data & 0x0F00) >> 8
			switch which {
			case 0xA1: //SKNP Vx
				if !keys[registers.general[data]] {
					registers.pc += 2
				}
			case 0x9E: //SKP Vx
				if keys[registers.general[data]] {
					registers.pc += 2
				}
			default:
				fmt.Println("MANCA 0xE", which)			
			}
		case 0xC000: //RND Vx, byte
			random := byte(rand.Intn(256)) & byte(data & 0x00FF)
			fmt.Println(random)
			registers.general[(data & 0x0F00) >> 8] = random
		default:
			fmt.Println("MANCA", opcode, data)
		}

		//Update timers
		if registers.dt > 0 {
			registers.dt--
		}
		if registers.st > 0 {
			registers.st--
			//TODO: play a sound
		}
		
		//Draw on screen

		//fmt.Print("\u001b[2J")
		//fmt.Print("\u001b[1;1H")
		//for i := 0; i < 32; i++ {
			//for j := 0; j < 64; j++ {
				//if screen[i][j] == 1 {
					//fmt.Print("\u25A0")
				//} else {
				//	fmt.Print(" ")
				//}
			//}
			//fmt.Println()
		//}
		//var lala string
		//fmt.Scan(&lala)
		//end := time.Now().UnixNano()
		
		time.Sleep(16 * time.Millisecond)
	}
}
