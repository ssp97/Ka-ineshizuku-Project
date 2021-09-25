package zero

import "github.com/wdvxdr1123/ZeroBot/message"

func Cardimage(data string) message.MessageSegment {
	return message.MessageSegment{
		Type: "cardimage",
		Data: map[string]string{
			"file": data,
		},
	}
}

func Share(url, title, image string)message.MessageSegment{
	return message.MessageSegment{
		Type: "share",
		Data: map[string]string{
			"url": url,
			"title": title,
			"image": image,
		},
	}
}