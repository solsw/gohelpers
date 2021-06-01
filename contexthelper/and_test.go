package contexthelper

import (
	"context"
	"testing"
	"time"
)

func ctxWithTimeout() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	return ctx
}

func ctxWithCancel() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Second)
		cancel()
	}()
	return ctx
}

func Test_andContext_Err(t *testing.T) {
	tests := []struct {
		name       string
		ctx        context.Context
		wantErrMsg string
	}{
		{name: "1",
			ctx:        NewAndContext(context.Background(), ctxWithCancel()),
			wantErrMsg: ":context canceled",
		},
		{name: "2",
			ctx:        NewAndContext(ctxWithTimeout(), context.Background()),
			wantErrMsg: "context deadline exceeded:",
		},
		{name: "3",
			ctx:        NewAndContext(ctxWithTimeout(), ctxWithCancel()),
			wantErrMsg: "context deadline exceeded:context canceled",
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
