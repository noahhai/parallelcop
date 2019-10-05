package main

import (
	"fmt"
	"math"
	"time"

	"github.com/getlantern/systray"
	"github.com/shirou/gopsutil/cpu"
	"github.com/skratchdot/open-golang/open"
)

const (
	INTERVAL_MS = 1000
)

func main() {
	onExit := func() {
		fmt.Println("Starting onExit")
		// Placeholder for logging
		fmt.Println("Finished onExit")
	}
	systray.Run(onReady, onExit)
}

func onReady() {
	ticker := time.NewTicker(INTERVAL_MS * time.Millisecond)

	go func() {
		systray.SetIcon(Data0)
		systray.SetTitle("Parallel Cop")
		systray.SetTooltip("Parallel Cop")
		systray.AddMenuItem("Score 100: all cores have same load", "")
		systray.AddMenuItem("Score 0: one core has all load", "")
		mUrl := systray.AddMenuItem("Open Parallel Cop Web Page", "https://placeholder")
		mQuit := systray.AddMenuItem("Quit", "Quit")

		for {
			select {
			case <-mUrl.ClickedCh:
				open.Run("https://placeholder")
			case <-mQuit.ClickedCh:
				ticker.Stop()
				systray.Quit()
				fmt.Println("Quitting now...")
				return
			case <- ticker.C:
				if percentages, err := cpu.Percent(0, true); err == nil {
					score := getParallelScore(percentages)
					tooltip := "Parallel Cop\n"
					for _, s := range percentages {
						tooltip += fmt.Sprintf("%d%%\n", int(s))
					}
					systray.SetTooltip(tooltip)

					scoreInt := int64(score)
					if scoreInt >= 0 && scoreInt <= 100 {
						icon := getIconForScore(scoreInt)
						systray.SetIcon(icon)
					}
				}
			}
		}
	}()
}


func getParallelScore(percentages []float64) float64 {
	sum := 0.0
	for _, e := range percentages {
		sum += e
	}
	mean := sum / float64(len(percentages))

	ss := 0.0
	for _, e := range percentages {
		ss += math.Pow(mean - e, 2)
	}
	mse := math.Sqrt(ss)/float64(len(percentages)-1)
	maxScore := 100.0
	score := maxScore / ( 1 + mse/4.0)
	return score
}

func getIconForScore(score int64) []byte {
	if score < 0 || score > 100 {
		panic("invalid score")
	}
	switch score {
	case 0:
		return Data0
	case 1:
		return Data1
	case 2:
		return Data2
	case 3:
		return Data3
	case 4:
		return Data4
	case 5:
		return Data5
	case 6:
		return Data6
	case 7:
		return Data7
	case 8:
		return Data8
	case 9:
		return Data9
	case 10:
		return Data10
	case 11:
		return Data11
	case 12:
		return Data12
	case 13:
		return Data13
	case 14:
		return Data14
	case 15:
		return Data15
	case 16:
		return Data16
	case 17:
		return Data17
	case 18:
		return Data18
	case 19:
		return Data19
	case 20:
		return Data20
	case 21:
		return Data21
	case 22:
		return Data22
	case 23:
		return Data23
	case 24:
		return Data24
	case 25:
		return Data25
	case 26:
		return Data26
	case 27:
		return Data27
	case 28:
		return Data28
	case 29:
		return Data29
	case 30:
		return Data30
	case 31:
		return Data31
	case 32:
		return Data32
	case 33:
		return Data33
	case 34:
		return Data34
	case 35:
		return Data35
	case 36:
		return Data36
	case 37:
		return Data37
	case 38:
		return Data38
	case 39:
		return Data39
	case 40:
		return Data40
	case 41:
		return Data41
	case 42:
		return Data42
	case 43:
		return Data43
	case 44:
		return Data44
	case 45:
		return Data45
	case 46:
		return Data46
	case 47:
		return Data47
	case 48:
		return Data48
	case 49:
		return Data49
	case 50:
		return Data50
	case 51:
		return Data51
	case 52:
		return Data52
	case 53:
		return Data53
	case 54:
		return Data54
	case 55:
		return Data55
	case 56:
		return Data56
	case 57:
		return Data57
	case 58:
		return Data58
	case 59:
		return Data59
	case 60:
		return Data60
	case 61:
		return Data61
	case 62:
		return Data62
	case 63:
		return Data63
	case 64:
		return Data64
	case 65:
		return Data65
	case 66:
		return Data66
	case 67:
		return Data67
	case 68:
		return Data68
	case 69:
		return Data69
	case 70:
		return Data70
	case 71:
		return Data71
	case 72:
		return Data72
	case 73:
		return Data73
	case 74:
		return Data74
	case 75:
		return Data75
	case 76:
		return Data76
	case 77:
		return Data77
	case 78:
		return Data78
	case 79:
		return Data79
	case 80:
		return Data80
	case 81:
		return Data81
	case 82:
		return Data82
	case 83:
		return Data83
	case 84:
		return Data84
	case 85:
		return Data85
	case 86:
		return Data86
	case 87:
		return Data87
	case 88:
		return Data88
	case 89:
		return Data89
	case 90:
		return Data90
	case 91:
		return Data91
	case 92:
		return Data92
	case 93:
		return Data93
	case 94:
		return Data94
	case 95:
		return Data95
	case 96:
		return Data96
	case 97:
		return Data97
	case 98:
		return Data98
	case 99:
		return Data99
	case 100:
		return Data100
	}
	return nil
}
