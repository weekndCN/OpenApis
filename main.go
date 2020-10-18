package main

import (
	"log"

	dingtalk "github.com/OpenApis/m/dingtalk"
)

func main() {
	// You should replace the webhook here with your own.
	webhook := "https://oapi.dingtalk.com/robot/send?access_token=d55abc1615863bb85db09857132ca76c14e760d67caa0c76b8f4f69b3b5a4ea6"
	robot := dingtalk.NewRobot(webhook)

	title := "乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身"
	text := "![screenshot](@lADOpwk3K80C0M0FoA) \n #### 乔布斯 20 年前想打造的苹果咖啡厅 \n\n Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划"
	singleTitle := "阅读全文"
	singleURL := "https://www.dingtalk.com/"
	btnOrientation := "0"
	hideAvatar := "0"

	err := robot.SendActionCard(title, text, singleTitle, singleURL, btnOrientation, hideAvatar)
	if err != nil {
		log.Fatal(err)
	}
}
