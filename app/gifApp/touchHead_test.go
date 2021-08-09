package gifApp

import (
	"github.com/ssp97/ZeroBot-Plugin/pkg/OicqUtils"
	"testing"
)

func TestName(t *testing.T) {
	faceImg := OicqUtils.GetQQFaceImg(38263547)
	make(faceImg, "test.gif")
}
