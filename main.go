package main

import (
	"fmt"
	"html"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func printMemu() {
	fmt.Println("\n➕➖ MATHEMATICIAN IN WONDERLAND ✖️ ➗")
	fmt.Println("-----------------------------------")
	fmt.Println("이상한 나라의 수학자 게임에 오신 것을 환영합니다.")
	fmt.Println("이 게임은 짧은 시간 안에 많은 문제를 풀어야 합니다.")
	fmt.Printf("당신의 능력을 보여주세요!\n\n")
	fmt.Println("1. 🎲 help 🎲")
	fmt.Println("2. 🔐 start test 🔐")
	fmt.Println("3. 🛫 Exit 🛬")
}

func printFlag() {
	fmt.Println("\n🎉 CONGRATULATIONS 🎉")
	fmt.Println("-----------------------------------")
	fmt.Printf("축하합니다! 플래그는 다음과 같습니다.\n\n")
	fmt.Printf("FLAG{1s_th1s_4_fl4g?}\n\n")
}

func randomEmoji() string {
	rand.Seed(time.Now().UnixNano())
	// http://apps.timwhitlock.info/emoji/tables/unicode
	emoji := [][]int{
		// Emoticons icons
		{128513, 128591},
		// Transport and map symbols
		{128640, 128704},
	}
	r := emoji[rand.Int()%len(emoji)]
	min := r[0]
	max := r[1]
	n := rand.Intn(max-min+1) + min
	return html.UnescapeString("&#" + strconv.Itoa(n) + ";")
}

func startGame() {
	var randomMap [10]struct {
		emoji  string
		number int
	}

	for i := 0; i < 10; i++ {
		randomMap[i] = struct {
			emoji  string
			number int
		}{randomEmoji(), rand.Intn(100)}
	}
	for {

		printMemu()
		var input int
		fmt.Printf("\nSelect menu >> ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			fmt.Println("\n🎲 HELP 🎲")
			fmt.Println("-----------------------------------")
			fmt.Printf("아마... 도움이 될꺼예요...?\n\n")

			for i := 0; i < 10; i++ {
				fmt.Printf("%s -> %d\n", randomMap[i].emoji, randomMap[i].number)
			}
		case 2:
			fmt.Printf("\n🔐 THE WONDERLAND TEST 🔐\n")
			fmt.Println("-----------------------------------")
			fmt.Println("500개의 문제에 대한 500개의 답이 있습니다.")
			fmt.Println("각 질문에 주어진 시간은 단 몇 ms 입니다.")
			fmt.Printf("당신은 이곳에서 어떤 답을 찾을 수 있을까요?\n\n")

			for i := 0; i < 200; i++ {
				//랜덤 이모지와 숫자 출력

				randIdx1 := rand.Intn(10)
				randIdx2 := rand.Intn(10)

				fmt.Printf("%d번째 문제: %s + %s = ", i+1, randomMap[randIdx1].emoji, randomMap[randIdx2].emoji)

				//사용자 입력
				var input int
				fmt.Scanln(&input)

				if input == randomMap[randIdx1].number+randomMap[randIdx2].number {
					fmt.Printf("🎉%d번째 문제 정답🎉\n", i+1)
				} else {
					fmt.Printf("🚫%d번째 문제 오답🚫\n", i+1)
					//프로그램 종료
					os.Exit(0)
				}
			}

			for i := 0; i < 200; i++ {

				randIdx1 := rand.Intn(10)
				randIdx2 := rand.Intn(10)
				randIdx3 := rand.Intn(10)

				fmt.Printf("%d번째 문제: %s + %s - %s = ", i+201, randomMap[randIdx1].emoji, randomMap[randIdx2].emoji, randomMap[randIdx3].emoji)

				//사용자 입력
				var input int
				fmt.Scanln(&input)

				if input == randomMap[randIdx1].number+randomMap[randIdx2].number-randomMap[randIdx3].number {
					fmt.Printf("🎉%d번째 문제 정답🎉\n", i+201)
				} else {
					fmt.Printf("🚫%d번째 문제 오답🚫\n", i+201)
					os.Exit(0)
				}
			}
			for i := 0; i < 100; i++ {
				randIdx1 := rand.Intn(10)
				randIdx2 := rand.Intn(10)
				randIdx3 := rand.Intn(10)
				randIdx4 := rand.Intn(10)

				fmt.Printf("%d번째 문제: %s + %s / (%s * %s) = ", i+401, randomMap[randIdx1].emoji, randomMap[randIdx2].emoji, randomMap[randIdx3].emoji, randomMap[randIdx4].emoji)

				//사용자 입력
				var input int
				fmt.Scanln(&input)

				if input == randomMap[randIdx1].number+randomMap[randIdx2].number/(randomMap[randIdx3].number*randomMap[randIdx4].number) {
					fmt.Printf("🎉%d번째 문제 정답🎉\n", i+401)
				} else {
					fmt.Printf("🚫%d번째 문제 오답🚫\n", i+401)
					os.Exit(0)
				}
				printFlag()
				os.Exit(0)
			}
		case 3:
			fmt.Println("\n🛫 EXIT 🛬")
			fmt.Println("-----------------------------------")
			fmt.Printf("Hava a nice day :)\n\n")
			os.Exit(0)
		default:
			fmt.Println("🎲 Invalid input 🎲")
		}
	}
}

func main() {
	go startGame()

	time.Sleep(time.Second * 90)
	fmt.Println("\n\n\n⏳ TIMEOUT ⌛")
	fmt.Println("-----------------------------------")
	fmt.Printf("시간 초과입니다.\n다시 시도해 주세요.\n\n")
}
