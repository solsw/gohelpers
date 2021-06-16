package contexthelper

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func Test_andContext_Deadline(t *testing.T) {
	tests := []struct {
		name string
		ctx  context.Context
		want time.Time
		ok   bool
	}{
		{name: "1",
			ctx: NewAndContext(context.Background(), context.Background()),
			ok:  false,
		},
		{name: "2",
			ctx:  NewAndContext(ctxWithTimeout(500*time.Millisecond), context.Background()),
			want: time.Now().Add(500 * time.Millisecond).Round(time.Millisecond),
			ok:   true,
		},
		{name: "3",
			ctx:  NewAndContext(context.Background(), ctxWithTimeout(250*time.Millisecond)),
			want: time.Now().Add(250 * time.Millisecond).Round(time.Millisecond),
			ok:   true,
		},
		{name: "4",
			ctx:  NewAndContext(ctxWithTimeout(2*time.Second), ctxWithTimeout(4*time.Second)),
			want: time.Now().Add(4 * time.Second).Round(time.Millisecond),
			ok:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotOk := tt.ctx.Deadline()
			got = got.Round(time.Millisecond)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("andContext.Deadline() got = %v, want %v", got, tt.want)
			}
			if gotOk != tt.ok {
				t.Errorf("andContext.Deadline() gotOk = %v, want %v", gotOk, tt.ok)
			}
		})
	}
}

func Test_andContext_Err(t *testing.T) {
	tests := []struct {
		name       string
		ctx        context.Context
		wantErrMsg string
	}{
		{name: "0",
			ctx:        NewAndContext(context.Background(), context.Background()),
			wantErrMsg: "",
		},
		{name: "1",
			ctx:        NewAndContext(context.Background(), ctxWithCancel(250*time.Millisecond)),
			wantErrMsg: ":" + context.Canceled.Error(),
		},
		{name: "2",
			ctx:        NewAndContext(ctxWithTimeout(500*time.Millisecond), context.Background()),
			wantErrMsg: context.DeadlineExceeded.Error() + ":",
		},
		{name: "3",
			ctx:        NewAndContext(ctxWithTimeout(500*time.Millisecond), ctxWithCancel(250*time.Millisecond)),
			wantErrMsg: context.DeadlineExceeded.Error() + ":" + context.Canceled.Error(),
		},
		{name: "4",
			ctx:        NewAndContext(ctxWithCancel(250*time.Millisecond), ctxWithTimeout(500*time.Millisecond)),
			wantErrMsg: context.Canceled.Error() + ":" + context.DeadlineExceeded.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.ctx.Done() == nil {
				if len(tt.wantErrMsg) != 0 {
					t.Errorf("errMsg0 '', wantErrMsg '%s'", tt.wantErrMsg)
				}
				return
			}
			<-tt.ctx.Done()
			if errMsg1 := tt.ctx.Err().Error(); errMsg1 != tt.wantErrMsg {
				t.Errorf("errMsg1 '%s', wantErrMsg '%s'", errMsg1, tt.wantErrMsg)
			}
			if errMsg2 := tt.ctx.Err().Error(); errMsg2 != tt.wantErrMsg {
				t.Errorf("errMsg2 '%s', wantErrMsg '%s'", errMsg2, tt.wantErrMsg)
			}
		})
	}
}
