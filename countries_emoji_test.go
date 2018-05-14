package countries_emoji

import (
	"fmt"
	"testing"
)

func TestChinaEmoji(t *testing.T) {
	fmt.Println(GetEmojiFlag("CN") == "🇨🇳")
}

func TestUSEmoji(t *testing.T) {
	fmt.Println(GetEmojiFlag("US") == "🇺🇸")
}

func TestSgpEmoji(t *testing.T) {
	fmt.Println(GetEmojiFlag("SG") == "🇸🇬")
}
