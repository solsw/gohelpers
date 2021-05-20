package contexthelper

import (
	"context"
	"testing"
	"time"
)

func Test_combinedContext_Err(t *testing.T) {
	ctx1, _ := context.WithTimeout(context.Background(), 2*time.Second)
	ctx2, cancel2 := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Second)
		cancel2()
	}()
	tests := []struct {
		name       string
		ctx        context.Context
		wantErrMsg string
	}{
		{name: "1",
			ctx:        NewCombinedContext(ctx1, ctx2),
			wantErrMsg: "context deadline exceeded : context canceled",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			<-tt.ctx.Done()
			if errMsg := tt.ctx.Err().Error(); errMsg != tt.wantErrMsg {
				t.Errorf("errMsg '%s', wantErrMsg '%s'", errMsg, tt.wantErrMsg)
			}
		})
	}
}
